function main() {
    const startTime = performance.now()

    const foundPrimes = []
    const limit = 250001

    for (let i = 2; i < limit; i++) {
        if (isPrime(i)) {
            foundPrimes.push(i)
        }
    }

    const endTime = performance.now()
    console.log(`Found: ${foundPrimes.length} prime numbers`)
    console.log(`Time: ${(endTime - startTime) / 1000} seconds`)

}

function isPrime(number) {
    for (let i = 2; i < number; i++) {
        if (number % i === 0) {
            return false
        }
    }
    return true
}

main()