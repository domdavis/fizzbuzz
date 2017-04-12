package main

import (
	"flag"

	"github.com/domdavis/fizzbuzz/handlers"
	"github.com/domdavis/fizzbuzz/microservice"
)

func main() {
	var service, url, fizz, buzz, number string
	var port int

	flag.StringVar(&service, "service", "",
		"The service to run [fizz, buzz, number or fizzbuzz] (Required).")
	flag.StringVar(&url, "url", "",
		"The url to run on (default is the service name")
	flag.IntVar(&port, "port", 8000, "The port to run on (default 8000)")

	flag.StringVar(&fizz, "fizz", "localhost:8000/fizz",
		"The url for the fizz service.")
	flag.StringVar(&buzz, "buzz", "localhost:8000/buzz",
		"The url for the buzz service.")
	flag.StringVar(&number, "number", "localhost:8000/number",
		"The url for the number service.")

	flag.Parse()

	if url == "" {
		url = service
	}

	switch {
	case service == "fizzbuzz":
		microservice.NewMicroservice(
			url, handlers.FizzBuzz(fizz, buzz, number), port)
	case handlers.Handlers[service] != nil:
		microservice.NewMicroservice(url, handlers.Handlers[service], port)
	default:
		flag.Usage()
	}
}
