package infrastructure

import (
    "encoding/json"
    "net/http"
    "datos/src/internal/domain/entities"
    "datos/src/internal/application"
)

type CardController struct {
    validateCardUseCase *application.ValidateCardUseCase
    rabbitMQProducer    *RabbitMQProducer
}

func NewCardController(
    validateCardUseCase *application.ValidateCardUseCase,
    rabbitMQProducer *RabbitMQProducer,
) *CardController {
    return &CardController{
        validateCardUseCase: validateCardUseCase,
        rabbitMQProducer:    rabbitMQProducer,
    }
}

func (c *CardController) ValidateCard(w http.ResponseWriter, r *http.Request) {
    var card entities.Card
    if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    isValid, err := c.validateCardUseCase.Execute(card)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if !isValid {
        c.rabbitMQProducer.SendCardVerifiedMessage("Datos incorrectos: " + card.Number)
        http.Error(w, "Datos incorrectos", http.StatusUnauthorized)
        return
    }

    c.rabbitMQProducer.SendCardVerifiedMessage("Datos correctos: " + card.Number)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Datos correctos"))
}