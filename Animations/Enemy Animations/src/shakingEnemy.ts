/// <reference path="enemyBase.ts" />

namespace main {

    export class ShakingEnemy extends EnemyBase {

        // ------------------------------------------------------------------------------------------------------------

        // The movement algorithm for the sprite. 

        public calculateMove(): void {
            this.toMoveX = Math.random() * 5 - 2.5
            this.toMoveY = Math.random() * 5 - 2.5
        }

        // ------------------------------------------------------------------------------------------------------------
    }
}