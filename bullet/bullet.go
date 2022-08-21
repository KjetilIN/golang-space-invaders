package bullet


import (
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerBullet struct {
	X float64
	Y float64
	SpeedY float64
	img *ebiten.Image
	op ebiten.DrawImageOptions
}

//Update the bullet postition - Moves the bullet up, if the bullet is on screen
func (b *PlayerBullet) Update() {
	if(b.Y > -40){
		b.Y -= b.SpeedY
	}
	
}

//Reset bullet 
func (b *PlayerBullet) Reset() {
	b.X = -1
	b.Y = -41
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

//Draw a bullet to the screen
func DrawBullet(screen *ebiten.Image, bl *PlayerBullet, img *ebiten.Image) {
	blOp := &ebiten.DrawImageOptions{}
	blOp.GeoM.Translate(bl.X,bl.Y)
	screen.DrawImage(img,blOp)
}

//Set the bullet with the lowest y cord, in front of the ship
func ResetHighBullet(bl1 *PlayerBullet, bl2 *PlayerBullet,bl3 *PlayerBullet,bl4 *PlayerBullet,bl5 *PlayerBullet, newX float64, newY float64 ){
	if (bl1.Y <= bl2.Y){
		bl1.X = newX
		bl1.Y = newY -40
	}else if (bl2.Y <= bl3.Y){
		bl2.X = newX
		bl2.Y = newY -40
	}else if (bl3.Y <= bl4.Y){
		bl3.X = newX
		bl3.Y = newY -40
	}else if (bl4.Y <= bl5.Y){
		bl4.X = newX
		bl4.Y = newY -40
	}else{
		bl5.X = newX
		bl5.Y = newY -40
	}
}


//Update each given bullet, using the update method
func UpdateEachBulletGiven(bl1 *PlayerBullet, bl2 *PlayerBullet,bl3 *PlayerBullet,bl4 *PlayerBullet,bl5 *PlayerBullet){
	bl1.Update()
	bl2.Update()
	bl3.Update()
	bl4.Update()
	bl5.Update()
}