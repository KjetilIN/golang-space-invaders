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

func (enemy *DownEnemy) reset(){
	enemy.x = 0
	enemy.y = -1 -bulletHeight
}


//Check if the enemy is hit by the given bullet
func (enemy *DownEnemy)checkAndUpdateBulletHit(bl *bullet.PlayerBullet, height int) bool {
	hitOnXLevel := bl.X >= float64(enemy.x) && bl.X + bulletHeight <= float64(enemy.x) 
	hitOnYLevel := bl.Y >= float64(enemy.y) && bl.Y + bulletHeight <= float64(enemy.y) 

	hasBulletInHitbox := hitOnXLevel && hitOnYLevel

	if(hasBulletInHitbox){
		//Reset the bullet and enemy 
		bl.Reset() 
		enemy.reset()
		return true
	}else{
		//Move enemy down as long as it is within the screen
		if(enemy.y > height){
			enemy.y += enemy.speed
		}else{
			//Reset the enemy or end the game
			enemy.reset()
		}
		
	}
	return false
}

func (enemy *DownEnemy)checkAllBullets(bl1 *bullet.PlayerBullet, bl2 *bullet.PlayerBullet,
	bl3 *bullet.PlayerBullet,bl4 *bullet.PlayerBullet,bl5 *bullet.PlayerBullet, screenHeight int){

	enemy.checkAndUpdateBulletHit(bl1,screenHeight)
	enemy.checkAndUpdateBulletHit(bl2,screenHeight)
	enemy.checkAndUpdateBulletHit(bl3,screenHeight)
	enemy.checkAndUpdateBulletHit(bl4,screenHeight)
	enemy.checkAndUpdateBulletHit(bl5,screenHeight)

}

//Update function
func Update(enemy1 *DownEnemy, enemy2 *DownEnemy, enemy3 *DownEnemy, 
	bl1 *bullet.PlayerBullet, bl2 *bullet.PlayerBullet,
	bl3 *bullet.PlayerBullet,bl4 *bullet.PlayerBullet,bl5 *bullet.PlayerBullet, screenHeight int){
	
	enemy1.checkAllBullets(bl1,bl2,bl3,bl4,bl5,screenHeight)
	enemy2.checkAllBullets(bl1,bl2,bl3,bl4,bl5,screenHeight)
	enemy3.checkAllBullets(bl1,bl2,bl3,bl4,bl5,screenHeight)
}
