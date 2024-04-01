package main

import (
	"fmt"
	"time"
	"goTowerDefens/objekts"
)

func main() {
	//test := objekts.Vektor{X: 0,Y: 0}
	//player := new(objekts.Player)
	player := objekts.NewPlayer(0,0,10,10,"brasse")
	gWorld := NewGameWorld()
	gWorld.addEntaty(&player.PlayerEnataty)
	fmt.Println(gWorld)
	//fmt.Println(player)
	ticker := time.NewTicker(1 * time.Second)
	go loop(ticker, gWorld)
	inputLoop(ticker, player)
}