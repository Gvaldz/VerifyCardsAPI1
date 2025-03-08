package infrastructure

import (
    "database/sql"
    "datos/src/internal/application"
    "datos/src/core"
)

type CardDependencies struct {
    DB      *sql.DB
    RabbitMQ *core.RabbitMQConnection
}

func NewCardDependencies(db *sql.DB, rabbitMQ *core.RabbitMQConnection) *CardDependencies {
    return &CardDependencies{DB: db, RabbitMQ: rabbitMQ}
}

func (d *CardDependencies) GetRoutes() *CardRoutes {
    cardRepo := NewCardRepository(d.DB)
    validateCardUseCase := application.NewValidateCardUseCase(cardRepo)

    rabbitMQProducer := NewRabbitMQProducer(d.RabbitMQ)
    cardController := NewCardController(validateCardUseCase, rabbitMQProducer)

    return NewCardRoutes(cardController)
}