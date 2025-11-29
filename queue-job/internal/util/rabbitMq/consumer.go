package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)


func ConsumeMessages(ch *amqp.Channel, queue string) {
	msgs, err := ch.Consume(
		queue,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		log.Fatalf("❌ Xabarlarni olishda xatolik: %v", err)
	}

	log.Println("👂 Consumer ishga tushdi, xabar kutyapti...")

	for msg := range msgs {
		fmt.Printf("📥 Xabar keldi: %s\n", msg.Body)
	}
}
