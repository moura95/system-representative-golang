package smtp

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type updateRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Server       string `json:"server"`
	Port         int32  `json:"port"`
	Cryptography string `json:"cryptography"`
}

func (s *Smtp) update(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	if req.Password != "" {
		hashedPassword, err := util.HashPassword(req.Password)
		if err != nil {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
			return
		}
		req.Password = hashedPassword
	}

	_, err := s.Db.UpdateSmtpByID(ctx, repository.UpdateSmtpByIDParams{
		RepresentativeID: representativeID,
		Email:            sql.NullString{String: req.Email, Valid: req.Email != ""},
		Password:         sql.NullString{String: req.Password, Valid: req.Password != ""},
		Server:           sql.NullString{String: req.Server, Valid: req.Server != ""},
		Port:             sql.NullInt32{Int32: req.Port, Valid: req.Port != 0},
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "OK", ""))
	return
}
