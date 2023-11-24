package factory

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lucas-gaitzsch/pdf-turtle-client-golang/models"
	"github.com/lucas-gaitzsch/pdf-turtle-client-golang/pdfturtleclient"

	"my-orders/internal/reports"
	"my-orders/internal/repository"
	"my-orders/internal/util"
)

// OrderReportFactory generates order reports
type OrderReportFactory struct {
	Db           repository.Querier
	PdfClientURL string
	OrderID      int32
	Ctx          *gin.Context
}

// CreateReport creates an order report
func (f *OrderReportFactory) CreateReport() (reports.Report, error) {
	var order dataOrderInfo
	var iData []dataOrderItems

	dataOrder, err := f.Db.GetOrderByID(f.Ctx, f.OrderID)
	if err != nil {
		return nil, util.ErrorDatabaseRead
	}
	dataBuyer, err := f.Db.GetCompanyByID(f.Ctx, dataOrder.CustomerID)
	if err != nil {
		return nil, util.ErrorDatabaseRead
	}
	dataFactory, err := f.Db.GetCompanyByID(f.Ctx, dataOrder.FactoryID)
	if err != nil {
		return nil, util.ErrorDatabaseRead
	}
	orderItems, err := f.Db.ListOrdersItemsByOrderID(f.Ctx, f.OrderID)
	if err != nil {
		return nil, util.ErrorDatabaseRead
	}
	for _, item := range orderItems {
		iData = append(iData, dataOrderItems{
			Code:        item.Code,
			Description: item.Description.String,
			Discount:    item.Discount,
			Ipi:         item.Ipi.String,
			OrderId:     item.OrderID,
			Price:       item.Price,
			ProductId:   item.ProductID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Total:       item.Total,
		})
	}
	var OrderStatusHtml string
	if dataOrder.Status == "Cotacao" {
		OrderStatusHtml = "Cotação"
	} else if dataOrder.Status == "Concluido" {
		OrderStatusHtml = "Pedido"
	}

	order = dataOrderInfo{
		Buyer:           dataOrder.Buyer.String,
		City:            dataBuyer.City.String,
		CreatedAt:       dataOrder.CreatedAt.Format("02/01/2006"),
		CustomerCnpj:    dataBuyer.Cnpj.String,
		CustomerEmail:   dataBuyer.Email.String,
		CustomerId:      dataBuyer.ID,
		CustomerName:    dataBuyer.Name,
		ExpireOrder:     dataOrder.ExpiredAt.Format("02/01/2006"),
		FactoryCnpj:     dataFactory.Cnpj.String,
		FactoryEmail:    dataFactory.Email.String,
		FactoryId:       dataFactory.ID,
		FactoryLogoUrl:  dataFactory.LogoUrl.String,
		FactoryName:     dataFactory.Name,
		FactoryNumber:   dataFactory.Number.String,
		FactoryPhone:    dataFactory.Phone.String,
		FactoryStreet:   dataFactory.Street.String,
		FactoryZipcode:  dataFactory.ZipCode.String,
		FantasyName:     dataFactory.FantasyName.String,
		FormPaymentId:   dataOrder.FormPaymentID.Int32,
		FormPaymentName: dataOrder.FormPaymentName.String,
		Ie:              dataFactory.Ie.String,
		IsActive:        dataOrder.IsActive,
		Items:           iData,
		OrderId:         int(dataOrder.ID),
		OrderNumber:     int(dataOrder.OrderNumber),
		Phone:           dataBuyer.Phone.String,
		PortageId:       int(dataOrder.PortageID),
		PortageName:     dataOrder.PortageName,
		SellerEmail:     dataOrder.SellerEmail.String,
		SellerId:        int(dataOrder.SellerID),
		SellerName:      dataOrder.SellerName,
		Shipping:        string(dataOrder.Shipping),
		State:           dataBuyer.State.String,
		Status:          string(dataOrder.Status),
		Street:          dataBuyer.Street.String,
		Total:           dataOrder.Total,
		ZipCode:         dataBuyer.ZipCode.String,
		OrderStatusHtml: OrderStatusHtml,
	}

	body, err := os.ReadFile("templates/pdf-turtle-bundle-midas/index.html")
	if err != nil {
		return nil, nil
	}
	bodyString := string(body)

	footer, err := os.ReadFile("templates/pdf-turtle-bundle-midas/footer.html")
	if err != nil {
		return nil, nil
	}

	pdfModel := models.RenderTemplateData{
		HtmlTemplate:       &bodyString,
		FooterHtmlTemplate: string(footer),
		Model:              order,
		TemplateEngine:     "golang",
		RenderOptions: models.RenderOptions{
			PageFormat: "A4",
			Margins: &models.RenderOptionsMargins{
				Top:    10,
				Right:  5,
				Bottom: 5,
				Left:   5,
			},
		},
	}

	return &orderReport{
		pdfModel:  pdfModel,
		pdfClient: f.PdfClientURL,
	}, nil
}

