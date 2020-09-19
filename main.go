package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

	"fmt"
	"time"
)

var posX int = 100
var posY int = 100

var a bool
var b bool

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Snake",
		Bounds: pixel.R(0, 0, 720, 720),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	go pos()

	for !win.Closed() {
		win.Clear(colornames.Darkgreen)
		imd.Clear()

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			mousePos := win.MousePosition()

			if mousePos.X < 360 {
				if mousePos.Y < 360 {
					if mousePos.X < mousePos.Y {
						a = true
						b = false
					} else {
						a = true
						b = true
					}
				} else {
					if mousePos.X+mousePos.Y < 720 {
						a = true
						b = false
					} else {
						a = false
						b = true
					}
				}
			} else {
				if mousePos.Y < 360 {
					if mousePos.X+mousePos.Y < 720 {
						a = true
						b = true
					} else {
						a = false
						b = false
					}
				} else {
					if mousePos.X < mousePos.Y {
						a = false
						b = true
					} else {
						a = false
						b = false
					}
				}
			}
		}

		imd.Color = pixel.RGB(0, 1, 1)
		imd.Push(pixel.V(float64(posX), float64(posY)))
		imd.Push(pixel.V(float64(posX+45), float64(posY+45)))
		imd.Rectangle(0)

		imd.Draw(win)

		win.Update()
	}
}

func pos() {
	for true {
		if a == true && b == true {
			posY -= 45
		} else if a == false && b == false {
			posX += 45
		} else if a == false && b == true {
			posY += 45
		} else {
			posX -= 45
		}

		fmt.Printf("x: %d, y: %d, %d%d\n", posX, posY, a, b)
		time.Sleep(1000000000)
	}
}

func main() {
	pixelgl.Run(run)
}
