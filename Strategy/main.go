package main

import (
	"strategy/strategies"
)

func main() {
	o := strategies.NewOrder("order 1", "sku 1", strategies.NewCsvExportService())
	o.Export()
	o.SetExportService(strategies.NewExcelExportService())
	o.Export()
}
