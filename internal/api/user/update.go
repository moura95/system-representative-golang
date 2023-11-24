package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type updateRequest struct {
	ID        int32  `json:"id"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	CPF       string `json:"cpf"`
}

func (u *User) update(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	_, err = u.Db.UpdateUserByID(ctx, repository.UpdateUserByIDParams{
		ID:        req.ID,
		Password:  sql.NullString{String: hashedPassword, Valid: req.Password != ""},
		FirstName: sql.NullString{String: req.FirstName, Valid: req.FirstName != ""},
		LastName:  sql.NullString{String: req.LastName, Valid: req.LastName != ""},
		Phone:     sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Cpf:       sql.NullString{String: req.CPF, Valid: req.CPF != ""},
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
				return
			}
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
