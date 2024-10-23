/// <reference path="enemyManager.ts" />

namespace main {

    export class Game {
        private canvas: HTMLCanvasElement
        private context: CanvasRenderingContext2D | null
        private canvasHeight: number
        private canvasWidth: number
        private targetFPS: number
        private targetUpdateDelta: number
        private lastUpdate: number
        private enemyManager: EnemyManager

        constructor() {
            this.canvasHeight = 800
            this.canvasWidth = 1500
            this.targetFPS = 60
            this.targetUpdateDelta = 1000 / this.targetFPS
            this.createCanvas()
            this.createContext()
            this.lastUpdate = new Date().getTime()
            this.enemyManager = new EnemyManager(this.canvasWidth, this.canvasHeight, this.context)
        }

        // ------------------------------------------------------------------------------------------------------------

        // Animates all the elements

        public animate() {
            if (this.updateFrame()) {
                this.context?.clearRect(0, 0, this.canvasWidth, this.canvasHeight)
                this.enemyManager.animate()
                this.lastUpdate = new Date().getTime()
            }

            requestAnimationFrame(() => {
                this.animate()
            })
        }

        // ------------------------------------------------------------------------------------------------------------

        // Create a new canvas element to draw on.

        private createCanvas(): HTMLCanvasElement {
            let newCanvas = document.createElement("canvas")

            newCanvas.width = this.canvasWidth
            newCanvas.height = this.canvasHeight

            newCanvas.style.border = "3px solid"
            newCanvas.style.position = "absolute"
            newCanvas.style.top = "50%"
            newCanvas.style.left = "50%"
            newCanvas.style.transform = "translate(-50%, -50%)"
            newCanvas.style.height = `${this.canvasHeight}px`
            newCanvas.style.width = `${this.canvasWidth}px`

            document.body.appendChild(newCanvas)

            return this.canvas = newCanvas
        }

        // ------------------------------------------------------------------------------------------------------------

        // Get the canvas context.

        private createContext(): CanvasRenderingContext2D | null {
            return this.context = this.canvas.getContext("2d")
        }

        // ------------------------------------------------------------------------------------------------------------

        // Updates calculates if enough time has passed between updates. 

        private updateFrame(): boolean {
            let now = new Date().getTime()
            let timeDelta = Math.abs(now - this.lastUpdate)

            if (timeDelta > this.targetUpdateDelta) {
                return true
            }
            return false
        }

        // ------------------------------------------------------------------------------------------------------------
    }
}

