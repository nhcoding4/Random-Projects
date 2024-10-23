// --------------------------------------------------------------------------------------------------------------------
var Game = /** @class */ (function () {
    function Game() {
        this.height = 600;
        this.spriteHeight = 523;
        this.spriteWidth = 575;
        this.staggerFrame = 5;
        this.width = 600;
        this.defaultAnimation = "idle";
        this.spirteAnimations = [];
    }
    // --------------------------------------------------------------------------------------------------------------------
    Game.prototype.animate = function (gameFrame, animationState) {
        var _this = this;
        var _a, _b;
        var dropDown = document.getElementById("animations");
        dropDown === null || dropDown === void 0 ? void 0 : dropDown.addEventListener("change", function (e) {
            animationState = e.target.value;
        });
        var currentLength = this.spirteAnimations[animationState].loc.length;
        var position = Math.floor(gameFrame / this.staggerFrame) % currentLength;
        var frameX = this.spriteWidth * position;
        var frameY = this.spirteAnimations[animationState].loc[position].y;
        (_a = this.context) === null || _a === void 0 ? void 0 : _a.clearRect(0, 0, this.width, this.height);
        (_b = this.context) === null || _b === void 0 ? void 0 : _b.drawImage(this.imageData, frameX, frameY, this.spriteWidth, this.spriteHeight, 0, 0, this.spriteWidth, this.spriteHeight);
        gameFrame++;
        if (gameFrame > currentLength * 10) {
            gameFrame = 1;
        }
        requestAnimationFrame(function () {
            _this.animate(gameFrame, animationState);
        });
    };
    // --------------------------------------------------------------------------------------------------------------------
    // Creates a box that allows animation selection
    Game.prototype.animationSelector = function () {
        var root = document.createElement("div");
        root.setAttribute("class", "controls");
        root.style.position = "absolute";
        root.style.zIndex = "10";
        root.style.top = "10px";
        root.style.left = "50%";
        root.style.transform = "translateX(-50%)";
        var label = document.createElement("label");
        label.setAttribute("for", "animations");
        label.textContent = "Choose Animation:";
        label.style.fontSize = "25px";
        var select = document.createElement("select");
        select.setAttribute("id", "animations");
        select.setAttribute("name", "animations");
        select.style.fontSize = "25px";
        var options = ["idle", "jump", "fall", "run", "dizzy", "sit", "roll", "bite", "ko", "getHit"];
        for (var i = 0; i < options.length; i++) {
            var newOption = document.createElement("option");
            var value = options[i];
            newOption.setAttribute("value", value);
            newOption.textContent = value;
            newOption.style.fontSize = "25px";
            select.appendChild(newOption);
        }
        root.appendChild(label);
        root.appendChild(select);
        document.body.appendChild(root);
        return this.animationDropdown = root;
    };
    // --------------------------------------------------------------------------------------------------------------------
    // Create a base canvas.
    Game.prototype.createCanvas = function () {
        var canvas = document.createElement("canvas");
        canvas.setAttribute("id", "canvas");
        canvas.width = this.width;
        canvas.height = this.height;
        canvas.style.width = "".concat(this.width, "px");
        canvas.style.height = "".concat(this.height, "px");
        canvas.style.border = "3px solid black";
        canvas.style.position = "absolute";
        canvas.style.top = "50%";
        canvas.style.left = "50%";
        canvas.style.transform = "translate(-50%, -50%)";
        document.body.appendChild(canvas);
        return this.canvas = canvas;
    };
    // --------------------------------------------------------------------------------------------------------------------
    // Load sprite sheet
    Game.prototype.loadImage = function () {
        var image = new Image();
        var imageLocation = "./shadow_dog.png";
        image.src = imageLocation;
        return this.imageData = image;
    };
    // --------------------------------------------------------------------------------------------------------------------
    // Generate sprite sheet animation locations
    Game.prototype.generateAnimations = function () {
        var _this = this;
        var animationStates = [
            {
                name: "idle",
                frames: 7,
            },
            {
                name: "jump",
                frames: 7,
            },
            {
                name: "fall",
                frames: 7,
            },
            {
                name: "run",
                frames: 9,
            },
            {
                name: "dizzy",
                frames: 11,
            },
            {
                name: "sit",
                frames: 5,
            },
            {
                name: "roll",
                frames: 7,
            },
            {
                name: "bite",
                frames: 7,
            },
            {
                name: "ko",
                frames: 12,
            },
            {
                name: "getHit",
                frames: 4,
            },
        ];
        animationStates.forEach(function (state, index) {
            var frames = {
                loc: []
            };
            for (var i = 0; i < state.frames; i++) {
                var positionX = i * _this.spriteWidth;
                var positionY = index * _this.spriteHeight;
                var newFrame = { x: positionX, y: positionY };
                frames.loc.push(newFrame);
            }
            _this.spirteAnimations[state.name] = frames;
        });
    };
    // --------------------------------------------------------------------------------------------------------------------
    Game.prototype.run = function () {
        this.loadImage();
        this.createCanvas();
        this.context = this.canvas.getContext("2d");
        this.animationSelector();
        this.generateAnimations();
        this.animate(0, this.defaultAnimation);
    };
    return Game;
}());
// --------------------------------------------------------------------------------------------------------------------
function main() {
    var game = new Game();
    game.run();
}
main();
