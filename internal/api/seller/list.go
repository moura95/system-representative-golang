package seller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type listRequest struct {
	IsActive bool `form:"is_active"`
}

func (s *Seller) list(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)
	var req listRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	arg := repository.ListSellersByRepresentativeIDParams{
		RepresentativeID: representativeID,
		IsActive:         req.IsActive,
	}
	listSeller, err := s.Db.ListSellersByRepresentativeID(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	if len(listSeller) == 0 {
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, listSeller, ""))
		return
	}

	var listResponse []GetResponse
	for _, seller := range listSeller {
		listResponse = append(listResponse, GetResponse{
			ID:    seller.ID,
			Cpf:   seller.Cpf,
			Name:  seller.Name,
			Email: seller.Email.String,
			Phone: seller.Phone.String,
		})
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, listResponse, ""))
	return
}
