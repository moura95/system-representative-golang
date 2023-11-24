package imports

import (
	"database/sql"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createProductsRequest struct {
	ID int32 `uri:"factoryId"`
}

type createProductsResponse struct {
	Success int32           `json:"success"`
	Failure productsFailure `json:"productsFailure"`
}

type productsFailure struct {
	Total   int       `json:"total"`
	Product []product `json:"data"`
}

type product struct {
	Name          string `csv:"nome" json:"name"`
	Code          string `csv:"codigo" json:"code"`
	Price         string `csv:"preco" json:"price"`
	IPI           string `csv:"ipi" json:"ipi"`
	Reference     string `csv:"referencia" json:"reference"`
	Description   string `csv:"descricao" json:"description"`
	FailedCsvLine int    `json:"failedCsvLine"`
}

func (i *Import) importProducts(ctx *gin.Context) {
	var (
		reprId       = ctx.Keys["representativeID"].(int32)
		req          createProductsRequest
		listProducts []*product
	)

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	// Check that the request contains a CSV file
	if ctx.GetHeader("Content-Type") != "text/csv" || ctx.GetHeader("Content-Length") == "0" {
		ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	// Read the request body
	requestBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	err = unmarshalCsvRequestBody(requestBody, &listProducts)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	fac, err := i.Db.GetCompanyUserByID(ctx, repository.GetCompanyUserByIDParams{RepresentativeID: reprId, CompanyID: req.ID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	var resp createProductsResponse
	for j, lp := range listProducts {
		p, errCreate := i.Db.CreateProduct(ctx, repository.CreateProductParams{
			RepresentativeID: fac.RepresentativeID,
			FactoryID:        fac.ID,
			Name:             lp.Name,
			Code:             lp.Code,
			Price:            lp.Price,
			Ipi:              sql.NullString{String: lp.IPI, Valid: lp.IPI != ""},
			Reference:        sql.NullString{String: lp.Reference, Valid: lp.Reference != ""},
			Description:      sql.NullString{String: lp.Description, Valid: lp.Description != ""},
		})
		if errCreate != nil {
			failedCsvLine := j + 2
			resp.Failure.Product = append(resp.Failure.Product, product{
				Name:          p.Name,
				Code:          p.Code,
				Price:         p.Price,
				IPI:           p.Ipi.String,
				Reference:     p.Reference.String,
				Description:   p.Description.String,
				FailedCsvLine: failedCsvLine,
			})
			resp.Failure.Total = len(resp.Failure.Product)
			continue
		}
		resp.Success++
	}
	if resp.Success != 0 {
		resp.Success++
	}

	ctx.AbortWithStatusJSON(http.StatusOK, util.SuccessResponse(200, resp, ""))
}
