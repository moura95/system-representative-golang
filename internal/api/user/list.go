package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type listRequest struct {
	IsActive bool `form:"is_active"`
}

func (u *User) list(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)
	var req listRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	listUser, err := u.Db.ListUsersByRepresentativeID(ctx, repository.ListUsersByRepresentativeIDParams{
		RepresentativeID: representativeID,
		IsActive:         req.IsActive,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	if len(listUser) == 0 {
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, listUser, ""))
		return
	}

	var listResponse []getResponse
	for _, user := range listUser {
		listResponse = append(listResponse, getResponse{
			ID:        user.ID,
			Cpf:       user.Cpf.String,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone.String,
		})
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, listResponse, ""))
	return
}
