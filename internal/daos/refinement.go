package daos

import (
	"github.com/iamnande/cardmod/internal/models"
)

// RefinementDAO is the DAO for the Refinement model.
//go:generate mockgen -source refinement.go -destination=./mocks/refinement.go -package mocks
type RefinementDAO interface {

	// Lists a collection of refinements.
	ListRefinements() []models.Refinement

	// Gets a refinement.
	GetRefinement(source, target string) (models.Refinement, error)
}