package main

import (
  "fmt"
  "time"
)

func something(timer <-chan time.Time, done chan bool) {
  <-timer
  fmt.Println("Waited!")
  done <- true
}

func main() {

  timer1 := time.NewTimer(time.Second * 2)

  done := make(chan bool)
  go something(timer1.C, done)
  <-done

  timer2 := time.NewTimer(time.Second * 4)
  go something(timer2.C, done)
  timer2.Stop()

  fmt.Println("Stopped")
}
