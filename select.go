package main

import (
  "fmt"
  "time"
)

func heavyLift(done chan bool) {
  time.Sleep(time.Second * 6)
  done <- true
}

func lightWork(done chan bool) {
  time.Sleep(time.Second * 2)
  done <- true
}

func main() {

  heavyChannel := make(chan bool)
  lightChannel := make(chan bool)

  noOfTasks := 4

  go heavyLift(heavyChannel)
  go lightWork(lightChannel)
  go lightWork(lightChannel)
  go lightWork(lightChannel)

  for {
    select {
    case <-heavyChannel:
      fmt.Println("Heavy dude done")
    case <-lightChannel:
      fmt.Println("Light dude done")
    }
    noOfTasks--
    if noOfTasks == 0 {
      break
    }
  }


}
