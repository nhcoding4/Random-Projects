import pyray as rl
from particle import Particle
from mouse import Mouse


class Sunray:
    def __init__(
        self,
        particles: list[Particle],
        mouse: Mouse,
        ray_divisor: int,
        opacity: float,
        thickness: float,
    ) -> None:
        self.particles = particles
        self.start_x: list[rl.Vector2] = []
        self.mouse = mouse
        self.ray_divisor = ray_divisor
        self.thickness = thickness
        self.color = rl.Color(255, 255, 255, int(255 * opacity))

    def calculate_rays(self) -> None:
        self.start_x = []
        for i, particle in enumerate(self.particles):
            if i % self.ray_divisor == 0:
                self.start_x.append(rl.Vector2(particle.x, particle.y))

    def draw_rays(self) -> None:
        endpoint = rl.Vector2(self.mouse.x, self.mouse.y)

        for ray_start in self.start_x:
            rl.draw_line_ex(
                ray_start,
                endpoint,
                self.thickness,
                self.color,
            )
