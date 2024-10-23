// 0.76s multi threaded
// 4.5s single threaded.

package main

import (
	"fmt"
	"sync"
	"time"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {
	start := time.Now()

	multi_threaded()
	//single_threaded()

	end := time.Since(start).Seconds()
	fmt.Println("Took:", end)
}

// --------------------------------------------------------------------------------------------------------------------

func multi_threaded() {
	ch := make(chan int32)
	pool := Pool{to_find: 250_001, workers: 24, ch: ch, waitgroup: sync.WaitGroup{}}
	go pool.Spin_Up()

	found_values := []int32{}
	for {
		value, open := <-ch
		if !open {
			break
		}
		found_values = append(found_values, value)
	}

	fmt.Println("Found primes:", len(found_values))
}

// --------------------------------------------------------------------------------------------------------------------

func is_prime(n int32) bool {
	// checks if a number is prime or not.
	for i := int32(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// --------------------------------------------------------------------------------------------------------------------

type Pool struct {
	to_find   int32
	workers   int32
	ch        chan int32
	waitgroup sync.WaitGroup
}

// --------------------------------------------------------------------------------------------------------------------

func (p *Pool) Process(start, stop int32) {
	// Work that each thread is going to do.
	defer p.waitgroup.Done()

	for i := start; i < stop; i++ {
		if is_prime(i) {
			p.ch <- i
		}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (p *Pool) Spin_Up() {
	// Spins up X workers and assigns them tasks.
	current_min := int32(0)
	tasks_per_worker := int32(p.to_find / p.workers)
	p.waitgroup.Add(int(p.workers))

	for i := int32(0); i < p.workers; i++ {
		go p.Process(current_min, current_min+tasks_per_worker)
		current_min += tasks_per_worker
	}

	p.waitgroup.Wait()
	close(p.ch)

}

// --------------------------------------------------------------------------------------------------------------------
