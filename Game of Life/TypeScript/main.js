var OFFSET = [[-1, 0], [1, 0], [0, -1], [0, 1], [-1, -1], [-1, 1], [1, -1], [1, 1]];
var main;
(function (main) {
    var Grid = /** @class */ (function () {
        function Grid(context, rows, columns, cellSize) {
            this.context = context;
            this.rows = rows;
            this.columns = columns;
            this.cellSize = cellSize;
            this.populateGrid();
        }
        // ------------------------------------------------
        // Counts the live neighbors.
        Grid.prototype.countLive = function (row, column) {
            var livecount = 0;
            for (var _i = 0, OFFSET_1 = OFFSET; _i < OFFSET_1.length; _i++) {
                var pair = OFFSET_1[_i];
                var neighborRow = (row + pair[1] + this.totalRows) % this.totalRows;
                var neighborColumn = (column + pair[0] + this.totalColumns) % this.totalColumns;
                livecount += this.cells[neighborRow][neighborColumn];
            }
            return livecount;
        };
        // ------------------------------------------------
        // Draws the grid on the screen.
        Grid.prototype.draw = function () {
            var _a;
            var totalRows = this.cells.length;
            var totalColumns = this.cells[0].length;
            for (var i = 0; i < totalRows; i++) {
                var y = i * this.cellSize;
                for (var j = 0; j < totalColumns; j++) {
                    var x = j * this.cellSize;
                    if (this.context != null) {
                        var cellValue = this.cells[i][j];
                        var fillstyle = (cellValue == 1) ? "white" : "black";
                        this.context.fillStyle = fillstyle;
                    }
                    (_a = this.context) === null || _a === void 0 ? void 0 : _a.fillRect(y, x, this.cellSize - 1, this.cellSize - 1);
                }
            }
        };
        // ------------------------------------------------
        // Create a grid of integers. 1 = alive. 0 = dead.
        Grid.prototype.populateGrid = function () {
            var newGrid = [];
            for (var i = 0; i < this.columns; i++) {
                var newColumn = [];
                for (var j = 0; j < this.rows; j++) {
                    var aliveStatus = Math.floor(Math.random() * 10);
                    var choice = (aliveStatus < 2) ? 1 : 0;
                    newColumn.push(choice);
                }
                newGrid.push(newColumn);
            }
            this.totalRows = newGrid.length;
            this.totalColumns = newGrid[0].length;
            return this.cells = newGrid;
        };
        // ------------------------------------------------
        // Update the state of the cells. 
        Grid.prototype.update = function () {
            var newGrid = [];
            for (var i = 0; i < this.totalRows; i++) {
                var newRow = [];
                for (var j = 0; j < this.totalColumns; j++) {
                    var liveNeighbors = this.countLive(i, j);
                    if (liveNeighbors == 3) {
                        newRow.push(1);
                    }
                    else if (liveNeighbors == 2) {
                        newRow.push(this.cells[i][j]);
                    }
                    else {
                        newRow.push(0);
                    }
                }
                newGrid.push(newRow);
            }
            return this.cells = newGrid;
        };
        return Grid;
    }());
    main.Grid = Grid;
})(main || (main = {}));
/// <reference path="grid.ts" />
var main;
(function (main) {
    var Simulation = /** @class */ (function () {
        function Simulation() {
            this.canvasWidth = 1900;
            this.canvasHeight = 1000;
            this.cellSize = 5;
            this.fps = 60;
            this.interval = 1000 / this.fps;
            this.createCanvas();
            this.createGrid();
            this.lastUpdate = Date.now();
        }
        // ------------------------------------------------
        // Create a drawing surface.
        Simulation.prototype.createCanvas = function () {
            var container = document.createElement("div");
            container.setAttribute("id", "container");
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
            container.appendChild(canvas);
            var resetButton = this.resetButton();
            container.appendChild(resetButton);
            document.body.appendChild(container);
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
        // Update and redraw the simulation every tick. 
        Simulation.prototype.draw = function () {
            var _this = this;
            var _a;
            this.now = Date.now();
            if (this.now - this.lastUpdate > this.interval) {
                (_a = this.context) === null || _a === void 0 ? void 0 : _a.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
                this.grid.draw();
                this.grid.update();
                this.lastUpdate = Date.now();
            }
            this.button.addEventListener("click", function () {
                _this.grid.populateGrid();
            });
            requestAnimationFrame(function () {
                _this.draw();
            });
        };
        // ------------------------------------------------
        // A reset button. Calls resets the grid state.
        Simulation.prototype.resetButton = function () {
            var newButton = document.createElement("button");
            newButton.innerText = "Reset";
            this.button = newButton;
            return newButton;
        };
        return Simulation;
    }());
    main.Simulation = Simulation;
})(main || (main = {}));
/// <reference path="simulation.ts" />
var main;
(function (main_1) {
    function main() {
        var newSimulation = new main_1.Simulation();
        newSimulation.draw();
    }
    main();
})(main || (main = {}));
