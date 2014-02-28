package main

import "fmt"

func main() {

  s := make([]string, 3)
  fmt.Println(s)

  fmt.Println("-----")

  s[0] = "asdf"
  s[1] = "qwer"
  s[2] = "foo"
  fmt.Println(s, s[2])
  fmt.Println(len(s))

  fmt.Println("-----")

  s = append(s, "bar")
  fmt.Println(s)

  fmt.Println("-----")

  c := make([]string, len(s))
  copy(c, s)
  fmt.Println(c,s)

  fmt.Println("-----")

  d := s[2:]
  fmt.Println(d)

  fmt.Println("-----")

  e := s[:1]
  fmt.Println(e)

  fmt.Println("-----")

  f := []string{"a","b","c"}
  fmt.Println(f)

  fmt.Println("-----")

  twos := make([][]int, 3)
  for i:=0; i<3; i++ {
    innerLen := i+1
    twos[i] = make([]int, innerLen)
    for j:=0; j<innerLen; j++ {
      twos[i][j] = i+j
    }
  }
  fmt.Println(twos)

  fmt.Println("-----")

  capa := make([]int, 5, 6)
  capb := append(capa,1,2,3,4)
  fmt.Println(capa, capb)

  fmt.Println("-----")

  originalSlice := []string{"a","b","c","d"}
  slicedSlice := originalSlice[:2]
  fmt.Println("Before:\n",originalSlice, len(originalSlice), cap(originalSlice))
  fmt.Println(slicedSlice, len(slicedSlice), cap(slicedSlice))
  slicedSlice = slicedSlice[:cap(slicedSlice)]
  fmt.Println(slicedSlice, len(slicedSlice), cap(slicedSlice))
  slicedSlice[0] = "x"
  fmt.Println("After: ",originalSlice, slicedSlice)

  fmt.Println("-----")

  // make a small slice
  smallSlice := []int{1,2,3,4,5}
  fmt.Println(smallSlice, len(smallSlice), cap(smallSlice))

  // make a large slice twice the size of the smaller one
  largeSlice := make([]int, len(smallSlice), (cap(smallSlice)+1)*2)
  fmt.Println(largeSlice, len(largeSlice), cap(largeSlice))

  // copy the smaller one into the larger
  copy(largeSlice, smallSlice)

  // alter a value in the small
  smallSlice[0] = -1
  fmt.Println("modified")
  fmt.Println(largeSlice, len(largeSlice), cap(largeSlice))
  fmt.Println(smallSlice, len(smallSlice), cap(smallSlice))
  fmt.Println("/modified")

  // make the small slice large
  smallSlice = largeSlice

  // modify the small slice again
  smallSlice[0] = -2
  fmt.Println(smallSlice, len(smallSlice), cap(smallSlice))
  fmt.Println(largeSlice, len(largeSlice), cap(largeSlice))

  fmt.Println("-----")

  ////////

  firstSlice := []int{1,2,3}
  secondSlice := []int{4,5,6}
  finalSlice := append(firstSlice, secondSlice...)
  fmt.Println(finalSlice)

  fmt.Println("-----")

}
