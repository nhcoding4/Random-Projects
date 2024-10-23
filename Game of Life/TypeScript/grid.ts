const OFFSET = [[-1, 0], [1, 0], [0, -1], [0, 1], [-1, -1], [-1, 1], [1, -1], [1, 1]]

namespace main {
    export class Grid {
        context: CanvasRenderingContext2D | null
        rows: number
        columns: number
        cellSize: number
        cells: number[][]
        totalRows: number
        totalColumns: number

        constructor(context: CanvasRenderingContext2D | null, rows: number, columns: number, cellSize: number) {
            this.context = context
            this.rows = rows
            this.columns = columns
            this.cellSize = cellSize
            this.populateGrid()
        }

        // ------------------------------------------------

        // Counts the live neighbors.

        countLive(row: number, column: number): number {
            let livecount = 0

            for (let pair of OFFSET) {
                let neighborRow = (row + pair[1] + this.totalRows) % this.totalRows
                let neighborColumn = (column + pair[0] + this.totalColumns) % this.totalColumns
                livecount += this.cells[neighborRow][neighborColumn]
            }

            return livecount
        }

        // ------------------------------------------------


        // Draws the grid on the screen.

        draw() {
            let totalRows = this.cells.length
            let totalColumns = this.cells[0].length

            for (let i = 0; i < totalRows; i++) {
                let y = i * this.cellSize

                for (let j = 0; j < totalColumns; j++) {
                    let x = j * this.cellSize

                    if (this.context != null) {
                        let cellValue = this.cells[i][j]
                        let fillstyle = (cellValue == 1) ? "white" : "black"
                        this.context.fillStyle = fillstyle
                    }
                    this.context?.fillRect(y, x, this.cellSize - 1, this.cellSize - 1)
                }
            }
        }

        // ------------------------------------------------

        // Create a grid of integers. 1 = alive. 0 = dead.

        populateGrid(): number[][] {
            let newGrid: number[][] = []

            for (let i = 0; i < this.columns; i++) {
                let newColumn: number[] = []

                for (let j = 0; j < this.rows; j++) {
                    let aliveStatus = Math.floor(Math.random() * 10)
                    let choice = (aliveStatus < 2) ? 1 : 0
                    newColumn.push(choice)
                }
                newGrid.push(newColumn)
            }

            this.totalRows = newGrid.length
            this.totalColumns = newGrid[0].length
            return this.cells = newGrid
        }

        // ------------------------------------------------

        // Update the state of the cells. 

        update(): number[][] {
            let newGrid: number[][] = []

            for (let i = 0; i < this.totalRows; i++) {
                let newRow: number[] = []
                for (let j = 0; j < this.totalColumns; j++) {
                    let liveNeighbors = this.countLive(i, j)

                    if (liveNeighbors == 3) {
                        newRow.push(1)
                    } else if (liveNeighbors == 2) {
                        newRow.push(this.cells[i][j])
                    } else {
                        newRow.push(0)
                    }
                }
                newGrid.push(newRow)
            }
            return this.cells = newGrid
        }

        // ------------------------------------------------
    }
}