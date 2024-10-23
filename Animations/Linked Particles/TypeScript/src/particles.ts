import { Mouse } from "./mouse"

export class Particle {
    private x!: number
    private y!: number
    private radius = Math.floor((Math.random() * 15) + 1)
    private dx = (Math.random() * 2) - 1
    private dy = (Math.random() * 2) - 1
    private pushForceX = 0
    private pushForceY = 0
    private friction: number

    constructor(
        private context: CanvasRenderingContext2D | null,
        private width: number,
        private height: number,
        private mouse: Mouse,
    ) {
        this.setPosition()
        this.friction = this.calculateFriction()
    }

    // --------------------------------------------------------------------------------------------

    // Keeps particles within the canvas. 

    private bounds(): void {
        if (this.x < 0 + this.radius) {
            this.x = this.radius
            this.dx *= -1
        }
        if (this.x > this.width - this.radius) {
            this.x = this.width - this.radius
            this.dx *= -1
        }

        if (this.y < 0 + this.radius) {
            this.y = this.radius
            this.dy *= -1
        }
        if (this.y > this.height - this.radius) {
            this.y = this.height - this.radius
            this.dy *= - 1
        }
    }

    // --------------------------------------------------------------------------------------------

    // Set a friction value base on size of particle.

    private calculateFriction(): number {
        if (this.radius <= 7) {
            return 0.99
        } else if (this.radius <= 10) {
            return 0.975
        } else {
            return 0.95
        }
    }

    // --------------------------------------------------------------------------------------------

    // Draw the individual particle on the screen.

    public draw(): void {
        if (this.context != null) {
            this.context.beginPath()
            this.context.arc(this.x, this.y, this.radius, 0, 2 * Math.PI, false)
            //this.context.fillStyle = `hsl(${this.x}, 100%, 50%)`
            this.context.fill()
            this.context.stroke()
        }
    }

    // --------------------------------------------------------------------------------------------

    // Returns the X and Y values of the particle.

    public getPosition(): number[] {
        return [this.x, this.y]
    }

    // --------------------------------------------------------------------------------------------

    // Returns the radius of the particle.

    public getRadius(): number {
        return this.radius
    }

    // --------------------------------------------------------------------------------------------

    // Changers the particle coordinates. 

    private move(): void {
        this.x += this.dx + this.pushForceX
        this.y += this.dy + this.pushForceY
    }

    // --------------------------------------------------------------------------------------------

    private pushWithMouse(): void {
        if (this.mouse.getPressedStatus()) {
            let mousePositions = this.mouse.getPosition()
            let dx = this.x - mousePositions[0]
            let dy = this.y - mousePositions[1]
            let distance = Math.hypot(dx, dy)

            if (distance < this.mouse.getRadius()) {
                let force = this.mouse.getRadius() / distance
                let angle = Math.atan2(dy, dx)
                this.pushForceX = Math.cos(angle) * force
                this.pushForceY = Math.sin(angle) * force
            }
        }
    }

    // --------------------------------------------------------------------------------------------

    // Changes width and height variables and resets the particle postion.

    public resizeCanvas(width: number, height: number): void {
        this.width = width
        this.height = height
        this.setPosition()
    }


    // --------------------------------------------------------------------------------------------

    // Set the x and y values of the particles inside the limits of the canvas.

    private setPosition(): void {
        this.x = (Math.random() * this.width - (this.radius * 2)) + this.radius * 2
        this.y = (Math.random() * this.height - (this.radius * 2)) + this.radius * 2
    }

    // --------------------------------------------------------------------------------------------

    // Update particle state.

    public update(): void {
        this.pushWithMouse()
        this.move()
        this.bounds()
        this.updateForce()
    }

    // --------------------------------------------------------------------------------------------

    // Adds friction to the force applied to particles.

    private updateForce(): void {
        this.pushForceX *= this.friction
        this.pushForceY *= this.friction
    }

    // --------------------------------------------------------------------------------------------

}