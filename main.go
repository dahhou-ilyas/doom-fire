package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var COLORS = []rl.Color{rl.White, rl.Yellow, rl.Orange, rl.Red, rl.Black}


var SCREEN_WIDTH , SCREEN_HIEGHT = 1000,600

var PARTITIO_BETWAN_COLORS =20

var posX , posY= 0 , 580

var PIXEL_SIZE = 5


func drawFire(palette []rl.Color , height float32){
	colums := SCREEN_WIDTH / PIXEL_SIZE

	nbRectangls := int(height) / PIXEL_SIZE

	for j := 0; j < colums; j++ {
		for i := 0; i < nbRectangls; i++ {	
			t := float32(i) / float32(nbRectangls)
			idx := int(t * float32(len(palette)-1))

			fmt.Println(idx)
			x := posX + j*PIXEL_SIZE
			y := posY - (i+1)*PIXEL_SIZE

			if idx>40 {
				rl.DrawRectangle(int32(x), int32(y), int32(PIXEL_SIZE-1), int32(PIXEL_SIZE-1), palette[idx-rand.Intn(10)])
			}else{
				rl.DrawRectangle(int32(x), int32(y), int32(PIXEL_SIZE-1), int32(PIXEL_SIZE-1), palette[idx])

			}
		}
	}
}

func ColorLerp(start rl.Color, end rl.Color , amount float32) rl.Color{
	if amount > 1 { amount = 1}
	if amount < 0 { amount = 0 }

	return rl.Color{
		R : uint8(rl.Lerp(float32(start.R),float32(end.R),amount)),
		G : uint8(rl.Lerp(float32(start.G),float32(end.G),amount)),
		B: uint8(rl.Lerp(float32(start.B),float32(end.B),amount)),
		A: uint8(rl.Lerp(float32(start.A),float32(end.A),amount)),
	}
}

func PaletPartion() []rl.Color {
	var partitionColor []rl.Color

	for i := 0; i < len(COLORS)-1; i++ {
		c0 := COLORS[i]
		c1 := COLORS[i+1]

		for step := 0; step < PARTITIO_BETWAN_COLORS; step++ {
			currentColor := ColorLerp(c0,c1,(float32(step)+0.5)/float32(PARTITIO_BETWAN_COLORS))

			partitionColor = append(partitionColor, currentColor)
		}
	}

	return partitionColor
}


func main() {
  rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HIEGHT), "raylib [core] example - basic window")
  defer rl.CloseWindow()

  rl.SetTargetFPS(60)


  palette := PaletPartion()

  height := float32(200)
  target := height  


  for !rl.WindowShouldClose() {

	dt := rl.GetFrameTime()

	diff := target - height

	if diff < 2 && diff > -2 {
		target = float32(rand.Intn(300) + 200)
	}

	height = rl.Lerp(height, target, 3*dt)
    
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	drawFire(palette, height)
	PaletPartion()
	rl.EndDrawing()
  }
}