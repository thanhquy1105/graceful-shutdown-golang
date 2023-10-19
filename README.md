# graceful-shutdown-golang
This repo is the graceful shutdown sample in Golang

## How to run
```go run .```

## How to test
1. Run the service
2. Open the other terminal, request to the API using this command line ```curl -X GET http://localhost:8080/api/v1/test-graceful-shutdown``` or open Postman, use get request to URL: http://localhost:8080/api/v1/test-graceful-shutdown
3. Back to the service terminal and stop the service: ```Ctrl + C```
4. Waiting the response from api
