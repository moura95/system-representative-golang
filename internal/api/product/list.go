package product

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type listRequest struct {
	FactoryID int32 `form:"factory_id"`
	IsActive  bool  `form:"is_active"`
}

func (p *Product) list(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req listRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	products, err := p.Db.ListProductsByRepresentativeID(ctx, repository.ListProductsByRepresentativeIDParams{
		IsActive:         req.IsActive,
		RepresentativeID: representativeID,
		FactoryID:        sql.NullInt32{Int32: req.FactoryID, Valid: req.FactoryID != 0},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseDelete.Error()))
		return
	}

	if len(products) == 0 {
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, products, ""))
		return
	}

	var listRespnse []getResponse
	for _, product := range products {
		listRespnse = append(listRespnse, getResponse{
			ID:          product.ID,
			FactoryName: product.FactoryName,
			FactoryID:   product.FactoryID,
			Name:        product.Name,
			Code:        product.Code,
			Price:       product.Price,
			Ipi:         product.Ipi.String,
			Reference:   product.Reference.String,
			Description: product.Description.String,
			ImageUrl:    product.ImageUrl.String,
			IsActive:    product.IsActive,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, listRespnse, ""))
	return
}
