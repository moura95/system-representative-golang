package smtp

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

func (s *Smtp) delete(ctx *gin.Context) {
	representativeID := ctx.Keys["representativeID"].(int32)

	_, err := s.Db.DeleteSmtpByRepresentativeID(ctx, representativeID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseDelete.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseDelete.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
