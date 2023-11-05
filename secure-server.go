package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for name, values := range r.Header {
			for _, value := range values {
				fmt.Fprintf(w, "%v: %v\n", name, value)
			}
		}
	})

	err := http.ListenAndServeTLS(":443", "fullchain.pem", "privkey.pem", nil)
	if err != nil {
		fmt.Println(err)
	}
}
