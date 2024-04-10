package main

import (
	"GoTowerDefense/objekts"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/quartercastle/vector"
)

var img *ebiten.Image
var spritMap = map[string]Vec{
	"landL":     {0, 1},
	"landR":     {2, 1},
	"landU":     {1, 0},
	"landD":     {1, 3},
	"land":      {3, 0},
	"landGrass": {4, 0},
	"Sine":      {5, 0},
	"watter":    {1, 1},
	"houseL":    {1, 3},
	"houseR":    {1, 4},
	"path|":     {2, 3},
	"path-":     {2, 4},
	"path-endR": {2, 5},
	"path-endL": {1, 5},
}

type Vec = vector.Vector

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("image/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

func loadImageFromFile(pathToFile string) *ebiten.Image {
	var err error
	var tempImg *ebiten.Image
	tempImg, _, err = ebitenutil.NewImageFromFile(pathToFile)
	if err != nil {
		log.Fatal(err)
		tempImg = nil
	}
	return tempImg
}

type Game struct {
	GWorld *GameWorld
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	g.GWorld.DrawMap(screen)
	g.GWorld.DrawEntaty(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func runGui(gWorld *GameWorld) {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{GWorld: gWorld}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//test := objekts.Vektor{X: 0,Y: 0}
	//player := new(objekts.Player)
	player := objekts.NewPlayer(0, 0, 10, 10, "brasse")
	gWorld := NewGameWorld()
	gWorld.addEntaty(&player.PlayerEnataty)
	gWorld.MapSprit = loadImageFromFile("image/map-tails.png")
	imageSize := objekts.Vektor{X: (256 / 8), Y: (256 / 8)}

	sx := 5 * (256 / 8)
	sy := 0 * (256 / 8)
	gWorld.Entatys[0].Sprit = gWorld.MapSprit.SubImage(image.Rect(sx, sy, sx+(256/8), sy+(256/8))).(*ebiten.Image)
	gWorld.Entatys[0].ImageOption = &ebiten.DrawImageOptions{}
	gWorld.Entatys[0].ImageOption.GeoM.Scale(2, 2)
	gWorld.Entatys[0].ImageOption.GeoM.Translate(0, 0)
	gWorld.Size = &Vec{(256 / 8), (256 / 8)}
	tempMapVektor := [5][5]Vec{}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 0:
				tempMapVektor[j][i] = spritMap["landR"]
			case 3:
				tempMapVektor[j][i] = spritMap["path-endR"]
			case 4:
				tempMapVektor[j][i] = spritMap["land"]
			default:
				tempMapVektor[j][i] = spritMap["path-"]
			}
		}
	}
	//i := 0
	x := int(spritMap["houseR"].X() * float64(imageSize.X))
	y := int(spritMap["houseR"].Y() * float64(imageSize.Y))
	var px float64
	var py float64
	for i := 0; i < 5; i++ {
		fmt.Print("test: ", i)
		tempEntaty := &objekts.Entaty{}
		tempEntaty.Posison = objekts.Vektor{X: i, Y: 1}
		tempEntaty.Size = imageSize
		tempEntaty.Sprit = gWorld.MapSprit.SubImage(image.Rect(x, y, x+imageSize.X, y+imageSize.Y)).(*ebiten.Image)
		tempEntaty.ImageOption = &ebiten.DrawImageOptions{}
		tempEntaty.ImageOption.GeoM.Scale(2, 2)
		px = float64(2 * imageSize.Y)
		py = float64((i * 2) * imageSize.Y)
		tempEntaty.ImageOption.GeoM.Translate(px, py)
		gWorld.addEntaty(tempEntaty)
	}

	gWorld.loadeImages(tempMapVektor, 2)
	//fmt.Println(gWorld)
	fmt.Println(gWorld.Entatys[0])
	ticker := time.NewTicker(1 * time.Second)
	go runGui(gWorld)
	go loop(ticker, gWorld)
	inputLoop(ticker, player)
}
