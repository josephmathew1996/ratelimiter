# Rate Limiter Application
This application uses `fixed window` rate-limiting algorithm to implement rate-limiting on per client basis.

Various other methods are:

Token bucket algorithm
        In the token bucket algorithm technique, tokens are added to a bucket, which is some sort of storage, at a fixed rate per time. Each request the application processes consumes a token from the bucket. The bucket has a fixed size, so tokens can’t pile up infinitely. When the bucket runs out of tokens, new requests are rejected.


Leaky bucket algorithm
        In the leaky bucket algorithm technique, requests are added to a bucket as they arrive and removed from the bucket for processing at a fixed rate. If the bucket fills up, additional requests are either rejected or delayed.


Sliding window algorithm
        Like the fixed window algorithm, the sliding window algorithm technique tracks the number of requests over a sliding window of time. Although the size of the window is fixed, the start of the window is determined by the time of the user’s first request instead of arbitrary time intervals. If the number of requests in a window exceeds a preset limit, subsequent requests are either dropped or delayed.


## Project structure 
Project structure is inspired by the the `Standard Go Project Layout` explained in the below repository.

Reference: https://github.com/golang-standards/project-layout

## Prerequisites

1. Install Golang SDK: https://go.dev/dl/
2. Execute following commands:
     
### To install all the dependencies

        go mod tidy -v 
    
### To run the server
    
        go run cmd/api/main.go

## Libraries used

1. Zap for logging: https://github.com/uber-go/zap
2. Viper for managing environment variables: https://github.com/spf13/viper
3. Echo framework for setting up the server: https://echo.labstack.com/
4. Standard libraries for other basic purposes: https://pkg.go.dev/std

## Test server

### Sample Commands

curl -X GET http://localhost:8080/api/v1/users

for i in {1..10}; do curl -X GET http://localhost:8080/api/v1/users; sleep 1; done
