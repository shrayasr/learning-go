package main

import "fmt"

func main() {

  var a[5] int
  fmt.Println(a)

  a[4] = 10
  fmt.Println(a)
  fmt.Println(a[4])

  x := "some_string"
  fmt.Println(len(x))
  fmt.Println(len(a))

  b := [5]int{1,2,3,4,5}
  fmt.Println(b, len(b))

  var twoD[3][3] int

  for i:=0; i<3; i++ {
    for j:=0; j<3; j++ {
      twoD[i][j] = i+j
    }
  }

  fmt.Println(twoD)

}
