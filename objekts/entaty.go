package objekts

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vektor struct {
	X int
	Y int
}

type Entaty struct {
	Posison     Vektor
	Size        Vektor
	Sprit       *ebiten.Image
	ImageOption *ebiten.DrawImageOptions
}

//func NewEntaty(x int,y int, h int,w int) Entaty {
//	return Entaty{ Vektor{x,y}, Vektor{h,w}}
//}

func (entaty *Entaty) AMove(x int, y int) {
	entaty.Posison.X += x
	entaty.Posison.Y += y
	entaty.ImageOption.GeoM.Translate(float64(y), float64(x))
	fmt.Print(entaty.ImageOption)
}

func (entaty *Entaty) DrawEntaty(screen *ebiten.Image) {
	screen.DrawImage(entaty.Sprit, entaty.ImageOption)
}
