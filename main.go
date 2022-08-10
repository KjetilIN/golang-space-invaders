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
	shipLevelY = 430
)


var(
	shipImg *ebiten.Image
	shipX int 
	shipY int 
)

func init (){
	var err error
	shipImg,_, err = ebitenutil.NewImageFromFile("ship.jpg")
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
    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

    // Write your game's rendering.
	screen.Fill(color.RGBA{0,0,0,0})
	screen.DrawImage(shipImg,op)
}

//Size of screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 320, 240
}

func main() {
    game := &Game{}
    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("Space Invaders 2022 - TekKom")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}