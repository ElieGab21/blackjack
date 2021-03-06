# Blackjack
A blackjack game that uses a general purpose deck package in Golang
The deck import in blackjack.go may change depending on the folder 
Supports multiple players vs 1 dealer
Supports betting

RULES:
1. A blackjack wins the game for the player (unless dealer also has blackjack)
2. If the player has above 21, the dealer wins
3. At >= 17, the dealer stops playing, unless it is a soft 17 or above (with an Ace)
4. The player with the bigger hands win
5. Supports the following actions for a player: Hit/ Stand/ Double Down

# Running the code
    $ go run blackjack.go -decks=2 -players=2 -shuffle=true -money=100  

use --help for details on the flag

# The deck folder
Deck.go is a general purpose deck package that doesn't apply only to blackjack. Contains:
- Create a deck (of multiple decks)
- Filter cards out of a deck
- Shuffle a deck
- Add jokers
- Cards themselves with Card Rank and Values
