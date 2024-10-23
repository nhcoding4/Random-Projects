import pyray as rl


class Config:
    def __init__(self) -> None:
        self.window_width = 1000
        self.window_height = 1000
        self.title = "SunRay Effect Python"
        self.target_fps = 144
        self.particles = 300
        self.connection_thickness = 1.5
        self.connection_distance = 100
        self.current_fps = 0
        self.mouse_radius = 300
        self.ray_divisor = 3
        self.ray_opacity = 0.2
        self.init_effect()

    # ------------------------------------------------------------------------------------------------------------------

    def init_effect(self) -> None:
        rl.set_config_flags(rl.ConfigFlags.FLAG_WINDOW_RESIZABLE)
        rl.init_window(self.window_width, self.window_height, self.title)
        rl.set_target_fps(self.target_fps)

    # ------------------------------------------------------------------------------------------------------------------

    def fps_counter(self) -> None:
        rl.draw_text(f"{self.current_fps}", 0, 0, 40, rl.GREEN)

    # ------------------------------------------------------------------------------------------------------------------

    def update_fps(self) -> None:
        self.current_fps = rl.get_fps()

    # ------------------------------------------------------------------------------------------------------------------
