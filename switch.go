package main

import (
  "fmt"
  "time"
)

func main() {

  i := 2

  switch(i) {

  case 0:
    fmt.Println("0")
  case 1:
    fmt.Println("1")
  case 2:
    fmt.Println("2")

  }

  x := time.Now().Weekday()
  switch (x) {

  case time.Saturday, time.Sunday:
    fmt.Println("YAY! Weekend")
  case time.Friday:
    fmt.Println("TOMORROW WEEKEND!")
  case time.Monday:
    fmt.Println("And it begins again!")
  default:
    fmt.Println("just work man")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Before noon")
  default:
    fmt.Println("After noon")
  }

}
