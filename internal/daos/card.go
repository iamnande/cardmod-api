package daos

import (
	"context"

	"github.com/google/uuid"

	"github.com/iamnande/cardmod/internal/domains/card"
)

// CardDAO is the data access object (controller<->data layer) interface definition.
type CardDAO interface {

	// ListCards lists all card entities.
	ListCards(ctx context.Context) ([]card.Card, error)

	// CreateCard creates a card entity.
	CreateCard(ctx context.Context, name string) (card.Card, error)

	// DescribeCard describes a card entity.
	DescribeCard(ctx context.Context, id uuid.UUID) (card.Card, error)

	// DeleteCard deletes a card entity.
	DeleteCard(ctx context.Context, id uuid.UUID) error
}