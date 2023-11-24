package representative

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type getResponse struct {
	ID          int32                `json:"id"`
	Cnpj        string               `json:"cnpj"`
	Name        string               `json:"company_name"`
	FantasyName string               `json:"fantasy_name"`
	Ie          string               `json:"ie"`
	Phone       string               `json:"company_phone"`
	Email       string               `json:"company_email"`
	Website     string               `json:"web_site"`
	LogoUrl     string               `json:"logo_url"`
	ZipCode     string               `json:"zip_code"`
	State       string               `json:"state"`
	City        string               `json:"city"`
	Street      string               `json:"street"`
	Number      string               `json:"number"`
	Plan        repository.PlanTypes `json:"plan"`
	DataExpire  time.Time            `json:"data_expire"`
	IsActive    bool                 `json:"is_active"`
}

func (r *Representative) get(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	rep, err := r.Db.GetRepresentativesByID(ctx, representativeID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, getResponse{
		ID:          rep.ID,
		Cnpj:        rep.Cnpj.String,
		Name:        rep.Name.String,
		FantasyName: rep.FantasyName.String,
		Ie:          rep.Ie.String,
		Phone:       rep.Phone.String,
		Email:       rep.Email.String,
		Website:     rep.Website.String,
		LogoUrl:     rep.LogoUrl.String,
		ZipCode:     rep.ZipCode.String,
		State:       rep.State.String,
		City:        rep.City.String,
		Street:      rep.Street.String,
		Number:      rep.Number.String,
		Plan:        rep.Plan,
		DataExpire:  rep.DataExpire,
		IsActive:    rep.IsActive,
	}, ""))
}
