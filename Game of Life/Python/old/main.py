# Old attempt. Decent reference to look back at.

from turtle import *
from cell import Cell
from tkinter import *

# ----------------------------------------------------------------------------------------------------------------------


# Constants. Change to edit parameters.
CELLS_PER_ROW = 50
GAP = 25
HEIGHT = CELLS_PER_ROW * GAP
WIDTH = CELLS_PER_ROW * GAP
SIZE_OFFSET = 100
BACKGROUND_COLOUR = "black"
BUTTON_WIDTH = 10
BUTTON_HEIGHT = 0
BUTTON_FONT = ("Arial", 20, "bold")


# ----------------------------------------------------------------------------------------------------------------------


def populate_grid(grid_width, grid_height, cells, gap):
    """
    Responsible for creating the cell objects and placing them in the grid.
    :param grid_width: int,
    :param grid_height: int,
    :param cells: int,
    :param gap: int,
    :return cell_list: list<cell objects>
    """
    # Cell list will hold all the grid rows which are in turn full of cell objects.
    cell_list = []

    # Starting place for grid creation. We want to start from the (width/2) * -1 of the given width and height to
    # create 'centred' cell block.
    start_x = (grid_width / 2) * -1
    start_y = (grid_height / 2) * -1

    # Dictates space between cells. Too close means cells overlap.
    incriment = gap

    for x in range(cells):
        # Store rows of cells
        new_row = []
        for y in range(cells):
            # Create a cell object for each cell in the grid
            new_cell = Cell(x=y, y=x)
            new_cell.penup()

            # Send the cell to the starting location
            new_cell.goto(start_x, start_y)

            # Add the increment before creating the next
            start_x += incriment
            new_row.append(new_cell)

        # Add the completed row to the cell list
        cell_list.append(new_row)

        # Increment the Y value, reset the X value (we create cells from left to right)
        start_y += incriment
        start_x -= incriment * cells

    return cell_list


# ----------------------------------------------------------------------------------------------------------------------


def reset_cells_button(cell_list):
    """
    Gives the reset button functionality. Calls the set_status() function on every cell in the grid allowing for a
    reset of the game state.
    :param cell_list:
    :return:
    """
    for row in cell_list:
        for cell in row:
            cell.set_status()


# ----------------------------------------------------------------------------------------------------------------------


def main():

    # Screen attributes
    screen = Screen()
    screen.tracer(
        False
    )  # Turns off animations. Need to invoke a screen update every cycle if False.
    screensize(WIDTH, HEIGHT)
    screen.setup(WIDTH + SIZE_OFFSET, HEIGHT + SIZE_OFFSET)
    screen.bgcolor(BACKGROUND_COLOUR)

    # Populate entire space with cells
    list_of_cells = populate_grid(WIDTH, HEIGHT, CELLS_PER_ROW, GAP)

    # Populate the list of cells.
    for row in list_of_cells:
        for cell in row:
            cell.find_neighbours(list_of_cells, GAP)

    # Place a button on the screen that allows the reset of game state
    canvas = screen.getcanvas()
    button = Button(
        canvas.master,
        text="Reset",
        command=lambda: reset_cells_button(list_of_cells),
        width=BUTTON_WIDTH,
        height=BUTTON_HEIGHT,
        font=BUTTON_FONT,
    )
    button.pack()
    button.place()

    on = True
    while on:
        screen.update()

        # Check the status of the neighbours from each cell's perspective
        i: int = 0
        while i < len(list_of_cells):
            j: int = 0
            while j < len(list_of_cells[0]):
                list_of_cells[i][j].check_neighbours()
                j += 1
            i += 1

        # Update the cells based upon alive/dead condition. We call this in its own loop because we want the game
        # state to be fully assessed each for "tick" of the game.
        i = 0
        while i < len(list_of_cells):
            j = 0
            while j < len(list_of_cells[0]):
                list_of_cells[i][j].change_status()
                j += 1
            i += 1

    screen.mainloop()


# ----------------------------------------------------------------------------------------------------------------------


if __name__ == "__main__":
    main()
