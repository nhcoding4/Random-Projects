/// <reference path="enemyBase.ts" />

namespace main {

    export class FlyingLeftEnemy extends EnemyBase {

        movement = (Math.random() * 6 + 1) * -1
        angle = Math.random() * 2
        angleSpeed = Math.random() * 0.2
        curve = Math.random() * 7


        // ------------------------------------------------------------------------------------------------------------

        // Allows the spite to move off the screen. Is put back at the other edge when not visable. 

        public boundariesLogic(): void {
            if (this.x < 0 - this.width) {
                this.x = this.canvasWidth + this.width
            }

            if (this.y < 0) {
                this.y = 1
            }
            if (this.y + this.height > this.canvasHeight) {
                this.y = this.canvasHeight - this.height
            }
        }

        // ------------------------------------------------------------------------------------------------------------

        // Moves the sprite from the right to the left of the screen at a unique speed. Adds a sin wave motion. 

        public calculateMove(): void {
            this.toMoveX = this.movement
            this.toMoveY = this.curve * Math.sin(this.angle)
            this.angle += this.angleSpeed

        }

        // ------------------------------------------------------------------------------------------------------------
    }
}