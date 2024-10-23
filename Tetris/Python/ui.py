import pyray as pr
import os
from helpers import colors
from blocks import *

# ----------------------------------------------------------------------------------------------------------------------


class UI:
    def __init__(self):
        self.base_font_size = 64
        self.font = pr.load_font_ex(
            (os.getcwd() + "/Fonts/font.ttf").encode(), self.base_font_size, None, 0
        )
        self.colors: list[pr.Color] = colors()
        self.font_size = self.base_font_size / 2

    # -------------------------------------------------------------------------

    def draw(self, game_over: bool, score: int, next_block: Block) -> None:
        """Draws the UI elements on the screen."""

        self.scoreboard(score)
        self.next_block(next_block)
        if game_over:
            self.game_over_text()

    # -------------------------------------------------------------------------

    def game_over_text(self) -> None:
        """Draws the game over text."""

        pr.draw_text_ex(
            self.font, "GAME OVER", pr.Vector2(320, 450), self.font_size, 2, pr.WHITE
        )

    # -------------------------------------------------------------------------

    def next_block(self, next_block: Block) -> None:
        """Elements used for the next block."""

        pr.draw_text_ex(
            self.font, "Next", pr.Vector2(365, 175), self.font_size, 2, pr.WHITE
        )
        pr.draw_rectangle_rounded(
            pr.Rectangle(320, 215, 170, 180), 0.3, 6, self.colors[8]
        )

        match next_block.id_code:
            case 3:
                next_block.draw(255, 290)
            case 4:
                next_block.draw(255, 280)
            case _:
                next_block.draw(270, 270)

    # -------------------------------------------------------------------------

    def scoreboard(self, score: int) -> None:
        """Elements used for the scoreboard."""

        pr.draw_text_ex(
            self.font, "Score", pr.Vector2(365, 15), self.font_size, 2, pr.WHITE
        )
        pr.draw_rectangle_rounded(
            pr.Rectangle(320, 55, 170, 60), 0.3, 6, self.colors[8]
        )

        score_text = str(score)
        score_length: pr.Vector2 = pr.measure_text_ex(
            self.font, score_text, self.font_size, 2
        )
        pr.draw_text_ex(
            self.font,
            score_text,
            pr.Vector2(320 + (170 - score_length.x) / 2, 65),
            self.font_size,
            2,
            pr.WHITE,
        )

    # -------------------------------------------------------------------------
