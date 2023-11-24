package user

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	SessionID             uuid.UUID            `json:"session_id"`
	UserID                int32                `json:"user_id"`
	AccessToken           string               `json:"access_token"`
	AccessTokenExpiresAt  time.Time            `json:"access_token_expires_at"`
	RefreshToken          string               `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time            `json:"refresh_token_expires_at"`
	Plan                  repository.PlanTypes `json:"plan"`
	Permissions           []string             `json:"permissions"`
}

func (u *User) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	user, err := u.Db.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorLoginInvalidDatabase.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorLoginInvalidDatabase.Error()))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorLoginDatabase.Error()))
		return
	}

	err = u.Db.UpdateLastLogin(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorLoginDatabase.Error()))
		return
	}

	accessToken, accessPayload, err := u.TokenMaker.CreateToken(
		user.ID,
		user.RepresentativeID,
		u.Config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	refreshToken, refreshPayload, err := u.TokenMaker.CreateToken(
		user.ID,
		user.RepresentativeID,
		u.Config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	session, err := u.Db.CreateSession(ctx, repository.CreateSessionParams{
		ID:               refreshPayload.ID,
		UserID:           user.ID,
		RepresentativeID: user.RepresentativeID,
		RefreshToken:     refreshToken,
		UserAgent:        ctx.Request.UserAgent(),
		ClientIp:         ctx.ClientIP(),
		IsBlocked:        false,
		ExpiresAt:        refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}
	permissions, _ := u.Db.GetUserPermissionAndName(ctx, user.ID)

	plan, err := u.Db.GetPlanByRepresentativesID(ctx, user.RepresentativeID)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorPlanExpired.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, loginResponse{
		SessionID:             session.ID,
		UserID:                user.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		Plan:                  plan,
		Permissions:           permissions,
	}, ""))
	return
}
