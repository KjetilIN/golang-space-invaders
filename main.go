package main

import (
	"fmt"
	"image/color"
	_ "image/jpeg"
	"log"
	"sync"

	"github.com/KjetilIN/golang-space-invaders/utils"
	"github.com/KjetilIN/golang-space-invaders/bullet"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	height = 600
	width = 600
	scaleFactor = 1 //Scale the images to this factor

	//SHIP
	shipWidth = 45
	shipLevelY = height*8/9 
	speed = 9
)


var (
	score uint16 

	bl1 *bullet.PlayerBullet
	bl2 *bullet.PlayerBullet
	bl3 *bullet.PlayerBullet
	bl4 *bullet.PlayerBullet
	bl5 *bullet.PlayerBullet

	
	Ship ship = *NewShip(0,shipLevelY)

)

//var Ship ship = *NewShip(0,0)
//var shipImg *ebiten.Image



//Implement ebiten game engine
type Game struct{

	//Images
	shipImg *ebiten.Image
	blt *ebiten.Image
	downEnemyImg *ebiten.Image

	//Sync tool
	once sync.Once

}

//Method that loads an image 
func loadImage (path string) *ebiten.Image{
	var err error 
	var img *ebiten.Image
	img,_, err = ebitenutil.NewImageFromFile(path)
	if err != nil {
		fmt.Printf("ERROR DID NOT LOAD: %v\n",path)
		return nil
	}
	return img
}

func init(){
	score = 0 
	bl1 = bullet.NewPlayerBullet(-1,-1,10)
	bl2 = bullet.NewPlayerBullet(-1,-1,10)
	bl3 = bullet.NewPlayerBullet(-1,-1,10)
	bl4 = bullet.NewPlayerBullet(-1,-1,10)
	bl5 = bullet.NewPlayerBullet(-1,-1,10)


}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
    // Write your game's logical update.

	//Load all images once 
	g.once.Do(func() {
		g.shipImg = loadImage("assets/figures/ship.jpg")
		g.blt = loadImage("assets/figures/bullet.jpg")
		g.downEnemyImg = loadImage("assets/figures/pink_left.jpg")

	})

	keys := utils.GetKeyPressed()
	
	
	if(len(keys) != 0){
		keyPressed := keys[0]

		//Left
		if(keyPressed == ebiten.KeyA || keyPressed == ebiten.KeyLeft){
			if (Ship.x - speed >= 0){
				Ship.x -= speed
			}
		}

		//Right
		if(keyPressed == ebiten.KeyD || keyPressed == ebiten.KeyRight){
			if(Ship.x + speed + shipWidth <= width){
				Ship.x += speed
			}
		}

		if (inpututil.IsKeyJustPressed(ebiten.KeySpace)){
			//Set a new bullet at the front og the ship 
			bullet.ResetHighBullet(bl1,bl2,bl3,bl4,bl5,Ship.x, shipLevelY)

			//Add 1 for the score temp
			score += 1
		}

	}

	//Update all bullets 
	bullet.UpdateEachBulletGiven(bl1,bl2,bl3,bl4,bl5)

    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.RGBA{0,0,0,0})

	shipOp := &ebiten.DrawImageOptions{}
	shipOp.GeoM.Translate(Ship.x,Ship.y)
	screen.DrawImage(g.shipImg,shipOp)
	
	//Draw bullets
	bullet.DrawBullet(screen,bl1, g.blt)
	bullet.DrawBullet(screen,bl2, g.blt)
	bullet.DrawBullet(screen,bl3, g.blt)
	bullet.DrawBullet(screen,bl4, g.blt)
	bullet.DrawBullet(screen,bl5, g.blt)

	utils.DrawScore(screen,int(score),width,height)

	
}

//Size of screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return width, height
}

func main(){
    game := &Game{}
    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(width, height)
    ebiten.SetWindowTitle("Space Invaders 2022 - TekKom")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}