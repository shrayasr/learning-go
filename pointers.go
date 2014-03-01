package main

import "fmt"

func byVal(num int) {
  num += 5
  fmt.Println("Inside fxn: ",num)
}

func byRef(num *int) {
  *num += 5
  fmt.Println("Inside fxn: ",*num)
}

func main() {

  num := 5

  fmt.Println("Before byVal: ",num)
  byVal(num)
  fmt.Println("After byVal: ",num)

  fmt.Println("Before byRef: ",num)
  byRef(&num)
  fmt.Println("After byRef: ",num)

}
