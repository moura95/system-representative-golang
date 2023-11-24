// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package repository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CompanyTypes string

const (
	CompanyTypesFactory  CompanyTypes = "Factory"
	CompanyTypesCustomer CompanyTypes = "Customer"
	CompanyTypesPortage  CompanyTypes = "Portage"
)

func (e *CompanyTypes) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CompanyTypes(s)
	case string:
		*e = CompanyTypes(s)
	default:
		return fmt.Errorf("unsupported scan type for CompanyTypes: %T", src)
	}
	return nil
}

type NullCompanyTypes struct {
	CompanyTypes CompanyTypes
	Valid        bool // Valid is true if CompanyTypes is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCompanyTypes) Scan(value interface{}) error {
	if value == nil {
		ns.CompanyTypes, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CompanyTypes.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCompanyTypes) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CompanyTypes), nil
}

type OriginLeadsEnum string

const (
	OriginLeadsEnumFacebook  OriginLeadsEnum = "Facebook"
	OriginLeadsEnumInstagram OriginLeadsEnum = "Instagram"
	OriginLeadsEnumGoogle    OriginLeadsEnum = "Google"
	OriginLeadsEnumLinkedin  OriginLeadsEnum = "Linkedin"
	OriginLeadsEnumOutros    OriginLeadsEnum = "Outros"
)

func (e *OriginLeadsEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OriginLeadsEnum(s)
	case string:
		*e = OriginLeadsEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for OriginLeadsEnum: %T", src)
	}
	return nil
}

type NullOriginLeadsEnum struct {
	OriginLeadsEnum OriginLeadsEnum
	Valid           bool // Valid is true if OriginLeadsEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOriginLeadsEnum) Scan(value interface{}) error {
	if value == nil {
		ns.OriginLeadsEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OriginLeadsEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOriginLeadsEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OriginLeadsEnum), nil
}

type PaymentReceiptFormType string

const (
	PaymentReceiptFormTypePix        PaymentReceiptFormType = "Pix"
	PaymentReceiptFormTypeInvoice    PaymentReceiptFormType = "Invoice"
	PaymentReceiptFormTypeTransfer   PaymentReceiptFormType = "Transfer"
	PaymentReceiptFormTypeCreditCard PaymentReceiptFormType = "CreditCard"
	PaymentReceiptFormTypeDebitCard  PaymentReceiptFormType = "DebitCard"
	PaymentReceiptFormTypeCash       PaymentReceiptFormType = "Cash"
	PaymentReceiptFormTypeCheque     PaymentReceiptFormType = "Cheque"
	PaymentReceiptFormTypeOutros     PaymentReceiptFormType = "Outros"
)

func (e *PaymentReceiptFormType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PaymentReceiptFormType(s)
	case string:
		*e = PaymentReceiptFormType(s)
	default:
		return fmt.Errorf("unsupported scan type for PaymentReceiptFormType: %T", src)
	}
	return nil
}

type NullPaymentReceiptFormType struct {
	PaymentReceiptFormType PaymentReceiptFormType
	Valid                  bool // Valid is true if PaymentReceiptFormType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPaymentReceiptFormType) Scan(value interface{}) error {
	if value == nil {
		ns.PaymentReceiptFormType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PaymentReceiptFormType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPaymentReceiptFormType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PaymentReceiptFormType), nil
}

type PaymentReceiptStatus string

const (
	PaymentReceiptStatusPending PaymentReceiptStatus = "Pending"
	PaymentReceiptStatusPaid    PaymentReceiptStatus = "Paid"
	PaymentReceiptStatusExpired PaymentReceiptStatus = "Expired"
)

func (e *PaymentReceiptStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PaymentReceiptStatus(s)
	case string:
		*e = PaymentReceiptStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for PaymentReceiptStatus: %T", src)
	}
	return nil
}

type NullPaymentReceiptStatus struct {
	PaymentReceiptStatus PaymentReceiptStatus
	Valid                bool // Valid is true if PaymentReceiptStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPaymentReceiptStatus) Scan(value interface{}) error {
	if value == nil {
		ns.PaymentReceiptStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PaymentReceiptStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPaymentReceiptStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PaymentReceiptStatus), nil
}

type PaymentReceiptType string

const (
	PaymentReceiptTypePayment PaymentReceiptType = "Payment"
	PaymentReceiptTypeReceipt PaymentReceiptType = "Receipt"
)

func (e *PaymentReceiptType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PaymentReceiptType(s)
	case string:
		*e = PaymentReceiptType(s)
	default:
		return fmt.Errorf("unsupported scan type for PaymentReceiptType: %T", src)
	}
	return nil
}

type NullPaymentReceiptType struct {
	PaymentReceiptType PaymentReceiptType
	Valid              bool // Valid is true if PaymentReceiptType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPaymentReceiptType) Scan(value interface{}) error {
	if value == nil {
		ns.PaymentReceiptType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PaymentReceiptType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPaymentReceiptType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PaymentReceiptType), nil
}

type PlanTypes string

const (
	PlanTypesTrial  PlanTypes = "Trial"
	PlanTypesSilver PlanTypes = "Silver"
	PlanTypesGold   PlanTypes = "Gold"
)

func (e *PlanTypes) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PlanTypes(s)
	case string:
		*e = PlanTypes(s)
	default:
		return fmt.Errorf("unsupported scan type for PlanTypes: %T", src)
	}
	return nil
}

type NullPlanTypes struct {
	PlanTypes PlanTypes
	Valid     bool // Valid is true if PlanTypes is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPlanTypes) Scan(value interface{}) error {
	if value == nil {
		ns.PlanTypes, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PlanTypes.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPlanTypes) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PlanTypes), nil
}

