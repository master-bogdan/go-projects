package main

import (
	"flag"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"sync"
	"time"
)

type Job struct {
	id    int
	value int
}

type Result struct {
	jobID  int
	result int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.id)
		time.Sleep(time.Second) // simulate work
		results <- Result{jobID: job.id, result: job.value * 2}
	}
}

func main() {
	supportedHTTPMethods := []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete}
	supportedHTTPHeaders := []string{"Content-Type", "Accept", "User-Agent", "Authorization"}

	concurrency := flag.Int("c", 0, "Number of workers")
	requests := flag.Int("n", 1, "Number of requests")
	method := flag.String("m", "POST", "HTTP method")
	headers := flag.String("H", "", "HTTP headers")
	body := flag.String("b", "{}", "Request body")

	flag.Parse()

	url := flag.Args()[0]

	if *concurrency < 0 {
		panic("Number of workers can't be less than 0")
	}

	if *requests <= 0 {
		panic("Number of requests can't be less than 1")
	}

	if !slices.Contains(supportedHTTPMethods, strings.ToUpper(*method)) {
		notSupportedMsg := fmt.Sprintf("Not supported HTTP method: %s \n Supported HTTP methods: %s", strings.ToUpper(*method), strings.Join(supportedHTTPMethods, ""))

		panic(notSupportedMsg)
	}

	if !slices.Contains(supportedHTTPHeaders, *headers) {
		notSupportedMsg := fmt.Sprintf("Not supported HTTP headers: %s \n Supported HTTP headers: %s", *headers, strings.Join(supportedHTTPHeaders, ""))

		panic(notSupportedMsg)
	}

	startMessage := fmt.Sprintf(`Running HTTP benchmark with:
- concurrency workers: %d
- number of requests: %d
- HTTP method: %s
- HTTP headers: %s
- HTTP body: %s
- URL: %s
`, *concurrency, *requests, *method, *headers, *body, url)

	fmt.Println(startMessage)

	// fmt.Printf(`
	// 	Requests: %d
	// 	Success: %d | Failures: %d
	// 	Avg Latency: %s
	// 	P95 Latency: %s
	// 	Throughput: %d req/s
	// `)

	var wg sync.WaitGroup

	jobs := make(chan Job, *requests)
	results := make(chan Result, *requests)

	for index := 1; index <= *concurrency; index++ {
		wg.Add(1)
		go worker(index, jobs, results, &wg)
	}

	for j := 1; j <= *requests; j++ {
		jobs <- Job{id: j, value: j}
	}
	close(jobs)
	wg.Wait()

	close(results)

	for res := range results {
		fmt.Printf("Result: job %d, value %d\n", res.jobID, res.result)
	}

	fmt.Println("All jobs completed.")
}
