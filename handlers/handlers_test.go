package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"net/url"
	"path"
	"github.com/domdavis/fizzbuzz/handlers"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/domdavis/fizzbuzz/microservice"
	"strings"
)

type handlerTest struct {
	input        string
	expectedOut  string
	expectedCode int
}

//TODO: test failure cases

func TestFizzHandler(t *testing.T) {
	tests := []handlerTest{
		{"0", "Fizz", http.StatusOK },
		{"1", "", http.StatusOK },
		{"2", "", http.StatusOK },
		{"3", "Fizz", http.StatusOK },
		{"4", "", http.StatusOK },
		{"5", "", http.StatusOK },
		{"6", "Fizz", http.StatusOK },
		{"", "404 page not found", http.StatusNotFound },
	}
	handlerTestRunner("fizz", tests, t)
}

func TestBuzzHandler(t *testing.T) {
	tests := []handlerTest{
		{"0", "Buzz", http.StatusOK },
		{"1", "", http.StatusOK },
		{"2", "", http.StatusOK },
		{"3", "", http.StatusOK },
		{"4", "", http.StatusOK },
		{"5", "Buzz", http.StatusOK },
		{"10", "Buzz", http.StatusOK },
		{"", "404 page not found", http.StatusNotFound },
	}
	handlerTestRunner("buzz", tests, t)
}

func TestNumberHandler(t *testing.T) {
	tests := []handlerTest{
		{"0", "", http.StatusOK },
		{"1", "1", http.StatusOK },
		{"2", "2", http.StatusOK },
		{"3", "", http.StatusOK },
		{"4", "4", http.StatusOK },
		{"5", "", http.StatusOK },
		{"6", "", http.StatusOK },
		{"7", "7", http.StatusOK },
		{"", "404 page not found", http.StatusNotFound },
	}
	handlerTestRunner("number", tests, t)
}

// private

func handlerTestRunner(key string, tests []handlerTest, t *testing.T) {
	h := handlers.Handlers[key]

	r := mux.NewRouter()
	uri := fmt.Sprintf("/%s/{%s}", key, microservice.In)
	r.HandleFunc(uri, h).Methods(http.MethodGet)
	// We do not need http.Handle("/", r) because we are are not running ListenAndServe (e.g. not actually starting the server)

	for _, test := range tests {

		testUrl := prepUrl(key, test.input)
		req, err := http.NewRequest("GET", testUrl, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		// check response code
		if status := rr.Code; status != test.expectedCode {
			t.Errorf("handler '%s' with path '%s' returned status code '%d' when '%d' was expected", key, testUrl, status, test.expectedCode)
		}

		// check response body
		if body := strings.Trim(rr.Body.String(), "\n"); body != test.expectedOut {
			t.Errorf("handler '%s' with path '%s' returned body '%s' when '%s' was expected", key, testUrl, body, test.expectedOut)
		}
	}
}

// helper function to build our url
func prepUrl(key string, id string) string {
	u, _ := url.Parse("/")
	u.Path = path.Join(u.Path, key, id)
	return u.String()
}