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
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

type CardOptions func(c []Card) []Card
type LessFn func(cards []Card) func(i, j int) bool

func New(options ...CardOptions) []Card {
	var deck []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			deck = append(deck, Card{suit, rank})
		}
	}

	for _, op := range options {
		deck = op(deck)
	}

	return deck
}

func Deck(n int) CardOptions {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}

func Filter(f func(card Card) bool) CardOptions {
	return func(cards []Card) []Card {
		var ret []Card

		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}

		return ret
	}
}

func Jokers(n int) CardOptions {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Joker, Rank(i)})
		}
		return cards
	}
}

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i, j := range r.Perm(len(cards)) {
		ret[i] = cards[j]
	}
	return ret
}

func Sort(less LessFn) CardOptions {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
