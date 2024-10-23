import pyray as pr

# ----------------------------------------------------------------------------------------------------------------------


def colors() -> list[pr.Color]:
    """All the colors that are used by the grid and the blocks."""

    dark_grey = pr.Color(26, 31, 40, 255)
    green = pr.Color(47, 230, 23, 255)
    red = pr.Color(232, 18, 18, 255)
    orange = pr.Color(226, 116, 17, 255)
    yellow = pr.Color(237, 234, 4, 255)
    purple = pr.Color(116, 0, 247, 255)
    cyan = pr.Color(21, 204, 209, 255)
    blue = pr.Color(13, 64, 216, 255)
    light_blue = pr.Color(59, 85, 162, 255)
    dark_blue = pr.Color(44, 44, 127, 255)

    return [
        dark_grey,
        green,
        red,
        orange,
        yellow,
        purple,
        cyan,
        blue,
        light_blue,
        dark_blue,
    ]


# ----------------------------------------------------------------------------------------------------------------------
