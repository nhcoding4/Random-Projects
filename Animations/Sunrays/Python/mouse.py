import pyray as rl


class Mouse:
    def __init__(self, radius: int) -> None:
        self.x = 0
        self.y = 0
        self.radius = radius
        self.pressed = False

    # ------------------------------------------------------------------------------------------------------------------

    def update(self) -> None:
        self.x = rl.get_mouse_x()
        self.y = rl.get_mouse_y()

    # ------------------------------------------------------------------------------------------------------------------

    def update_pressed(self) -> None:
        if rl.is_mouse_button_down(rl.MouseButton.MOUSE_BUTTON_LEFT):
            self.pressed = True
        else:
            self.pressed = False

    # ------------------------------------------------------------------------------------------------------------------
