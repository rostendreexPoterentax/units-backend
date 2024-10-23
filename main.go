package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

}

func main()  {

	er := http.ListenAndServe(":4000", nil)
	if er != nil {

	}
}
