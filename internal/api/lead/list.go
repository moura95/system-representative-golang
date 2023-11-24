package lead

import (
	"github.com/gin-gonic/gin"
	"my-orders/internal/util"
	"net/http"
)

func (l *Lead) list(ctx *gin.Context) {

	leads, err := l.Db.ListLeads(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseRead.Error()))
		return
	}

	var listResponse []getResponse
	for _, lead := range leads {
		listResponse = append(listResponse, getResponse{
			ID:        lead.ID,
			Name:      lead.Name,
			Email:     lead.Email,
			Phone:     lead.Phone,
			Origin:    lead.Origin,
			IsActive:  lead.IsActive,
			CreatedAt: lead.CreatedAt,
			UpdatedAt: lead.UpdatedAt,
		})

	}
	ctx.JSON(http.StatusOK, util.SuccessResponse(200, listResponse, ""))
}
