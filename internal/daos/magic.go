package daos

import (
	"context"

	"github.com/iamnande/cardmod/internal/models"
)

// MagicDAO is the DAO for the Magic model.
//go:generate mockgen -source magic.go -destination=./mocks/magic.go -package mocks
type MagicDAO interface {

	// Gets a card.
	GetMagic(ctx context.Context, name string) (models.Magic, error)

	// Lists a collection of cards.
	ListMagics(ctx context.Context) ([]models.Magic, error)
}
