package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for work := 1; work <= 3; work++ {
		go worker(work, jobs, results)
	}

	for job := 1; job <= 5; job++ {
		jobs <- job
	}
	close(jobs)

	for result := 1; result <= 5; result++ {
		<-results
	}

}
