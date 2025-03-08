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

func (p *RabbitMQProducer) SendCardVerifiedMessage(cardNumber string) {
    q, err := p.conn.Ch.QueueDeclare(
        "cards", // name
        true,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare a queue: %s", err)
    }

    err = p.conn.Ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(cardNumber),
        })
    if err != nil {
        log.Fatalf("Failed to publish a message: %s", err)
    }

    log.Printf(" [x] Sent card verified: %s", cardNumber)
}