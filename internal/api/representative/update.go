package representative

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type updateRequest struct {
	CompanyName  string `json:"company_name"`
	FantasyName  string `json:"fantasy_name"`
	Ie           string `json:"ie"`
	CompanyPhone string `json:"company_phone"`
	CompanyEmail string `json:"company_email"`
	Website      string `json:"web_site"`
	LogoUrl      string `json:"logo_url"`
	ZipCode      string `json:"zip_code"`
	State        string `json:"state"`
	City         string `json:"city"`
	Street       string `json:"street"`
	Number       string `json:"number"`
}

func (r *Representative) update(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := r.Db.UpdateRepresentativeByID(ctx, repository.UpdateRepresentativeByIDParams{
		ID:          representativeID,
		Name:        sql.NullString{String: req.CompanyName, Valid: req.CompanyName != ""},
		FantasyName: sql.NullString{String: req.FantasyName, Valid: req.FantasyName != ""},
		Ie:          sql.NullString{String: req.Ie, Valid: req.Ie != ""},
		Phone:       sql.NullString{String: req.CompanyPhone, Valid: req.CompanyPhone != ""},
		Email:       sql.NullString{String: req.CompanyEmail, Valid: req.CompanyEmail != ""},
		Website:     sql.NullString{String: req.Website, Valid: req.Website != ""},
		LogoUrl:     sql.NullString{String: req.LogoUrl, Valid: req.LogoUrl != ""},
		ZipCode:     sql.NullString{String: req.ZipCode, Valid: req.ZipCode != ""},
		State:       sql.NullString{String: req.State, Valid: req.State != ""},
		City:        sql.NullString{String: req.City, Valid: req.City != ""},
		Street:      sql.NullString{String: req.Street, Valid: req.Street != ""},
		Number:      sql.NullString{String: req.Number, Valid: req.Number != ""},
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
				return
			}
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
