package imports

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createCompaniesResponse struct {
	Success int32          `json:"success"`
	Failure companyFailure `json:"failure"`
}

type companyFailure struct {
	Total   int       `json:"total"`
	Company []company `json:"data"`
}

type company struct {
	CNPJ          string `csv:"cnpj" json:"cnpj"`
	Name          string `csv:"nome" json:"name"`
	Type          string `csv:"tipo" json:"type"`
	FantasyName   string `csv:"nome_fantasia" json:"fantasyName"`
	IE            string `csv:"ie" json:"ie"`
	Phone         string `csv:"telefone" json:"phone"`
	Email         string `csv:"email" json:"email"`
	ZipCode       string `csv:"cep" json:"zipCode"`
	State         string `csv:"estado" json:"state"`
	City          string `csv:"cidade" json:"city"`
	Street        string `csv:"rua" json:"street"`
	Number        string `csv:"numero" json:"number"`
	FailedCsvLine int    `json:"failedCsvLine"`
}

func (i *Import) importCompanies(ctx *gin.Context) {
	var (
		reprId      = ctx.Keys["representativeID"].(int32)
		listCompany []*company
	)

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

	err = unmarshalCsvRequestBody(requestBody, &listCompany)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	mapCompany := make(map[string]string)

	// Adding key-value pairs to the map
	mapCompany["Fabrica"] = "Factory"
	mapCompany["Transportadora"] = "Portage"
	mapCompany["Cliente"] = "Customer"

	var resp createCompaniesResponse
	for j, lc := range listCompany {
		typeCompany := mapCompany[lc.Type]

		p, errCreate := i.Db.CreateCompany(ctx, repository.CreateCompanyParams{
			RepresentativeID: reprId,
			Type:             repository.CompanyTypes(typeCompany),
			Name:             lc.Name,
			Cnpj:             sqlNullString(lc.CNPJ, true),
			Email:            sqlNullString(lc.Email, lc.Email != ""),
			Street:           sqlNullString(lc.Street, lc.Street != ""),
			Number:           sqlNullString(lc.Number, lc.Number != ""),
			City:             sqlNullString(lc.City, lc.City != ""),
			State:            sqlNullString(lc.State, lc.State != ""),
			ZipCode:          sqlNullString(lc.ZipCode, lc.ZipCode != ""),
			FantasyName:      sqlNullString(lc.FantasyName, lc.FantasyName != ""),
			Ie:               sqlNullString(lc.IE, lc.IE != ""),
			Phone:            sqlNullString(lc.Phone, lc.Phone != ""),
		})
		if errCreate != nil {
			failedCsvLine := j + 2
			resp.Failure.Company = append(resp.Failure.Company, company{
				CNPJ:          lc.CNPJ,
				Name:          p.Name,
				FantasyName:   lc.FantasyName,
				IE:            lc.IE,
				Phone:         lc.Phone,
				Email:         lc.Email,
				ZipCode:       lc.ZipCode,
				State:         lc.State,
				City:          lc.City,
				Street:        lc.Street,
				Number:        lc.Number,
				FailedCsvLine: failedCsvLine,
			})
			resp.Failure.Total = len(resp.Failure.Company)
			log.Printf(errCreate.Error())
			continue
		}
		resp.Success++
	}
	if resp.Success != 0 {
		resp.Success++
	}

	ctx.AbortWithStatusJSON(http.StatusOK, util.SuccessResponse(200, resp, ""))
	return
}
