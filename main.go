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
