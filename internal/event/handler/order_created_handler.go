package handler

import (
	"encoding/json"
	"fmt"
	"github.com/rodrigoachilles/clean-architecture/pkg/events"
	"github.com/rodrigoachilles/clean-architecture/pkg/log"
	"sync"

	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Logger.Info().Msg(fmt.Sprintf("Order created: %v", event.GetPayload()))

	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	_ = h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}