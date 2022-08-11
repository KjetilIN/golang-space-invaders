package bullet

type PlayerBullet struct {
	X float64
	Y float64
	SpeedY float64
}

//Update the bullet postition - Moves the bullet up
func (b *PlayerBullet) Update() {
	b.Y -= b.SpeedY
}

//Create a new player bullet
func NewPlayerBullet(x float64, y float64, speedy float64) *PlayerBullet{
	var playerBullet = PlayerBullet{
		X: x,
		Y: y,
		SpeedY: speedy,
	}

	return &playerBullet
}

//Makes an empty list of bullets 
func NewBulletList() []PlayerBullet{
	return make([]PlayerBullet,0)
}
