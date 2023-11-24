package seller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type updateRequest struct {
	ID          int32  `json:"id"`
	Cpf         string `json:"cpf"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Pix         string `json:"pix"`
	Observation string `json:"observation"`
}

func (s *Seller) update(ctx *gin.Context) {

	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := s.Db.UpdateSellerByID(ctx, repository.UpdateSellerByIDParams{
		ID:          req.ID,
		Cpf:         sql.NullString{String: req.Cpf, Valid: req.Cpf != ""},
		Name:        sql.NullString{String: req.Name, Valid: req.Name != ""},
		Phone:       sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Email:       sql.NullString{String: req.Email, Valid: req.Email != ""},
		Pix:         sql.NullString{String: req.Pix, Valid: req.Pix != ""},
		Observation: sql.NullString{String: req.Observation, Valid: req.Observation != ""},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
