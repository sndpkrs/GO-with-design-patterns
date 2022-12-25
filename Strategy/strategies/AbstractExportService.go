package strategies

import "fmt"

type Order struct {
	name string
	sku  string
	aes  AbstractExportService
}

func NewOrder(name, sku string, serv AbstractExportService) *Order {
	return &Order{
		name: name,
		sku:  sku,
		aes:  serv,
	}
}

func (o *Order) SetExportService(serv AbstractExportService) {
	o.aes = serv
}

func (o *Order) Export() {
	o.aes.Export(*o)
}

type AbstractExportService interface {
	Export(Order)
}

type CsvExport struct {
}

func NewCsvExportService() *CsvExport {
	return &CsvExport{}
}

func (c *CsvExport) Export(o Order) {
	fmt.Printf("Exporting from csv %v and %v", o.name, o.sku)
}

type ExcelExport struct {
}

func NewExcelExportService() *ExcelExport {
	return &ExcelExport{}
}

func (c *ExcelExport) Export(o Order) {
	fmt.Printf("Exporting from excel %v and %v", o.name, o.sku)
}
