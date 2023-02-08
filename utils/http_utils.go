package utils

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func PrintRequest(r *http.Request) string {
	//https://go.dev/src/net/http/httputil/example_test.go
	dump, err := httputil.DumpRequest(r, true)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Printf("Request:\n%s\n", string(dump))

	return string(dump)
}
