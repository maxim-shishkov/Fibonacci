package apiREST

import (
	f "Fibonacci/src"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func WebServer() {
	r := mux.NewRouter()
	r.HandleFunc("/src/", HandlGet).Methods("GET").
		Queries("x", "{x}").
		Queries("y", "{y}")

	log.Println("Start Server")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func HandlGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	x, err := strconv.Atoi(params["x"])
	if err != nil {
		fmt.Fprintf(w, "Invalid input X")
		return
	}

	y, err := strconv.Atoi(params["y"])
	if err != nil {
		fmt.Fprintf(w, "Invalid input Y")
		return
	}

	if x >= y {
		fmt.Fprintf(w, "Invalid input  x >= y")
		return
	}

	fibonacci := f.GetFibonacci(x, y)
	json.NewEncoder(w).Encode(fibonacci)
}
