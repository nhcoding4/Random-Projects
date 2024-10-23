/// <reference path="shakingEnemy.ts" />
/// <reference path="flyingLeftEnemy.ts" />
/// <reference path="sideToSideEnemy.ts" />
/// <reference path="circularEnemy.ts" />

namespace main {

    export class EnemyManager {
        images: HTMLImageElement[]
        shakingEnemy: ShakingEnemy[]
        flyingLeftEnemy: FlyingLeftEnemy[]
        sideToSideEnemy: SideToSideEnemy[]
        circularEnemy: CircularEnemy[]

        constructor(
            private canvasWidth: number,
            private canvasHeight: number,
            private context: CanvasRenderingContext2D | null
        ) {
            this.loadImages()
            this.createShakingEnemy(1)
            this.createFlyingLeftEnemy(1)
            this.createSideToSideEnemy(1)
            this.createCircularEnemy(10)
        }

        // ------------------------------------------------------------------------------------------------------------

        // Triggers update and movement for all sprites.

        public animate(): void {
            for (let enemy of this.shakingEnemy) {
                enemy.animate()
            }

            for (let enemy of this.flyingLeftEnemy) {
                enemy.animate()
            }

            for (let enemy of this.sideToSideEnemy) {
                enemy.animate()
            }

            for (let enemy of this.circularEnemy) {
                enemy.animate()
            }
        }

        // ------------------------------------------------------------------------------------------------------------

        // Creates an enemy that looks like a buzzsaw. 

        public createCircularEnemy(total: number): CircularEnemy[] {
            let spriteWidth = 213
            let spriteHeight = 213
            let totalFrames = 5
            let width = spriteWidth / 3
            let height = spriteHeight / 3

            let newCircularEnemies: CircularEnemy[] = []            

            for (let i = 0; i < total; i++) {
                let newEnemy = new CircularEnemy(
                    (Math.random() * this.canvasWidth + width) - width,
                    (Math.random() * this.canvasHeight + height) - height,
                    this.context,
                    this.canvasWidth,
                    this.canvasHeight,
                    this.images[3],
                    spriteWidth,
                    spriteHeight,
                    totalFrames,
                    width,
                    height,
                )
                newCircularEnemies.push(newEnemy)
            }

            return this.circularEnemy = newCircularEnemies
        }

        // ------------------------------------------------------------------------------------------------------------

        // Creates an enemy that goes all sorts of crazy around the screen.

        public createSideToSideEnemy(total: number): SideToSideEnemy[] {
            let spriteWidth = 218
            let spriteHeight = 177
            let totalFrames = 5
            let width = spriteWidth / 3
            let height = spriteHeight / 2.7

            let newSideToSideEnemies: SideToSideEnemy[] = []

            for (let i = 0; i < total; i++) {
                let newEnemy = new SideToSideEnemy(
                    (Math.random() * this.canvasWidth + width) - width,
                    (Math.random() * this.canvasHeight + height) - height,
                    this.context,
                    this.canvasWidth,
                    this.canvasHeight,
                    this.images[2],
                    spriteWidth,
                    spriteHeight,
                    totalFrames,
                    width,
                    height,
                )
                newSideToSideEnemies.push(newEnemy)
            }

            return this.sideToSideEnemy = newSideToSideEnemies
        }

        // ------------------------------------------------------------------------------------------------------------


        // Creates enemies that shake around the same position. 

        public createShakingEnemy(total: number): ShakingEnemy[] {
            let spriteWidth = 293
            let spriteHeight = 155
            let totalFrames = 4
            let width = spriteWidth / 3
            let height = spriteHeight / 2.5

            let newShakingEnemies: ShakingEnemy[] = []

            for (let i = 0; i < total; i++) {
                let newEnemy = new ShakingEnemy(
                    (Math.random() * this.canvasWidth + width) - width,
                    (Math.random() * this.canvasHeight + height) - height,
                    this.context,
                    this.canvasWidth,
                    this.canvasHeight,
                    this.images[0],
                    spriteWidth,
                    spriteHeight,
                    totalFrames,
                    width,
                    height,
                )
                newShakingEnemies.push(newEnemy)
            }

            return this.shakingEnemy = newShakingEnemies
        }

        // ------------------------------------------------------------------------------------------------------------

        // Creates enemies that fly from the right of the screen to the left.

        public createFlyingLeftEnemy(total: number): FlyingLeftEnemy[] {
            let spriteWidth = 266
            let spriteHeight = 188
            let totalFrames = 5
            let width = spriteWidth / 3
            let height = spriteWidth / 3.5

            let newflyingEnemies: FlyingLeftEnemy[] = []

            for (let i = 0; i < total; i++) {
                let newEnemy = new FlyingLeftEnemy(
                    this.canvasWidth + width + (Math.random() * (width * 3)),
                    (Math.random() * this.canvasHeight + height) - height,
                    this.context,
                    this.canvasWidth,
                    this.canvasHeight,
                    this.images[1],
                    spriteWidth,
                    spriteHeight,
                    totalFrames,
                    width,
                    height,
                )
                newflyingEnemies.push(newEnemy)
            }

            return this.flyingLeftEnemy = newflyingEnemies
        }

        // ------------------------------------------------------------------------------------------------------------

        // Loads enemy sprite assets into memory.

        private loadImages(): HTMLImageElement[] {
            let prefix = "../enemies/"
            let paths = ["enemy1", "enemy2", "enemy3", "enemy4"]
            let suffix = ".png"
            let loadedImages: HTMLImageElement[] = []

            for (let i = 0; i < paths.length; i++) {
                let enemyImage = new Image()
                enemyImage.src = `${prefix}${paths[i]}${suffix}`
                loadedImages.push(enemyImage)
            }

            return this.images = loadedImages
        }

        // ------------------------------------------------------------------------------------------------------------
    }
}