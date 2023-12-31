// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Querier interface {
	AddUserPermission(ctx context.Context, arg AddUserPermissionParams) (UserPermission, error)
	ChangePasswordUserByID(ctx context.Context, arg ChangePasswordUserByIDParams) error
	CreateActivity(ctx context.Context, arg CreateActivityParams) (Activity, error)
	CreateCalendar(ctx context.Context, arg CreateCalendarParams) (Calendar, error)
	CreateCompany(ctx context.Context, arg CreateCompanyParams) (Company, error)
	CreateLead(ctx context.Context, arg CreateLeadParams) (Lead, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateOrderItems(ctx context.Context, arg CreateOrderItemsParams) (OrderItem, error)
	CreatePaymentForm(ctx context.Context, arg CreatePaymentFormParams) (FormPayment, error)
	CreatePaymentReceipt(ctx context.Context, arg CreatePaymentReceiptParams) (PaymentReceipt, error)
	CreatePermission(ctx context.Context, name string) (Permission, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateRepresentative(ctx context.Context, arg CreateRepresentativeParams) (Representative, error)
	CreateSellers(ctx context.Context, arg CreateSellersParams) (Seller, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateSmtp(ctx context.Context, arg CreateSmtpParams) (Smtp, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCalendarByID(ctx context.Context, id int32) (Calendar, error)
	DeleteCompanyByID(ctx context.Context, id int32) (Company, error)
	DeleteFilePaymentReceiptByID(ctx context.Context, id int32) (FilesPaymentReceipt, error)
	DeleteLeadByID(ctx context.Context, id int32) (Lead, error)
	DeleteOrderByID(ctx context.Context, id int32) (Order, error)
	DeleteOrderItemsByID(ctx context.Context, arg DeleteOrderItemsByIDParams) error
	DeletePaymentFormByID(ctx context.Context, id int32) (FormPayment, error)
	DeletePaymentReceiptByID(ctx context.Context, id int32) (PaymentReceipt, error)
	DeletePermissionByID(ctx context.Context, id int32) (Permission, error)
	DeleteProductByID(ctx context.Context, id int32) (Product, error)
	DeleteRepresentativeByID(ctx context.Context, id int32) (Representative, error)
	DeleteSellerByID(ctx context.Context, id int32) (Seller, error)
	DeleteSmtpByRepresentativeID(ctx context.Context, representativeID int32) (Smtp, error)
	DeleteUserByID(ctx context.Context, id int32) (User, error)
	GetAllPermissions(ctx context.Context) ([]Permission, error)
	GetAllRepresentativesByID(ctx context.Context) ([]Representative, error)
	GetCalendarByID(ctx context.Context, id int32) (Calendar, error)
	GetCompanyByID(ctx context.Context, id int32) (Company, error)
	GetCompanyUserByID(ctx context.Context, arg GetCompanyUserByIDParams) (GetCompanyUserByIDRow, error)
	GetEmailAndNameByRepresentativeID(ctx context.Context, id int32) (GetEmailAndNameByRepresentativeIDRow, error)
	GetLastOrderByRepresentativeID(ctx context.Context, representativeID int32) (int32, error)
	GetLeadByEmail(ctx context.Context, email string) (Lead, error)
	GetLeadByID(ctx context.Context, id int32) (Lead, error)
	GetOrderByID(ctx context.Context, id int32) (GetOrderByIDRow, error)
	GetOrderItemsByID(ctx context.Context, arg GetOrderItemsByIDParams) (GetOrderItemsByIDRow, error)
	GetPaymentFormByID(ctx context.Context, id int32) (FormPayment, error)
	GetPaymentReceiptByID(ctx context.Context, arg GetPaymentReceiptByIDParams) (GetPaymentReceiptByIDRow, error)
	GetPermissionByID(ctx context.Context, id int32) (Permission, error)
	GetPlanByRepresentativesID(ctx context.Context, id int32) (PlanTypes, error)
	GetProductByID(ctx context.Context, id int32) (GetProductByIDRow, error)
	GetRepresentativeDateExpByID(ctx context.Context, id int32) (time.Time, error)
	GetRepresentativesByID(ctx context.Context, id int32) (Representative, error)
	GetSellerByID(ctx context.Context, id int32) (Seller, error)
	GetSessionByID(ctx context.Context, id uuid.UUID) (Session, error)
	GetSmtpByRepresentativeID(ctx context.Context, representativeID int32) (Smtp, error)
	GetTotalUsersByRepresentativesID(ctx context.Context, representativeID int32) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	GetUserPasswordByID(ctx context.Context, id int32) (string, error)
	GetUserPermissionAndName(ctx context.Context, userID int32) ([]string, error)
	ListActivity(ctx context.Context) ([]Activity, error)
	ListActivityByRepresentativeID(ctx context.Context, representativeID int32) ([]Activity, error)
	ListActivityByUserID(ctx context.Context, userID int32) ([]Activity, error)
	ListCalendarsByUserID(ctx context.Context, userID int32) ([]Calendar, error)
	ListCompaniesByRepresentativeID(ctx context.Context, arg ListCompaniesByRepresentativeIDParams) ([]Company, error)
	ListLeads(ctx context.Context) ([]Lead, error)
	ListOrdersByRepresentativeID(ctx context.Context, arg ListOrdersByRepresentativeIDParams) ([]ListOrdersByRepresentativeIDRow, error)
	ListOrdersItemsByOrderID(ctx context.Context, orderID int32) ([]ListOrdersItemsByOrderIDRow, error)
	ListPaymentFormsByRepresentativeID(ctx context.Context, representativeID int32) ([]FormPayment, error)
	ListPaymentReceiptByRepresentativeID(ctx context.Context, representativeID int32) ([]ListPaymentReceiptByRepresentativeIDRow, error)
	ListProductsByRepresentativeID(ctx context.Context, arg ListProductsByRepresentativeIDParams) ([]ListProductsByRepresentativeIDRow, error)
	ListSellersByRepresentativeID(ctx context.Context, arg ListSellersByRepresentativeIDParams) ([]Seller, error)
	ListUsersByRepresentativeID(ctx context.Context, arg ListUsersByRepresentativeIDParams) ([]User, error)
	RemoveCompanyByID(ctx context.Context, id int32) (Company, error)
	RemoveOrderByID(ctx context.Context, id int32) (Order, error)
	RemoveProductByID(ctx context.Context, id int32) (Product, error)
	RemoveRepresentativeByID(ctx context.Context, id int32) (Representative, error)
	RemoveSellerByID(ctx context.Context, id int32) (Seller, error)
	RemoveUserByID(ctx context.Context, id int32) (User, error)
	RemoveUserPermissionByID(ctx context.Context, arg RemoveUserPermissionByIDParams) (UserPermission, error)
	RestoreCompanyByID(ctx context.Context, id int32) (Company, error)
	RestoreOrderByID(ctx context.Context, id int32) (Order, error)
	RestoreProductByID(ctx context.Context, id int32) (Product, error)
	RestoreRepresentativeByID(ctx context.Context, id int32) (Representative, error)
	RestoreSellerByID(ctx context.Context, id int32) (Seller, error)
	RestoreUserByID(ctx context.Context, id int32) (User, error)
	TopBuyerByRepresentativeID(ctx context.Context, arg TopBuyerByRepresentativeIDParams) ([]TopBuyerByRepresentativeIDRow, error)
	TopFactoryByRepresentativeID(ctx context.Context, arg TopFactoryByRepresentativeIDParams) ([]TopFactoryByRepresentativeIDRow, error)
	TopSalesPerProductByRepresentativeID(ctx context.Context, arg TopSalesPerProductByRepresentativeIDParams) ([]TopSalesPerProductByRepresentativeIDRow, error)
	TotalSalesPerDayByRepresentativeID(ctx context.Context, arg TotalSalesPerDayByRepresentativeIDParams) ([]TotalSalesPerDayByRepresentativeIDRow, error)
	UpdateCalendarByID(ctx context.Context, arg UpdateCalendarByIDParams) (Calendar, error)
	UpdateCompanyByID(ctx context.Context, arg UpdateCompanyByIDParams) (Company, error)
	UpdateLastLogin(ctx context.Context, id int32) error
	UpdateOrderByID(ctx context.Context, arg UpdateOrderByIDParams) (Order, error)
	UpdateOrderItemByID(ctx context.Context, arg UpdateOrderItemByIDParams) (OrderItem, error)
	UpdatePaymentFormByID(ctx context.Context, arg UpdatePaymentFormByIDParams) (FormPayment, error)
	UpdatePaymentPaymentReceiptByID(ctx context.Context, arg UpdatePaymentPaymentReceiptByIDParams) (PaymentReceipt, error)
	UpdatePermissionByID(ctx context.Context, arg UpdatePermissionByIDParams) (Permission, error)
	UpdatePlanByID(ctx context.Context, arg UpdatePlanByIDParams) (Representative, error)
	UpdateProductByID(ctx context.Context, arg UpdateProductByIDParams) (Product, error)
	UpdateRepresentativeByID(ctx context.Context, arg UpdateRepresentativeByIDParams) (Representative, error)
	UpdateSellerByID(ctx context.Context, arg UpdateSellerByIDParams) (Seller, error)
	UpdateSmtpByID(ctx context.Context, arg UpdateSmtpByIDParams) (Smtp, error)
	UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) (User, error)
	UploadFilePaymentReceipt(ctx context.Context, arg UploadFilePaymentReceiptParams) (FilesPaymentReceipt, error)
}

var _ Querier = (*Queries)(nil)
