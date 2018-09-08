package main

import (
	"crypto/sha256"
  "fmt"
)

func main() {
  data := "testdata"
  testSha := sha256.Sum256([]byte(data))
  fmt.Println(testSha)
}
