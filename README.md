# blackjack
A blackjack game that uses a general purpose deck package in Golang
The deck import in blackjack.go may change depending on the folder 
Only supports 1 player and 1 dealer so far

RULES:
1. A blackjack wins the game for the player (unless dealer also has blackjack)
2. If the player has above 21, the dealer wins
3. At >= 17, the dealer stops playing, unless it is a soft 17 or above (with an Ace)
4. The player with the bigger hands win
5. Supports the following actions for a player: Hit/ Stand/ Double Down
