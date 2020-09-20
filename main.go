package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"

	"image/color"
	"time"
)

var snakePosX int = 270
var snakePosY int = 360
var snakePosSvX int
var snakePosSvY int

var a bool
var b bool

var dead bool

type block struct {
	x int
	y int
}

var apple = new(block)
var snake = []*block{}

func drawPixel(imd *imdraw.IMDraw, posX int, posY int, color color.RGBA) {
	imd.Color = color
	imd.Push(pixel.V(float64(posX), float64(posY)))
	imd.Push(pixel.V(float64(posX+45), float64(posY+45)))
	imd.Rectangle(0)
}

func updatePos() {
	for i := len(snake) - 1; i > 0; i-- {
		snake[i].x = snake[i-1].x
		snake[i].y = snake[i-1].y
	}
}

func newBody() {
	body := new(block)

	body.x = -45
	body.y = -45

	snake = append(snake, body)
}

func randPos() {
	rand.Seed(time.Now().UnixNano())

	apple.x = rand.Intn(16) * 45
	apple.y = rand.Intn(16) * 45
}

func makeSnake() {
	snake = []*block{}

	head := new(block)

	head.x = 270
	head.y = 360

	snake = append(snake, head)

	newBody()
	newBody()
}

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

	makeSnake()

	go pos()

	updatePos()

	for !win.Closed() {
		if dead {
			win.Clear(colornames.Darkred)
		} else {
			win.Clear(colornames.Darkgreen)
		}

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

		if win.JustPressed(pixelgl.KeyUp) {
			a = false
			b = true
		} else if win.JustPressed(pixelgl.KeyDown) {
			a = true
			b = true
		} else if win.JustPressed(pixelgl.KeyRight) {
			a = false
			b = false
		} else if win.JustPressed(pixelgl.KeyLeft) {
			a = true
			b = false
		}

		for _, v := range snake {
			drawPixel(imd, v.x, v.y, colornames.Darkcyan)
		}

		drawPixel(imd, apple.x, apple.y, colornames.Indianred)

		imd.Draw(win)
		win.Update()
	}
}

// pos stuff goes here
func pos() {
	var die int = -2

	for true {
		updatePos()

		if a == true && b == true {
			snake[0].y -= 45
		} else if a == false && b == false {
			snake[0].x += 45
		} else if a == false && b == true {
			snake[0].y += 45
		} else {
			snake[0].x -= 45
		}

		for i, v := range snake {
			for s, z := range snake {
				if z.x == v.x && z.y == v.y && i != s {
					die++
				}
			}
		}

		if dead {
			dead = false
		}

		if die > 1 {
			dead = true

			makeSnake()

			updatePos()
			updatePos()

			a = false
			b = false

			die = -2
		}

		if snake[0].x == apple.x && snake[0].y == apple.y {
			newBody()
			randPos()
		}

		time.Sleep(1000000000)
	}
}

func main() {
	pixelgl.Run(run)
}
