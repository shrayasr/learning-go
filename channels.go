package main

import "fmt"

type person struct {
  id int
  name string
}

func doSomethingSimple(c chan string) {
  c <- "foo"
}

func doSomethingComplex(c chan person) {
  x := person{1, "shrayas"}
  c <- x
}

func main() {

  simpleChannel := make(chan string)
  complexChannel := make(chan person)

  go doSomethingSimple(simpleChannel)
  go doSomethingComplex(complexChannel)

  fmt.Println(<-simpleChannel)
  x := <-complexChannel
  fmt.Printf("%d - %s\n", x.id, x.name)

}
