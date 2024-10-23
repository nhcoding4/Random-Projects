// Manages go routines to allow for concurrent building of data.

package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"strings"
	"sync"
)

// --------------------------------------------------------------------------------------------------------------------

func collect_data() Person_CV {
	// Manages go routines, calls functions that loads/collects data.

	// -----------

	// Check for errors
	errch := make(chan error)
	go func() {
		for {
			err := <-errch
			if err != nil {
				log.Fatal(err)
			}

		}
	}()

	// -----------

	// Waitgroup and worker total
	var waitgroup sync.WaitGroup
	waitgroup.Add(5)

	// -----------

	// Values populated by go routines.
	var full_name string
	var sex string
	var age int
	var social_security_number string
	var location string
	var zipcode int
	var jobs []string
	var hobby_slice []string

	// -----------

	// Social security number
	go func(waitgroup *sync.WaitGroup) {

		defer waitgroup.Done()

		ch := make(chan string)

		go social_security(ch)

		social_security_number = <-ch

	}(&waitgroup)

	// -----------

	// City and zipcode
	go func(waitgroup *sync.WaitGroup) {

		defer waitgroup.Done()

		ch := make(chan string)
		intch := make(chan int)

		go city(ch, errch, intch)

		location = <-ch
		zipcode = <-intch

	}(&waitgroup)

	// -----------

	// Hobbies
	go func(waitgroup *sync.WaitGroup) {

		defer waitgroup.Done()

		ch := make(chan string)
		amount := rand.IntN(5)

		go hobbies(amount, ch, errch)

		for {
			hobby, open := <-ch
			if !open {
				break
			}
			hobby_slice = append(hobby_slice, hobby)
		}

	}(&waitgroup)

	// -----------

	// Name and sex
	go func(waitgroup *sync.WaitGroup) {

		defer waitgroup.Done()

		ch := make(chan string)

		go name(ch, errch)

		first := true
		for {
			data, open := <-ch
			if !open {
				break
			}
			if first {
				sex += data
				first = false
				continue
			}
			full_name += fmt.Sprintf("%v ", data)
		}

		full_name = strings.TrimSpace(full_name)

	}(&waitgroup)

	// -----------

	// Make a job history
	go func(waitgroup *sync.WaitGroup) {

		defer waitgroup.Done()

		ch := make(chan string)
		intch := make(chan int)

		go job_history(ch, intch, errch)

		age = <-intch

		for {
			job, open := <-ch
			if !open {
				break
			}
			jobs = append(jobs, job)
		}

	}(&waitgroup)

	// -----------

	// Wait for tasks to complete
	waitgroup.Wait()

	// -----------

	// Populate the person struct with the collected data
	new_person := Person_CV{
		name:            full_name,
		sex:             sex,
		age:             age,
		social_Security: social_security_number,
		location:        location,
		zipcode:         zipcode,
		hobbies:         hobby_slice,
		job_history:     jobs,
	}

	// -----------

	return new_person
}

// --------------------------------------------------------------------------------------------------------------------
