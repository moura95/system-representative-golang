package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	CPF       string `json:"cpf"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

func (u *User) create(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)
	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	if !util.ValidateEmail(req.Email) {
		ctx.JSON(400, gin.H{
			"error": "Email Invalido",
		})
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	user, err := u.Db.CreateUser(ctx, repository.CreateUserParams{
		RepresentativeID: representativeID,
		Cpf:              sql.NullString{String: req.CPF, Valid: req.CPF != ""},
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Email:            req.Email,
		Password:         hashedPassword,
		Phone:            sql.NullString{String: req.Phone, Valid: req.Phone != ""},
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErroEmailInUse.Error()))
				return
			}
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	_, err = u.Db.AddUserPermission(ctx, repository.AddUserPermissionParams{
		UserID:       user.ID,
		PermissionID: 6,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
				return
			}
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
