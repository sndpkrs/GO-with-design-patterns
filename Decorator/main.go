package main

import (
	components "decorator/Components"
	"fmt"
)

func main() {
	var azms components.AbstractMailService = components.NewAzureMailService()
	azms.SendMail()
	var gms components.AbstractMailService = components.NewGcpMailService()
	gms.SendMail()
	fmt.Println("---------------------------------------------")

	var dazms components.AbstractMailService = components.NewMailStatisticsDecorator(azms)
	var dgms components.AbstractMailService = components.NewMailStatisticsDecorator(gms)
	dazms.SendMail()
	dgms.SendMail()

	fmt.Println("---------------------------------------------")

	dazms = components.NewMailLogsDecorator(azms)
	dgms = components.NewMailLogsDecorator(gms)
	dazms.SendMail()
	dgms.SendMail()
}
