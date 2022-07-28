package main

import (
	"fmt"
	"map/models"
)

type student struct {
	Name string
	Age  int
}

func BuildMap() {
	//定义map
	m := make(map[string]*student)

	//定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	// 遍历结构体数组，依次赋值给map
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

	//打印map
	for _, v := range m {
		fmt.Println(v.Name, v.Age)
	}
}

func IterateMap() {
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		m[1] = i
	}

	fmt.Println(m)
}

//LongQueueRemind 长龙提醒
func LongQueueRemind() {

	lotteries := models.GetLotteries()

	var resMap = make(map[int64]map[string]map[string]int)
	for _, lottery := range lotteries {
		DaXiao := lottery.DaXiao

		sec := lottery.BlockInfo.BlockTime / 1000 % 60
		if resMap[sec] == nil {
			resMap[sec] = map[string]map[string]int{
				"dx": {
					"d": 0, // 大
					"x": 0, // 小
				}, // 1
			}
		}

		if DaXiao == 1 {
			// 小
			resMap[sec]["dx"]["x"]++
			resMap[sec]["dx"]["d"] = 0
		}
		if DaXiao == 2 {
			// 大
			resMap[sec]["dx"]["d"]++
			resMap[sec]["dx"]["x"] = 0
		}
		fmt.Println(resMap[sec]["dx"])
	}

	fmt.Println("----------------------")
	for key1 := range resMap {
		for key2 := range resMap[key1] {
			for key3 := range resMap[key1][key2] {
				num := resMap[key1][key2][key3]
				if num >= 3 {
					fmt.Println(key1, key2, key3, num)
				}
			}
		}
	}
}
