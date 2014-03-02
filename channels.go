package main

import (
  "fmt"
  "time"
)

func sleep(num int, done chan bool) {
  time.Sleep(time.Duration(num)*time.Second)
  fmt.Println(num)
  done <- true
}

func main() {

  isDone := make(chan bool)
  toSort := []int{7,6,1,2,9,2,3}
  count := len(toSort)

  for _, num := range toSort {
    go sleep(num, isDone)
  }

  for {
    <-isDone
    count--;
    if count == 0 {
      break
    }
  }
}
