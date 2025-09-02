package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/master-bogdan/http-benchmark/benchmark"
	"github.com/master-bogdan/http-benchmark/flags"
)

func main() {
	f := flags.Flags{}

	flags := f.ParseFlags()
	err := flags.ValidateFlags()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	startMessage := fmt.Sprintf(
		`Running HTTP benchmark with:
- concurrency workers: %d
- number of requests: %d
- HTTP method: %s
- HTTP headers: %s
- HTTP body: %s
- URL: %s
`, flags.Concurrency, flags.Requests, flags.Method, strings.Join(flags.Headers, ", "), flags.Body, flags.Url)

	fmt.Println(startMessage)

	// fmt.Printf(`
	// 	Requests: %d
	// 	Success: %d | Failures: %d
	// 	Avg Latency: %s
	// 	P95 Latency: %s
	// 	Throughput: %d req/s
	// `)

	benchmark := &benchmark.Benchmark{
		JobsNumber:    flags.Requests,
		WorkersNumber: flags.Concurrency,
		Job: benchmark.Job{
			Url:     flags.Url,
			Method:  flags.Method,
			Headers: []string{"Content-Type:application/json"},
			Body:    flags.Body,
		},
	}

	benchmark.Run()

	fmt.Println("All jobs completed.")
}
