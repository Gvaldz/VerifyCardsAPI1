package application

import (
	"datos/src/internal/domain/entities"
    "datos/src/internal/domain"
)

type ValidateCardUseCase struct {
    repo domain.CardRepository
}

func NewValidateCardUseCase(repo domain.CardRepository) *ValidateCardUseCase {
    return &ValidateCardUseCase{repo: repo}
}

func (uc *ValidateCardUseCase) Execute(card entities.Card) (bool, error) {
    return uc.repo.VerifyCard(card)
}