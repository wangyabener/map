package reminder

import (
	"fmt"
	"map/models"
)

func Build() {
	lotteries := models.GetLotteries()

	// game list
	games := models.GetGames()

	// secends
	secs := [20]int{}
	for i, j := 0, 0; i < len(secs); i++ {
		secs[i] = j
		j = j + 3
	}

	// init data
	data := make(map[int]map[string]models.Game)
	for _, sec := range secs {
		items := make(map[string]models.Game)
		for _, game := range games {
			items[game.Code] = models.Game{
				Name: game.Name,
				Code: game.Code,
			}
		}
		data[sec] = items
	}

	// rep := make(map[int64]map[string]models.Game)
	// rep := make(map[int64]map[string]map[string]int)
	rep := map[int64]map[string]map[string]int{
		3: {
			"dx": {
				"d": 0,
				"x": 0,
			},
		},
	}

	for _, lottery := range lotteries {
		sec := lottery.BlockInfo.BlockTime / 1000 % 60
		for _, game := range games {
			rep[sec][game.Code]["d"]++
		}
	}

	fmt.Printf("%v", rep)
}
