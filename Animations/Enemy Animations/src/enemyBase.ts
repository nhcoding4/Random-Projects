namespace main {

    // Base class for every enemy to inherit from.

    export class EnemyBase {
        public lastUpdate: number
        public updateDelta: number
        public toMoveX: number
        public toMoveY: number
        public currentFrame: number

        constructor(
            public x: number,
            public y: number,
            public context: CanvasRenderingContext2D | null,
            public canvasWidth: number,
            public canvasHeight: number,
            public image: HTMLImageElement,
            public spriteWidth: number,
            public spriteHeight: number,
            public totalSpriteFrames: number,
            public width: number,
            public height: number,
        ) {

            this.currentFrame = 0
            this.lastUpdate = new Date().getTime()
            this.updateDelta = (Math.random() * 60 + 10) - 10
            this.toMoveX = 0
            this.toMoveY = 0
        }

        // ------------------------------------------------------------------------------------------------------------

        // Animate the sprite and calls the function that prepares movement for the next animation frame.

        public animate(): void {
            this.context?.drawImage(
                this.image,
                this.currentFrame * this.spriteWidth,
                0,
                this.spriteWidth,
                this.spriteHeight,
                this.x,
                this.y,
                this.width,
                this.height
            )
            this.update()
        }

        // ------------------------------------------------------------------------------------------------------------

        // Responsible for dealing with what happens to sprites as they attempt to move out of the canvas view.

        public boundariesLogic(): void {
            if (this.x < 0) {
                this.x = 1
            }
            if (this.x + this.width > this.canvasWidth) {
                this.x = this.canvasWidth - this.width
            }

            if (this.y < 0) {
                this.y = 1
            }
            if (this.y + this.height > this.canvasHeight) {
                this.y = this.canvasHeight - this.height
            }
        }

        // ------------------------------------------------------------------------------------------------------------

        /*
            Each sprite has a specific movement pattern. This function should control the x and y movement. As a result
            large variation is expected between specific sprites.
        */
        public calculateMove(): void { }

        // ------------------------------------------------------------------------------------------------------------

        /* 
            Update the animation sprite movement for the next frame . Should tie together x and y movement and 
            bounds logic. Should be overloaded to fit the needs of the specific animation. 
        */

        public update(): void {
            this.calculateMove()
            this.x += this.toMoveX
            this.y += this.toMoveY
            this.boundariesLogic()

            this.updateFlyingAnimation()
        }

        // ------------------------------------------------------------------------------------------------------------

        // Should be called by update. Updates the flying animation of the sprite based upon a randomized time delta.

        public updateFlyingAnimation(): void {
            let checkDelta = () => {
                let now = new Date().getTime()
                if (now - this.lastUpdate > this.updateDelta) {
                    return true
                }
                return false
            }

            if (checkDelta()) {
                this.currentFrame++
                this.lastUpdate = new Date().getTime()
            }
            if (this.currentFrame > this.totalSpriteFrames) {
                this.currentFrame = 0
            }
        }

        // ------------------------------------------------------------------------------------------------------------

    }
}