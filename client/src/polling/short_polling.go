package polling

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func ShortPollingClient(url string) {
	for {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Short Polling Response:", string(body))
		resp.Body.Close()
		time.Sleep(5 * time.Second)
	}
}
