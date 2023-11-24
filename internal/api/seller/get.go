package seller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

type getRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type GetResponse struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Cpf         string `json:"cpf"`
	Pix         string `json:"pix"`
	Observation string `json:"observation"`
	IsActive    bool   `json:"is_active"`
}

func (s *Seller) get(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	seller, err := s.Db.GetSellerByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, GetResponse{
		ID:          seller.ID,
		Name:        seller.Name,
		Email:       seller.Email.String,
		Phone:       seller.Phone.String,
		Cpf:         seller.Cpf,
		Pix:         seller.Pix.String,
		Observation: seller.Observation.String,
		IsActive:    seller.IsActive,
	}, ""))
	return
}
