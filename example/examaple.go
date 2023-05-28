package main

import (
	"io"
	"net/http"
)

/*
We have to basically replicate the functionality of the net/http
package, using only the net package.
*/

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bye", bar)

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello!")
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bye!")
}
