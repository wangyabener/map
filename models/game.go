package models

type result struct {
	Key string
	Val int64
}

type Game struct {
	Name string
	Code string
	Res  []result
}

func GetGames() [4]Game {
	data := [4]Game{
		{
			Name: "大小", Code: "dx", Res: []result{
				{Key: "d", Val: 0}, {Key: "x", Val: 0},
			},
		},
		{
			Name: "单双", Code: "ds", Res: []result{
				{Key: "d", Val: 0}, {Key: "s", Val: 0},
			},
		},
		{
			Name: "双尾", Code: "sw", Res: []result{
				{Key: "t", Val: 0}, {Key: "f", Val: 0},
			},
		},
		{
			Name: "牛牛", Code: "nn", Res: []result{
				{Key: "z", Val: 0}, {Key: "x", Val: 0},
			},
		},
	}

	return data
}
