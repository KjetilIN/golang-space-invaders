package main

import (
	_ "image/jpeg"
)

type ship struct{
	x float64
	y float64
}

//Creates and returnes a new ship struct
func NewShip(x float64, y float64) *ship {

	var ship = ship{
		x: x,
		y:y,
	}

	return &ship

}