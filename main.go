package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// formatRequest generates ascii representation of a request
// copied from https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}


func requestHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf(" --- New request from %v at %v --- \n%v\n ---       End of request      --- \n\n\n",
		req.Host,
		time.Now().Format("15:04:05"),
		formatRequest(req))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", requestHandler)
	fmt.Println("HTTP Dummy running at http://localhost:9000, handling all requests now")
	if err := http.ListenAndServe(":9000", mux); err != nil {
		fmt.Println("Something went wrong:", err.Error())
	}
}
