package paseto

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (t *Token) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	refreshPayload, err := t.TokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorTokenExpired.Error()))
		return
	}

	session, err := t.Db.GetSessionByID(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	if session.IsBlocked {
		_ = fmt.Errorf("blocked session")
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorSessionBlocked.Error()))
		return
	}

	if session.UserID != refreshPayload.UserID {
		_ = fmt.Errorf("incorrect session user")
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorSessionBlocked.Error()))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		_ = fmt.Errorf("mismatched session token")
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorSessionBlocked.Error()))
		return
	}

	if time.Now().After(session.ExpiresAt) {
		_ = fmt.Errorf("expired session")
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorTokenExpired.Error()))
		return
	}

	accessToken, accessPayload, err := t.TokenMaker.CreateToken(
		refreshPayload.UserID,
		refreshPayload.RepresentativeID,
		t.Config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}, ""))
}
