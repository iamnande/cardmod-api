package main

import (
	"context"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/repositories"
)

// Card represents a card.
type Card struct {
	name  string
	level int32
}

func (c *Card) Name() string {
	return c.name
}

func (c *Card) Level() int32 {
	return c.level
}

var (
	cards = []*Card{
		{name: "Geezard", level: 1},
		{name: "Funguar", level: 1},
		{name: "Bite Bug", level: 1},
		{name: "Red Bat", level: 1},
		{name: "Gayla", level: 1},
		{name: "Gesper", level: 1},
		{name: "Fastitocalon-F", level: 1},
		{name: "Blood Soul", level: 1},
		{name: "Caterchipillar", level: 1},
		{name: "Cockatrice", level: 1},
		{name: "Grat", level: 2},
		{name: "Buel", level: 2},
		{name: "Mesmerize", level: 2},
		{name: "Glacial Eye", level: 2},
		{name: "Belhelmel", level: 2},
		{name: "Thrustaevis", level: 2},
		{name: "Anacondaur", level: 2},
		{name: "Creeps", level: 2},
		{name: "Grendel", level: 2},
		{name: "Jelleye", level: 2},
		{name: "Grand Mantis", level: 2},
		{name: "Forbidden", level: 3},
		{name: "Armadodo", level: 3},
		{name: "Tri-Face", level: 3},
		{name: "Fastitocalon", level: 3},
		{name: "Snow Lion", level: 3},
		{name: "Ochu", level: 3},
		{name: "Death Claw", level: 3},
		{name: "Cactuar", level: 3},
		{name: "Tonberry", level: 3},
		{name: "Abyss Worm", level: 3},
		{name: "Turtapod", level: 4},
		{name: "Vysage", level: 4},
		{name: "T-Rexaur", level: 4},
		{name: "Bomb", level: 4},
		{name: "Blitz", level: 4},
		{name: "Wendigo", level: 4},
		{name: "Torama", level: 4},
		{name: "Imp", level: 4},
		{name: "Blue Dragon", level: 4},
		{name: "Adamantoise", level: 4},
		{name: "Hexadragon", level: 4},
		{name: "Iron Giant", level: 5},
		{name: "Behemoth", level: 5},
		{name: "Chimera", level: 5},
		{name: "PuPu", level: 5},
		{name: "Elastoid", level: 5},
		{name: "GIM47N", level: 5},
		{name: "Malboro", level: 5},
		{name: "Elnoyle", level: 5},
		{name: "Tonberry King", level: 5},
		{name: "Wedge, Biggs", level: 5},
		{name: "Fujin, Raijin", level: 6},
		{name: "Elvoret", level: 6},
		{name: "X-ATM092", level: 6},
		{name: "Granaldo", level: 6},
		{name: "Gerogero", level: 6},
		{name: "Iguion", level: 6},
		{name: "Abadon", level: 6},
		{name: "Trauma", level: 6},
		{name: "Oilboyle", level: 6},
		{name: "Shumi Tribe", level: 6},
		{name: "Krysta", level: 6},
		{name: "Propagator", level: 7},
		{name: "Jumbo Cactuar", level: 7},
		{name: "Tri-Point", level: 7},
		{name: "Gargantua", level: 7},
		{name: "Mobile Type 8", level: 7},
		{name: "Sphinxara", level: 7},
		{name: "Tiamat", level: 7},
		{name: "BGH251F2", level: 7},
		{name: "Red Giant", level: 7},
		{name: "Catoblepas", level: 7},
		{name: "Ultima Weapon", level: 7},
		{name: "Chubby Chocobo", level: 8},
		{name: "Angelo", level: 8},
		{name: "Gilgamesh", level: 8},
		{name: "MiniMog", level: 8},
		{name: "Chicobo", level: 8},
		{name: "Quezacotl", level: 8},
		{name: "Shiva", level: 8},
		{name: "Ifrit", level: 8},
		{name: "Siren", level: 8},
		{name: "Sacred", level: 8},
		{name: "Minotaur", level: 8},
		{name: "Carbuncle", level: 9},
		{name: "Diablos", level: 9},
		{name: "Leviathan", level: 9},
		{name: "Odin", level: 9},
		{name: "Pandemona", level: 9},
		{name: "Cerberus", level: 9},
		{name: "Alexander", level: 9},
		{name: "Phoenix", level: 9},
		{name: "Bahamut", level: 9},
		{name: "Doomtrain", level: 9},
		{name: "Eden", level: 9},
		{name: "Ward", level: 10},
		{name: "Kiros", level: 10},
		{name: "Laguna", level: 10},
		{name: "Selphie", level: 10},
		{name: "Quistis", level: 10},
		{name: "Irvine", level: 10},
		{name: "Zell", level: 10},
		{name: "Rinoa", level: 10},
		{name: "Edea", level: 10},
		{name: "Seifer", level: 10},
		{name: "Squall", level: 10},
	}
)

func seedCards(ctx context.Context, db *database.Client) error {

	// initialize client
	cardDB := repositories.NewCardRepository(db)

	// create the diff
	for i := 0; i < len(cards); i++ {
		card, err := cardDB.GetCard(ctx, cards[i].Name())
		if err != nil {
			if !cerrors.IsAPIError(err) {
				return err
			}
		}

		if card == nil {
			_, err = cardDB.CreateCard(ctx, cards[i].Name(), cards[i].Level())
			if err != nil {
				return err
			}
		}
	}

	// success
	return nil

}
