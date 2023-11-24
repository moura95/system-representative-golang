package lead

import (
	"my-orders/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"my-orders/internal/util"
)

type createRequest struct {
	FirstName string                     `json:"first_name"`
	LastName  string                     `json:"last_name"`
	Email     string                     `json:"email"`
	Phone     string                     `json:"phone"`
	Origin    repository.OriginLeadsEnum `json:"origin"`
}

func (u *Lead) create(ctx *gin.Context) {
	var req createRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorInvalidRequest.Error()))
		return
	}
	name := req.FirstName + " " + req.LastName

	lead, err := u.Db.CreateLead(ctx, repository.CreateLeadParams{
		Name:   name,
		Email:  req.Email,
		Phone:  req.Phone,
		Origin: req.Origin,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, util.ErrorResponse(400, "", util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse(200, lead.ID, ""))

	return
}
