package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(tracer.WithAgentAddr("localhost:8126"))
	defer tracer.Stop()
	r := muxtrace.NewRouter(muxtrace.WithServiceName("manual-tracing"))
	r.HandleFunc("/", manual).Methods("GET")
	var port = ":3010"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func manual(w http.ResponseWriter, r *http.Request) {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Another service was called")
	w.WriteHeader(http.StatusOK)

}
