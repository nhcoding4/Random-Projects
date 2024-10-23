import { Effect } from "./effect"
import { Mouse } from "./mouse"

class Simulation {
    private width = window.innerWidth
    private height = window.innerHeight
    private totalParticles = 750
    private canvas: HTMLCanvasElement
    private context: CanvasRenderingContext2D | null
    private mouse: Mouse
    public effect: Effect

    constructor() {
        this.canvas = this.makeCanvas()
        this.context = this.getContext()
        this.setContextProperties()
        this.mouse = this.makeMouse()
        this.effect = this.makeEffect()

        // Resize window event listener.
        window.addEventListener("resize", _ => {
            let width = window.innerWidth
            let height = window.innerHeight

            if (Math.abs(width - this.width) != 0 || Math.abs(height - this.height) != 0) {
                this.width = width
                this.height = height
                this.canvas.width = this.width
                this.canvas.height = this.height

                this.effect.resize(this.width, this.height)
                this.setContextProperties()
            }
        })
    }

    // --------------------------------------------------------------------------------------------

    // Get canvas context.

    private getContext(): CanvasRenderingContext2D | null {
        return this.canvas.getContext("2d")
    }

    // --------------------------------------------------------------------------------------------

    // Create a new Canvas element. 

    private makeCanvas(): HTMLCanvasElement {
        let newCanvas = document.createElement("canvas")

        newCanvas.setAttribute("id", "canvas1")
        newCanvas.width = this.width
        newCanvas.height = this.height
        newCanvas.style.background = "black"
        newCanvas.style.position = "abosolute"
        newCanvas.style.left = "0"
        newCanvas.style.top = "0"

        document.body.appendChild(newCanvas)

        return newCanvas
    }

    // --------------------------------------------------------------------------------------------

    // Init the effect class.

    private makeEffect(): Effect {
        let connectionDistance = 100
        let lineWidth = 1.5

        return new Effect(
            this.context,
            this.width,
            this.height,
            this.totalParticles,
            connectionDistance,
            lineWidth,
            this.mouse
        )
    }

    // --------------------------------------------------------------------------------------------

    // Init the mouse object. 

    private makeMouse(): Mouse {
        let radius = 300

        return new Mouse(radius)
    }

    // --------------------------------------------------------------------------------------------

    // Updates and draws objects.

    public run(): void {
        if (this.context != null) {
            this.context.clearRect(0, 0, this.width, this.height)
        }

        this.effect.update()
        this.effect.draw()

        requestAnimationFrame(() => {
            this.run()
        })
    }

    // --------------------------------------------------------------------------------------------

    // Sets the context properties.

    private setContextProperties() {
        if (this.context != null) {
            let gradient = this.context.createLinearGradient(0, 0, this.width, this.height)
            gradient.addColorStop(0, "cyan")
            gradient.addColorStop(0.5, "purple")
            gradient.addColorStop(1, "orangered")

            this.context.fillStyle = gradient
            this.context.strokeStyle = "black"
        }
    }

    // --------------------------------------------------------------------------------------------
}

// ------------------------------------------------------------------------------------------------

let simulation = new Simulation()
simulation.run()