# Constants

MIN_NUMBER = 1
MAX_NUMBER = 9


# ----------------------------------------------------------------------------------------------------------------------


def main() -> None:
    unsolved_board: list[list[int]] = board()
    print_board(unsolved_board)
    solve(unsolved_board)
    print_board(unsolved_board)


# ----------------------------------------------------------------------------------------------------------------------


def solve(sodoku_board: list[list[int]]) -> bool:
    """A recursive function that will place numbers in a valid state."""

    find: tuple[int, int] = find_empty_cell(sodoku_board)

    # Base case. The board is full (no empty cells).
    if not find:
        return True
    else:
        row, column = find

    for i in range(MIN_NUMBER, MAX_NUMBER + 1):
        # Add valid entries to the board.
        if valid(sodoku_board, i, (row, column)):
            sodoku_board[row][column] = i

            # Keep recursively trying to find a valid solution.
            if solve(sodoku_board):
                return True

            # Undo the previous change to the board state as it is not resulting in a valid board.
            sodoku_board[row][column] = 0

    return False


# ----------------------------------------------------------------------------------------------------------------------


def find_empty_cell(sodoku_board: list[list[int]]) -> tuple[int, int] | None:
    """Returns the x and y coordinates of an empty cell which is denoted by a 0."""

    for i, row in enumerate(sodoku_board):
        for j, number in enumerate(row):
            if number == 0:
                return i, j

    return None


# ----------------------------------------------------------------------------------------------------------------------


def valid(
    sodoku_board: list[list[int]], number: int, positions: tuple[int, int]
) -> bool:
    """Checks for a valid board state. Returns false if the action is invalid."""

    # Check each element to see if the same already exists on the X axis.
    for i, _ in enumerate(sodoku_board[0]):
        if sodoku_board[positions[0]][i] == number and positions[1] != i:
            return False

    # Check the column to see if the number already exists on the Y axis.
    for i, _ in enumerate(sodoku_board):
        if sodoku_board[i][positions[1]] == number and positions[0] != i:
            return False

    # Check the 3 x 3 grid to see if the number already exists within it.
    box_x: int = positions[1] // 3
    box_y: int = positions[0] // 3

    # Keep our search within the 3 x 3 grid.
    for i in range(box_y * 3, box_y * 3 + 3):
        for j in range(box_x * 3, box_x * 3 + 3):
            if sodoku_board[i][j] == number and (i, j) != positions:
                return False

    return True


# ----------------------------------------------------------------------------------------------------------------------


def board() -> list[list[int]]:
    """A sodoku board. Empty cells are represented by a 0."""

    soduku_board: list[list[int]] = [
        [7, 8, 0, 4, 0, 0, 1, 2, 0],
        [6, 0, 0, 0, 7, 5, 0, 0, 9],
        [0, 0, 0, 6, 0, 1, 0, 7, 8],
        [0, 0, 7, 0, 4, 0, 2, 6, 0],
        [0, 0, 1, 0, 5, 0, 9, 3, 0],
        [9, 0, 4, 0, 6, 0, 0, 0, 5],
        [0, 7, 0, 3, 0, 0, 0, 1, 2],
        [1, 2, 0, 0, 0, 7, 4, 0, 0],
        [0, 4, 9, 2, 0, 6, 0, 0, 7],
    ]
    return soduku_board


# ----------------------------------------------------------------------------------------------------------------------


def print_board(sodoku_board: list[list[int]]) -> None:
    """Prints the contents of the sodoku board."""

    # -----------------------------------------------------

    # Helpers

    def print_dashes(row_length: int) -> None:
        for i in range(row_length):
            if i == 0:
                print(" -", end="")
            print("---", end="")
        print()

    def print_pipe() -> None:
        print(" | ", end="")

    # -----------------------------------------------------

    # Print Board

    row_length: int = len(sodoku_board[0]) + len(sodoku_board[0]) // 3

    for i, row in enumerate(sodoku_board):
        if i % 3 == 0:
            print_dashes(row_length)

        for j, number in enumerate(row):
            if j % 3 == 0:
                print_pipe()
            if number == 0:
                print(" - ", end="")
            else:
                print(f" {number} ", end="")

        print_pipe()
        print()

    print_dashes(row_length)


# ----------------------------------------------------------------------------------------------------------------------


if __name__ == "__main__":
    main()
