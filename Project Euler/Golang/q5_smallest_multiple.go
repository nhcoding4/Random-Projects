package main

import (
	"fmt"
	"sync"
	"time"
)

// --------------------------------------------------------------------------------------------------------------------

// Quickest solution only using a single thead. Around 1-3s

func q5_smallest_multiple_single_thread() {
	// A solution only using a single threaded approach.

	start := time.Now()
	number := 1
	found := false

	// Continue checking until we find a number that meets the criteria.
	for !found {

		check(number, &found)
		number++

	}
	fmt.Println("Go 1.22.3 single-threaded:", time.Since(start))
}

func check(number int, found *bool) {
	// Check if the number can be divided evenly between all the relevant numbers.

	// Check for 11 to 20 is the same for checking for 1 - 20. 1-10 is composite of 11-20.
	for i := 11; i <= 20; i++ {
		// Stop checking if it doesn't divide evenly.
		if number%i != 0 {
			return
		}
	}

	// We have found a number that is evenly divided.
	fmt.Println("Found value:", number)
	*found = true
}

// --------------------------------------------------------------------------------------------------------------------
// Slower multi-threaded solution. 10-20s

var found bool = false
var max_threads int = 12
var buffer_size int = 1000

type Worker_Group struct {
	Jobs_in       chan int
	Total_workers int
	Wait_group    sync.WaitGroup
}

func (wg *Worker_Group) worker() {

	for {
		// Continue taking numbers from the work channel until it is closed
		number, open := <-wg.Jobs_in

		// If the channel is closed (number is found), signal the worker is done and break out the loop.
		if !open {
			wg.Wait_group.Done()
			break
		}

		func() {
			for i := 11; i <= 20; i++ {
				// Stop checking this number if it cannot be evenly divided
				if number%i != 0 {
					return
				}
			}
			// Found a number that can. Report it to the rest of the workers.
			fmt.Println("Found value:", number)
			found = true
		}()
	}
}

func (wg *Worker_Group) run() {
	// Add workers to the wait group
	wg.Wait_group.Add(wg.Total_workers)

	// Spin up X workers
	for i := 0; i < wg.Total_workers; i++ {
		go wg.worker()
	}

	// Wait for the workers to finish.
	wg.Wait_group.Wait()
}

// --------------------------------------------------------------------------------------------------------------------

func q5_smallest_multiple_multi_thread() {
	start := time.Now()

	// Buffer the channel to make sure we aren't filling up the job pool with an absurd amount of numbers
	// Around 1000 is the sweet spot after playing around with it.
	jobs_chan := make(chan int, buffer_size)

	// Create our worker group
	worker_group := Worker_Group{Jobs_in: jobs_chan, Total_workers: max_threads}

	// Pump numbers into the workgroup's job channel until we find a solution.
	go func() {
		number := 1
		for !found {

			jobs_chan <- number
			number++

		}
		close(jobs_chan)
	}()

	worker_group.run()
	fmt.Println("Multi-threaded:", time.Since(start))
}

// --------------------------------------------------------------------------------------------------------------------
