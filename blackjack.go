package main

import (
	"flag"
	"fmt"
	d "github/blackjack/deck"
	"strings"
)

const (
	blackjack = 21
)

/* Hand type + methods */

//Hand type, a slice of cards
type Hand []d.Card

//printing a hand
func (h Hand) String() string {
	var temp []string

	for _, card := range h {
		temp = append(temp, card.String())
	}

	return strings.Join(temp, ", ")
}

//To calculate the value of the hand
func (h Hand) value() int {
	var value, aces int

	for _, card := range h {
		switch card.Rank {
		case d.Ace: //Ace value varies between 1 and 11
			aces++
		case d.Jack, d.Queen, d.King: //All have the same value of 10
			value += 10
		default:
			value += int(card.Rank)
		}
	}

	if aces > 0 {
		var possibleSolutions []int

		for i := 0; i <= aces; i++ {
			possibleSolutions = append(possibleSolutions, value+(aces-i)+(11*i))
		}

		value = func(v []int) (m int) {
			m = v[0]

			for i := 1; i < len(v); i++ {
				if v[i] > m && v[i] <= blackjack {
					m = v[i]
				}
			}
			return
		}(possibleSolutions)
	}

	return value
}

/* Players type + methods */
type players []Hand

func (p *players) initialise(deck *[]d.Card, playersNmb int) {
	for i := 0; i < playersNmb; i++ {
		var player Hand
		*deck, player = draw(*deck, 2, player)
		*p = append(*p, player)
	}

	//return *deck
}

/* Main game functions */
func draw(deck []d.Card, nmbOfCards int, hand Hand) ([]d.Card, Hand) {
	for i := 0; i < nmbOfCards; i++ {
		hand = append(hand, deck[0])
		deck = deck[1:]
	}

	return deck, hand
}

func dealerTurn(deck []d.Card, dealer Hand) int {

	for {
		fmt.Printf("Dealer: %v  (%d) \n\n", dealer, dealer.value())

		if dealer.value() == blackjack {
			break

		} else if dealer.value() >= 17 && dealer.value() < blackjack {
			for _, card := range dealer {
				if card.Rank == d.Ace {
					deck, dealer = draw(deck, 1, dealer)
					break
				} else {
					return dealer.value()
				}
			}
		} else if dealer.value() > blackjack {
			break

		} else {
			deck, dealer = draw(deck, 1, dealer)
		}

	}

	return dealer.value()

}

func playerTurn(deck *[]d.Card, player, dealer Hand, playerNumber int) {

	var action int
	doubleDown := false
	turn := 0

	for {
		fmt.Printf("\nDealer: %v \n\n", dealer[0])
		fmt.Printf("Player %d: %v  (%d) \n\n", playerNumber, player, player.value())

		if doubleDown {
			action = 2 //Goes directly to dealer phase with a double down
		} else {
			fmt.Printf("Player %d turn: \n", playerNumber)
			fmt.Println("1. Hit\n2. Stand\n3. Double down")
			fmt.Scanf("%d", &action)
		}

		switch action {
		case 1:
			*deck, player = draw(*deck, 1, player) //Player draws 1
			fmt.Printf("\nDealer: %v \n\n", dealer[0])
			fmt.Printf("Player %d: %v  (%d) \n\n", playerNumber, player, player.value())
			turn++

			if player.value() > blackjack {
				return
			}

		case 2:
			return
		case 3:
			if turn > 0 {
				fmt.Print("Cannot double down after hitting once\n\n")
				break
			}
			*deck, player = draw(*deck, 1, player) //Player draws 1
			fmt.Printf("\nDealer: %v \n\n", dealer[0])
			fmt.Printf("Player %d: %v  (%d) \n\n", playerNumber, player, player.value())

			if player.value() > blackjack {
				return
			} else {
				doubleDown = true
			}
		default:
			fmt.Println("Must be a number from the options")
			continue
		}
	}
}

func start(deck *[]d.Card, dealer Hand, allPlayers players) {

	for i := 0; i < 2; i++ {
		*deck, dealer = draw(*deck, 1, dealer)
	}

	for i, player := range allPlayers {
		if player.value() == blackjack {
			fmt.Printf("Player %d: %v  (%d) \n\n", i+1, player, player.value())
			continue
		} else {
			playerTurn(deck, player, dealer, i+1)
		}
	}

	dealerValue := dealerTurn(*deck, dealer)

	for i, player := range allPlayers {
		switch {
		case dealerValue > blackjack:
			if player.value() > blackjack {
				fmt.Printf("Both dealer and player %d are busted \n", i+1)
			} else {
				fmt.Printf("Player %d wins ! \n", i+1)
			}
		case dealerValue < blackjack:
			if player.value() == dealerValue {
				fmt.Printf("It is a push. Player %d keeps his bet \n", i+1)
			} else if player.value() > blackjack || dealerValue > player.value() {
				fmt.Printf("Dealer beats player %d \n", i+1)
			} else if player.value() > dealerValue {
				fmt.Printf("Player %d wins ! \n", i+1)
			}
		case dealerValue == blackjack:
			if dealerValue == player.value() {
				fmt.Printf("It is a push. Player %d keeps his bet \n", i+1)
			} else {
				fmt.Printf("Dealer beats player %d \n", i+1)
			}
		}
	}

}

func main() {
	numberOfDecks := flag.Uint("decks", 3, "The number of decks at the start of the game. Default 3 and cannot be < 1")
	shuffle := flag.Bool("shuffle", true, "Suffle the deck at the start. Default true. Must be boolean")
	numberOfPlayers := flag.Uint("players", 1, "Number of players vs the dealer (default 1, cannot be < 1)")

	flag.Parse()
	var dealer Hand
	var option int

	for {
		deck := func(nmb uint, shuff bool) []d.Card {
			if shuff {
				return d.New(d.Deck(int(nmb)), d.Shuffle)
			} else {
				return d.New(d.Deck(int(nmb)))
			}
		}(*numberOfDecks, *shuffle)

		var allPlayers players
		allPlayers.initialise(&deck, int(*numberOfPlayers))

		fmt.Println("\n\nWhat do you want to do?")
		fmt.Println("1. Play the game\n2. Exit the game")
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			fmt.Print("Let's start! \n")
			start(&deck, dealer, allPlayers)

		case 2:
			fmt.Println("Thanks for playing!")
			return
		}
	}

}
