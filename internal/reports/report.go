package reports

// Report interface defines the method for generating the report
type Report interface {
	GenerateReport() ([]byte, error)
}
