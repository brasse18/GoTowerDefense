package main

import (
	"fmt"
	"time"
	"GoTowerDefense/objekts"
	"log"
	_ "image/png"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{
	GWorld *GameWorld
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(g.GWorld.Entatys[0].Posison.Y), float64(g.GWorld.Entatys[0].Posison.X))
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func runGui(gWorld *GameWorld) {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{GWorld: gWorld}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//test := objekts.Vektor{X: 0,Y: 0}
	//player := new(objekts.Player)
	player := objekts.NewPlayer(0,0,10,10,"brasse")
	gWorld := NewGameWorld()
	gWorld.addEntaty(&player.PlayerEnataty)
	fmt.Println(gWorld)
	//fmt.Println(player)
	ticker := time.NewTicker(1 * time.Second)
	go runGui(gWorld)
	go loop(ticker, gWorld)
	inputLoop(ticker, player)
}