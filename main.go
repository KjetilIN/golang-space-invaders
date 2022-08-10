package main

import (
	"fmt"
	"image/color"
	"log"

	_ "image/jpeg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	height = 600
	width = 600

	//SHIP
	shipLevelY = height*8/9 
	speed = 10
)


var(
	//Ship
	shipImg *ebiten.Image
	shipX int 
)

func init (){

	//Load the image 
	var err error
	shipImg,_, err = ebitenutil.NewImageFromFile("assets/figures/ship.jpg")
	if err != nil {
		fmt.Println("Error loading the ship")
	}

}

//Implement ebiten game engine
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
    // Write your game's logical update.
	if(shipX + 52 > width ){
		shipX = 0
	}
	shipX += speed
    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {

	shipOp := &ebiten.DrawImageOptions{}

	//Resize the ship
	shipOp.GeoM.Scale(0.5, 0.5)
	shipOp.GeoM.Translate(float64(shipX), shipLevelY)

    // Write your game's rendering.
	screen.Fill(color.RGBA{0,0,0,0})
	
	screen.DrawImage(shipImg,shipOp)
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