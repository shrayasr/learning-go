package main

import "fmt"

func main() {

  channel := make (chan string, 2)

  channel <- "asdf"
  channel <- "qwer"

  fmt.Println(<-channel)
  fmt.Println(<-channel)

}
