package main

import (
	"fmt"
	"image/color"
	"log"
	"sync"
	_ "image/jpeg"
	"github.com/KjetilIN/golang-space-invaders/controlls"
	"github.com/KjetilIN/golang-space-invaders/bullet"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	Ship ship = *NewShip(0,shipLevelY)
	plBullets []bullet.PlayerBullet = bullet.NewBulletList()

)

//var Ship ship = *NewShip(0,0)
//var shipImg *ebiten.Image



//Implement ebiten game engine
type Game struct{

	shipImg *ebiten.Image
	blt *ebiten.Image

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

//Function for drawing asset 
func DrawAsset(givenX int, givenY int, img *ebiten.Image, op *ebiten.DrawImageOptions, screen *ebiten.Image){
	if (op == nil){
		op = &ebiten.DrawImageOptions{}
	}
	op.GeoM.Scale(scaleFactor,scaleFactor)
	op.GeoM.Translate(float64(givenX),float64(givenY))
	screen.DrawImage(img, op)

}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
    // Write your game's logical update.

	//Load all images once 
	g.once.Do(func() {
		g.shipImg = loadImage("assets/figures/ship.jpg")
		g.blt = loadImage("assets/figures/bullet.jpg")

	})


	keys := controlls.GetKeyPressed()
	
	
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

		if (keyPressed == ebiten.KeySpace){
			if (len(plBullets)<1){
				var newBlt = bullet.NewPlayerBullet(Ship.x,shipLevelY-40,10)
				plBullets = append(plBullets, *newBlt)
			}else{
				plBullets = bullet.NewBulletList()
			}
		}

	}


	//Update bullet movement 
	for _, bullet := range plBullets{
		bullet.Y -= speed
	}

    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	screen.Fill(color.RGBA{0,0,0,0})

	//Draw ship
	DrawAsset(int(Ship.x),int(Ship.y),g.shipImg,nil,screen)

	//Draw each bullet

	for _,bullet := range plBullets{
		DrawAsset(int(bullet.X),int(bullet.Y),g.blt,nil,screen)
	}
	

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