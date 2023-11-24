package calendar

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/util"
)

func (c *Calendar) list(ctx *gin.Context) {
	userID := ctx.Keys["userID"].(int32)

	listCalendars, err := c.Db.ListCalendarsByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	if len(listCalendars) == 0 {
		ctx.JSON(http.StatusOK, util.SuccessResponse(200, listCalendars, ""))
		return
	}

	var listResponse []GetResponse
	for _, calendar := range listCalendars {
		listResponse = append(listResponse, GetResponse{
			ID:         calendar.ID,
			Title:      calendar.Title,
			VisitStart: calendar.VisitStart,
			VisitEnd:   calendar.VisitEnd,
			AllDay:     calendar.Allday,
		})
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, listResponse, ""))
	return
}
