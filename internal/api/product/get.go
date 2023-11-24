package product

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

type GetRequest struct {
	ID int32 `uri:"id" binding:"required,numeric"`
}

type getResponse struct {
	ID          int32     `json:"id"`
	FactoryName string    `json:"factory_name"`
	FactoryID   int32     `json:"factory_id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Price       string    `json:"price"`
	Ipi         string    `json:"ipi"`
	Reference   string    `json:"reference"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) get(ctx *gin.Context) {
	var req GetRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	product, err := p.Db.GetProductByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, getResponse{
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
	}, ""))
}
