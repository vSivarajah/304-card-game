package game_logic


func importGameRules() map[int]float32{
	gameRules := make(map[int]float32)
	gameRules[60] = 14.5
	gameRules[70] = 13.5
	gameRules[80] = 12.5
	gameRules[90] = 11.5
	gameRules[100] = 10.5
	gameRules[110] = 9.5
	gameRules[250] = 5.5
	gameRules[260] = 5.5

	return gameRules
}