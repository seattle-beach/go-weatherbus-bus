# Getting started

1. Clone the repo
2. Install the testing tools

  ```
  go get github.com/onsi/ginkgo/ginkgo
  go get github.com/onsi/gomega
  github.com/golang/protobuf/proto
  ```

3. Run the tests

  ```
  ginkgo -r
  ```
  
4. Run the service

  ```
  go run main.go  # will listen on port 9092
  ````
