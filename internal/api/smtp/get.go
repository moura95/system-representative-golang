package smtp

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

type getResponse struct {
	Email        string `json:"email,omitempty"`
	Server       string `json:"server,omitempty"`
	Port         int32  `json:"port,omitempty"`
	Cryptography string `json:"cryptography,omitempty"`
}

func (s *Smtp) get(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	smtp, err := s.Db.GetSmtpByRepresentativeID(ctx, representativeID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.SuccessResponse(200, []string{}, ""))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, getResponse{
		Email:  smtp.Email,
		Server: smtp.Server,
		Port:   smtp.Port,
	}, ""))
	return
}
