/// <reference path="grid.ts" />

namespace main {
    export class Simulation {
        button: HTMLButtonElement
        canvas: HTMLCanvasElement
        context: CanvasRenderingContext2D | null
        canvasHeight: number
        canvasWidth: number
        cellSize: number
        grid: Grid
        fps: number
        interval: number
        now: number
        lastUpdate: number

        constructor() {
            this.canvasWidth = 1900
            this.canvasHeight = 1000
            this.cellSize = 5
            this.fps = 60
            this.interval = 1000 / this.fps
            this.createCanvas()
            this.createGrid()
            this.lastUpdate = Date.now()
        }

        // ------------------------------------------------

        // Create a drawing surface.

        createCanvas(): HTMLCanvasElement {
            let canvas = document.createElement("canvas")
            canvas.setAttribute("id", "canvas")

            canvas.height = this.canvasHeight
            canvas.width = this.canvasWidth

            canvas.style.border = "3px black"
            canvas.style.position = "absolute"
            canvas.style.top = "50%"
            canvas.style.left = "50%"
            canvas.style.transform = "translate(-50%, -50%)"
            canvas.style.height = `${this.canvasHeight}px`
            canvas.style.width = ` ${this.canvasWidth}px`

            this.context = canvas.getContext("2d")

            document.body.appendChild(canvas)

            return this.canvas = canvas
        }

        // ------------------------------------------------

        // Creates an instance of this grid object.

        createGrid(): Grid {
            let totalCellsX = Math.floor(this.canvasWidth / this.cellSize)
            let totalCellsY = Math.floor(this.canvasHeight / this.cellSize)
            let grid = new Grid(this.context, totalCellsY, totalCellsX, this.cellSize)

            return this.grid = grid
        }

        // ------------------------------------------------

        // Update and redraw the simulation every tick. 

        draw() {
            this.now = Date.now()

            if (this.now - this.lastUpdate > this.interval) {
                this.context?.clearRect(0, 0, this.canvasWidth, this.canvasHeight)
                this.grid.draw()
                this.grid.update()
                this.lastUpdate = Date.now()
            }

            this.button.addEventListener("click", () => {
                this.grid.populateGrid()
            })

            requestAnimationFrame(() => {
                this.draw()
            })
        }

        // ------------------------------------------------

        // A reset button. Calls resets the grid state.

        resetButton(): HTMLButtonElement {
            let newButton = document.createElement("button")

            newButton.innerText = "Reset"

            this.button = newButton
            document.body.appendChild(newButton)

            return newButton
        }

        // ------------------------------------------------
    }
}