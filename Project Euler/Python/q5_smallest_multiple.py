from datetime import datetime

# Only a mere 15-20 times slower than Go/C on replacing for loops with while loops.
# Same(ish) speed as C when using pypy3. I have my doubts the timing accurate given it take some time to display an output.


def main():
    start: datetime = datetime.now()

    found: bool = False
    number: int = 1

    while not found:
        found = check_number(number)
        number += 1

    end: datetime = datetime.now()
    print(f"Python 3.10.14 pypy3: {end-start}")


def check_number(number: int) -> bool:
    i: int = 11
    while i <= 20:
        if number % i != 0:
            return False
        i += 1

    print("Found value:", number)
    return True


if __name__ == "__main__":
    main()
