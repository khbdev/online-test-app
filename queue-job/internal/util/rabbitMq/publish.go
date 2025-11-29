package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)


func PublishMessage(ch *amqp.Channel, exchange, routingKey, message string) error {
	err := ch.Publish(
		exchange,   // direct_exchange
		routingKey, // queue_key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
			Priority:    5, // o‘rtacha priority
		},
	)
	if err != nil {
		return fmt.Errorf("xabar yuborishda xatolik: %w", err)
	}
	log.Printf("📨 Xabar yuborildi: %s\n", message)
	return nil
}
