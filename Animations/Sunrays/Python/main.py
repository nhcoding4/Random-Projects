import pyray as rl
from config import Config
from effect import Effect
from sunray import Sunray


def main() -> None:
    config = Config()
    effect = Effect(
        config.window_width,
        config.window_height,
        config.particles,
        config.connection_distance,
        config.connection_thickness,
        config.mouse_radius,
    )
    sunray = Sunray(
        effect.particles,
        effect.mouse,
        config.ray_divisor,
        config.ray_opacity,
        config.connection_thickness,
    )

    while not rl.window_should_close():
        if rl.is_window_resized():
            config.window_width = rl.get_screen_width()
            config.window_height = rl.get_screen_height()
            effect.resize_screen(config.window_width, config.window_height)

        effect.update_particles()
        config.update_fps()
        sunray.calculate_rays()

        rl.begin_drawing()
        rl.clear_background(rl.BLACK)
        sunray.draw_rays()
        effect.draw_particles()
        config.fps_counter()
        rl.end_drawing()

    rl.close_window()


if __name__ == "__main__":
    main()
