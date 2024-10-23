package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Config struct {
	width, height, target_fps int32
	title                     string
}

// --------------------------------------------------------------------------------------------------------------------

func (c *Config) init_window() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	rl.InitWindow(c.width, c.height, c.title)
	rl.SetTargetFPS(c.target_fps)
}

// --------------------------------------------------------------------------------------------------------------------

func (c *Config) resize_window() {
	if rl.IsWindowResized() {
		c.width = int32(rl.GetScreenWidth())
		c.height = int32(rl.GetScreenHeight())
	}
}

// --------------------------------------------------------------------------------------------------------------------
