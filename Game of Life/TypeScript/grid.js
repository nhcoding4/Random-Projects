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
        // Draws the grid on the screen.
        Grid.prototype.draw = function () {
            var _a, _b;
            var totalRows = this.cells.length;
            var totalColumns = this.cells[0].length;
            for (var i = 0; i < totalRows; i++) {
                var y = i * this.cellSize;
                for (var j = 0; j < totalColumns; j++) {
                    var x = j * this.cellSize;
                    if (this.cells[i][j] != 1 && this.context != null) {
                        this.context.fillStyle == "white";
                    }
                    else {
                        ((_a = this.context) === null || _a === void 0 ? void 0 : _a.fillStyle) == "black";
                    }
                    (_b = this.context) === null || _b === void 0 ? void 0 : _b.fillRect(y, x, this.cellSize - 1, this.cellSize - 1);
                }
            }
        };
        // ------------------------------------------------
        // Create a grid of integers using a 2d array.
        Grid.prototype.populateGrid = function () {
            var newGrid = [];
            for (var i = 0; i < this.columns; i++) {
                var newColumn = [];
                for (var j = 0; j < this.rows; j++) {
                    newColumn.push(0);
                }
                newGrid.push(newColumn);
            }
            return this.cells = newGrid;
        };
        return Grid;
    }());
    main.Grid = Grid;
})(main || (main = {}));
