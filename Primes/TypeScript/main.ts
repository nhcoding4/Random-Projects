function main() {
    const startTime = performance.now()
    const foundPrimes: Array<number> = []
    const limit = 250001

    for (let i = 2; i < limit; i++) {
        if (isPrime(i)) {
            foundPrimes.push(i)
        }
    }

    const endTime = performance.now()
    console.log(`Found ${foundPrimes.length} prime numbers`)
    console.log(`Time ${(endTime - startTime) / 1000} seconds`)

}


function isPrime(integer: number): boolean {
    for (let i = 2; i < integer; i++) {
        if (integer % i == 0) {
            return false
        }
    }
    return true
}

main()
