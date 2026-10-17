package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var COLORS = []rl.Color{rl.Black, rl.Red, rl.Orange, rl.Yellow, rl.White}
var FiRE_PALLET = PaletPartion()

const SCREEN_WIDTH , SCREEN_HIEGHT = 1000,600
const PARTITIO_BETWAN_COLORS =15
const PIXEL_SIZE = 5

const FIRE_WIDTH , FIRE_HIEGHT = SCREEN_WIDTH / PIXEL_SIZE , SCREEN_HIEGHT / PIXEL_SIZE

var FIRE_ARRAY [FIRE_HIEGHT][FIRE_WIDTH]int = getFireArray();




func getFireArray() [FIRE_HIEGHT][FIRE_WIDTH]int {
	fire_array := [FIRE_HIEGHT][FIRE_WIDTH]int{};
	for i := 0; i < FIRE_WIDTH; i++ {
		fire_array[FIRE_HIEGHT-1][i] = len(FiRE_PALLET)-1
	}
	return fire_array
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

func doFire(){
	for x := 0; x < FIRE_WIDTH; x++ {
		for y := 1; y < FIRE_HIEGHT; y++ {
			color_index := FIRE_ARRAY[y][x];
			if color_index != 0 {
				random := rand.Intn(2)
				FIRE_ARRAY[y-1][((x+random)%FIRE_WIDTH+FIRE_WIDTH)%FIRE_WIDTH]=color_index-random
			}
		}
	}
}



func drawFire(){
	for y, row := range FIRE_ARRAY {
		for x, color_index := range row  {
			if color_index != 0 {
				rl.DrawRectangle(int32(x*PIXEL_SIZE),int32(y*PIXEL_SIZE),PIXEL_SIZE-1,PIXEL_SIZE-1,FiRE_PALLET[color_index])
			}
		}
	}
}



func main() {
  rl.InitWindow(int32(SCREEN_WIDTH), int32(SCREEN_HIEGHT), "raylib [core] example - basic window")
  defer rl.CloseWindow()

  rl.SetTargetFPS(60)


  for !rl.WindowShouldClose() {	
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	drawFire()
	doFire()
	rl.EndDrawing()
  }
}