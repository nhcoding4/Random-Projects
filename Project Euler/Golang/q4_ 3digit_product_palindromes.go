package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

// --------------------------------------------------------------------------------------------------------------------

type Job struct {
	first_number  int
	second_number int
	total         int
}

func (job *Job) Calculate(ch chan Job, wg *sync.WaitGroup) {

	// Waitgroup
	defer wg.Done()

	// Get the total
	job.total = job.first_number * job.second_number

	// Convert the total to a string
	converted := strconv.Itoa(job.total)

	// Reverse the string
	var reversed string
	for i := len(converted) - 1; i >= 0; i-- {
		reversed += string(converted[i])
	}

	if converted == reversed {
		// Output the result over the channel
		ch <- Job{job.first_number, job.second_number, job.total}
	}
}

// --------------------------------------------------------------------------------------------------------------------

func question_4_palindrome() {
	// Highest Palindrome from the product of 3 numbers.

	job_channel := make(chan Job)

	go func() {
		var waitgroup sync.WaitGroup
		for i := 999; i >= 100; i-- {
			for j := 999; j >= 100; j-- {
				new_job := Job{first_number: i, second_number: j}
				waitgroup.Add(1)
				go new_job.Calculate(job_channel, &waitgroup)
			}
		}
		waitgroup.Wait()
		close(job_channel)
	}()

	results := []Job{}

	for {
		value, open := <-job_channel
		if !open {
			break
		}
		results = append(results, value)
	}

	// Sort results
	sort.Slice(results, func(i, j int) bool {
		return results[i].total < results[j].total
	})

	// Print the highest
	fmt.Println(results[len(results)-1])

}

// --------------------------------------------------------------------------------------------------------------------
