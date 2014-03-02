package main

import (
  "fmt"
  "math/rand"
)

func main() {

  count := 10

  for i:=0; i<count; i++ {
    fmt.Println(rand.Intn(10))
  }
}
