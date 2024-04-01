package main

import (
	"GoTowerDefense/objekts"
)

type GameWorld struct {
	Entatys [10]*objekts.Entaty
	MapSize int
}

func (gWorld *GameWorld) addEntaty(inEntaty *objekts.Entaty) {
	//gWorld.Entatys = append(gWorld.Entatys, inEntaty)
	gWorld.Entatys[0] = inEntaty
}

func NewGameWorld() *GameWorld {
	return &GameWorld{Entatys: [10]*objekts.Entaty{}, MapSize: 0}
}