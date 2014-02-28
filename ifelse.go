package main

import "fmt"

func main() {

  if 7%2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

  if 4%2 == 0 {
    fmt.Println("IDC")
  }

  if x := 2; x < 5 {
    fmt.Println(x, " is less than 5")
  } else if x > 0 {
    fmt.Println(x, " is greater than 0")
  } else if x == 2 {
    fmt.Println(x, " is equal to 2")
  }

}
