package components

import (
	"fmt"
)

type MailStatisticsDecorator struct {
	amsb AbstractMailService
}

func NewMailStatisticsDecorator(ams AbstractMailService) *MailStatisticsDecorator {
	return &MailStatisticsDecorator{
		amsb: ams,
	}
}

func (msd *MailStatisticsDecorator) SendMail() {
	fmt.Println("Statistics Decorator call started")
	msd.amsb.SendMail()
	fmt.Println("Statistics Decorator call ended")
}

type MailLogsDecorator struct {
	amsb AbstractMailService
}

func NewMailLogsDecorator(ams AbstractMailService) *MailLogsDecorator {
	return &MailLogsDecorator{
		amsb: ams,
	}
}

func (mld *MailLogsDecorator) SendMail() {
	fmt.Println("Logs Decorator call started")
	mld.amsb.SendMail()
	fmt.Println("Logs Decorator call ended")
}
