package company

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	Type        string `json:"type"`
	Cnpj        string `json:"cnpj"`
	Name        string `json:"name"`
	FantasyName string `json:"fantasy_name"`
	Phone       string `json:"phone"`
	Ie          string `json:"ie"`
	Email       string `json:"email"`
	WebSite     string `json:"web_site"`
	LogoUrl     string `json:"logo_url"`
	ZipCode     string `json:"zip_code"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street      string `json:"street"`
	Number      string `json:"number"`
}

func (c *Company) create(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := c.Db.CreateCompany(ctx, repository.CreateCompanyParams{
		RepresentativeID: representativeID,
		Type:             repository.CompanyTypes(req.Type),
		Cnpj:             sql.NullString{String: req.Cnpj, Valid: true},
		Name:             req.Name,
		FantasyName:      sql.NullString{String: req.FantasyName, Valid: req.FantasyName != ""},
		Ie:               sql.NullString{String: req.Ie, Valid: req.Ie != ""},
		Phone:            sql.NullString{String: req.Phone, Valid: req.Phone != ""},
		Email:            sql.NullString{String: req.FantasyName, Valid: req.FantasyName != ""},
		Website:          sql.NullString{String: req.WebSite, Valid: req.WebSite != ""},
		LogoUrl:          sql.NullString{String: req.LogoUrl, Valid: req.LogoUrl != ""},
		ZipCode:          sql.NullString{String: req.ZipCode, Valid: req.ZipCode != ""},
		State:            sql.NullString{String: req.State, Valid: req.State != ""},
		City:             sql.NullString{String: req.City, Valid: req.City != ""},
		Street:           sql.NullString{String: req.Street, Valid: req.Street != ""},
		Number:           sql.NullString{String: req.Number, Valid: req.Number != ""},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
