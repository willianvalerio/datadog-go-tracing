package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	httptracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start()
	defer tracer.Stop()

	r := muxtrace.NewRouter(muxtrace.WithServiceName("go-tracing"))

	r.HandleFunc("/auto", auto).Methods("GET")
	r.HandleFunc("/manual", manual).Methods("GET")

	var port = ":3001"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}

/*This func intrumenting the code manually*/
func manual(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get endpoint, manual-tracing")
	time.Sleep(200 * time.Millisecond)
	span, ctx := tracer.StartSpanFromContext(r.Context(), "gateway.manual")
	defer span.Finish()
	time.Sleep(100 * time.Millisecond)

	req, err := http.NewRequest("GET", "http://localhost:3010/", nil)
	req = req.WithContext(ctx)

	err = tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(req.Header))
	if err != nil {
		fmt.Println(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		w.WriteHeader(res.StatusCode)
	}

}

/*This func intrumenting the code automatically*/
func auto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get endpoint, auto-tracing")
	time.Sleep(200 * time.Millisecond)
	httpClient := httptracer.WrapClient(http.DefaultClient)
	res, _ := httpClient.Get("http://localhost:3010")
	w.WriteHeader(res.StatusCode)

}
