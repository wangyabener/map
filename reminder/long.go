package reminder

import (
	"fmt"
	"map/models"
	"map/service"
)

func Build() {
	lotteries := models.GetLotteries()

	// game list
	games := models.GetGames()

	fmt.Println(games)

	// secends
	secs := [1]int{}
	for i, j := 0, 0; i < len(secs); i++ {
		secs[i] = j
		j = j + 3
	}

	data := make(map[int]map[string]models.Game)

	for _, lottery := range lotteries {
		for _, sec := range secs {
			items := make(map[string]models.Game)
			for _, game := range games {
				// reflect
				f, _ := service.InvokeObjectMethod(new(service.Lucky), game.Func, lottery.BlockNum)
				// result
				s := f[0].String()

				results := make([]*models.Result, 0, len(game.Res))
				for _, res := range game.Res {
					if s == res.Key {
						res.Val++
					} else {

					}

					results = append(results, &models.Result{
						Key: res.Key,
						Val: res.Val,
					})
				}
				items[game.Code] = models.Game{
					Code: game.Code,
					Res:  results,
				}
			}
			data[sec] = items
		}
	}

	for _, items := range data {
		for _, item := range items {
			for _, v := range item.Res {
				fmt.Println(v.Key, v.Val)
			}
		}
	}
}
