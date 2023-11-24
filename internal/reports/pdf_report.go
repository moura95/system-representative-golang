package reports

// PDFReport is a struct for generating PDF reports
type PDFReport struct{}

// CreatePDFReport creates a PDF report using the given report factory
func (p *PDFReport) CreatePDFReport(factory ReportFactory) ([]byte, error) {

	report, err := factory.CreateReport()
	if err != nil {
		return nil, err
	}

	pdf, err := report.GenerateReport()
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
