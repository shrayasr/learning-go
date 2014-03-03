package main

import (
  "fmt"
  "time"
)

func worker1 (ticker <-chan time.Time) {
  for t := range ticker {
    fmt.Println("W1 | tick @ ", t)
  }
}

func worker2 (ticker <-chan time.Time) {
  for t := range ticker {
    fmt.Println("W2 | tick @ ", t)
  }
}

func main() {

  ticker1 := time.NewTicker(time.Second * 1)

  go worker1(ticker1.C)
  go worker2(ticker1.C)

  time.Sleep(time.Second * 3)
  ticker1.Stop()

  fmt.Println("ticker1 stopped")
}
