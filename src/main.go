package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	txt = ""
)

func update() {
	pos := rl.GetMousePosition()
	txt = fmt.Sprintf("%d : %d", int32(pos.X), int32(pos.Y))
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	rl.DrawText(txt, 20, 20, 20, rl.LightGray)

	rl.EndDrawing()
}

func main() {
	rl.InitWindow(450, 450, "raylib [core] example - basic window")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		render()
		update()
	}

	rl.CloseWindow()
}
