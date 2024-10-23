import random
from typing import List

# --------------------------------------------------------------------------------------------------------------------


class Person:
    def __init__(self, name: str):
        self.Name = name
        self.Number = 0

    def get_number(self):
        self.Number = random.randint(1, 101)


# --------------------------------------------------------------------------------------------------------------------


def main():

    # Populate the player list.
    players = populate_players(take_input("Enter the number of players:"))

    # Keep running rounds until there is only 1 player remaining.
    current_round = 1
    while len(players) > 1:
        print(f"Round {current_round}\n-------")
        players = round(players)
        current_round += 1

    # Display winner
    print(f"\n{players[0].Name} is the winner!\n----------------------")


# --------------------------------------------------------------------------------------------------------------------


def round(players: List[Person]) -> List[Person]:
    """Logic for each round. Pits players against each other and returns the winners."""

    winners = []

    # Until there is 1 or 0 players left in the players list, pair off players and get a winner.
    while True:
        if len(players) <= 1:
            break
        winners.append(winner(players.pop(), players.pop()))

    # Add any leftover players (needed for non multiples of 4 as it creates a non symmetrical tree).
    while len(players) > 0:
        winners.append(players.pop())

    return winners


# --------------------------------------------------------------------------------------------------------------------


def populate_players(amount: int) -> List[Person]:
    """Creates a list of X amount of Person objects"""

    player_list = []

    # Add player to the list
    for i in range(amount):
        player_list.append(Person(f"Player {i + 1}"))

    return player_list


# --------------------------------------------------------------------------------------------------------------------


def winner(player_1: Person, player_2: Person) -> Person:
    """Selects a winner (who has the highest number) out of 2 players"""

    # Keep selecting numbers until the numbers are not equal to one another.
    while True:
        player_1.get_number()
        player_2.get_number()

        if player_1.Number != player_2.Number:
            break

    print(
        f"{player_1.Name} rolls {player_1.Number} VS {player_2.Name} rolls {player_2.Number}"
    )

    # Return the player with the highest number
    if player_1.Number > player_2.Number:
        print(f"{player_1.Name} wins!\n")
        return player_1

    print(f"{player_2.Name} wins!\n")
    return player_2


# --------------------------------------------------------------------------------------------------------------------


def take_input(prompt: str) -> int:
    """Prompts user for input until a valid integer has been entered"""

    while True:
        try:
            return int(input(prompt))
        except Exception as _:
            print("Invalid input. Please enter a number.")


# ---------------------------------------------------------------------------------------------------------------------


if __name__ == "__main__":
    main()
