package company

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type getRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getResponse struct {
	ID          int32                   `json:"id"`
	Type        repository.CompanyTypes `json:"type"`
	Cnpj        string                  `json:"cnpj"`
	Name        string                  `json:"name"`
	FantasyName string                  `json:"fantasy_name"`
	Ie          string                  `json:"ie"`
	Phone       string                  `json:"phone"`
	Email       string                  `json:"email"`
	Website     string                  `json:"web_site"`
	LogoUrl     string                  `json:"logo_url"`
	ZipCode     string                  `json:"zip_code"`
	State       string                  `json:"state"`
	City        string                  `json:"city"`
	Street      string                  `json:"street"`
	Number      string                  `json:"number"`
	IsActive    bool                    `json:"is_active"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
}

func (c *Company) get(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	company, err := c.Db.GetCompanyByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, getResponse{
		ID:          company.ID,
		Type:        company.Type,
		Cnpj:        company.Cnpj.String,
		Name:        company.Name,
		FantasyName: company.FantasyName.String,
		Ie:          company.Ie.String,
		Phone:       company.Phone.String,
		Email:       company.Email.String,
		Website:     company.Website.String,
		LogoUrl:     company.LogoUrl.String,
		ZipCode:     company.ZipCode.String,
		State:       company.State.String,
		City:        company.City.String,
		Street:      company.Street.String,
		Number:      company.Number.String,
		IsActive:    company.IsActive,
		CreatedAt:   company.CreatedAt,
		UpdatedAt:   company.UpdatedAt,
	}, ""))
	return
}
