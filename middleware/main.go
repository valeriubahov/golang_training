package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Logger struct {
	handler http.Handler
}

type Person interface{}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("LOG: \n- METHOD: %s\n- PATH: %s\n- DURATION: %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func main() {

	s := mux.NewRouter()

	s.HandleFunc("/", WelcomeMessage)

	s.HandleFunc("/user", GetRandomUser)

	test := NewLogger(s)

	log.Printf("server is listening at %s", "http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", test))

}

func WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	data, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		log.Printf("Error: %s", err)
		panic(err)
	}

	u, e := io.ReadAll(data.Body)

	if e != nil {
		log.Printf("Error: %s", e)

		panic(e)
	}

	fmt.Fprint(w, string(u))
}
