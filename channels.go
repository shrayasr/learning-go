package main

import "fmt"

func putToChannel (c chan string) {
  c <- "asdf"
}

func getFromChannel (c chan string) {
  fmt.Println(<-c)
}

func main() {

  channel := make (chan string)

  go getFromChannel(channel)
  channel <- "qwer"
  go getFromChannel(channel)
  putToChannel(channel)
}
