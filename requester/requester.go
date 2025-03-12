package requester

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"task/models"
	"time"
)

const (
	concurrentLimit = 5
)

var globalCounter int32

func sendRequest(url string) models.RequestResult {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)

	count := atomic.AddInt32(&globalCounter, 1)
	result := models.RequestResult{
		URL:    url,
		Error:  err,
		Number: count,
	}

	if err != nil {
		return result
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	return result
}

func worker(jobs <-chan string, results chan<- models.RequestResult, wg *sync.WaitGroup) {
	for url := range jobs {
		result := sendRequest(url)
		results <- result
		wg.Done()
	}
}

func ProcessRequests(urls []string) {
	var wg sync.WaitGroup
	jobs := make(chan string, concurrentLimit)
	results := make(chan models.RequestResult, len(urls))

	for i := 0; i < concurrentLimit; i++ {
		go worker(jobs, results, &wg)
	}

	for i, url := range urls {
		wg.Add(1)
		jobs <- url
		fmt.Printf("Sending request number: %d...\n", i+1)

		if (i+1)%concurrentLimit == 0 {
			fmt.Printf("Waiting for batch to complete, current request waiting: %d...\n", i+2)
			wg.Wait()
		}
	}

	close(jobs)
	wg.Wait()
	close(results)

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error sending request to %s: %v\n", result.URL, result.Error)
		} else if result.StatusCode == 200 {
			fmt.Printf("Request number %d to %s completed successfully.\n", result.Number, result.URL)
		} else {
			fmt.Printf("Request number %d to %s failed with status %d (ignored).\n", result.Number, result.URL, result.StatusCode)
		}
	}
}
