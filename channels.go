package main

import (
  "fmt"
  "time"
)

func doSomething(val chan bool) {
  <-val
  fmt.Println("doing something")
}

func doSomethingElse(val chan bool) {
  <-val
  fmt.Println("doing something else")
}

func khallas(val chan bool) {
  time.Sleep(time.Second)
  val <- true
  val <- true
  fmt.Println("khallas")
}

func main() {

  channel := make(chan bool)

  go doSomething(channel)
  go doSomethingElse(channel)
  go khallas(channel)

  var input string
  fmt.Scanln(&input)
}
