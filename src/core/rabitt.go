package core

import (
    "log"
    "os"
    "github.com/streadway/amqp"
)

type RabbitMQConnection struct {
    Conn *amqp.Connection
    Ch   *amqp.Channel
}

func NewRabbitMQConnection() (*RabbitMQConnection, error) {
    rabbitMQURL := os.Getenv("RABBITMQ_URL")
    if rabbitMQURL == "" {
        log.Fatal("La variable de entorno RABBITMQ_URL no está configurada")
    }

    conn, err := amqp.Dial(rabbitMQURL)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    return &RabbitMQConnection{Conn: conn, Ch: ch}, nil
}

func (r *RabbitMQConnection) Close() {
    if r.Ch != nil {
        r.Ch.Close()
    }
    if r.Conn != nil {
        r.Conn.Close()
    }
}