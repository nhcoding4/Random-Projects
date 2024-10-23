# Basic FizzBuzz Program

# ---------------------------------------------------------------------------------------------------------------------


def main():

    # Prompt for input
    limit = take_input("Enter a maximum number:\n")

    # Fizzbuzzing
    for i in range(1, limit + 1):
        to_print = ""
        if i % 5 == 0:
            if i % 3 == 0:
                to_print += "FizzBuzz"
            else:
                to_print += "Buzz"
        elif i % 3 == 0:
            to_print += "Fizz"
        else:
            to_print = i

        print(to_print)


# ---------------------------------------------------------------------------------------------------------------------


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
