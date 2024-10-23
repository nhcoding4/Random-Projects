declare namespace main {
    class EnemyBase {
        x: number;
        y: number;
        context: CanvasRenderingContext2D | null;
        canvasWidth: number;
        canvasHeight: number;
        image: HTMLImageElement;
        spriteWidth: number;
        spriteHeight: number;
        totalSpriteFrames: number;
        width: number;
        height: number;
        lastUpdate: number;
        updateDelta: number;
        toMoveX: number;
        toMoveY: number;
        currentFrame: number;
        constructor(x: number, y: number, context: CanvasRenderingContext2D | null, canvasWidth: number, canvasHeight: number, image: HTMLImageElement, spriteWidth: number, spriteHeight: number, totalSpriteFrames: number, width: number, height: number);
        animate(): void;
        boundariesLogic(): void;
        calculateMove(): void;
        update(): void;
        updateFlyingAnimation(): void;
    }
}
declare namespace main {
    class CircularEnemy extends EnemyBase {
        recalculateGameTicks: number;
        passedTicks: number;
        calculateMove(): void;
        update(): void;
    }
}
declare namespace main {
    class ShakingEnemy extends EnemyBase {
        calculateMove(): void;
    }
}
declare namespace main {
    class FlyingLeftEnemy extends EnemyBase {
        movement: number;
        angle: number;
        angleSpeed: number;
        curve: number;
        boundariesLogic(): void;
        calculateMove(): void;
    }
}
declare namespace main {
    class SideToSideEnemy extends EnemyBase {
        angle: number;
        angleSpeed: number;
        curve: number;
        calculateMove(): void;
        update(): void;
    }
}
declare namespace main {
    class EnemyManager {
        private canvasWidth;
        private canvasHeight;
        private context;
        images: HTMLImageElement[];
        shakingEnemy: ShakingEnemy[];
        flyingLeftEnemy: FlyingLeftEnemy[];
        sideToSideEnemy: SideToSideEnemy[];
        circularEnemy: CircularEnemy[];
        constructor(canvasWidth: number, canvasHeight: number, context: CanvasRenderingContext2D | null);
        animate(): void;
        createCircularEnemy(total: number): CircularEnemy[];
        createSideToSideEnemy(total: number): SideToSideEnemy[];
        createShakingEnemy(total: number): ShakingEnemy[];
        createFlyingLeftEnemy(total: number): FlyingLeftEnemy[];
        private loadImages;
    }
}
declare namespace main {
    class Game {
        private canvas;
        private context;
        private canvasHeight;
        private canvasWidth;
        private targetFPS;
        private targetUpdateDelta;
        private lastUpdate;
        private enemyManager;
        constructor();
        animate(): void;
        private createCanvas;
        private createContext;
        private updateFrame;
    }
}
declare namespace main {
}
