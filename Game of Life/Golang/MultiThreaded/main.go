package main

/*
Multithreaded vs Single threaded benchmarks: 10600k | Rtx 3070.
	Multi 2px size = 17fps | Single 8fps | 112% speed increase.
	Multi 5px size = 105fps | Single 55fps | 90% speed increase.
	7px size = 210fps | Single 105 fps | 100% speed increase.
	10px size = 430fps | single 200fps | 115% speed increase.
*/

import "math"

func main() {
	var cellSize int32 = 2
	simulation := Simulation{
		rows:      int32(math.Floor(1080 / float64(cellSize))),
		columns:   int32(math.Floor(1920 / float64(cellSize))),
		cellSize:  cellSize,
		targetFPS: 1000,
		title:     "Game of Life Concurrent.",
		fontSize:  40,
	}
	simulation.Init()
	simulation.run()
}
