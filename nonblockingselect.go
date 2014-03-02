package main

import "fmt"

func main() {

  ch := make(chan bool,5)

  go func() {
    fmt.Println(<-ch)
  }()

  for i:=0; i<5; i++ {
    select
    {
      case ch <- true:
      default:
    }
  }

}
