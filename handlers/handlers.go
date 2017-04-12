package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/domdavis/fizzbuzz/client"
	"github.com/domdavis/fizzbuzz/microservice"
	"github.com/gorilla/mux"
)

type processor func(int) (int, string)

const (
	Fizz = "Fizz"
	Buzz = "Buzz"
)

var Handlers = map[string]microservice.Handler{
	"fizz":   fizzHandler,
	"buzz":   buzzHandler,
	"number": numberHandler,
}

func FizzBuzz(urls ...string) microservice.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, func(i int) (int, string) {
			if s, err := client.Get(urls, i); err != nil {
				return http.StatusInternalServerError, "Internal Server Error" + err.Error()
			} else {
				return 200, s
			}
		})
	}
}

func handler(w http.ResponseWriter, r *http.Request, f processor) {
	in := mux.Vars(r)[microservice.In]

	if in == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}

	if i, err := strconv.Atoi(in); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Invalid input: %s", in)))
	} else {
		status, msg := f(i)
		w.WriteHeader(status)
		w.Write([]byte(msg))
	}
}

func fizzHandler(w http.ResponseWriter, r *http.Request) {
	handler(w, r, func(i int) (int, string) {
		if i%3 == 0 {
			return http.StatusOK, Fizz
		}

		return http.StatusOK, ""
	})
}

func buzzHandler(w http.ResponseWriter, r *http.Request) {
	handler(w, r, func(i int) (int, string) {
		if i%5 == 0 {
			return http.StatusOK, Buzz
		}

		return http.StatusOK, ""
	})
}

func numberHandler(w http.ResponseWriter, r *http.Request) {
	handler(w, r, func(i int) (int, string) {
		if i%3 != 0 && i%5 != 0 {
			return http.StatusOK, strconv.Itoa(i)
		}

		return http.StatusOK, ""
	})
}
