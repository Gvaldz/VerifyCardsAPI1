package infrastructure

import (
	"datos/src/internal/domain/entities"
    "datos/src/internal/domain"
    "database/sql"
)

type cardRepository struct {
    db *sql.DB
}

func NewCardRepository(db *sql.DB) domain.CardRepository {
    return &cardRepository{db: db}
}

func (r *cardRepository) VerifyCard(card entities.Card) (bool, error) {
    var storedCard entities.Card
    err := r.db.QueryRow("SELECT number, name, expiry, cvv FROM cards WHERE number = ?", card.Number).
        Scan(&storedCard.Number, &storedCard.Name, &storedCard.Expiry, &storedCard.CVV)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, nil
        }
        return false, err
    }

    if storedCard.Number == card.Number &&
       storedCard.Name == card.Name &&
       storedCard.Expiry == card.Expiry &&
       storedCard.CVV == card.CVV {
        return true, nil
    }

    return false, nil
}