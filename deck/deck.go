//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit int

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank int

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

//Returns a new deck
func New(opts ...func([]Card) []Card) []Card {
	var deck []Card

	for suit := Spade; suit < Joker; suit++ {
		for rank := minRank; rank <= maxRank; rank++ {
			deck = append(deck, Card{suit, rank})

		}
	}

	for _, opt := range opts {
		deck = opt(deck)
	}

	return deck
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

//Sort with a custom less function
func Sort(less func(cards []Card) func(i, j int) bool) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}

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

func Jokers(n int) func(cards []Card) []Card {

	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) func(cards []Card) []Card {
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

func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card

		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}

		return ret
	}
}
