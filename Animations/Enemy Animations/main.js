var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var main;
(function (main) {
    // Base class for every enemy to inherit from.
    var EnemyBase = /** @class */ (function () {
        function EnemyBase(x, y, context, canvasWidth, canvasHeight, image, spriteWidth, spriteHeight, totalSpriteFrames, width, height) {
            this.x = x;
            this.y = y;
            this.context = context;
            this.canvasWidth = canvasWidth;
            this.canvasHeight = canvasHeight;
            this.image = image;
            this.spriteWidth = spriteWidth;
            this.spriteHeight = spriteHeight;
            this.totalSpriteFrames = totalSpriteFrames;
            this.width = width;
            this.height = height;
            this.currentFrame = 0;
            this.lastUpdate = new Date().getTime();
            this.updateDelta = (Math.random() * 60 + 10) - 10;
            this.toMoveX = 0;
            this.toMoveY = 0;
        }
        // ------------------------------------------------------------------------------------------------------------
        // Animate the sprite and calls the function that prepares movement for the next animation frame.
        EnemyBase.prototype.animate = function () {
            var _a;
            (_a = this.context) === null || _a === void 0 ? void 0 : _a.drawImage(this.image, this.currentFrame * this.spriteWidth, 0, this.spriteWidth, this.spriteHeight, this.x, this.y, this.width, this.height);
            this.update();
        };
        // ------------------------------------------------------------------------------------------------------------
        // Responsible for dealing with what happens to sprites as they attempt to move out of the canvas view.
        EnemyBase.prototype.boundariesLogic = function () {
            if (this.x < 0) {
                this.x = 1;
            }
            if (this.x + this.width > this.canvasWidth) {
                this.x = this.canvasWidth - this.width;
            }
            if (this.y < 0) {
                this.y = 1;
            }
            if (this.y + this.height > this.canvasHeight) {
                this.y = this.canvasHeight - this.height;
            }
        };
        // ------------------------------------------------------------------------------------------------------------
        /*
            Each sprite has a specific movement pattern. This function should control the x and y movement. As a result
            large variation is expected between specific sprites.
        */
        EnemyBase.prototype.calculateMove = function () { };
        // ------------------------------------------------------------------------------------------------------------
        /*
            Update the animation sprite movement for the next frame . Should tie together x and y movement and
            bounds logic. Should be overloaded to fit the needs of the specific animation.
        */
        EnemyBase.prototype.update = function () {
            this.calculateMove();
            this.x += this.toMoveX;
            this.y += this.toMoveY;
            this.boundariesLogic();
            this.updateFlyingAnimation();
        };
        // ------------------------------------------------------------------------------------------------------------
        // Should be called by update. Updates the flying animation of the sprite based upon a randomized time delta.
        EnemyBase.prototype.updateFlyingAnimation = function () {
            var _this = this;
            var checkDelta = function () {
                var now = new Date().getTime();
                if (now - _this.lastUpdate > _this.updateDelta) {
                    return true;
                }
                return false;
            };
            if (checkDelta()) {
                this.currentFrame++;
                this.lastUpdate = new Date().getTime();
            }
            if (this.currentFrame > this.totalSpriteFrames) {
                this.currentFrame = 0;
            }
        };
        return EnemyBase;
    }());
    main.EnemyBase = EnemyBase;
})(main || (main = {}));
/// <reference path="enemyBase.ts" />
var main;
(function (main) {
    var CircularEnemy = /** @class */ (function (_super) {
        __extends(CircularEnemy, _super);
        function CircularEnemy() {
            var _this = _super !== null && _super.apply(this, arguments) || this;
            _this.recalculateGameTicks = (Math.random() * 120) + 60;
            _this.passedTicks = 0;
            return _this;
            // ------------------------------------------------------------------------------------------------------------
        }
        // ------------------------------------------------------------------------------------------------------------
        // Sets a new sprite position every 60 - 120 game ticks. 
        CircularEnemy.prototype.calculateMove = function () {
            if (this.passedTicks >= this.recalculateGameTicks) {
                var newX = Math.floor(Math.random() * (this.canvasWidth - this.width));
                var newY = Math.floor(Math.random() * (this.canvasHeight - this.height));
                this.toMoveX = this.x - newX;
                this.toMoveY = this.y - newY;
                this.passedTicks = 0;
            }
            this.passedTicks++;
        };
        // ------------------------------------------------------------------------------------------------------------
        // Updates the position of the sprite a small percantage of the distance each game tick.
        CircularEnemy.prototype.update = function () {
            this.calculateMove();
            this.x -= this.toMoveX / 100;
            this.y -= this.toMoveY / 100;
            this.boundariesLogic();
            this.updateFlyingAnimation();
        };
        return CircularEnemy;
    }(main.EnemyBase));
    main.CircularEnemy = CircularEnemy;
})(main || (main = {}));
/// <reference path="enemyBase.ts" />
var main;
(function (main) {
    var ShakingEnemy = /** @class */ (function (_super) {
        __extends(ShakingEnemy, _super);
        function ShakingEnemy() {
            return _super !== null && _super.apply(this, arguments) || this;
        }
        // ------------------------------------------------------------------------------------------------------------
        // The movement algorithm for the sprite. 
        ShakingEnemy.prototype.calculateMove = function () {
            this.toMoveX = Math.random() * 5 - 2.5;
            this.toMoveY = Math.random() * 5 - 2.5;
        };
        return ShakingEnemy;
    }(main.EnemyBase));
    main.ShakingEnemy = ShakingEnemy;
})(main || (main = {}));
/// <reference path="enemyBase.ts" />
var main;
(function (main) {
    var FlyingLeftEnemy = /** @class */ (function (_super) {
        __extends(FlyingLeftEnemy, _super);
        function FlyingLeftEnemy() {
            var _this = _super !== null && _super.apply(this, arguments) || this;
            _this.movement = (Math.random() * 6 + 1) * -1;
            _this.angle = Math.random() * 2;
            _this.angleSpeed = Math.random() * 0.2;
            _this.curve = Math.random() * 7;
            return _this;
            // ------------------------------------------------------------------------------------------------------------
        }
        // ------------------------------------------------------------------------------------------------------------
        // Allows the spite to move off the screen. Is put back at the other edge when not visable. 
        FlyingLeftEnemy.prototype.boundariesLogic = function () {
            if (this.x < 0 - this.width) {
                this.x = this.canvasWidth + this.width;
            }
            if (this.y < 0) {
                this.y = 1;
            }
            if (this.y + this.height > this.canvasHeight) {
                this.y = this.canvasHeight - this.height;
            }
        };
        // ------------------------------------------------------------------------------------------------------------
        // Moves the sprite from the right to the left of the screen at a unique speed. Adds a sin wave motion. 
        FlyingLeftEnemy.prototype.calculateMove = function () {
            this.toMoveX = this.movement;
            this.toMoveY = this.curve * Math.sin(this.angle);
            this.angle += this.angleSpeed;
        };
        return FlyingLeftEnemy;
    }(main.EnemyBase));
    main.FlyingLeftEnemy = FlyingLeftEnemy;
})(main || (main = {}));
/// <reference path="enemyBase.ts" />
var main;
(function (main) {
    var SideToSideEnemy = /** @class */ (function (_super) {
        __extends(SideToSideEnemy, _super);
        function SideToSideEnemy() {
            var _this = _super !== null && _super.apply(this, arguments) || this;
            _this.angle = 0;
            _this.angleSpeed = Math.random() * 2;
            _this.curve = Math.random() * 200 + 50;
            return _this;
            // ------------------------------------------------------------------------------------------------------------
        }
        // ------------------------------------------------------------------------------------------------------------
        // Moves the sprites in circular movement. Change the MAth.PI / X  numbers to get different effects
        SideToSideEnemy.prototype.calculateMove = function () {
            this.angle += this.angleSpeed;
            this.x = (this.canvasWidth - this.width) / 2 * Math.sin(this.angle * Math.PI / 200) + (this.canvasWidth / 2 - this.width / 2);
            this.y = (this.canvasHeight - this.height) / 2 * Math.cos(this.angle * Math.PI / 250) + (this.canvasHeight / 2 - this.height / 2);
        };
        // ------------------------------------------------------------------------------------------------------------     
        // Update the elements.
        SideToSideEnemy.prototype.update = function () {
            this.calculateMove();
            this.boundariesLogic();
            this.updateFlyingAnimation();
        };
        return SideToSideEnemy;
    }(main.EnemyBase));
    main.SideToSideEnemy = SideToSideEnemy;
})(main || (main = {}));
/// <reference path="shakingEnemy.ts" />
/// <reference path="flyingLeftEnemy.ts" />
/// <reference path="sideToSideEnemy.ts" />
/// <reference path="circularEnemy.ts" />
var main;
(function (main) {
    var EnemyManager = /** @class */ (function () {
        function EnemyManager(canvasWidth, canvasHeight, context) {
            this.canvasWidth = canvasWidth;
            this.canvasHeight = canvasHeight;
            this.context = context;
            this.loadImages();
            this.createShakingEnemy(1);
            this.createFlyingLeftEnemy(1);
            this.createSideToSideEnemy(1);
            this.createCircularEnemy(10);
        }
        // ------------------------------------------------------------------------------------------------------------
        // Triggers update and movement for all sprites.
        EnemyManager.prototype.animate = function () {
            for (var _i = 0, _a = this.shakingEnemy; _i < _a.length; _i++) {
                var enemy = _a[_i];
                enemy.animate();
            }
            for (var _b = 0, _c = this.flyingLeftEnemy; _b < _c.length; _b++) {
                var enemy = _c[_b];
                enemy.animate();
            }
            for (var _d = 0, _e = this.sideToSideEnemy; _d < _e.length; _d++) {
                var enemy = _e[_d];
                enemy.animate();
            }
            for (var _f = 0, _g = this.circularEnemy; _f < _g.length; _f++) {
                var enemy = _g[_f];
                enemy.animate();
            }
        };
        // ------------------------------------------------------------------------------------------------------------
        // Creates an enemy that looks like a buzzsaw. 
        EnemyManager.prototype.createCircularEnemy = function (total) {
            var spriteWidth = 213;
            var spriteHeight = 213;
            var totalFrames = 5;
            var width = spriteWidth / 3;
            var height = spriteHeight / 3;
            var newCircularEnemies = [];
            for (var i = 0; i < total; i++) {
                var newEnemy = new main.CircularEnemy((Math.random() * this.canvasWidth + width) - width, (Math.random() * this.canvasHeight + height) - height, this.context, this.canvasWidth, this.canvasHeight, this.images[3], spriteWidth, spriteHeight, totalFrames, width, height);
                newCircularEnemies.push(newEnemy);
            }
            return this.circularEnemy = newCircularEnemies;
        };
        // ------------------------------------------------------------------------------------------------------------
        // Creates an enemy that goes all sorts of crazy around the screen.
        EnemyManager.prototype.createSideToSideEnemy = function (total) {
            var spriteWidth = 218;
            var spriteHeight = 177;
            var totalFrames = 5;
            var width = spriteWidth / 3;
            var height = spriteHeight / 2.7;
            var newSideToSideEnemies = [];
            for (var i = 0; i < total; i++) {
                var newEnemy = new main.SideToSideEnemy((Math.random() * this.canvasWidth + width) - width, (Math.random() * this.canvasHeight + height) - height, this.context, this.canvasWidth, this.canvasHeight, this.images[2], spriteWidth, spriteHeight, totalFrames, width, height);
                newSideToSideEnemies.push(newEnemy);
            }
            return this.sideToSideEnemy = newSideToSideEnemies;
        };
        // ------------------------------------------------------------------------------------------------------------
        // Creates enemies that shake around the same position. 
        EnemyManager.prototype.createShakingEnemy = function (total) {
            var spriteWidth = 293;
            var spriteHeight = 155;
            var totalFrames = 4;
            var width = spriteWidth / 3;
            var height = spriteHeight / 2.5;
            var newShakingEnemies = [];
            for (var i = 0; i < total; i++) {
                var newEnemy = new main.ShakingEnemy((Math.random() * this.canvasWidth + width) - width, (Math.random() * this.canvasHeight + height) - height, this.context, this.canvasWidth, this.canvasHeight, this.images[0], spriteWidth, spriteHeight, totalFrames, width, height);
                newShakingEnemies.push(newEnemy);
            }
            return this.shakingEnemy = newShakingEnemies;
        };
        // ------------------------------------------------------------------------------------------------------------
        // Creates enemies that fly from the right of the screen to the left.
        EnemyManager.prototype.createFlyingLeftEnemy = function (total) {
            var spriteWidth = 266;
            var spriteHeight = 188;
            var totalFrames = 5;
            var width = spriteWidth / 3;
            var height = spriteWidth / 3.5;
            var newflyingEnemies = [];
            for (var i = 0; i < total; i++) {
                var newEnemy = new main.FlyingLeftEnemy(this.canvasWidth + width + (Math.random() * (width * 3)), (Math.random() * this.canvasHeight + height) - height, this.context, this.canvasWidth, this.canvasHeight, this.images[1], spriteWidth, spriteHeight, totalFrames, width, height);
                newflyingEnemies.push(newEnemy);
            }
            return this.flyingLeftEnemy = newflyingEnemies;
        };
        // ------------------------------------------------------------------------------------------------------------
        // Loads enemy sprite assets into memory.
        EnemyManager.prototype.loadImages = function () {
            var prefix = "../enemies/";
            var paths = ["enemy1", "enemy2", "enemy3", "enemy4"];
            var suffix = ".png";
            var loadedImages = [];
            for (var i = 0; i < paths.length; i++) {
                var enemyImage = new Image();
                enemyImage.src = "".concat(prefix).concat(paths[i]).concat(suffix);
                loadedImages.push(enemyImage);
            }
            return this.images = loadedImages;
        };
        return EnemyManager;
    }());
    main.EnemyManager = EnemyManager;
})(main || (main = {}));
/// <reference path="enemyManager.ts" />
var main;
(function (main) {
    var Game = /** @class */ (function () {
        function Game() {
            this.canvasHeight = 800;
            this.canvasWidth = 1500;
            this.targetFPS = 60;
            this.targetUpdateDelta = 1000 / this.targetFPS;
            this.createCanvas();
            this.createContext();
            this.lastUpdate = new Date().getTime();
            this.enemyManager = new main.EnemyManager(this.canvasWidth, this.canvasHeight, this.context);
        }
        // ------------------------------------------------------------------------------------------------------------
        // Animates all the elements
        Game.prototype.animate = function () {
            var _this = this;
            var _a;
            if (this.updateFrame()) {
                (_a = this.context) === null || _a === void 0 ? void 0 : _a.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
                this.enemyManager.animate();
                this.lastUpdate = new Date().getTime();
            }
            requestAnimationFrame(function () {
                _this.animate();
            });
        };
        // ------------------------------------------------------------------------------------------------------------
        // Create a new canvas element to draw on.
        Game.prototype.createCanvas = function () {
            var newCanvas = document.createElement("canvas");
            newCanvas.width = this.canvasWidth;
            newCanvas.height = this.canvasHeight;
            newCanvas.style.border = "3px solid";
            newCanvas.style.position = "absolute";
            newCanvas.style.top = "50%";
            newCanvas.style.left = "50%";
            newCanvas.style.transform = "translate(-50%, -50%)";
            newCanvas.style.height = "".concat(this.canvasHeight, "px");
            newCanvas.style.width = "".concat(this.canvasWidth, "px");
            document.body.appendChild(newCanvas);
            return this.canvas = newCanvas;
        };
        // ------------------------------------------------------------------------------------------------------------
        // Get the canvas context.
        Game.prototype.createContext = function () {
            return this.context = this.canvas.getContext("2d");
        };
        // ------------------------------------------------------------------------------------------------------------
        // Updates calculates if enough time has passed between updates. 
        Game.prototype.updateFrame = function () {
            var now = new Date().getTime();
            var timeDelta = Math.abs(now - this.lastUpdate);
            if (timeDelta > this.targetUpdateDelta) {
                return true;
            }
            return false;
        };
        return Game;
    }());
    main.Game = Game;
})(main || (main = {}));
/// <reference path="game.ts" />
var main;
(function (main_1) {
    function main() {
        var game = new main_1.Game();
        game.animate();
    }
    main();
})(main || (main = {}));
//# sourceMappingURL=main.js.map