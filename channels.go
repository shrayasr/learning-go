package main

import "fmt"

type person struct {
  id int
  name string
}

func main() {

  simpleChannel := make(chan string)
  complexChannel := make(chan person)

  go func() {
    simpleChannel <- "foo"
  }()

  go func() {
    x := person{1,"shrayas"}
    complexChannel <- x
  }()

  fmt.Println(<-simpleChannel)
  x := <-complexChannel
  fmt.Printf("%d - %s\n", x.id, x.name)

}
