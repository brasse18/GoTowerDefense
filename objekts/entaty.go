package objekts

type Vektor struct {
	X int
	Y int
}

type Entaty struct {
	Posison Vektor
	size Vektor
}

func NewEntaty(x int,y int, h int,w int) Entaty {
	return Entaty{ Vektor{x,y}, Vektor{h,w}}
}

func (entaty *Entaty) AMove(x int, y int) {
	entaty.Posison.X += x
	entaty.Posison.Y += y
}