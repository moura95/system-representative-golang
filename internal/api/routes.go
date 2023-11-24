package api

import (
	"my-orders/internal/api/lead"
	paymentReceipt "my-orders/internal/api/payment_receipt"
	"net/http"

	"github.com/gin-gonic/gin"

	"my-orders/internal/api/file"
	"my-orders/internal/api/imports"

	"my-orders/internal/api/calendar"
	"my-orders/internal/api/company"
	"my-orders/internal/api/dashboard"
	"my-orders/internal/api/middleware"
	"my-orders/internal/api/order"
	"my-orders/internal/api/order_items"
	"my-orders/internal/api/paseto"
	"my-orders/internal/api/payment_form"
	"my-orders/internal/api/product"
	"my-orders/internal/api/representative"
	"my-orders/internal/api/seller"
	"my-orders/internal/api/smtp"
	"my-orders/internal/api/stripe"
	"my-orders/internal/api/user"
)

func (s *Server) createRoutesV1(router *gin.Engine) {
	router.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	routes := router.Group("/")

	authRoutes := router.Group("/", middleware.AuthMiddleware(*s.tokenMaker))
	planRoutes := router.Group("/", middleware.AuthMiddleware(*s.tokenMaker), middleware.PlanMiddleware(*s.store))

	calendar.ICalendar(&calendar.Calendar{Db: *s.store}).SetupCalendarRoute(planRoutes)
	company.ICompany(&company.Company{Db: *s.store}).SetupCompanyRoute(planRoutes)
	dashboard.IDashboard(&dashboard.Dashboard{Db: *s.store}).SetupDashboardRoute(planRoutes)
	order.IOrder(&order.Order{Db: *s.store, Config: *s.config}).SetupOrderRoute(planRoutes)
	orderItems.IOrderItem(&orderItems.OrderItem{Db: *s.store}).SetupOrderItemRoute(planRoutes)
	paseto.IToken(&paseto.Token{Db: *s.store, TokenMaker: *s.tokenMaker, Config: *s.config}).SetupTokenRoute(routes)
	paymentForm.IPaymentForm(&paymentForm.PaymentForm{Db: *s.store}).SetupPaymentFormRoute(planRoutes)
	paymentReceipt.IPaymentReceipt(&paymentReceipt.PaymentReceipt{Db: *s.store}).SetupPaymentReceiptRoute(planRoutes)
	product.IProduct(&product.Product{Db: *s.store}).SetupProductRoute(planRoutes)
	representative.IRepresentative(&representative.Representative{Db: *s.store}).SetupRepresentativeRoute(routes, authRoutes)
	seller.ISeller(&seller.Seller{Db: *s.store}).SetupSellerRoute(planRoutes)
	smtp.ISmtp(&smtp.Smtp{Db: *s.store}).SetupSmtpRoute(routes)
	stripe.IStripe(&stripe.Stripe{Db: *s.store, Config: *s.config}).SetupStripeRoute(routes, authRoutes)
	user.IUser(&user.User{Db: *s.store, TokenMaker: *s.tokenMaker, Config: *s.config}).SetupUserRoute(routes, authRoutes)
	file.IFile(&file.File{Db: *s.store, Config: *s.config}).SetupFileRoute(routes, planRoutes)
	imports.IImports(&imports.Import{Db: *s.store}).SetupImportRoute(planRoutes)
	lead.ILead(&lead.Lead{Db: *s.store}).SetupLeadRoute(routes, authRoutes)
}
