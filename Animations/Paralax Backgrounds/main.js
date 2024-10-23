// --------------------------------------------------------------------------------------------------------------------
var Layer = /** @class */ (function () {
    // --------------------------------------------------------
    // Object setup.
    function Layer(image, speedModifier, gameSpeed, context, height, width) {
        this.context = context;
        this.gameFrame = 0;
        this.height = height;
        this.image = image;
        this.speed = gameSpeed * speedModifier;
        this.speedModifier = speedModifier;
        this.width = width;
        this.x2 = width;
        this.x = 0;
        this.y = 0;
    }
    // --------------------------------------------------------
    // Draws images to the canvas.
    Layer.prototype.draw = function () {
        var _a, _b;
        (_a = this.context) === null || _a === void 0 ? void 0 : _a.drawImage(this.image, this.x, this.y, this.width, this.height);
        (_b = this.context) === null || _b === void 0 ? void 0 : _b.drawImage(this.image, this.x2, this.y, this.width, this.height);
    };
    // --------------------------------------------------------
    // Change the x position of the layer
    Layer.prototype.update = function () {
        var _this = this;
        var moveBackground = (function (xValue, offset) {
            if (xValue < -_this.width) {
                xValue = _this.width - _this.speed + offset;
            }
            else {
                xValue -= _this.speed;
            }
            return xValue;
        });
        this.x = Math.floor(moveBackground(this.x, this.x2));
        this.x2 = Math.floor(moveBackground(this.x2, this.x));
    };
    // --------------------------------------------------------        
    // Updates the speed of the layer
    Layer.prototype.updateSpeed = function (newSpeed) {
        this.speed = newSpeed * this.speedModifier;
    };
    return Layer;
}());
// --------------------------------------------------------------------------------------------------------------------
var Backgrounds = /** @class */ (function () {
    // --------------------------------------------------------
    // Object setup.
    function Backgrounds() {
        document.body.style.background = "black";
        this.canvasHeight = 700;
        this.canvasWidth = 800;
        this.gamespeed = 5;
        this.x = 0;
        this.x2 = 2400;
        this.createElements();
        this.loadImages();
    }
    // --------------------------------------------------------
    // Draw objects.
    Backgrounds.prototype.animate = function () {
        var _this = this;
        var _a;
        (_a = this.context) === null || _a === void 0 ? void 0 : _a.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
        this.gamespeed = parseInt(this.slider.value);
        this.paragraph.textContent = "Game Speed: ".concat(this.gamespeed);
        for (var _i = 0, _b = this.layers; _i < _b.length; _i++) {
            var layer = _b[_i];
            layer.updateSpeed(this.gamespeed);
            layer.update();
            layer.draw();
        }
        requestAnimationFrame(function () {
            _this.animate();
        });
    };
    // --------------------------------------------------------
    // Create object to draw on. 
    Backgrounds.prototype.createCanvas = function () {
        var canvas = document.createElement("canvas");
        canvas.setAttribute("id", "canvas");
        canvas.height = this.canvasHeight;
        canvas.width = this.canvasWidth;
        canvas.style.position = "relative";
        canvas.style.width = "".concat(this.canvasWidth, "px");
        canvas.style.height = "".concat(this.canvasHeight, "px");
        document.body.appendChild(canvas);
        this.context = canvas.getContext("2d");
        this.canvas = canvas;
        return canvas;
    };
    // --------------------------------------------------------
    // Creates a container for all the other objects.
    Backgrounds.prototype.createContainer = function () {
        var container = document.createElement("div");
        container.style.position = "absolute";
        container.style.width = "".concat(this.canvasWidth, "px");
        container.style.transform = "translate(-50%, -50%)";
        container.style.top = "50%";
        container.style.left = "50%";
        container.style.border = "3px solid white";
        container.style.fontSize = "25px";
        return container;
    };
    // --------------------------------------------------------
    // Creates the canvas and slider elements.
    Backgrounds.prototype.createElements = function () {
        var container = this.createContainer();
        document.body.appendChild(container);
        var canvas = this.createCanvas();
        container.appendChild(canvas);
        var paragraph = document.createElement("p");
        paragraph.textContent = "Game Speed: ".concat(this.gamespeed);
        paragraph.style.color = "white";
        container.appendChild(paragraph);
        this.paragraph = paragraph;
        var span = document.createElement("span");
        span.setAttribute("id", "showGameSpeed");
        paragraph.appendChild(span);
        var slider = this.createSlider();
        slider.style.width = "100%";
        container.appendChild(slider);
    };
    // --------------------------------------------------------
    // Creates a slider element.
    Backgrounds.prototype.createSlider = function () {
        var input = document.createElement("input");
        input.setAttribute("type", "range");
        input.setAttribute("min", "0");
        input.setAttribute("max", "20");
        input.setAttribute("value", "5");
        input.setAttribute("class", "slider");
        input.setAttribute("id", "slider");
        this.slider = input;
        this.slider.value = "".concat(this.gamespeed);
        return input;
    };
    // --------------------------------------------------------
    // Loads images from files and creates layer objects out of them.
    Backgrounds.prototype.loadImages = function () {
        var base = "./backgroundLayers/";
        var images = [];
        for (var i = 0; i < 5; i++) {
            var backgroundLayer = new Image();
            backgroundLayer.src = "".concat(base, "layer-").concat(i + 1, ".png");
            images.push(backgroundLayer);
        }
        var layers = [];
        var speedModifier = 0.1;
        for (var _i = 0, images_1 = images; _i < images_1.length; _i++) {
            var image = images_1[_i];
            var newLayer = new Layer(image, speedModifier, this.gamespeed, this.context, this.canvasHeight, this.x2);
            layers.push(newLayer);
            speedModifier += 0.1;
        }
        return this.layers = layers;
    };
    // --------------------------------------------------------
    // Main loop.
    Backgrounds.prototype.run = function () {
        this.animate();
    };
    return Backgrounds;
}());
// --------------------------------------------------------------------------------------------------------------------
window.addEventListener("load", function () {
    var background = new Backgrounds();
    background.run();
});
