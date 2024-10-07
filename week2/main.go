package main

import "fmt"

type Player struct {
	Name         string
	Age          int
	JerseyNumber int
	Team         string
}

func main() {
	foppa := Player{Name: "John", Age: 36, JerseyNumber: 11, Team: "Group2"}
	zata := Player{Name: "Mia", Age: 26, JerseyNumber: 12, Team: "Group1"}
	kata := Player{Name: "Teo", Age: 46, JerseyNumber: 13, Team: "Group2"}

	var players = [...]Player{foppa, zata, kata}
	for _, player := range players {

		fmt.Println(player.Name)
	}
}
