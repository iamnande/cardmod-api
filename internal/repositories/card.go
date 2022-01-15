package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/iamnande/cardmod/internal/daos"
	"github.com/iamnande/cardmod/internal/database"
	dbCard "github.com/iamnande/cardmod/internal/database/card"
	"github.com/iamnande/cardmod/internal/domains/card"
)

// cardRepository is a repository layer for accessing card entities in the data layer.
//go:generate mockgen -source ../daos/card.go -destination=mocks/card.go -package mocks github.com/iamnande/cardmod/internal/daos CardDAO
type cardRepository struct {
	client *database.Client
}

// this is a build/compile time check to ensure our implementation satisfies the DAO interface.
var _ daos.CardDAO = (*cardRepository)(nil)

// NewCardRepository initializes a new card repository instance.
func NewCardRepository(client *database.Client) *cardRepository {
	return &cardRepository{
		client: client,
	}
}

// ListCards lists all available card entities.
func (repo *cardRepository) ListCards(ctx context.Context, search string) ([]card.Card, error) {

	// list: base query
	query := repo.client.Card.Query()

	// list: if there's a search limiter, include it
	if search != "" {
		query = query.Where(dbCard.Or(
			dbCard.NameContainsFold(search),
		))
	}

	// list: list the cards
	cards, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	// list: return the list of cards
	return marshalCardContainers(cards), nil

}

// CreateCard creates a new card entity.
func (repo *cardRepository) CreateCard(ctx context.Context, name string) (card.Card, error) {

	// create: initialize transaction
	tx, err := repo.client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	// create: create the card
	card, err := tx.Card.Create().
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	// create: return the newly created card
	return &cardContainer{card}, tx.Commit()

}

// GetCard gets a card entity.
func (repo *cardRepository) GetCard(ctx context.Context, id uuid.UUID) (card.Card, error) {

	// get: gets the card
	card, err := repo.client.Card.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// get: return the card to be retrieved
	return &cardContainer{card}, nil

}

// DeleteCard deletes an existing card entity.
func (repo *cardRepository) DeleteCard(ctx context.Context, id uuid.UUID) error {

	// delete: initialize transaction
	tx, err := repo.client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	// delete: delete the card
	if err := repo.client.Card.DeleteOneID(id).Exec(ctx); err != nil {
		return rollback(tx, err)
	}

	// delete: return "success"
	return tx.Commit()

}
