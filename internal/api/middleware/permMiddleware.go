package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/repository"
	"my-orders/internal/util"
)

var accessRules = map[string][]string{
	"/calendar":       {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/company":        {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/dashboard":      {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/orders":         {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/token":          {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/product":        {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/representative": {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/seller":         {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/smtp":           {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/order":          {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/payment_form":   {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/user":           {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/stripe":         {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/import":         {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
	"/file":           {"Admin", "Staff", "Admin_Representative", "Director", "Financial", "Secretary", "Seller"},
}

func PermMiddleware(db repository.Querier) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Keys["userID"].(int32)
		perm, err := db.GetUserPermissionAndName(ctx, userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(403, "", util.ErrorInvalidRequest.Error()))
			return
		}

		baseRoute := ctx.Request.URL.Path
		allowedRoles := accessRules[baseRoute]

		for _, role := range perm {
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					ctx.Next()
					return
				}
			}
		}
		ctx.AbortWithStatusJSON(http.StatusOK, util.ErrorResponse(403, "", util.ErrorPermNotEnough.Error()))
		return
	}
}
