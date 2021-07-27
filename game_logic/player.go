package game_logic

import (
	"github.com/vSivarajah/304-card-game/deck"
)

type Player struct {
	Trump bool
	Name string
	Team int
	Cards []deck.Card
	TrumpPlacer TrumpPlacer
}

type TrumpPlacer struct {
	Bid string
	PlacedTrump deck.Card
}

func NewGame(players Players) []deck.Card {
	cards := deck.New(deck.DefaultSort, deck.Shuffle)
	return cards
}



func (p Player) DealHand(cards []deck.Card, numberOfCards int) []deck.Card{
	deckPlayer := []deck.Card{}


	for i := 0; i < numberOfCards; i++ {
		deckPlayer = append(deckPlayer, cards[0])
		deck.Remove(cards, 0)
	}
	return deckPlayer
}

type Players []Player


