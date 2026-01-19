package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	id int
}

func worker(id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {

		fmt.Printf("Worker %d: Processing image %d\n", id, job.ID)
		time.Sleep(time.Millisecond * 500)
		results <- fmt.Sprintf("Result: Image %d processed by Worker %d", job.ID, id)
	}

}

func testworkers() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan string, numJobs)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {

		wg.Add(1) // start the reciever before hand

		worker(i, jobs, results, &wg)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- Job{id: j} // send it to through the jobs channel
	}

	close(jobs)

	go func() {

		defer wg.Wait()
		close(results)

	}()

	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("All work completed!")
}
