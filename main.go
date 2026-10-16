package main

import (

	rl "github.com/gen2brain/raylib-go/raylib"
)

var COLORS = []rl.Color{rl.Black,rl.Red,rl.Orange,rl.Yellow,rl.White}

var PARTITIO_BETWAN_COLORS = 5

var posX,posY = 10 , 200

var width , hight = 20,20

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

func drawRectanglesPaletPartinio(){
	colors := PaletPartion()


	Shift := 0

	for i := 0; i < len(colors); i++ {
		rl.DrawRectangle(int32(posX+Shift),int32(posY),int32(width),int32(hight),colors[i])
		Shift = Shift + width
	}
}

func main() {
  rl.InitWindow(800, 450, "raylib [core] example - basic window")
  defer rl.CloseWindow()

  rl.SetTargetFPS(60)


  for !rl.WindowShouldClose() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)
    drawRectanglesPaletPartinio()

    rl.EndDrawing()
  }
}