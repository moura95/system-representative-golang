package product

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type updateRequest struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Price       string `json:"price"`
	Ipi         string `json:"ipi"`
	Reference   string `json:"reference"`
	Unity       string `json:"unity"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

func (p *Product) update(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := p.Db.UpdateProductByID(ctx, repository.UpdateProductByIDParams{
		ID:          req.ID,
		Name:        sql.NullString{String: req.Name, Valid: req.Name != ""},
		Code:        sql.NullString{String: req.Code, Valid: req.Code != ""},
		Price:       sql.NullString{String: req.Price, Valid: req.Price != ""},
		Ipi:         sql.NullString{String: req.Ipi, Valid: req.Ipi != ""},
		Reference:   sql.NullString{String: req.Reference, Valid: req.Reference != ""},
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		ImageUrl:    sql.NullString{String: req.ImageUrl, Valid: req.ImageUrl != ""},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
