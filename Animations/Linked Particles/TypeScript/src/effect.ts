import { Particle } from "./particles"
import { Mouse } from "./mouse"

export class Effect {
    private particles: Particle[] = []

    constructor(
        private context: CanvasRenderingContext2D | null,
        private width: number,
        private height: number,
        private totalParticles: number,
        private connectionDistance: number,
        private lineWidth: number,
        private mouse: Mouse,
    ) {
        this.createParticles()        
    }

    // --------------------------------------------------------------------------------------------

    // Draws all the particles on the screen.

    public draw(): void {
        this.connectParticles()
        for (let particle of this.particles) {
            particle.draw()
        }
    }

    // --------------------------------------------------------------------------------------------

    // Create x amount of particles. 

    private createParticles(): void {
        for (let i = 0; i < this.totalParticles; i++) {
            let newParticle = new Particle(
                this.context,
                this.width,
                this.height,
                this.mouse,
            )
            this.particles.push(newParticle)
        }
    }

    // --------------------------------------------------------------------------------------------

    // Draws a line between particles that are inside X distance of one another. 

    private connectParticles(): void {
        for (let i = 0; i < this.totalParticles; i++) {
            let iParticlePositions: number[] = this.particles[i].getPosition()

            for (let j = i; j < this.totalParticles; j++) {
                if (i == j) {
                    continue
                }
                let jParticlePositions: number[] = this.particles[j].getPosition()

                let dx = iParticlePositions[0] - jParticlePositions[0]
                let dy = iParticlePositions[1] - jParticlePositions[1]
                let distance = Math.hypot(dx, dy)

                if (distance <= this.connectionDistance && this.context != null) {
                    let opacity = 1 - (distance / this.connectionDistance)

                    this.context.save()

                    this.context.beginPath()
                    this.context.moveTo(iParticlePositions[0], iParticlePositions[1])
                    this.context.lineTo(jParticlePositions[0], jParticlePositions[1])
                    this.context.lineWidth = this.lineWidth
                    this.context.globalAlpha = opacity
                    this.context.strokeStyle = "white"
                    this.context.stroke()

                    this.context.restore()
                }
            }
        }
    }

    // --------------------------------------------------------------------------------------------

    // Repopulates the width and height.

    public resize(width: number, height: number): void {
        this.width = width
        this.height = height

        for (let particle of this.particles) {
            particle.resizeCanvas(this.width, this.height)
        }
    }

    // --------------------------------------------------------------------------------------------

    // Update the position of all particles in the effect.

    public update(): void {
        for (let particle of this.particles) {
            particle.update()
        }
    }

    // --------------------------------------------------------------------------------------------
}