package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afzal/go-course/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("API - Learning in Golang")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	routes.CourseRouters(r)

	//listen to a port
	log.Fatal(http.ListenAndServe(":5500", r))

}

//serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API by golang</h1>"))
}
