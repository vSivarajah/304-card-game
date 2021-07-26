//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Seven
	Eight
	Queen
	King
	Ten
	Ace
	Nine
	Jack

)

const (
	minRank = Seven
	maxRank = Jack
)

type Card struct {
	Suit
	Rank
}

func(c Card) String() string {
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

func Remove(slice []Card, s int) []Card {
	return append(slice[:s], slice[s+1:]...)
}

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card

	// for each suit
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Less(cards []Card) func (i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit) * int(maxRank) + int(c.Rank)
}

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c:= range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}