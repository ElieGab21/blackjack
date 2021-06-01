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
	var value int

	for _, card := range h {

		switch card.Rank {
		case d.Ace: //Ace value varies between 1 and 11
			if value <= 10 {
				value += 11
			} else {
				value += 1
			}
		case d.Jack, d.Queen, d.King: //All have the same value of 10
			value += 10
		default:
			value += int(card.Rank)
		}
	}

	return value
}

func draw(deck []d.Card, nmbOfCards int, hand Hand) ([]d.Card, Hand) {
	for i := 0; i < nmbOfCards; i++ {
		hand = append(hand, deck[0])
		deck = deck[1:]
	}

	return deck, hand
}

func dealerChoice(deck []d.Card, dealer Hand) {

	for {
		fmt.Printf("Dealer: %v  (%d) \n\n", dealer, dealer.value())

		if dealer.value() == blackjack {
			fmt.Println("Dealer has won!")
			break

		} else if dealer.value() >= 17 && dealer.value() < blackjack {
			for _, card := range dealer {
				if card.Rank == d.Ace {
					fmt.Println("hi")
					deck, dealer = draw(deck, 1, dealer)
					break
				} else {
					return
				}
			}
		} else if dealer.value() > blackjack {
			fmt.Println("Dealer has busted. Player wins")
			break

		} else {
			deck, dealer = draw(deck, 1, dealer)
		}

	}

}

func start(deck []d.Card, player, dealer Hand) {

	for i := 0; i < 2; i++ {

		deck, player = draw(deck, 1, player)
		deck, dealer = draw(deck, 1, dealer)
	}

	fmt.Println()
	fmt.Printf("Dealer: %v \n\n", dealer[0])
	fmt.Printf("Player: %v  (%d) \n\n", player, player.value())

	if player.value() == blackjack {
		fmt.Println("The player has won")
		return
	}

	var action int

	for {
		fmt.Println("1. Hit")
		fmt.Println("2. Stand")
		fmt.Scanf("%d", &action)

		switch action {
		case 1:
			deck, player = draw(deck, 1, player) //Player draws 1
			fmt.Println()
			fmt.Printf("Dealer: %v \n\n", dealer[0])
			fmt.Printf("Player: %v  (%d) \n\n", player, player.value())
		case 2:
			if dealer.value() < blackjack {
				dealerChoice(deck, dealer)

				// fmt.Printf("Dealer: %v  (%d) \n\n", dealer, dealer.value())
				fmt.Printf("Player: %v  (%d) \n\n", player, player.value())

				if player.value() > blackjack {
					fmt.Println("You have been busted")
					return

				} else if player.value() == dealer.value() {
					fmt.Println("It is a push. Nobody wins")
					return

				} else if player.value() > dealer.value() {
					fmt.Println("The player has won")
					return

				} else if player.value() < dealer.value() {
					fmt.Println("The dealer has won")
					return
				}

			} else if dealer.value() == blackjack {
				fmt.Printf("Dealer: %v  (%d) \n\n", dealer, dealer.value())
				fmt.Printf("Player: %v  (%d) \n\n", player, player.value())
				fmt.Println("The dealer has won")
				return

			} else {
				fmt.Println("Dealer is busted!")
			}
		}

		if player.value() > blackjack {
			fmt.Println("You have been busted!")
			return
		}
	}

}

func main() {
	numberOfDecks := flag.Uint("decks", 3, "The number of decks at the start of the game. Default 3 and cannot be < 1")
	shuffle := flag.Bool("shuffle", true, "Suffle the deck at the start. Default true. Must be boolean")

	flag.Parse()

	var player, dealer Hand

	func(player, dealer Hand) {
		var option int

		for {
			deck := func(nmb uint, shuff bool) []d.Card {
				if shuff {
					return d.New(d.Deck(int(nmb)), d.Shuffle)
				} else {
					return d.New(d.Deck(int(nmb)))
				}
			}(*numberOfDecks, *shuffle)

			fmt.Println("What do you want to do?")
			fmt.Println("1. Play the game")
			fmt.Println("2. Exit the game")

			fmt.Scanf("%d", &option)

			switch option {
			case 1:
				fmt.Print("Let's start! \n\n")
				start(deck, player, dealer)

			case 2:
				fmt.Println("Thanks for playing!")
				return
			}
		}
	}(player, dealer)

	// for i := 0; i < 2; i++ {
	// 	deck, player = draw(deck, 1, player)
	// 	deck, dealer = draw(deck, 1, dealer)
	// }

	// fmt.Println()
	// fmt.Printf("Dealer: %v \n\n", dealer[0])
	// fmt.Printf("Player: %v  (%d) \n\n", player, player.value())

	// if player.value() == 21 {
	// 	fmt.Println("The player has won")
	// }

}
