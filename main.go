package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/vSivarajah/304-card-game/game_logic"
)

var trumpArray = []bool{false, false, false, true}


func main() {
	players := game_logic.Players{{
		Trump: false,
		Name:  "Viggi",
		Team:  0,
	},
		{
			Trump: false,
			Name:  "player 2",
			Team:  1,
		},
		{
			Trump: false,
			Name:  "Player 3",
			Team:  0,
		},
		{
			Trump: true,
			Name:  "Player 4",
			Team:  1,
		},
	}

	for i := 0; i < 12; i++ {
		fmt.Println("Round #", i)
		trumpPlacer := game_logic.TrumpPlacer{}
		cards := game_logic.NewGame(players)
		rotate(trumpArray)
		fmt.Println(trumpArray)
		for i, player := range players {
			player.Trump = trumpArray[i]
			player.Cards = player.DealHand(cards, 4)

			if player.Trump {
				fmt.Println("Trump holder: ", player.Name, ": ", player.Trump)
				promptBid := promptui.Select{
					Label: "It is your turn to place Bid",
					Items: []int{60, 70, 80, 90, 100},
				}
				_, Bid, err := promptBid.Run()

				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}

				promptTrump := promptui.Select{
					Label: "Select trump card",
					Items: player.Cards,
				}
				i, _, err := promptTrump.Run()

				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}

				fmt.Printf("You choose %d:  %s\n", i+1, player.Cards[i])
				trumpPlacer.PlacedTrump = player.Cards[i]
				trumpPlacer.Bid = Bid

			}
		}

		fmt.Println(trumpPlacer)
	}

}

func rotate(trumpHolders []bool) {
	result := append(trumpHolders[len(trumpHolders)-1:], trumpHolders[:len(trumpHolders)-1]...)
	for i := 0; i < len(trumpHolders); i++ {
		trumpHolders[i] = result[i]
	}
}