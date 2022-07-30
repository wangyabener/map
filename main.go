package main

import "map/reminder"

type M struct{}
type In struct{}
type Out struct{}

func (m *M) Example(in In) Out {
	return Out{}
}

func main() {
	// f, _ := service.InvokeObjectMethod(new(service.Lucky), "Daxiao", "234uy2a5")
	// fmt.Println(f[0].Int())
	reminder.Build()
	// BuildMap()
	// IterateMap()
	// LongQueueRemind()
}
