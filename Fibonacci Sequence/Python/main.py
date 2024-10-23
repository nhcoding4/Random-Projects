from typing import List


def main():
    total_numbers: int = take_integer_input("Please enter the amount of Fibonacci numbers you want to calculate: ")

    fib_sequence: List[int] = [1, 1]
    for i in range(total_numbers):
        fib_sequence.append(fib_sequence[i] + fib_sequence[i+1])

    print(f"Fibonacci numbers: {fib_sequence}")


def take_integer_input(prompt: str) -> int:
    """ Takes input from the user and makes sure they enter a valid integer. """
    while True:
        try:
            user_input = int(input(prompt))
            if user_input > 1:
                return user_input
            print("Invalid input. Please enter a positive integer over 1.")
        except ValueError:
            print("Invalid input. Please enter a positive integer over 1.")


if __name__ == "__main__":
    main()
