package company

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
	Email       string `json:"email"`
	WebSite     string `json:"web_site"`
	LogoUrl     string `json:"logo_url"`
	Street      string `json:"street"`
	Number      string `json:"number"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	Cnpj        string `json:"cnpj"`
	FantasyName string `json:"fantasy_name"`
	Ie          string `json:"ie"`
	Phone       string `json:"phone"`
}

func (c *Company) update(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := c.Db.UpdateCompanyByID(ctx, repository.UpdateCompanyByIDParams{
		ID:          req.ID,
		Name:        sql.NullString{String: req.Name, Valid: req.Name != ""},
		Email:       sql.NullString{String: req.Email, Valid: req.Email != ""},
		Website:     sql.NullString{String: req.WebSite, Valid: req.WebSite != ""},
		LogoUrl:     sql.NullString{String: req.LogoUrl, Valid: req.LogoUrl != ""},
		Street:      sql.NullString{String: req.Street, Valid: req.Street != ""},
		Number:      sql.NullString{String: req.Number, Valid: req.Number != ""},
		City:        sql.NullString{String: req.City, Valid: req.City != ""},
		State:       sql.NullString{String: req.State, Valid: req.State != ""},
		ZipCode:     sql.NullString{String: req.ZipCode, Valid: req.ZipCode != ""},
		Cnpj:        sql.NullString{String: req.Cnpj, Valid: req.Cnpj != ""},
		FantasyName: sql.NullString{String: req.FantasyName, Valid: req.FantasyName != ""},
		Ie:          sql.NullString{String: req.Ie, Valid: req.Ie != ""},
		Phone:       sql.NullString{String: req.Phone, Valid: req.Phone != ""},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
