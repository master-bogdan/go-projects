package benchmark

import (
	"fmt"
	"sync"
	"time"

	http_client "github.com/master-bogdan/http-benchmark/http-client"
)

type Result struct {
	JobID      int
	Latency    int
	StatusCode int
	Error      string
}

type Worker struct {
	ID      int
	Jobs    chan Job
	Results chan Result
	Wg      *sync.WaitGroup
}

func (w Worker) New() {
	client := http_client.NewHttpClient()
	defer w.Wg.Done()

	for job := range w.Jobs {
		start := time.Now()
		fmt.Printf("Worker %d processing job %d\n", w.ID, job.ID)

		newJob := job
		statusCode, err := newJob.New(client)

		end := time.Now()
		latency := end.Sub(start).Milliseconds()

		if err != nil {
			w.Results <- Result{JobID: job.ID, Latency: int(latency), StatusCode: statusCode, Error: err.Error()}

			continue
		}

		w.Results <- Result{JobID: job.ID, Latency: int(latency), StatusCode: statusCode}
	}
}
