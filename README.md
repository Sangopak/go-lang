# Go (golang)
## Key Features
1. Simplicity
2. Strongly and statically types
3. Garbage collected
4. Fast complie times
5. Built-in concurrency
6. Compiled to standalone binaries
## Installtion
1. Go to https://go.dev/ and download the latest stable version
2. Follow the installtion guide to install
3. You can get a lot of examples https://gobyexample.com/
4. Setup the GOROOT and GOPATH env variables
## Run Go program
Run the below command from root folder
```
go run src/cmd/main.go
```
## Build Go program
Run the below command from root folder
```
go build -v ./...
```
## Run Tests
Run the below command from root folder
```
go test -v ./...
```
### For Benchmark Testing
Run the below command where you have the BenchMark test example
/Users/sangopakkundu/workspace/go-lang/src/http/client
```
go test -bench=.
```