package benchmark

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Job struct {
	ID      int
	Url     string
	Headers []string
	Method  string
	Body    string
}

func (j *Job) New(client *http.Client) (statusCode int, err error) {
	req, err := http.NewRequest(j.Method, j.Url, bytes.NewBuffer([]byte(j.Body)))
	if err != nil {
		fmt.Println("Error creating request:", err)

		return 0, err
	}

	if len(j.Headers) > 0 {
		for _, v := range j.Headers {
			array := strings.Split(v, ":")

			req.Header.Set(array[0], array[1])
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return 0, err
	}

	_, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error reading response:", err)
		return 0, err
	}

	return resp.StatusCode, nil
}
