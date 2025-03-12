package utils

import (
	"flag"
	"fmt"
	"os"
)

const (
	maxAmoumtEndpoints = 100
)

func ParseFlags() []string {
	numRequests := flag.Int("num", -1, "Total number of requests (optional, ignored if not provided)")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: No endpoints provided.")
		os.Exit(1)
	}

	if *numRequests == -1 {
		if len(args) > maxAmoumtEndpoints || len(args) <= 0 {
			fmt.Printf("Error: Cannot exceed max limit of 100 endpoints or have none.\n")
			os.Exit(1)
		}
		*numRequests = len(args)
	} else if *numRequests > maxAmoumtEndpoints {
		fmt.Printf("Error: Cannot exceed max limit of 100 requests.\n")
		os.Exit(1)
	} else if *numRequests <= 0 {
		fmt.Println("Error: Cannot have 0 or less requests.")
		os.Exit(1)
	} else if *numRequests > len(args) {
		fmt.Println("Error: Cannot have more requests than endpoints.")
		os.Exit(1)
	}

	return args[:*numRequests]
}
