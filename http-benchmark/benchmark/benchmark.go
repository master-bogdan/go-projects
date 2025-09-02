package benchmark

import (
	"fmt"
	"sync"
)

type Benchmark struct {
	JobsNumber    int
	WorkersNumber int
	Job
}

func (b *Benchmark) Run() {
	var wg sync.WaitGroup

	jobs := make(chan Job, b.JobsNumber)
	results := make(chan Result, b.JobsNumber)

	for index := 1; index <= b.WorkersNumber; index++ {
		wg.Add(1)
		worker := Worker{ID: index, Jobs: jobs, Results: results, Wg: &wg}
		go worker.New()
	}

	for index := 1; index <= b.JobsNumber; index++ {
		jobs <- Job{ID: index, Url: b.Url, Method: b.Method, Headers: b.Headers, Body: b.Body}
	}
	close(jobs)
	wg.Wait()

	close(results)

	for res := range results {
		fmt.Printf("Result: bob %d, latency: %dms, statusCode: %d\n", res.JobID, res.Latency, res.StatusCode)
	}
}
