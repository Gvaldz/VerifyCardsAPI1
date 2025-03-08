package domain

import "datos/src/internal/domain/entities"

type CardRepository interface {
    VerifyCard(card entities.Card) (bool, error)
}

type CardService interface {
    ValidateCard(card entities.Card) (bool, error)
}