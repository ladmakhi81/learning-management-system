package pkgrabbitmqclient

import (
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	"github.com/streadway/amqp"
)

type RabbitmqClient struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRabbitmqClient(clientURL string) (*RabbitmqClient, error) {
	connection, connectionErr := amqp.Dial(clientURL)
	if connectionErr != nil {
		return nil, baseerror.NewServerErr(
			connectionErr.Error(),
			"RabbitmqClient.NewRabbitmqClient",
		)
	}
	channel, channelErr := connection.Channel()
	if channelErr != nil {
		return nil, baseerror.NewServerErr(
			channelErr.Error(),
			"RabbitmqClient.connect",
		)
	}
	return &RabbitmqClient{
		connection: connection,
		channel:    channel,
	}, nil
}

func (c RabbitmqClient) Close() {
	c.channel.Close()
	c.connection.Close()
}
