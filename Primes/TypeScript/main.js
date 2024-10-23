function main() {
    var startTime = performance.now();
    var foundPrimes = [];
    var limit = 250001;
    for (var i = 2; i < limit; i++) {
        if (isPrime(i)) {
            foundPrimes.push(i);
        }
    }
    var endTime = performance.now();
    console.log("Found ".concat(foundPrimes.length, " prime numbers"));
    console.log("Time ".concat((endTime - startTime) / 1000, " seconds"));
}
function isPrime(integer) {
    for (var i = 2; i < integer; i++) {
        if (integer % i == 0) {
            return false;
        }
    }
    return true;
}
main();
