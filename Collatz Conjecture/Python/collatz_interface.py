from turtle import *
from pen import Pen
from typing import List

# ----------------------------------------------------------------------------------------------------------------------


class CollatzInterface(Turtle):
    def __init__(self, screen_width: int, screen_height: int) -> None:
        super().__init__()

        self.width: int = screen_width
        self.sequences: List[List[int]] = []
        self.height: int = screen_height

        self.pen = Pen(screen_width=self.width)
        self.screen = Screen()
        self.screen.screensize(self.width, self.height)
        self.screen.bgcolor("grey")
        self.screen.tracer(False)

        self.make_collatz()
        self.draw()
        self.screen.mainloop()

    # ----------------------------------------------------------

    def make_collatz(self):
        """ Attempts to make an array of collatz sequences from a text file of numbers. """

        try:
            with open("collatz.txt", "r") as file:
                targets = file.readlines()

            for _, number in enumerate(targets):

                sequence: List[int] = []
                number = int(number)

                while number > 1:
                    if number % 2 == 0:
                        number /= 2
                    else:
                        number = number * 3 + 1
                    sequence.append(int(number))

                self.sequences.append(sequence)
                self.screen.update()

        except Exception as e:
            print(f"CollatzPlotter -> make_collatz: there was an error: {e}")

    # ----------------------------------------------------------

    def draw(self):
        """ Tells the pen object to draw the sequence. """

        for i, sequence in enumerate(self.sequences):
            self.pen.draw_sequence(sequence)

