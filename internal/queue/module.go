package queue

import (
	"fmt"

	queueservice "github.com/ladmakhi81/learning-management-system/internal/queue/service"
	pkgrabbitmqclient "github.com/ladmakhi81/learning-management-system/pkg/rabbitmq"
	"go.uber.org/dig"
)

type QueueModule struct {
	container *dig.Container
}

func NewQueueModule(
	container *dig.Container,
) QueueModule {
	return QueueModule{
		container: container,
	}
}

func (m QueueModule) LoadModule() {
	m.container.Invoke(func(rabbitmqClient *pkgrabbitmqclient.RabbitmqClient) {
		var err error

		pdfQueueService, pdfQueueServiceErr := queueservice.NewPDFQueueService(rabbitmqClient)
		if pdfQueueServiceErr != nil {
			err = pdfQueueServiceErr
		}

		m.container.Provide(func() *queueservice.PDFQueueService {
			return pdfQueueService
		})

		if err == nil {
			fmt.Println("------ Queue Module Load ------")
		} else {
			fmt.Println("------ Queue Module Not Load: Failed ------")
		}
	})
}
