package product

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	FactoryID   int32  `json:"factory_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Price       string `json:"price"`
	Ipi         string `json:"ipi"`
	Reference   string `json:"reference"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

func (p *Product) create(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := p.Db.CreateProduct(ctx, repository.CreateProductParams{
		RepresentativeID: representativeID,
		FactoryID:        req.FactoryID,
		Name:             req.Name,
		Code:             req.Code,
		Price:            req.Price,
		Ipi:              sql.NullString{String: req.Ipi, Valid: req.Ipi != ""},
		Reference:        sql.NullString{String: req.Reference, Valid: req.Reference != ""},
		Description:      sql.NullString{String: req.Description, Valid: req.Description != ""},
		ImageUrl:         sql.NullString{String: req.ImageUrl, Valid: req.ImageUrl != ""},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
