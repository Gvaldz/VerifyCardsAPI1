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
        http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
        return
    }

    isValid, err := c.validateCardUseCase.Execute(card)
    if err != nil {
        http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
        return
    }

    if !isValid {
        c.rabbitMQProducer.SendCardVerifiedMessage("Datos incorrectos: " + card.Number)
        http.Error(w, `{"error": "Datos incorrectos"}`, http.StatusUnauthorized)
        return
    }

    c.rabbitMQProducer.SendCardVerifiedMessage("Datos correctos: " + card.Number)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Datos correctos"})
}
