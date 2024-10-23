def main():
    limit: int = 100
    squared_total: int = 0
    i_value_total: int = 0

    for i in range(limit):
        i_value_total += i
        squared_total += i**2

    i_value_total *= i_value_total

    print(i_value_total - squared_total)


if __name__ == "__main__":
    main()
