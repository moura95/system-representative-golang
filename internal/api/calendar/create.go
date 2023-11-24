package calendar

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type createRequest struct {
	Title      string    `json:"title"`
	VisitStart time.Time `json:"visit_start"`
	VisitEnd   time.Time `json:"visit_end"`
	AllDay     bool      `json:"allDay"`
}

func (c *Calendar) create(ctx *gin.Context) {
	userID := ctx.Keys["userID"].(int32)
	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := c.Db.CreateCalendar(ctx, repository.CreateCalendarParams{
		UserID:     userID,
		Title:      req.Title,
		VisitStart: req.VisitStart,
		VisitEnd:   req.VisitEnd,
		Allday:     req.AllDay,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
