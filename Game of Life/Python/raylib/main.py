import pyray as rl
from grid import Grid


class Simulation:
    def __init__(self) -> None:
        self.window_width = 1920
        self.window_height = 1080
        self.cell_size = 5
        self.title = "Game of Life"
        self.target_fps = 144
        self.window_setup()

        self.background_color: rl.Color = rl.BLACK
        self.grid = Grid(self.window_width, self.window_height, self.cell_size)

    # -------------------------------------------------------------------------

    def window_setup(self) -> None:
        """Window setup. Must be ran before anything else."""

        rl.init_window(self.window_width, self.window_height, self.title)
        rl.set_target_fps(self.target_fps)

    # -------------------------------------------------------------------------

    def run(self) -> None:
        """Simulation mainloop."""

        while not rl.window_should_close():
            rl.begin_drawing()

            self.grid.update()

            rl.clear_background(self.background_color)
            self.grid.draw()
            rl.end_drawing()

        rl.close_window()

    # -------------------------------------------------------------------------


simulation = Simulation()
simulation.run()
