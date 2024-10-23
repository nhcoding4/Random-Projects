/// <reference path="enemyBase.ts" />

namespace main {

    export class SideToSideEnemy extends EnemyBase {
        angle = 0
        angleSpeed = Math.random() * 2
        curve = Math.random() * 200 + 50

        // ------------------------------------------------------------------------------------------------------------

        // Moves the sprites in circular movement. Change the MAth.PI / X  numbers to get different effects

        public calculateMove(): void {
            this.angle += this.angleSpeed
            this.x = (this.canvasWidth - this.width) / 2 * Math.sin(this.angle * Math.PI / 200) + (this.canvasWidth / 2 - this.width / 2)
            this.y = (this.canvasHeight - this.height) / 2 * Math.cos(this.angle * Math.PI / 250) + (this.canvasHeight / 2 - this.height / 2)
        }

        // ------------------------------------------------------------------------------------------------------------     

        // Update the elements.

        public update(): void {
            this.calculateMove()
            this.boundariesLogic()
            this.updateFlyingAnimation()
        }

        // ------------------------------------------------------------------------------------------------------------

    }
}