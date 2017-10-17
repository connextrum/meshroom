/*
A dummy deployable web service
*/

package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Hypha, %s", request.URL.Path[1:])
}

func postDatumHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "postDatumHandler called with %s", request.URL.Path[1:])
}

func queryCropHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "queryCropHandler called with %s", request.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}
