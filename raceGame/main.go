package racegame

import "fmt"

func RaceGame(players map[Drivable]int) {
	for range make([]int, 5) {
		for player := range players {
			players[player] += player.Drive()
		}
	}

	var winner Drivable
	var max = -1
	for driver := range players {
		if players[driver] > max {
			max = players[driver]
			winner = driver
		}
	}

	fmt.Print("The winner is: ", winner)
}
