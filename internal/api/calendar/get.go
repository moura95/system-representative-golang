package calendar

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

type getRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

type GetResponse struct {
	ID         int32     `json:"id"`
	Title      string    `json:"title"`
	VisitStart time.Time `json:"visit_start"`
	VisitEnd   time.Time `json:"visit_end"`
	AllDay     bool      `json:"allDay"`
}

func (c *Calendar) get(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	cal, err := c.Db.GetCalendarByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
			return
		}
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, GetResponse{
		ID:         cal.ID,
		Title:      cal.Title,
		VisitStart: cal.VisitStart,
		VisitEnd:   cal.VisitEnd,
		AllDay:     cal.Allday,
	}, ""))
	return
}
