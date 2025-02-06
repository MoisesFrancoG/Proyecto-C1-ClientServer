package main

import (
	"demo/client/src/polling"
	"time"
)

func main() {
	go polling.LongPollingClient()
	go polling.ShortPollingClient("http://localhost:8080/products")
	go polling.ShortPollingClient("http://localhost:8080/updated-products")

	for {
		time.Sleep(time.Hour)
	}
}
