package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Number struct {
	raw  int
	proc string
}

func (data Number) Proccess() Number {
	if data.raw%5 == 0 && data.raw%3 == 0 {
		data.proc = "FizzBuzz"
	} else if data.raw%5 == 0 {
		data.proc = "Buzz"
	} else if data.raw%3 == 0 {
		data.proc = "Fizz"
	} else {
		data.proc = strconv.Itoa(data.raw)
	}
	return data
}

type Jobs struct {
	raw_data          []Number
	processed_numbers []Number
	workers           int
	wg                sync.WaitGroup
	finish            chan Number
}

func (job *Jobs) Process(tasks []Number) {
	defer job.wg.Done()

	for _, task := range tasks {
		task := task.Proccess()
		job.finish <- task
	}
}

func (job *Jobs) Assign() {
	worker_jobs := len(job.raw_data) / job.workers

	for i := 0; i < job.workers; i++ {
		job.wg.Add(1)
		go job.Process(job.raw_data[:worker_jobs])
		job.raw_data = job.raw_data[worker_jobs:]
	}

	job.wg.Wait()
	close(job.finish)

}

func main() {
	no := []Number{}

	for i := 0; i < 100; i++ {
		no = append(no, Number{raw: i + 1})
	}

	jobs := Jobs{raw_data: no, workers: 10, finish: make(chan Number, len(no))}

	jobs.Assign()

	unsorted := []Number{}
	for {
		number, open := <-jobs.finish
		if !open {
			break
		}
		unsorted = append(unsorted, number)
	}

	sorted := merge_sort(unsorted)
	for i := 0; i < len(sorted); i++ {
		fmt.Println(sorted[i].proc)
	}

}

func merge_sort(data []Number) []Number {
	if len(data) < 2 {
		return data
	}

	split := len(data) / 2
	left := merge_sort(data[:split])
	right := merge_sort(data[split:])
	return sort(left, right)
}

func sort(left []Number, right []Number) []Number {
	i, j := 0, 0
	data := []Number{}

	for i < len(left) && j < len(right) {
		if left[i].raw <= right[j].raw {
			data = append(data, left[i])
			i++
		} else {
			data = append(data, right[j])
			j++
		}

	}

	for i < len(left) {
		data = append(data, left[i])
		i++

	}

	for j < len(right) {
		data = append(data, right[j])
		j++
	}

	return data
}