// orderReport is an implementation of the Report interface for generating order reports
type orderReport struct {
	pdfModel  models.RenderTemplateData
	pdfClient string
}

type dataOrderInfo struct {
	Buyer           string           `json:"buyer"`
	City            string           `json:"city"`
	CreatedAt       string           `json:"created_at"`
	CustomerCnpj    string           `json:"customer_cnpj"`
	CustomerEmail   string           `json:"customer_email"`
	CustomerId      int32            `json:"customer_id"`
	CustomerName    string           `json:"customer_name"`
	ExpireOrder     string           `json:"expire_order"`
	FactoryCnpj     string           `json:"factory_cnpj"`
	FactoryEmail    string           `json:"factory_email"`
	FactoryId       int32            `json:"factory_id"`
	FactoryLogoUrl  string           `json:"factory_logo_url"`
	FactoryName     string           `json:"factory_name"`
	FactoryNumber   string           `json:"factory_number"`
	FactoryPhone    string           `json:"factory_phone"`
	FactoryStreet   string           `json:"factory_street"`
	FactoryZipcode  string           `json:"factory_zipcode"`
	FantasyName     string           `json:"fantasy_name"`
	FormPaymentId   int32            `json:"form_payment_id"`
	FormPaymentName string           `json:"form_payment_name"`
	Ie              string           `json:"ie"`
	IsActive        bool             `json:"is_active"`
	Items           []dataOrderItems `json:"items"`
	OrderId         int              `json:"order_id"`
	OrderNumber     int              `json:"order_number"`
	Phone           string           `json:"phone"`
	PortageId       int              `json:"portage_id"`
	PortageName     string           `json:"portage_name"`
	SellerEmail     string           `json:"seller_email"`
	SellerId        int              `json:"seller_id"`
	SellerName      string           `json:"seller_name"`
	Shipping        string           `json:"shipping"`
	State           string           `json:"state"`
	Status          string           `json:"status"`
	Street          string           `json:"street"`
	Total           string           `json:"total"`
	UrlPdf          string           `json:"url_pdf"`
	ZipCode         string           `json:"zip_code"`
	OrderStatusHtml string
}

type dataOrderItems struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Discount    string `json:"discount"`
	Ipi         string `json:"ipi"`
	OrderId     int32  `json:"order_id"`
	Price       string `json:"price"`
	ProductId   int32  `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int32  `json:"quantity"`
	Total       string `json:"total"`
}

// GenerateReport generates an order report
func (r *orderReport) GenerateReport() ([]byte, error) {
	pdfClient := pdfturtleclient.NewPdfTurtleClient(r.pdfClient)
	resp, err := pdfClient.RenderTemplate(r.pdfModel)

	if err != nil {
		return nil, err
	}
	defer resp.Close()

	pdf, err := io.ReadAll(resp)
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
