package enemy

import(
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/KjetilIN/golang-space-invaders/bullet"
)

const(
	bulletWidht = 34
	bulletHeight = 45
)

//Downenemy struct 
type DownEnemy struct{
	y int 
	x int 
	speed int 
}


func NewDownEnemy(x int, y int, speed int) *DownEnemy{
	return &DownEnemy{x: x, y:y, speed: speed}
}

//Draw a enemy to the screen
func Draw(screen *ebiten.Image, enemy *DownEnemy, img *ebiten.Image) {
	enemyOp := &ebiten.DrawImageOptions{}
	enemyOp.GeoM.Translate(float64(enemy.x),float64(enemy.y))
	screen.DrawImage(img,enemyOp)
}

//Check if any bullet hits the cuttent enemy 
func (d *DownEnemy) isHit(bl1 bullet.PlayerBullet, bl2 bullet.PlayerBullet, bl3 bullet.PlayerBullet, bl4 bullet.PlayerBullet, bl5 bullet.PlayerBullet)bool{
	isAnyBulletOnSameX := isBulletOnX(&bl1,d.x) || isBulletOnX(&bl2,d.x) || isBulletOnX(&bl3,d.x) || isBulletOnX(&bl4,d.x) || isBulletOnX(&bl5,d.x)
	isAnyBulletOnSameY := isBulletOnY(&bl1,d.y) || isBulletOnY(&bl2,d.y) || isBulletOnY(&bl3,d.y) || isBulletOnY(&bl4,d.y) || isBulletOnY(&bl5,d.y)

	if(isAnyBulletOnSameX){
		println("IS ON SAME X")
	}

	if(isAnyBulletOnSameY){
		println("IS ON SAME Y")
	}



	if(isAnyBulletOnSameX && isAnyBulletOnSameY){
		return true
	}
	return false
}

//Check if the bullet is within the same x box 
func isBulletOnX(bl *bullet.PlayerBullet, x int)bool{
	return bl.X <= float64(x) && bl.X + bulletWidht >= float64(x) 
}

//Check if the bullet is within the same y box 
func isBulletOnY(bl *bullet.PlayerBullet, Y int)bool{
	return bl.Y <= float64(Y) && bl.Y + bulletHeight >= float64(Y) 
}

//Update all enemies
func UpdateAll(enemy1 *DownEnemy, enemy2 *DownEnemy, enemy3 *DownEnemy, 
	bl1 *bullet.PlayerBullet, bl2 *bullet.PlayerBullet,
	bl3 *bullet.PlayerBullet,bl4 *bullet.PlayerBullet,bl5 *bullet.PlayerBullet, screenHeight int){

	enemy1.updateEnemy(bl1,bl2,bl3,bl4,bl5, screenHeight)
	enemy1.updateEnemy(bl1,bl2,bl3,bl4,bl5, screenHeight)
	enemy1.updateEnemy(bl1,bl2,bl3,bl4,bl5, screenHeight)
}

//Update a single enemy 
func (e *DownEnemy)updateEnemy(bl1 *bullet.PlayerBullet, bl2 *bullet.PlayerBullet,
	bl3 *bullet.PlayerBullet,bl4 *bullet.PlayerBullet,bl5 *bullet.PlayerBullet , height int){

	if(e.isHit(*bl1,*bl2,*bl3,*bl4,*bl5)){
		//Reset the postition of the enemy
		e.x = 30
		e.y = -30
	}else{
		//Move enemy down
		if(e.y < height){
			e.y += e.speed
		}
		
		println(e.y)
	}
}