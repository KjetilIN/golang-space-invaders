package utils

import (
	"strconv"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	numberWidth = 18
	scoreLabelWith = 80
)

//Method that loads an image
func loadImage (path string) *ebiten.Image{
	var err error 
	var img *ebiten.Image
	img,_, err = ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil
	}
	return img
}

func DrawScoreLabel(screen *ebiten.Image,x int, y int){	
	scoreImg := loadImage("assets/score/scoretxt.jpg")

	op:= &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x),float64(y))
	screen.DrawImage(scoreImg,op)
}

func DrawNumber(x int, y int, img *ebiten.Image, screen *ebiten.Image){	
	op:= &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x),float64(y))
	screen.DrawImage(img,op)
}

//Get the image for the number 
func GetNumberImage(number string) *ebiten.Image{
	return loadImage("assets/numbers/"+ number + ".jpg")
}


func DrawScore(screen *ebiten.Image, score int, screenWitdh int, screenHeight int){

	numberString := getScoreString(score)

	currentXValue := screenWitdh - numberWidth

	//Reverse the loop
	for i:= len(numberString)-1 ; i >= 0 ; i--{
		DrawNumber(currentXValue,0,GetNumberImage(string([]rune(numberString)[i])),screen)
		currentXValue -= numberWidth
	}

	DrawScoreLabel(screen,currentXValue-scoreLabelWith,0)

	
}


//Return a 6 char long string that represent the score 
func getScoreString(score int ) string{
	numberString := strconv.Itoa(score)
	dif := 6 - len(numberString)

	if dif <= 0{
		dif = 0 
	}

	if(len(numberString) > 6){
		charsToRemove := len(numberString)-6
		numberString = numberString[:len(numberString)- charsToRemove]
	}else{
		for i:= 0 ; i < dif; i++{
			numberString = "0" + numberString
		}
	}

	return numberString
}