package company

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type listRequest struct {
	Type     string `form:"type"`
	IsActive bool   `form:"is_active"`
}

func (c *Company) list(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req listRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	listCompanies, err := c.Db.ListCompaniesByRepresentativeID(ctx, repository.ListCompaniesByRepresentativeIDParams{
		RepresentativeID: representativeID,
		Type:             repository.CompanyTypes(req.Type),
		IsActive:         req.IsActive,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	if len(listCompanies) == 0 {
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, listCompanies, ""))
		return
	}

	var listResponse []getResponse
	for _, seller := range listCompanies {
		listResponse = append(listResponse, getResponse{
			ID:          seller.ID,
			Type:        seller.Type,
			Cnpj:        seller.Cnpj.String,
			Name:        seller.Name,
			FantasyName: seller.FantasyName.String,
			Ie:          seller.Ie.String,
			Phone:       seller.Phone.String,
			Email:       seller.Email.String,
			Website:     seller.Website.String,
			LogoUrl:     seller.LogoUrl.String,
			ZipCode:     seller.ZipCode.String,
			State:       seller.State.String,
			City:        seller.City.String,
			Street:      seller.Street.String,
			Number:      seller.Number.String,
			IsActive:    seller.IsActive,
			CreatedAt:   seller.CreatedAt,
			UpdatedAt:   seller.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, listResponse, ""))
	return
}
