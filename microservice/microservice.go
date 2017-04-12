package microservice

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler func(http.ResponseWriter, *http.Request)

const In = "in"

func NewMicroservice(s string, f Handler, port int) {
	r := mux.NewRouter()
	uri := fmt.Sprintf("/%s/{%s}", s, In)
	r.HandleFunc(uri, f).Methods(http.MethodGet)

	http.Handle("/", r)

	log.Printf("Listening on localhost:%d/%s ...", port, s)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
