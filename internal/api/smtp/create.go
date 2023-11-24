package smtp

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     int32  `json:"port"`
}

func (s *Smtp) create(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := s.Db.CreateSmtp(ctx, repository.CreateSmtpParams{
		Email:            req.Email,
		Password:         req.Password,
		Server:           req.Server,
		Port:             req.Port,
		RepresentativeID: representativeID,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
