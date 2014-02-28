package main

import "fmt"

func main() {

  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  for j := 10; j >= 0; j-- {
    fmt.Println(j)
  }

  i = 1
  for {
    fmt.Println("BOOM")
    if i > 10 {
      break
    } else {
      i++
    }
  }

}
