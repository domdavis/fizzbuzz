# FizzBuzz Microservice

FizzBuzz split into a set of 4 microservices designed to be run as a fleet of docker images for highly concurrent, highly resilient deployment.

To run locally:

```bash
make docker
./start.sh
```

To stop locally:

```bash
./stop.sh
```

Get usage instructions:

```bash
go run fizzbuzz.go
```

Also checkout the `Makefile`, `Dockerfile`, and `start.sh` for more details.
