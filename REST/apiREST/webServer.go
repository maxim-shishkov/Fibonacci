package apiREST

import (
	f "Fibonacci/src"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func WebServer() {
	r := mux.NewRouter()
	r.HandleFunc("/fib/", HandleGet).Methods("GET").
		Queries("x", "{x}").
		Queries("y", "{y}")

	log.Println("Start Server")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func HandleGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	x := params["x"]
	y := params["y"]


	fibonacci,err := f.GetFibonacci(x, y)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(fibonacci)
	}
}
