package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	// a := &Actor{name, hp, speed}
	// return a

	a := Actor{name, hp, speed}
	return &a
}

func main() {
	var actor = NewActor("금토끼", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Printf(actor.Name)
	fmt.Println(*actor)
}