type ShippingEnum string

const (
	ShippingEnumCIF    ShippingEnum = "CIF"
	ShippingEnumFOB    ShippingEnum = "FOB"
	ShippingEnumOutros ShippingEnum = "Outros"
)

func (e *ShippingEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ShippingEnum(s)
	case string:
		*e = ShippingEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for ShippingEnum: %T", src)
	}
	return nil
}

type NullShippingEnum struct {
	ShippingEnum ShippingEnum
	Valid        bool // Valid is true if ShippingEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullShippingEnum) Scan(value interface{}) error {
	if value == nil {
		ns.ShippingEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ShippingEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullShippingEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ShippingEnum), nil
}

type StatusEnum string

const (
	StatusEnumRascunho  StatusEnum = "Rascunho"
	StatusEnumCotacao   StatusEnum = "Cotacao"
	StatusEnumCancelado StatusEnum = "Cancelado"
	StatusEnumConcluido StatusEnum = "Concluido"
)

func (e *StatusEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = StatusEnum(s)
	case string:
		*e = StatusEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for StatusEnum: %T", src)
	}
	return nil
}

type NullStatusEnum struct {
	StatusEnum StatusEnum
	Valid      bool // Valid is true if StatusEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatusEnum) Scan(value interface{}) error {
	if value == nil {
		ns.StatusEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.StatusEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatusEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.StatusEnum), nil
}

type Activity struct {
	ID               int32
	Action           string
	ReferenceUrl     string
	UserID           int32
	RepresentativeID int32
	CreatedAt        time.Time
}

type Calendar struct {
	ID         int32
	Title      string
	VisitStart time.Time
	VisitEnd   time.Time
	Allday     bool
	UserID     int32
}

type Company struct {
	ID               int32
	RepresentativeID int32
	Type             CompanyTypes
	Cnpj             sql.NullString
	Name             string
	FantasyName      sql.NullString
	Ie               sql.NullString
	Phone            sql.NullString
	Email            sql.NullString
	Website          sql.NullString
	LogoUrl          sql.NullString
	ZipCode          sql.NullString
	State            sql.NullString
	City             sql.NullString
	Street           sql.NullString
	Number           sql.NullString
	IsActive         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type FilesOrder struct {
	ID        int32
	OrderID   int32
	UrlFile   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FilesPaymentReceipt struct {
	ID               int32
	PaymentReceiptID int32
	UrlFile          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type FormPayment struct {
	ID               int32
	RepresentativeID int32
	Name             string
}

type Lead struct {
	ID        int32
	Name      string
	Email     string
	Phone     string
	Origin    OriginLeadsEnum
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Order struct {
	ID               int32
	RepresentativeID int32
	FactoryID        int32
	CustomerID       int32
	PortageID        int32
	SellerID         int32
	FormPaymentID    sql.NullInt32
	OrderNumber      int32
	UrlPdf           sql.NullString
	Buyer            sql.NullString
	Shipping         ShippingEnum
	Status           StatusEnum
	ExpiredAt        time.Time
	Total            string
	IsActive         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type OrderItem struct {
	OrderID   int32
	ProductID int32
	Quantity  int32
	Price     string
	Discount  string
	Total     string
}

type PaymentReceipt struct {
	ID               int32
	RepresentativeID int32
	TypePayment      PaymentReceiptType
	Status           PaymentReceiptStatus
	Description      string
	Amount           string
	ExpirationDate   sql.NullTime
	PaymentDate      sql.NullTime
	DocNumber        sql.NullString
	Recipient        sql.NullString
	PaymentForm      PaymentReceiptFormType
	IsActive         bool
	Installment      int32
	IntervalDays     int32
	CreatedAt        time.Time
	UpdatedAt        time.Time
	AdditionalInfo   sql.NullString
}

type Permission struct {
	ID   int32
	Name string
}

type Product struct {
	ID               int32
	RepresentativeID int32
	FactoryID        int32
	Name             string
	Code             string
	Price            string
	Ipi              sql.NullString
	Reference        sql.NullString
	Description      sql.NullString
	ImageUrl         sql.NullString
	IsActive         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Representative struct {
	ID          int32
	Cnpj        sql.NullString
	Name        sql.NullString
	FantasyName sql.NullString
	Ie          sql.NullString
	Phone       sql.NullString
	Email       sql.NullString
	Website     sql.NullString
	LogoUrl     sql.NullString
	ZipCode     sql.NullString
	State       sql.NullString
	City        sql.NullString
	Street      sql.NullString
	Number      sql.NullString
	Plan        PlanTypes
	StripeID    sql.NullString
	DataExpire  time.Time
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Seller struct {
	ID               int32
	RepresentativeID int32
	Cpf              string
	Name             string
	Phone            sql.NullString
	Email            sql.NullString
	Pix              sql.NullString
	Observation      sql.NullString
	IsActive         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Session struct {
	ID               uuid.UUID
	UserID           int32
	RepresentativeID int32
	RefreshToken     string
	UserAgent        string
	ClientIp         string
	IsBlocked        bool
	ExpiresAt        time.Time
	CreatedAt        time.Time
}

type Smtp struct {
	RepresentativeID int32
	IsActive         bool
	Email            string
	Password         string
	Server           string
	Port             int32
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type User struct {
	ID               int32
	RepresentativeID int32
	Cpf              sql.NullString
	FirstName        string
	LastName         string
	Email            string
	Password         string
	Phone            sql.NullString
	IsActive         bool
	LastLogin        time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type UserPermission struct {
	UserID       int32
	PermissionID int32
}