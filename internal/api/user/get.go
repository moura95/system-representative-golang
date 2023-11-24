package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

type getRequest struct {
	ID int32 `uri:"id" binding:"required,numeric"`
}

type getResponse struct {
	ID        int32  `json:"id"`
	Cpf       string `json:"cpf"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IsActive  bool   `json:"is_active"`
	Phone     string `json:"phone"`
}

func (u *User) get(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	user, err := u.Db.GetUserByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, getResponse{
		ID:        user.ID,
		Cpf:       user.Cpf.String,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		IsActive:  user.IsActive,
		Phone:     user.Phone.String,
	}, ""))
	return
}
