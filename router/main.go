package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	router := mux.NewRouter()

	// define the routes
	router.HandleFunc("/", createWelcomeMessage)
	router.HandleFunc("/user", getRandomUser)

	// listen on port 8080
	http.ListenAndServe(":8080", router)

}

// func that will run at route ?
func createWelcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my page")
}

// func that will run at /user route
func getRandomUser(w http.ResponseWriter, r *http.Request) {
	// using http package to call an endpoint and retrieve a random user
	data, err := http.Get("https://randomuser.me/api/")

	// check errors and exit if error exists
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// using ioutil to read the responde body
	val, e := ioutil.ReadAll(data.Body)

	// leave the goroutine if response has errors
	// panic will exit the goroutine
	if e != nil {
		panic(e)
	}

	// print the values on web
	fmt.Fprintln(w, string(val))
}
