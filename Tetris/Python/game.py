from grid import *
from blocks import *
from ui import *
from sounds import *
from datetime import datetime, timedelta
import random
import pyray as pr


# ----------------------------------------------------------------------------------------------------------------------


class Game:
    def __init__(self) -> None:
        self.window_width = 500
        self.window_height = 620
        self.title = "Tetris"
        self.background_color = pr.Color(44, 44, 127, 255)
        self.targetFPS = 144
        self.setup_window()

        self.sound = Sounds()

        self.ui = UI()

        self.cell_size = 30
        self.grid = Grid(self.cell_size)

        self.blocks: list[Block] = self.fill_block_list()
        self.current_block: Block = self.get_block()
        self.next_block: Block = self.get_block()

        self.time_delta = 1.0
        self.auto_move_time: datetime = datetime.now()
        self.calculate_move_time()

        self.move_delta = 0.125
        self.next_move_delta: datetime = datetime.now()
        self.calculate_move_delta()

        self.game_over = False
        self.score = 0
        self.rows_cleared = 0

    # -------------------------------------------------------------------------

    def automatic_move_down(self) -> None:
        """Moves the block down after X amount of time has passed."""

        now: datetime = datetime.now()
        if now >= self.auto_move_time:
            self.move_block_down()
            self.calculate_move_time()

    # -------------------------------------------------------------------------

    def block_fits(self) -> bool:
        """Checks if the block fits in the grid after X movement."""

        cells: list[Position] = self.current_block.get_cell_positions()

        for cell in cells:
            if not self.grid.is_cell_empty(cell.row, cell.column):
                return False

        return True

    # -------------------------------------------------------------------------

    def calculate_move_delta(self) -> None:
        """Dictates time between user inputs."""

        self.next_move_delta: datetime = datetime.now() + timedelta(
            seconds=self.move_delta
        )

    # -------------------------------------------------------------------------

    def calculate_move_time(self) -> None:
        """Calculates the time in the future that the next automatic down
        movement should take place."""

        self.auto_move_time: datetime = datetime.now() + timedelta(
            seconds=self.time_delta
        )

    # -------------------------------------------------------------------------

    def draw(self) -> None:
        self.grid.draw_grid()
        self.current_block.draw(11, 11)
        self.ui.draw(self.game_over, self.score, self.next_block)

    # -------------------------------------------------------------------------

    def fill_block_list(self) -> list[Block]:
        """Returns a list of valid blocks to the caller."""
        return [
            IBlock(),
            JBlock(),
            LBlock(),
            OBlock(),
            SBlock(),
            TBlock(),
            ZBlock(),
        ]

    # -------------------------------------------------------------------------

    def get_block(self) -> Block:
        """Returns a random block to the caller."""

        choice = random.randint(0, len(self.blocks) - 1)
        block = self.blocks[choice]
        del self.blocks[choice]

        if len(self.blocks) == 0:
            self.blocks = self.fill_block_list()

        return block

    # -------------------------------------------------------------------------

    def handle_input(self) -> None:
        """Handles user input"""

        key_pressed: int = pr.get_key_pressed()
        if self.game_over and key_pressed:
            self.game_over = False
            self.reset()

        match key_pressed:
            case pr.KeyboardKey.KEY_UP:
                self.rotate_block()
            case pr.KeyboardKey.KEY_SPACE:
                self.snap_block_down()
            case _:
                pass

        now: datetime = datetime.now()

        if pr.is_key_down(pr.KeyboardKey.KEY_LEFT) and now > self.next_move_delta:
            self.move_block_left()
            self.calculate_move_delta()
        if pr.is_key_down(pr.KeyboardKey.KEY_RIGHT) and now > self.next_move_delta:
            self.move_block_right()
            self.calculate_move_delta()
        elif pr.is_key_down(pr.KeyboardKey.KEY_DOWN) and now > self.next_move_delta:
            self.move_block_down()
            self.calculate_move_delta()
            self.update_score(0, 1)

    # -------------------------------------------------------------------------

    def is_block_outside_grid(self) -> bool:
        """Checks if any of the cells, when moved would be outside the grid."""

        cells: list[Position] = self.current_block.get_cell_positions()

        for cell in cells:
            if self.grid.is_cell_outside_grid(cell.row, cell.column):
                return True

        return False

    # -------------------------------------------------------------------------

    def lock_block(self) -> None:
        """Locks the block in place by setting the color on the grid to the
        same as that block and selects a new block."""

        current_block: list[Position] = self.current_block.get_cell_positions()

        for cell in current_block:
            self.grid.grid[cell.row][cell.column] = self.current_block.id_code

        self.current_block = self.next_block
        if not self.block_fits():
            self.game_over = True

        self.next_block = self.get_block()

        rows_cleared: int = self.grid.clear_full_rows()

        if rows_cleared > 0:
            self.sound.play_clear_sound()
            self.update_score(rows_cleared, 0)
            self.rows_cleared += rows_cleared

        if self.rows_cleared >= 3:
            self.time_delta *= 0.9
            self.rows_cleared = 0

    # -------------------------------------------------------------------------

    def move_block_down(self) -> None:
        """Moves the block down 1 square."""

        if not self.game_over:
            self.current_block.move(1, 0)

            if self.is_block_outside_grid() or not self.block_fits():
                self.current_block.move(-1, 0)
                self.lock_block()

    # -------------------------------------------------------------------------

    def move_block_left(self) -> None:
        """Moves the block left 1 square."""

        if not self.game_over:
            self.current_block.move(0, -1)

            if self.is_block_outside_grid() or not self.block_fits():
                self.current_block.move(0, 1)

    # -------------------------------------------------------------------------

    def move_block_right(self) -> None:
        """Move the block right 1 square."""

        if not self.game_over:
            self.current_block.move(0, 1)

            if self.is_block_outside_grid() or not self.block_fits():
                self.current_block.move(0, -1)

    # -------------------------------------------------------------------------

    def rotate_block(self) -> None:
        """Moves the block state to its next rotation."""

        if not self.game_over:
            self.current_block.rotate_block()

            if self.is_block_outside_grid() or not self.block_fits():
                self.current_block.undo_rotation()
            else:
                self.sound.play_rotate_sound()

    # -------------------------------------------------------------------------

    def setup_window(self) -> None:
        """Window initialisation. Needs to run before anything else."""

        pr.init_window(self.window_width, self.window_height, self.title)
        pr.set_target_fps(self.targetFPS)

    # -------------------------------------------------------------------------

    def reset(self) -> None:
        """Resets the game state to the starting state."""

        self.grid.grid_setup()
        self.fill_block_list()
        self.current_block = self.get_block()
        self.next_block = self.get_block()
        self.score = 0

    # -------------------------------------------------------------------------

    def run(self) -> None:
        """All drawing and game state updates must be called here."""

        while not pr.window_should_close():
            pr.update_music_stream(self.sound.music)
            self.handle_input()
            self.automatic_move_down()

            pr.begin_drawing()

            pr.clear_background(self.background_color)
            self.draw()

            pr.end_drawing()

        self.sound.close_audio_device()
        pr.close_window()

    # -------------------------------------------------------------------------

    def snap_block_down(self) -> None:
        """Snaps the current block down as far as possible"""

        if not self.game_over:
            while True:
                self.update_score(0, 2)
                self.current_block.move(1, 0)

                if self.is_block_outside_grid() or not self.block_fits():
                    self.current_block.move(-1, 0)
                    self.lock_block()
                    break

    # -------------------------------------------------------------------------

    def update_score(self, lines_cleared: int, move_down_points: int) -> None:

        match lines_cleared:
            case 1:
                self.score += 100
            case 2:
                self.score += 300
            case 3:
                self.score += 500
            case _:
                pass

        self.score += move_down_points

    # -------------------------------------------------------------------------
