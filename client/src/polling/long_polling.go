package polling

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func LongPollingClient() {
	for {
		resp, err := http.Get("http://localhost:8080/products/long-polling")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Long Polling Response:", string(body)) // Mostrar la notificación del servidor
		resp.Body.Close()
	}
}
