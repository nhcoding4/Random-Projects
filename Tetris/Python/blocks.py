import pyray as pr
from position import *
from helpers import colors


# ----------------------------------------------------------------------------------------------------------------------


class Block:
    def __init__(self) -> None:
        """Parent block for other blocks. Should not be used itself."""
        self.id_code = 0
        self.positions: dict[int, list[Position]] = {}
        self.cell_size = 30
        self.rotation = 0
        self.colors: list[pr.Color] = colors()
        self.offset_rows = 0
        self.offset_columns = 0
        self.move(0, 3)

    # -------------------------------------------------------------------------

    def draw(self, offset_x: int, offset_y: int) -> None:
        """Draws the block onto the grid."""

        for item in self.get_cell_positions():
            pr.draw_rectangle(
                item.column * self.cell_size + offset_x,
                item.row * self.cell_size + offset_y,
                self.cell_size - 1,
                self.cell_size - 1,
                self.colors[self.id_code],
            )

    # -------------------------------------------------------------------------

    def move(self, rows: int, columns: int) -> None:
        """Moves the block around the grid."""

        self.offset_rows += rows
        self.offset_columns += columns

    # -------------------------------------------------------------------------

    def get_cell_positions(self) -> list[Position]:
        """Calculates the tile positions considering the current offsets."""

        current_positions: list[Position] = self.positions[self.rotation]
        moved_tiles: list[Position] = []

        for tile in current_positions:
            new_position = Position(
                tile.row + self.offset_rows, tile.column + self.offset_columns
            )
            moved_tiles.append(new_position)

        return moved_tiles

    # -------------------------------------------------------------------------

    def rotate_block(self) -> None:
        """Changes the block to its next state."""

        self.rotation += 1
        if self.rotation > len(self.positions) - 1:
            self.rotation = 0

    # -------------------------------------------------------------------------

    def undo_rotation(self) -> None:
        """Reverts to the previous rotation."""

        self.rotation -= 1
        if self.rotation < 0:
            self.rotation = len(self.positions) - 1


# ----------------------------------------------------------------------------------------------------------------------


class LBlock(Block):
    def __init__(self) -> None:
        super().__init__()
        self.id_code = 1
        self.fill_positions()

    # -------------------------------------------------------------------------

    def fill_positions(self) -> None:
        """Defines the position of all the block's states."""

        self.positions[0] = [
            Position(0, 2),
            Position(1, 0),
            Position(1, 1),
            Position(1, 2),
        ]
        self.positions[1] = [
            Position(0, 1),
            Position(1, 1),
            Position(2, 1),
            Position(2, 2),
        ]
        self.positions[2] = [
            Position(1, 0),
            Position(1, 1),
            Position(1, 2),
            Position(2, 0),
        ]
        self.positions[3] = [
            Position(0, 0),
            Position(0, 1),
            Position(1, 1),
            Position(2, 1),
        ]


# ----------------------------------------------------------------------------------------------------------------------


class JBlock(Block):
    def __init__(self) -> None:
        super().__init__()
        self.id_code = 2
        self.fill_positions()

    # -------------------------------------------------------------------------

    def fill_positions(self) -> None:
        """Defines the position of all the block's states."""

        self.positions[0] = [
            Position(0, 0),
            Position(1, 0),
            Position(1, 1),
            Position(1, 2),
        ]
        self.positions[1] = [
            Position(0, 1),
            Position(0, 2),
            Position(1, 1),
            Position(2, 1),
        ]
        self.positions[2] = [
            Position(1, 0),
            Position(1, 1),
            Position(1, 2),
            Position(2, 2),
        ]
        self.positions[3] = [
            Position(0, 1),
            Position(1, 1),
            Position(2, 0),
            Position(2, 1),
        ]


# ----------------------------------------------------------------------------------------------------------------------


