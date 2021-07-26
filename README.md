# 304-card-game
Go implementation of the south asian card game 304

## Game rules
* Jack 30
* Nines 20
* Aces 11
* Ten 10
* King 3
* Queen 2
* Eight 0 
* Seven 0

The game is named after the total number of points in the pack, which is 304.

### Flow of game
* Four players maximum
* Two teams 
* The cards are dealt by the dealer to all four players in a counter-clockwise manner
* Each getting four cards in the first round.
* The player to the right of the dealer picks out a trump card from their hand and places it face down on the table.
* They then call a target score for their team, which they expect to be able to win judging by the cards in their own hand.
* A member of the opposing team can then pick their own trump, and call a target score higher than the first player's.
* The rest of the cards are distributed and the game is started.

* The game is played as a trick-taker, where the players follow the suit led.
* If they do not have the particular suit, then they can try to guess the trump by passing a face down card to the player who has closed a trump card. 
* If they guess it correctly then they get the hand; otherwise the game goes on.

Points are scored according to the points shown in the table. After all the hands are completed, if the team has scored the required score they win.