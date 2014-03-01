package main

import "fmt"

func main() {

  x := "ABC"

  fmt.Println(x[0:1])

  for index, runeValue := range x {
    fmt.Printf("%#U at %d\n", runeValue, index)
  }
}
