/// <reference path="grid.ts" />
var main;
(function (main) {
    var Simulation = /** @class */ (function () {
        function Simulation() {
            this.canvasWidth = 1000;
            this.canvasHeight = 1000;
            this.cellSize = 25;
            this.createCanvas();
            this.createGrid();
        }
        // ------------------------------------------------
        // Create a drawing surface.
        Simulation.prototype.createCanvas = function () {
            var canvas = document.createElement("canvas");
            canvas.setAttribute("id", "canvas");
            canvas.height = this.canvasHeight;
            canvas.width = this.canvasWidth;
            canvas.style.border = "3px black";
            canvas.style.position = "absolute";
            canvas.style.top = "50%";
            canvas.style.left = "50%";
            canvas.style.transform = "translate(-50%, -50%)";
            canvas.style.height = "".concat(this.canvasHeight, "px");
            canvas.style.width = " ".concat(this.canvasWidth, "px");
            this.context = canvas.getContext("2d");
            document.body.appendChild(canvas);
            return this.canvas = canvas;
        };
        // ------------------------------------------------
        // Creates an instance of this grid object.
        Simulation.prototype.createGrid = function () {
            var totalCellsX = Math.floor(this.canvasWidth / this.cellSize);
            var totalCellsY = Math.floor(this.canvasHeight / this.cellSize);
            var grid = new main.Grid(this.context, totalCellsY, totalCellsX, this.cellSize);
            return this.grid = grid;
        };
        // ------------------------------------------------
        Simulation.prototype.draw = function () {
            var _this = this;
            var _a;
            (_a = this.context) === null || _a === void 0 ? void 0 : _a.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
            this.grid.draw();
            requestAnimationFrame(function () {
                _this.draw();
            });
        };
        return Simulation;
    }());
    main.Simulation = Simulation;
})(main || (main = {}));
