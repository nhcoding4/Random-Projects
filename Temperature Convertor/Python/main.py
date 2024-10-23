import os


# ---------------------------------------------------------------------------------------------------------------------


def main():

    while True:        
        # Display options on the screen
        print("\nOptions\n-------")
        print("1) Convert Celsius to Fahrenheit")
        print("2) Convert Fahrenheit to Celsius")
        print("3) Exit Program")
        
        # Keep prompting user until they enter a valid option
        while True:
            user_choice = take_input("Select an option: ")
            if 0 < user_choice <= 3:
                break
            else:
                print("Invalid input")
        
        # Convert user input into Fahrenheit, clear the screen when done.
        if user_choice == 1:
            print(f"{celsius_farhrenheit(take_input("Enter a value in Celsius: "))}F")
            pause()
            clear_screen()

        # Convert user input into Celsius, clear the screen when done.
        elif user_choice == 2:
            print(f"{fahrenheit_celsius(take_input("Enter a value in Farhrenheit: "))}C")
            pause()
            clear_screen()

        # Exit the program, clear the screen when done.
        else:
            clear_screen()
            os._exit(1)


# ---------------------------------------------------------------------------------------------------------------------


def take_input(prompt: str) -> float:
    """Makes the use enter a valid number"""

    while True:
        try:
            return float(input(prompt))
        except Exception as _:
            print("Invalid input. Please enter a number.")


# ---------------------------------------------------------------------------------------------------------------------


def celsius_farhrenheit(temp: float) -> float:
    """Converts a number into celsius"""

    return (temp * 1.8) + 32.0


# ---------------------------------------------------------------------------------------------------------------------


def fahrenheit_celsius(temp: float) -> float:
    """Converts a number into Fahrenheit"""

    return (temp - 32) / 1.8


# ---------------------------------------------------------------------------------------------------------------------


def pause() -> None:
    """Used to pause the program until the use wants it to continue"""

    _ = input("Press 'ENTER' to continue")


# ---------------------------------------------------------------------------------------------------------------------


def clear_screen() -> None:
    """Clears the console of all text"""

    os.system("cls" if os.name == "nt" else "clear")


 # ---------------------------------------------------------------------------------------------------------------------


if __name__ == "__main__":
    main()
