/// <reference path="enemyBase.ts" />

namespace main {

    export class CircularEnemy extends EnemyBase {

        recalculateGameTicks = (Math.random() * 120) + 60
        passedTicks = 0

        // ------------------------------------------------------------------------------------------------------------

        // Sets a new sprite position every 60 - 120 game ticks. 

        public calculateMove(): void {
            if (this.passedTicks >= this.recalculateGameTicks) {
                let newX = Math.floor(Math.random() * (this.canvasWidth - this.width))
                let newY = Math.floor(Math.random() * (this.canvasHeight - this.height))

                this.toMoveX = this.x - newX
                this.toMoveY = this.y - newY
                this.passedTicks = 0
            }
            this.passedTicks++
        }

        // ------------------------------------------------------------------------------------------------------------

        // Updates the position of the sprite a small percantage of the distance each game tick.

        public update(): void {

            this.calculateMove()
            this.x -= this.toMoveX / 100
            this.y -= this.toMoveY / 100
            this.boundariesLogic()
            this.updateFlyingAnimation()
        }

        // ------------------------------------------------------------------------------------------------------------
    }

}