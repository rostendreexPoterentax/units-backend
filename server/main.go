package main

import (
	"fmt"
	"log"
	"net/http"
	"sever/modules/middle/auth"
	"sever/modules/middle/low/database"
	"time"

	"github.com/rotisserie/eris"
)

func handler(w http.ResponseWriter, r *http.Request) {
	id, er := auth.AddUser(&auth.DboUser{
		Name:           "Vladimir",
		Password:       "6811935566",
		Email:          "Poterentax93@gmail.com",
		CreationMoment: time.Now().UTC(),
	})
	if er != nil {
		eris.Wrap(er, "failed to create first test user")
	}
	fmt.Printf("%v", id)
	fmt.Print("\nserver handler started")
}

func main() {
	er := database.InitDB()
	if er != nil {
		log.Fatalf("Database not connected\nError:\n %v", eris.ToString(er, true))
	}

	http.HandleFunc("/", handler)
	fmt.Print("server started:\nhttp://localhost:4000")
	er = http.ListenAndServe(":4000", nil)
	if er != nil {
		log.Printf("Server not started.\nError:\n%v",
			eris.Wrap(er, "failed to start http.Listen"))
	}

}
