package objekts

type Player struct {
	PlayerEnataty Entaty
	Name string
	ReDraw bool
}

func (player *Player) AMove(x int, y int) {
	player.PlayerEnataty.AMove(x,y)
	player.ReDraw = true
}

func NewPlayer(x int, y int, h int, w int, name string) *Player {
	return &Player{
		PlayerEnataty: Entaty{
			Posison: Vektor{X: 0,Y: 0}, 
			size: Vektor{X: 10,Y: 10},
			},
		Name: name,
		ReDraw: false,
	}
}