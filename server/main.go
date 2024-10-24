package main

import (
	"fmt"
	"log"
	"net/http"

	"sever/modules/low/database"
	"sever/modules/middle/auth"

	"github.com/rotisserie/eris"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\nserver handler started")
}

func main() {
	er := database.InitDB()
	if er != nil {
		log.Fatalf("Database not connected\nError:\n %v", eris.ToString(er, true))
	}
	log.Print("Registration routes...")
	handlersInit()
	log.Print("Routers is registred")
	fmt.Print("server started:\nhttp://localhost:4000")
	er = http.ListenAndServe(":4000", nil)
	if er != nil {
		log.Printf("Server not started.\nError:\n%v",
			eris.Wrap(er, "failed to start http.Listen"))
	}

}

func handlersInit() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/api_user_registration", auth.ApiRegisterUser)
}
