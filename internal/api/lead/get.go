package lead

import (
	"github.com/gin-gonic/gin"
	"my-orders/internal/repository"
	"my-orders/internal/util"
	"net/http"
	"time"
)

type getRequest struct {
	ID int32 `uri:"id" binding:"required,numeric"`
}

type getResponse struct {
	ID        int32                      `json:"id"`
	Name      string                     `json:"name"`
	Email     string                     `json:"email"`
	Phone     string                     `json:"phone"`
	Origin    repository.OriginLeadsEnum `json:"origin"`
	IsActive  bool                       `json:"is_active"`
	CreatedAt time.Time                  `json:"created_at"`
	UpdatedAt time.Time                  `json:"updated_at"`
}

func (l *Lead) get(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	lead, err := l.Db.GetLeadByID(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, getResponse{
		ID:        lead.ID,
		Name:      lead.Name,
		Email:     lead.Email,
		Phone:     lead.Phone,
		Origin:    lead.Origin,
		IsActive:  lead.IsActive,
		CreatedAt: lead.CreatedAt,
		UpdatedAt: lead.UpdatedAt,
	}, ""))
}
