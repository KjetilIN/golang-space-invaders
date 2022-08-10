package main

import (
	"fmt"
	"image/color"
	"log"
	_ "image/jpeg"
	"github.com/KjetilIN/golang-space-invaders/controlls"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	height = 600
	width = 600
	scaleFactor = 0.5 //Scale the images to this factor

	//SHIP
	shipLevelY = height*8/9 
	speed = 9
)

var Ship ship = *NewShip(0,0)
var shipImg *ebiten.Image



//Implement ebiten game engine
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
    // Write your game's logical update.

	keys := controlls.GetKeyPressed()
	
	if(len(keys) != 0){
		keyPressed := keys[0]

		//Left
		if(keyPressed == ebiten.KeyA){
			Ship.x -= speed
		}

		//Right
		if(keyPressed == ebiten.KeyD){
			Ship.x += speed
		}

	}

    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	//Load the image
	var err error 
	shipImg,_, err = ebitenutil.NewImageFromFile("assets/figures/ship.jpg")
	if err != nil {
		fmt.Println("ERROR NO SHIP")
	}

	// Write your game's rendering.
	screen.Fill(color.RGBA{0,0,0,0})
	op.GeoM.Translate(Ship.x, shipLevelY)
	screen.DrawImage(shipImg,op)

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