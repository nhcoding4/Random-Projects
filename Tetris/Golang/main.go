package main

func main() {
	game := Game{
		cellSize:  30,
		targetFPS: 60,
		title:     "Tetris",
		rows:      20,
		columns:   10,
	}
	game.init()
	game.run()
}
