package main

import (
  "fmt"
  "time"
  "math/rand"
)

func doWork(done chan bool) {
  rand.Seed(42) //TODO Understand why

  secsToWork := rand.Intn(10)
  fmt.Println("Working...", secsToWork)

  time.Sleep(time.Second * time.Duration(secsToWork))
  done <- true
}

func main() {

  isWorkDone := make(chan bool)

  go doWork(isWorkDone)

  select {
  case <-isWorkDone:
    fmt.Println("Work done in time!")
  case <-time.After(time.Second * 5):
    fmt.Println("Work not done in time")
  }

}
