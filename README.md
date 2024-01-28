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
### Running Profileing tool
1. `cd` to the directory where Benchmark test is present, example `cd /Users/sangopakkundu/workspace/go-lang/src/http/client`
2. Run the below command that runs the benchmark test in current directory and creates CPU and Memory profile
`go test -cpuprofile cpu.prof -memprofile mem.prof -bench .`
3. Profile can be visualized using below command using Graphviz plugin
`go tool pprof --dot cpu.prof>cpu.dot`
`go tool pprof --dot mem.prof>mem.dot`
4. You can install Graphviz in Mac by `brew install graphviz` and then run below command to visualize profile in browser
`go tool pprof -http=":8080" cpu.prof`
`go tool pprof -http=":8080" mem.prof`

#### If you get TCP port already in use error on Mac
`lsof -i -P | grep LISTEN | grep 8080`
and then `kill -9 PID` 