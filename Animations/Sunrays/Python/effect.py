import math
from mouse import Mouse
from particle import Particle
import pyray as rl


class Effect:
    def __init__(
        self,
        window_width: int,
        window_height: int,
        total_particles: int,
        link_distance: int,
        line_thickness: float,
        mouse_radius: int,
    ) -> None:
        self.window_width = window_width
        self.window_height = window_height
        self.total_particles = total_particles
        self.particles: list[Particle] = []
        self.create_particles()

        self.link_distance = link_distance
        self.line_thickness = line_thickness
        self.line_start: list[rl.Vector2] = []
        self.line_end: list[rl.Vector2] = []
        self.line_color: list[rl.Color] = []

        self.mouse = Mouse(mouse_radius)

    # ------------------------------------------------------------------------------------------------------------------

    def calculate_connections(self) -> None:
        for i, _ in enumerate(self.particles):
            for j in range(i, len(self.particles)):
                if i == j:
                    continue
                dx = self.particles[i].x - self.particles[j].x
                dy = self.particles[i].y - self.particles[j].y
                distance = math.hypot(dx, dy)

                if distance < self.link_distance:
                    self.line_start.append(
                        rl.Vector2(self.particles[i].x, self.particles[i].y)
                    )
                    self.line_end.append(
                        rl.Vector2(self.particles[j].x, self.particles[j].y)
                    )
                    opacity = 1 - (distance / self.link_distance)
                    self.line_color.append(rl.Color(255, 255, 255, int(255 * opacity)))

    # ------------------------------------------------------------------------------------------------------------------

    def create_particles(self) -> None:
        for _ in range(self.total_particles):
            self.particles.append(Particle(self.window_width, self.window_height))

    # ------------------------------------------------------------------------------------------------------------------

    def draw_particles(self) -> None:
        for i, line in enumerate(self.line_start):
            rl.draw_line_ex(
                self.line_start[i],
                self.line_end[i],
                self.line_thickness,
                self.line_color[i],
            )

        for particle in self.particles:
            particle.draw()

    # ------------------------------------------------------------------------------------------------------------------

    def resize_screen(self, new_width: int, new_height: int) -> None:
        self.window_width = new_width
        self.window_height = new_height

        for particle in self.particles:
            particle.window_width = self.window_width
            particle.window_height = self.window_height
            particle.set_initial_values()

    # ------------------------------------------------------------------------------------------------------------------

    def update_particles(self) -> None:
        self.line_start = []
        self.line_end = []
        self.line_color = []

        self.calculate_connections()

        self.mouse.update_pressed()
        if self.mouse.pressed:
            self.mouse.update()

        for particle in self.particles:
            particle.update(
                self.mouse.x, self.mouse.y, self.mouse.radius, self.mouse.pressed
            )

    # ------------------------------------------------------------------------------------------------------------------
