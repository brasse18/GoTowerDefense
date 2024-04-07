package main

import (
	"image"
	"GoTowerDefense/objekts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quartercastle/vector"
)

type GameWorld struct {
	MapSprit *ebiten.Image
	Size *vector.Vector
	op [5][5]*ebiten.DrawImageOptions
	ImageMap [5][15]*ebiten.Image
	Entatys [10]*objekts.Entaty
}

func (gWorld *GameWorld) addEntaty(inEntaty *objekts.Entaty) {
	//gWorld.Entatys = append(gWorld.Entatys, inEntaty)
	gWorld.Entatys[0] = inEntaty
}

func RelativeCrop(source *ebiten.Image, r image.Rectangle) *ebiten.Image {
    rx, ry := source.Bounds().Min.X+r.Min.X, source.Bounds().Min.Y+r.Min.Y
    return source.SubImage(image.Rect(rx, ry, rx+r.Max.X, ry+r.Max.Y)).(*ebiten.Image)
}

func (gWorld *GameWorld) loadeImages(vec [5][5]vector.Vector) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			gWorld.op[i][j] = &ebiten.DrawImageOptions{}
			gWorld.op[i][j].GeoM.Scale(1, 1)
			gWorld.op[i][j].GeoM.Translate(gWorld.Size.X()*float64(i), gWorld.Size.Y()*float64((j)))
			//fix sönder
			//var rec = image.Rect(int(vec[i][j].X()*50), int(vec[i][j].Y()*50), 50, 50)
			//fungerar ej
			//var rec = image.Rect(0, 0, 50, 50)
			//--------------
			//gWorld.ImageMap[i][j] = gWorld.MapSprit.SubImage(rec).(*ebiten.Image)
			sx := int(vec[i][j].X() * gWorld.Size.X())
			sy := int(vec[i][j].Y() * gWorld.Size.Y())
			gWorld.ImageMap[i][j] = gWorld.MapSprit.SubImage(image.Rect(sx, sy, sx+int(gWorld.Size.X()), sy+int(gWorld.Size.Y()))).(*ebiten.Image)
			//gWorld.ImageMap[i][j] = RelativeCrop(gWorld.MapSprit, rec)
		}
	 }	
}

func (gWorld *GameWorld) DrawMap(screen *ebiten.Image) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			screen.DrawImage(gWorld.ImageMap[i][j], gWorld.op[i][j])
		}
	 }
}

func (gWorld *GameWorld) DrawEntaty(screen *ebiten.Image) {
	screen.DrawImage(gWorld.Entatys[0].Sprit, gWorld.Entatys[0].ImageOption)
}

func NewGameWorld() *GameWorld {
	
	return &GameWorld{Entatys: [10]*objekts.Entaty{}}

}