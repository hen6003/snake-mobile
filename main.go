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

var dirA bool
var dirB bool

var oldDirA bool
var oldDirB bool

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
	for same := true; same; {
		rand.Seed(time.Now().UnixNano())

		apple.x = rand.Intn(16) * 45
		apple.y = rand.Intn(16) * 45

		for _, v := range snake {
			if v.x == apple.x || v.y == apple.y {
				same = true
				return
			} else {
				same = false
			}
		}
	}
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

func setDirUp() {
	if oldDirA == true && oldDirB == true {
		return
	}
	dirA = false
	dirB = true
}

func setDirDown() {
	if oldDirA == false && oldDirB == true {
		return
	}
	dirA = true
	dirB = true
}

func setDirRight() {
	if oldDirA == true && oldDirB == false {
		return
	}
	dirA = false
	dirB = false
}

func setDirLeft() {
	if oldDirA == false && oldDirB == false {
		return
	}
	dirA = true
	dirB = false
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Snake Mobile",
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
	randPos()

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
						setDirLeft()
					} else {
						setDirDown()
					}
				} else {
					if mousePos.X+mousePos.Y < 720 {
						setDirLeft()
					} else {
						setDirUp()
					}
				}
			} else {
				if mousePos.Y < 360 {
					if mousePos.X+mousePos.Y < 720 {
						setDirDown()
					} else {
						setDirRight()
					}
				} else {
					if mousePos.X < mousePos.Y {
						setDirUp()
					} else {
						setDirRight()
					}
				}
			}
		}

		if win.JustPressed(pixelgl.KeyUp) {
			setDirUp()
		} else if win.JustPressed(pixelgl.KeyDown) {
			setDirDown()
		} else if win.JustPressed(pixelgl.KeyRight) {
			setDirRight()
		} else if win.JustPressed(pixelgl.KeyLeft) {
			setDirLeft()
		}

		for i, v := range snake {
			col := colornames.Darkcyan

			col.G = col.G - uint8(i*2)
			col.B = col.B - uint8(i*2)

			drawPixel(imd, v.x, v.y, col)
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

		oldDirA = dirA
		oldDirB = dirB

		if dirA == true && dirB == true {
			snake[0].y -= 45
		} else if dirA == false && dirB == false {
			snake[0].x += 45
		} else if dirA == false && dirB == true {
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

		if snake[0].y >= 720 || snake[0].y < 0 || snake[0].x >= 720 || snake[0].x < 0 {
			die++
			die++
		}

		if dead {
			dead = false
		}

		if die > 1 {
			dead = true

			makeSnake()

			updatePos()
			updatePos()

			dirA = false
			dirB = false

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
