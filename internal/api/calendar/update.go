package calendar

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

type updateRequest struct {
	ID         int32     `json:"id" binding:"required"`
	Title      string    `json:"title"`
	VisitStart time.Time `json:"visit_start"`
	VisitEnd   time.Time `json:"visit_end"`
	AllDay     bool      `json:"allDay"`
}

func (c *Calendar) update(ctx *gin.Context) {
	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}

	_, err := c.Db.UpdateCalendarByID(ctx, repository.UpdateCalendarByIDParams{
		ID:         req.ID,
		Title:      sql.NullString{String: req.Title, Valid: req.Title != ""},
		VisitStart: sql.NullTime{Time: req.VisitStart, Valid: !req.VisitStart.IsZero()},
		VisitEnd:   sql.NullTime{Time: req.VisitEnd, Valid: !req.VisitEnd.IsZero()},
		Allday:     sql.NullBool{Bool: req.AllDay, Valid: req.AllDay != false},
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseUpdate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, "ok", ""))
	return
}
