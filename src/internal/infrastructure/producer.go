package infrastructure

import (
	"log"
	"github.com/streadway/amqp"
	"datos/src/core"
)

type RabbitMQProducer struct {
	conn *core.RabbitMQConnection
}

func NewRabbitMQProducer(conn *core.RabbitMQConnection) *RabbitMQProducer {
	return &RabbitMQProducer{conn: conn}
}

func (p *RabbitMQProducer) DeclareQueue(queueName string) error {
	_, err := p.conn.Ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("Failed to declare queue '%s': %s", queueName, err)
		return err
	}

	log.Printf("Queue '%s' declared successfully", queueName)
	return nil
}

func (p *RabbitMQProducer) PublishMessage(queueName string, message string) error {
	err := p.conn.Ch.Publish(
		"",         // exchange
		queueName,  // routing key (queue name)
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Printf("Failed to publish message to queue '%s': %s", queueName, err)
		return err
	}

	log.Printf("Published message to queue '%s': %s", queueName, message)
	return nil
}

func (p *RabbitMQProducer) SendCardVerifiedMessage(cardNumber string) error {
	queueName := "cards"
	err := p.PublishMessage(queueName, cardNumber)
	if err != nil {
		log.Printf("Failed to send card verified message: %s", err)
		return err
	}

	log.Printf(" [x] Sent card verified: %s", cardNumber)
	return nil
}