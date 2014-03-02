package main

import "fmt"

func doSomething(count int) {

  for i:=0; i<count; i++ {
    fmt.Printf("doing something [%d/%d]\n", i, count)
  }
}

func main() {

  doSomething(5)

  go doSomething(100);
  go doSomething(100);

  fmt.Println("Go routines running simultaneously.")
  var input string
  fmt.Scanln(&input)

}
