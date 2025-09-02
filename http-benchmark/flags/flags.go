package flags

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"slices"
	"strings"
)

type headerSlice []string

func (h *headerSlice) String() string {
	return strings.Join(*h, ", ")
}

func (h *headerSlice) Set(value string) error {
	*h = append(*h, value)
	return nil
}

type Flags struct {
	Url         string
	Concurrency int
	Requests    int
	Method      string
	Headers     headerSlice
	Body        string
}

func (f *Flags) ParseFlags() *Flags {
	concurrency := flag.Int("c", 0, "Number of workers")
	requests := flag.Int("n", 1, "Number of requests")
	method := flag.String("m", "POST", "HTTP method")

	var headers headerSlice
	flag.Var(&headers, "H", "HTTP header to include, can be specified multiple times")
	body := flag.String("b", "{}", "Request body")

	flag.Parse()

	url := flag.Args()[0]

	return &Flags{
		Concurrency: *concurrency,
		Requests:    *requests,
		Method:      *method,
		Headers:     headers,
		Body:        *body,
		Url:         url,
	}
}

func (f *Flags) ValidateFlags() error {
	supportedHTTPMethods := []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete}
	supportedHTTPHeaders := []string{"Content-Type", "Accept", "User-Agent", "Authorization"}

	if f.Concurrency <= 0 {
		return errors.New("Number of workers can't be less than 0")
	}

	if f.Requests <= 0 {
		return errors.New("Number of requests can't be less than 1")
	}

	if !slices.Contains(supportedHTTPMethods, strings.ToUpper(f.Method)) {
		notSupportedMsg := fmt.Sprintf("Not supported HTTP method: %s \n Supported HTTP methods: %s", strings.ToUpper(f.Method), strings.Join(supportedHTTPMethods, ", "))

		return errors.New(notSupportedMsg)
	}

	for _, header := range f.Headers {
		if !slices.Contains(supportedHTTPHeaders, header) {
			notSupportedMsg := fmt.Sprintf("Not supported HTTP header: %s \n Supported HTTP headers: %s", header, strings.Join(supportedHTTPHeaders, ", "))

			return errors.New(notSupportedMsg)
		}
	}

	return nil
}
