package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Tobil"}

	gamer.AddGame("Mario")
	gamer.AddGame("PUBG")
	gamer.AddGame("ML")
	gamer.AddGame("Fifa 2020")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
