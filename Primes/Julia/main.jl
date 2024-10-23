function main()
    @time begin
        maximum = 250000
        foundPrimes::Vector{Int} = []

        for i = 2:maximum
            if isPrime(i)
                push!(foundPrimes, i)
            end
        end
        print(size(foundPrimes))
    end
end


function isPrime(number::Int)::Bool
    for i = 2:number-1
        if number % i == 0
            return false
        end
    end
    return true
end

main()


