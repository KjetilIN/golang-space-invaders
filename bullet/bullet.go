package bullet

type PlayerBullet struct {
	x int
	y int
	speedY int
}

func (b *PlayerBullet) update() {
	b.y += b.speedY
}

func NewPlayerBullet(x int, y int, speedy int) *PlayerBullet{
	var playerBullet = PlayerBullet{
		x: x,
		y: y,
		speedY: speedy,
	}

	return &playerBullet
}
