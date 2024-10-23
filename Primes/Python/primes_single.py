def isPrime(i: int) -> bool:
    for x in range(2, i):
        if i % x == 0:
            return False
    return True


def main() -> None:
    foundPrimes: list[int] = []

    for i in range(2, 250001):
        if isPrime(i):
            foundPrimes.append(i)

    print(f"Found: {len(foundPrimes)} primes")


main()
