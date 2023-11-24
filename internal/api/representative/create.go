package representative

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	Cnpj         string `json:"cnpj"`
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
	CPF          string `json:"cpf"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserEmail    string `json:"user_email"`
	Password     string `json:"password"`
}

func (r *Representative) create(ctx *gin.Context) {
	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	if !util.ValidateEmail(req.UserEmail) {
		ctx.JSON(400, gin.H{
			"error": "Email Invalido",
		})
		return
	}
	if !util.ValidateEmail(req.CompanyEmail) {
		ctx.JSON(400, gin.H{
			"error": "Email Invalido",
		})
		return
	}

	representative, err := r.Db.CreateRepresentative(ctx, repository.CreateRepresentativeParams{
		Cnpj:        sql.NullString{String: req.Cnpj, Valid: req.Cnpj != ""},
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
				ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
				return
			}
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	user, err := r.Db.CreateUser(ctx, repository.CreateUserParams{
		RepresentativeID: representative.ID,
		Cpf:              sql.NullString{String: req.CPF, Valid: req.CPF != ""},
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Email:            req.UserEmail,
		Password:         hashedPassword,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErroEmailInUse.Error()))
				return
			}
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	PaymentForm := []string{"A Vista", "5/10/15", "10/20/30", "15/30/45", "30/60/90"}

	for _, paymentForm := range PaymentForm {
		_, err = r.Db.CreatePaymentForm(ctx, repository.CreatePaymentFormParams{
			RepresentativeID: representative.ID,
			Name:             paymentForm,
		})
		if err != nil {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
			return
		}
	}

	_, err = r.Db.AddUserPermission(ctx, repository.AddUserPermissionParams{
		UserID:       user.ID,
		PermissionID: 3,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
				return
			}
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	PortageList := []string{"Outros", "Propria Fabrica"}

	for _, portage := range PortageList {
		_, err = r.Db.CreateCompany(ctx, repository.CreateCompanyParams{
			Name: portage,
			FantasyName: sql.NullString{
				String: portage,
				Valid:  true},
			RepresentativeID: representative.ID,
			Type:             "Portage",
			Cnpj:             sql.NullString{String: "", Valid: true},
			Ie: sql.NullString{
				String: "",
				Valid:  true},
			Phone: sql.NullString{String: "",
				Valid: true},

			Email: sql.NullString{String: "",
				Valid: true},

			Website: sql.NullString{String: "",
				Valid: true},
			ZipCode: sql.NullString{String: "",
				Valid: true},
			State: sql.NullString{String: "",
				Valid: true},
			City: sql.NullString{String: "",
				Valid: true},
			Street: sql.NullString{String: "",
				Valid: true},
			Number: sql.NullString{String: "", Valid: true},
		})
		if err != nil {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
			return
		}
	}
	_, err = r.Db.CreateSellers(ctx, repository.CreateSellersParams{
		RepresentativeID: representative.ID,
		Name:             req.FirstName + " " + req.LastName,
		Cpf:              req.CPF,
		Phone:            sql.NullString{String: "", Valid: true},
		Email:            sql.NullString{String: req.UserEmail, Valid: req.UserEmail != ""},
		Pix:              sql.NullString{String: "", Valid: true},
		Observation:      sql.NullString{String: "", Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
