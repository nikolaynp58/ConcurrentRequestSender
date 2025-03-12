// package main

// import (
// 	"flag"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"sync"
// 	"sync/atomic"
// 	"time"
// )

// const (
// 	maxRequests     = 100
// 	concurrentLimit = 5
// )

// var globalCounter int32

// type RequestResult struct {
// 	URL        string
// 	StatusCode int
// 	Error      error
// 	Number     int32
// }

// func sendRequest(url string) RequestResult {
// 	client := &http.Client{Timeout: 5 * time.Second}
// 	resp, err := client.Get(url)

// 	count := atomic.AddInt32(&globalCounter, 1)
// 	result := RequestResult{
// 		URL:    url,
// 		Error:  err,
// 		Number: count,
// 	}

// 	if err != nil {
// 		return result
// 	}
// 	defer resp.Body.Close()

// 	result.StatusCode = resp.StatusCode
// 	return result
// }

// func worker(jobs <-chan string, results chan<- RequestResult, wg *sync.WaitGroup) {
// 	for url := range jobs {
// 		result := sendRequest(url)
// 		results <- result
// 		wg.Done()
// 	}
// }

// func processRequests(urls []string) {
// 	var wg sync.WaitGroup
// 	jobs := make(chan string, concurrentLimit)
// 	results := make(chan RequestResult, len(urls))

// 	for i := 0; i < concurrentLimit; i++ {
// 		go worker(jobs, results, &wg)
// 	}

// 	for i, url := range urls {
// 		wg.Add(1)
// 		jobs <- url
// 		fmt.Printf("Sending request number: %d...\n", i+1)

// 		if (i+1)%concurrentLimit == 0 {
// 			fmt.Printf("Waiting for batch to complete, current request waiting: %d...\n", i+2)
// 			wg.Wait()
// 		}
// 	}

// 	close(jobs)
// 	wg.Wait()
// 	close(results)

// 	for result := range results {
// 		if result.Error != nil {
// 			fmt.Printf("Error sending request to %s: %v\n", result.URL, result.Error)
// 		} else if result.StatusCode == 200 {
// 			fmt.Printf("Request number %d to %s completed successfully.\n", result.Number, result.URL)
// 		} else {
// 			fmt.Printf("Request number %d to %s failed with status %d (ignored).\n", result.Number, result.URL, result.StatusCode)
// 		}
// 	}
// }

// func parseFlags() []string {
// 	numRequests := flag.Int("num", -1, "Total number of requests (optional, ignored if not provided)")
// 	flag.Parse()

// 	args := flag.Args()
// 	if len(args) == 0 {
// 		fmt.Println("Error: No endpoints provided.")
// 		os.Exit(1)
// 	}

// 	if *numRequests == -1 {
// 		if len(args) > maxRequests {
// 			fmt.Printf("Error: Cannot exceed max limit of %d requests.\n", maxRequests)
// 			os.Exit(1)
// 		}
// 		*numRequests = len(args)
// 	} else if *numRequests > maxRequests {
// 		fmt.Printf("Error: Cannot exceed max limit of %d requests.\n", maxRequests)
// 		os.Exit(1)
// 	} else if *numRequests <= 0 {
// 		fmt.Println("Error: Cannot have 0 or less requests.")
// 		os.Exit(1)
// 	} else if *numRequests > len(args) {
// 		fmt.Println("Error: Cannot have more requests than endpoints.")
// 		os.Exit(1)
// 	}

// 	return args[:*numRequests]
// }

// func main() {
// 	urls := parseFlags()

// 	processRequests(urls)

// 	fmt.Println("All requests completed.")
// }

package main

import (
	"fmt"
	"task/requester"
	"task/utils"
)

func main() {
	urls := utils.ParseFlags()

	requester.ProcessRequests(urls)

	fmt.Println("All requests completed.")
}
