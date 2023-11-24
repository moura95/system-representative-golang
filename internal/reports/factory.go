package reports

// ReportFactory interface defines the method for creating reports
type ReportFactory interface {
	CreateReport() (Report, error)
}
