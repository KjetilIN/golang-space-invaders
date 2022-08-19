package utils

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var keys []ebiten.Key

func GetKeyPressed() []ebiten.Key {
	keys = inpututil.AppendPressedKeys(keys[:0])
	return keys
}