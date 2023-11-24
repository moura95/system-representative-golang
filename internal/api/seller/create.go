package seller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	Cpf         string `json:"cpf"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Pix         string `json:"pix"`
	Observation string `json:"observation"`
}

func (s *Seller) create(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := s.Db.CreateSellers(ctx, repository.CreateSellersParams{
		RepresentativeID: representativeID,
		Cpf:              req.Cpf,
		Name:             req.Name,
		Email:            sql.NullString{String: req.Email, Valid: req.Email != ""},
		Phone:            sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Pix:              sql.NullString{String: req.Pix, Valid: req.Pix != ""},
		Observation:      sql.NullString{String: req.Observation, Valid: req.Observation != ""},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
