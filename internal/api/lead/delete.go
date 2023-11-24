package lead

import (
	"github.com/gin-gonic/gin"
	"my-orders/internal/util"
	"net/http"
)

type deleteRequest struct {
	ID int32 `uri:"id" binding:"required,numeric"`
}

func (l *Lead) delete(ctx *gin.Context) {
	var req deleteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := l.Db.DeleteLeadByID(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
}