class IBlock(Block):
    def __init__(self) -> None:
        super().__init__()
        self.id_code = 3
        self.fill_positions()
        self.move(-1, 0)

    # -------------------------------------------------------------------------

    def fill_positions(self) -> None:
        """Defines the position of all the block's states."""

        self.positions[0] = [
            Position(1, 0),
            Position(1, 1),
            Position(1, 2),
            Position(1, 3),
        ]
        self.positions[1] = [
            Position(0, 2),
            Position(1, 2),
            Position(2, 2),
            Position(3, 2),
        ]
        self.positions[2] = [
            Position(2, 0),
            Position(2, 1),
            Position(2, 2),
            Position(2, 3),
        ]
        self.positions[3] = [
            Position(0, 1),
            Position(1, 1),
            Position(2, 1),
            Position(3, 1),
        ]


# ----------------------------------------------------------------------------------------------------------------------


class OBlock(Block):
    def __init__(self) -> None:
        super().__init__()
        self.id_code = 4
        self.fill_positions()
        self.move(0, 1)

    # -------------------------------------------------------------------------

    def fill_positions(self) -> None:
        """Defines the position of all the block's states."""

        self.positions[0] = [
            Position(0, 0),
            Position(0, 1),
            Position(1, 0),
            Position(1, 1),
        ]


# ----------------------------------------------------------------------------------------------------------------------


class SBlock(Block):
    def __init__(self) -> None:
        super().__init__()
        self.id_code = 5
        self.fill_positions()

    # -------------------------------------------------------------------------

    def fill_positions(self) -> None:
        """Defines the position of all the block's states."""

        self.positions[0] = [
            Position(0, 1),
            Position(0, 2),
            Position(1, 0),
            Position(1, 1),
        ]
        self.positions[1] = [
            Position(0, 1),
            Position(1, 1),
            Position(1, 2),
            Position(2, 2),
        ]
        self.positions[2] = [
            Position(1, 1),
            Position(1, 2),
            Position(2, 0),
            Position(2, 1),
        ]
        self.positions[3] = [
            Position(0, 0),
            Position(1, 0),
            Position(1, 1),
            Position(2, 1),
        ]


# ----------------------------------------------------------------------------------------------------------------------


class TBlock(Block):
    def __init__(self) -> None:
        super().__init__()
        self.id_code = 6
        self.fill_positions()

    # -------------------------------------------------------------------------

    def fill_positions(self) -> None:
        """Defines the position of all the block's states."""

        self.positions[0] = [
            Position(0, 1),
            Position(1, 0),
            Position(1, 1),
            Position(1, 2),
        ]
        self.positions[1] = [
            Position(0, 1),
            Position(1, 1),
            Position(1, 2),
            Position(2, 1),
        ]
        self.positions[2] = [
            Position(1, 0),
            Position(1, 1),
            Position(1, 2),
            Position(2, 1),
        ]
        self.positions[3] = [
            Position(0, 1),
            Position(1, 0),
            Position(1, 1),
            Position(2, 1),
        ]


# ----------------------------------------------------------------------------------------------------------------------


class ZBlock(Block):
    def __init__(self) -> None:
        super().__init__()
        self.id_code = 7
        self.fill_positions()

    # -------------------------------------------------------------------------

    def fill_positions(self) -> None:
        """Defines the position of all the block's states."""

        self.positions[0] = [
            Position(0, 0),
            Position(0, 1),
            Position(1, 1),
            Position(1, 2),
        ]
        self.positions[1] = [
            Position(0, 2),
            Position(1, 1),
            Position(1, 2),
            Position(2, 1),
        ]
        self.positions[2] = [
            Position(1, 0),
            Position(1, 1),
            Position(2, 1),
            Position(2, 2),
        ]
        self.positions[3] = [
            Position(0, 1),
            Position(1, 0),
            Position(1, 1),
            Position(2, 0),
        ]


# ----------------------------------------------------------------------------------------------------------------------
