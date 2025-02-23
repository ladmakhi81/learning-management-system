package queueservice

import (
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	pkgrabbitmqclient "github.com/ladmakhi81/learning-management-system/pkg/rabbitmq"
)

type PDFQueueService struct {
	QueueService *pkgrabbitmqclient.RabbitmqService
}

func NewPDFQueueService(
	rabbitmqClient *pkgrabbitmqclient.RabbitmqClient,
) (*PDFQueueService, error) {
	service := pkgrabbitmqclient.NewRabbitmqService(rabbitmqClient)
	if err := service.InitQueue("PDF-Queue", "pdf-compression"); err != nil {
		return nil, baseerror.NewServerErr(
			err.Error(),
			"PDFQueueService.NewPDFQueueService",
		)
	}
	return &PDFQueueService{
		QueueService: service,
	}, nil
}
