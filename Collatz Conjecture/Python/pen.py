from typing import List
from turtle import *
import random


# ----------------------------------------------------------------------------------------------------------------------


class Pen(Turtle):
    def __init__(self, screen_width: int) -> None:
        super().__init__()

        self.screen_width: int = screen_width
        self.starting_place = (self.screen_width//2) * -1, 0

        self.hideturtle()
        self.set_colour()
        self.starting_pos()

        # ----------------------------------------------------------

    def draw_sequence(self, collatz_sequence: List[int]) -> None:
        """ Graphs the collatz sequence """

        x_move: int = self.screen_width // len(collatz_sequence)
        last_pos: int = self.starting_place[0]

        for i, value in enumerate(collatz_sequence):
            last_pos = last_pos + x_move
            self.goto(last_pos, collatz_sequence[i] * 2)

        self.starting_pos()
        self.set_colour()

        # ----------------------------------------------------------

    def set_colour(self) -> None:
        """ Sets the colour of the pen. """

        colours: List[str] = ["red", "green", "blue", "orange", "yellow", "purple", "black"]
        self.pencolor(random.choice(colours))

        # ----------------------------------------------------------

    def starting_pos(self):
        """ Sets the starting position of the pen. """

        self.penup()
        self.goto(self.starting_place)
        self.pendown()
