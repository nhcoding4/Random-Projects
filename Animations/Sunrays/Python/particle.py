import math

import pyray as rl
import random


class Particle:
    def __init__(self, window_width: int, window_height: int) -> None:
        self.window_width = window_width

        self.window_height = window_height
        self.radius = random.randint(5, 15)
        self.x = 0
        self.y = 0
        self.movementX = 0
        self.movementY = 0
        self.pushX = 0
        self.pushY = 0
        self.friction = 0.95
        self.power = 3

        self.set_initial_values()
        self.color = rl.color_from_hsv(self.x + self.y, 1.0, 1.0)

    # ------------------------------------------------------------------------------------------------------------------

    def draw(self) -> None:
        rl.draw_circle(self.x, self.y, self.radius, self.color)

    # ------------------------------------------------------------------------------------------------------------------

    def move(self) -> None:
        self.x += int(round(self.movementX + self.pushX, 0))
        self.y += int(round(self.movementY + self.pushY, 0))

    # ------------------------------------------------------------------------------------------------------------------

    def check_bounds(self) -> None:
        if self.x < self.radius:
            self.x = self.radius
            self.movementX *= -1

        if self.x > self.window_width - self.radius:
            self.x = self.window_width - self.radius
            self.movementX *= -1

        if self.y < self.radius:
            self.y = self.radius
            self.movementY *= -1

        if self.y > self.window_height - self.radius:
            self.y = self.window_height - self.radius
            self.movementY *= -1

    # ------------------------------------------------------------------------------------------------------------------

    def push_particle(self, mouse_x: int, mouse_y: int, mouse_radius: int) -> None:
        dx: float = self.x - mouse_x
        dy: float = self.y - mouse_y
        distance: float = math.hypot(dx, dy)

        if distance < mouse_radius:
            angle: float = math.atan2(dy, dx)
            self.pushX = math.cos(angle) * self.power
            self.pushY = math.sin(angle) * self.power

    # ------------------------------------------------------------------------------------------------------------------

    def set_initial_values(self) -> None:
        self.x = random.randint(self.radius, self.window_width - self.radius)
        self.y = random.randint(self.radius, self.window_height - self.radius)
        self.movementX = random.randint(-4, 4)
        self.movementY = random.randint(-4, 4)

    # ------------------------------------------------------------------------------------------------------------------

    def update(
        self, mouse_x: int, mouse_y: int, mouse_radius: int, mouse_pressed: bool
    ) -> None:

        if mouse_pressed:
            self.push_particle(mouse_x, mouse_y, mouse_radius)
        self.move()
        self.check_bounds()
        self.color = rl.color_from_hsv(self.x + self.y, 1.0, 1.0)
        self.update_push()

    # ------------------------------------------------------------------------------------------------------------------

    def update_push(self) -> None:
        self.pushX *= self.friction
        self.pushY *= self.friction

    # ------------------------------------------------------------------------------------------------------------------
