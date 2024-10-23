import pyray as rl
import random


class Grid:
    def __init__(self, screen_width: int, screen_height: int, cell_size: int) -> None:
        self.cell_size = cell_size
        self.screen_width = screen_width
        self.screen_height = screen_height
        self.rows = int(self.screen_height / self.cell_size)
        self.columns = int(self.screen_width / self.cell_size)
        self.colors: list[rl.Color] = [rl.BLACK, rl.WHITE]

        self.grid: list[list[int]] = []
        self.grid_setup()

        self.offsets = [
            (-1, 0),
            (1, 0),
            (0, -1),
            (0, 1),
            (-1, -1),
            (-1, 1),
            (1, -1),
            (1, 1),
        ]

        self.set_initial_status()

    # -------------------------------------------------------------------------

    def count_live_neighbors(self, row: int, column: int) -> int:
        """Sums the total alive neighbors in for a given cell."""

        live_neighbors = 0

        for offset in self.offsets:
            neighbor_row = (row + offset[1] + self.rows) % self.rows
            neighbor_column = (column + offset[0] + self.columns) % self.columns
            live_neighbors += self.grid[neighbor_row][neighbor_column]

        return live_neighbors

    # -------------------------------------------------------------------------

    def draw(self) -> None:
        """Draws the cells. Black if dead, white if alive."""

        for i, row in enumerate(self.grid):
            for j, column in enumerate(row):
                rl.draw_rectangle(
                    self.cell_size * j,
                    self.cell_size * i,
                    self.cell_size - 1,
                    self.cell_size - 1,
                    self.colors[column],
                )

    # -------------------------------------------------------------------------

    def set_initial_status(self) -> None:
        """Gives the grid a random status"""

        for i, row in enumerate(self.grid):
            for j, _ in enumerate(row):
                choice = random.randint(0, 5)
                if choice < 2:
                    self.grid[i][j] = 1
                else:
                    self.grid[i][j] = 0

    # -------------------------------------------------------------------------

    def grid_setup(self) -> None:
        """Fills the grid with 0's"""

        for _ in range(self.rows):
            new_row: list[int] = []

            for _ in range(self.columns):
                new_row.append(0)

            self.grid.append(new_row)

    # -------------------------------------------------------------------------

    def update(self) -> None:
        """Creates a new grid state."""

        new_grid: list[list[int]] = []

        for i, row in enumerate(self.grid):
            new_row: list[int] = []

            for j, row in enumerate(row):
                alive_neighbors = self.count_live_neighbors(i, j)

                match alive_neighbors:
                    case 3:
                        new_row.append(1)
                    case 2:
                        new_row.append(self.grid[i][j])
                    case _:
                        new_row.append(0)

            new_grid.append(new_row)

        self.grid = new_grid

    # -------------------------------------------------------------------------
