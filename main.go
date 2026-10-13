package main

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var COLORS = []rl.Color{rl.Black,rl.Red,rl.Orange,rl.Yellow,rl.White}

var PARTITIO_BETWAN_COLORS = 5


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

func drawingPalet() []rl.Color {
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
  rl.InitWindow(800, 450, "raylib [core] example - basic window")
  defer rl.CloseWindow()

  rl.SetTargetFPS(60)

  for !rl.WindowShouldClose() {
    rl.BeginDrawing()
	fmt.Println("rendring")
    rl.ClearBackground(rl.RayWhite)
    rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

    rl.EndDrawing()
  }
}