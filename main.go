package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/manifoldco/promptui"
	"github.com/vSivarajah/304-card-game/game_logic"
)

func main(){
	playerOne := game_logic.Player{
		Trump: false,
		Name:  "Viggi",
		Team:  0,
	}
	playerTwo := game_logic.Player{
		Trump: false,
		Name:  "Test1",
		Team:  1,
	}
	playerThree := game_logic.Player{
		Trump: false,
		Name:  "Test2",
		Team:  0,
	}
	playerFour := game_logic.Player{
		Trump: true,
		Name:  "Test3",
		Team:  1,
	}


	players := game_logic.Players{playerOne, playerTwo, playerThree, playerFour}
	cards := game_logic.NewGame(players)
	playerOne.Cards = playerOne.DealHand(cards,4)
	playerTwo.Cards = playerTwo.DealHand(cards,4)
	playerThree.Cards = playerThree.DealHand(cards,4)
	playerFour.Cards = playerFour.DealHand(cards,4)



	players = game_logic.Players{playerOne, playerTwo, playerThree, playerFour}
	trumpPlacer := game_logic.TrumpPlacer{}
	for _, player := range players {
		if player.Trump {
			var qs = []*survey.Question{
				{
					Name:     "bid",
					Prompt:   &survey.Input{Message: "Please place a bid"},
					Validate: survey.Required,
					Transform: survey.Title,
				},
			}

			answers := struct {
				Bid string
			}{}

			err := survey.Ask(qs, &answers)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			prompt := promptui.Select{
				Label:             "Select trump card",
				Items: 				player.Cards,
			}
			i, _, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Printf("You choose %d:  %s\n", i+1, player.Cards[i])
			trumpPlacer.PlacedTrump = player.Cards[i]
			trumpPlacer.Bid = answers.Bid
		}
	}

	fmt.Println(trumpPlacer)


}