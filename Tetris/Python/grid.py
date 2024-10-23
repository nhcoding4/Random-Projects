import pyray as pr
from helpers import *


# ----------------------------------------------------------------------------------------------------------------------


class Grid:
    def __init__(self, cell_size: int) -> None:
        self.rows = 20
        self.columns = 10
        self.cell_size = cell_size
        self.grid: list[list[int]] = []
        self.colors: list[pr.Color] = colors()

        self.grid_setup()

    # -------------------------------------------------------------------------

    def clear_full_rows(self) -> int:
        """Clears the full rows and returns the amount of rows the user cleared."""

        completed = 0
        for i in range(self.rows - 1, -1, -1):
            if self.is_row_full(i):
                self.clear_row(i)
                completed += 1
            elif completed > 0:
                self.move_row_down(i, completed)

        return completed

    # -------------------------------------------------------------------------

    def clear_row(self, row: int) -> None:
        """Sets all the values in a row to 0."""

        for i in range(self.columns):
            self.grid[row][i] = 0

    # -------------------------------------------------------------------------

    def draw_grid(self) -> None:
        """Draws the grid. Color changes depending on the value held in a cell."""

        for i, row in enumerate(self.grid):
            for j, cell_value in enumerate(row):
                pr.draw_rectangle(
                    j * self.cell_size + 11,
                    i * self.cell_size + 11,
                    self.cell_size - 1,
                    self.cell_size - 1,
                    self.colors[cell_value],
                )

    # -------------------------------------------------------------------------

    def grid_setup(self) -> None:
        """Fills the list representing the grid with 0's."""

        new_grid: list[list[int]] = []

        for _ in range(self.rows):
            new_row: list[int] = []
            for _ in range(self.columns):
                new_row.append(0)
            new_grid.append(new_row)

        self.grid = new_grid

    # -------------------------------------------------------------------------

    def is_cell_empty(self, row: int, column: int) -> bool:
        """Checks if the next cell is empty."""

        if self.grid[row][column] == 0:
            return True

        return False

    # -------------------------------------------------------------------------

    def is_cell_outside_grid(self, row: int, column: int) -> bool:
        """Checks if any cell is outside the grid."""

        def valid_rows(y: int) -> bool:
            return 0 <= y < self.rows

        def valid_columns(x: int) -> bool:
            return 0 <= x < self.columns

        if valid_rows(row) and valid_columns(column):
            return False

        return True

    # -------------------------------------------------------------------------

    def is_row_full(self, row: int) -> bool:
        """Checks if a row is full or not."""

        for i in range(self.columns):
            if self.grid[row][i] == 0:
                return False

        return True

    # -------------------------------------------------------------------------

    def move_row_down(self, row: int, number_of_rows: int) -> None:
        """Moves the row down X number of rows."""

        for i in range(self.columns):
            self.grid[row + number_of_rows][i] = self.grid[row][i]
            self.grid[row][i] = 0

    # -------------------------------------------------------------------------
