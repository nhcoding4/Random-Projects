package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	width               int32
	height              int32
	title               string
	targetFPS           int32
	backgroundColor     *rl.Color
	effect              *Effect
	mouse               *Mouse
	sunRay              *SunRay
	mouseRadius         int32
	totalParticles      int32
	powerPushMultiplier float64
	linkDistance        int32
	linkThickness       float32
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) init() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(g.width, g.height, g.title)
	rl.SetTargetFPS(g.targetFPS)
	g.createMouse()
	g.createEffect()
	g.createSunRay()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) createEffect() {
	g.effect = &Effect{
		width:               g.width,
		height:              g.height,
		totalParticles:      g.totalParticles,
		mouse:               g.mouse,
		powerPushMultiplier: g.powerPushMultiplier,
		linkDistance:        g.linkDistance,
		linkThickness:       g.linkThickness,
	}
	g.effect.init()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) createMouse() {
	g.mouse = &Mouse{
		x:      g.width / 2,
		y:      g.height / 2,
		radius: g.mouseRadius,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) createSunRay() {
	g.sunRay = &SunRay{
		mouse:     g.mouse,
		effect:    g.effect,
		divisor:   3,
		thickness: 1.5,
		opacity:   0.4,
	}
	g.sunRay.init()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) displayFps() {
	fps := rl.GetFPS()
	rl.DrawText(fmt.Sprintf("%v", fps), 0, 0, 40, rl.Green)
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) run() {
	for !rl.WindowShouldClose() {
		g.updateState()

		rl.BeginDrawing()
		rl.ClearBackground(*g.backgroundColor)
		g.sunRay.draw()
		g.effect.draw()
		g.displayFps()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) updateState() {
	if rl.IsWindowResized() {
		g.updateWindowSize()
		g.effect.setPosition()
	}

	g.mouse.checkLeftClickDown()
	if g.mouse.activated {
		g.mouse.update()
	}

	g.effect.update()
	g.sunRay.calculateLocations()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) updateWindowSize() {
	g.width = int32(rl.GetScreenWidth())
	g.height = int32(rl.GetScreenHeight())
	g.effect.updateWindowSize()
}

// ---------------------------------------------------------------------------------------------------------------------
