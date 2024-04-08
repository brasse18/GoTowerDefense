package objekts

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	Posison Vektor
	size Vektor
	Sprit *ebiten.Image
	ImageOption *ebiten.DrawImageOptions
}

func (button *Button) DrawButton(screen *ebiten.Image) {
	if button.ImageOption == nil || button.Sprit == nil {

	}
	screen.DrawImage(button.Sprit, button.ImageOption)
}