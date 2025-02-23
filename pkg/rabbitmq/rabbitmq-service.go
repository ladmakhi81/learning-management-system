package pkgrabbitmqclient

import (
	"encoding/json"

	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	"github.com/streadway/amqp"
)

type RabbitmqService struct {
	client     *RabbitmqClient
	queueName  string
	routingKey string
}

const exchangeName = "LMS_EXCHANGE"

func NewRabbitmqService(client *RabbitmqClient) *RabbitmqService {
	return &RabbitmqService{
		client: client,
	}
}

func (svc *RabbitmqService) InitQueue(queueName, routingKey string) error {
	if svc.queueName != "" && svc.routingKey != "" {
		return nil
	}

	exchangeErr := svc.client.channel.ExchangeDeclare(exchangeName, amqp.ExchangeDirect, true, false, false, false, nil)
	if exchangeErr != nil {
		return baseerror.NewServerErr(
			exchangeErr.Error(),
			"RabbitmqClient.connect",
		)
	}

	_, queueErr := svc.client.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if queueErr != nil {
		return baseerror.NewServerErr(
			queueErr.Error(),
			"RabbitmqClient.connect",
		)
	}

	queueBindErr := svc.client.channel.QueueBind(queueName, routingKey, exchangeName, false, nil)
	if queueBindErr != nil {
		return baseerror.NewServerErr(
			queueBindErr.Error(),
			"RabbitmqClient.connect",
		)
	}
	svc.queueName = queueName
	svc.routingKey = routingKey

	return nil
}

func (svc RabbitmqService) Publish(data any) error {
	convertedData, convertedErr := json.Marshal(data)
	if convertedErr != nil {
		return baseerror.NewServerErr(
			convertedErr.Error(),
			"RabbitmqClient.publish",
		)
	}
	publishErr := svc.client.channel.Publish(
		exchangeName,
		svc.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        convertedData,
		},
	)
	if publishErr != nil {
		return baseerror.NewServerErr(
			publishErr.Error(),
			"RabbitmqClient.publish",
		)
	}
	return nil
}

func (svc RabbitmqService) Receive() (<-chan amqp.Delivery, error) {
	return svc.client.channel.Consume(svc.queueName, "", false, false, false, false, nil)
}
