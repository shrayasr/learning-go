package main

import "fmt"

func add2nos (a int, b int) (int, bool) {

  neg := false
  c := a+b

  if c < 0 {
    neg = true
  }

  return c,neg
}

func addManyNos (a ...int) int {

  fmt.Println(a)
  fmt.Print("= ")

  sum := 0
  for _, num := range a {
    sum += num
  }

  return sum

}

func counter() func() int {
  count := 0

  return func() int {
    count++;
    return count;
  }
}


func main() {

  sum, neg := add2nos(1,2)
  fmt.Println(sum, neg)

  sum, neg = add2nos(1,-5)
  fmt.Println(sum, neg)

  fmt.Println("-----")

  fmt.Println(addManyNos(1,2,3,4,5))

  manyNos := []int{11,22,33,44,55}
  fmt.Println(addManyNos(manyNos...))

  counter1 := counter()
  counter2 := counter()

  fmt.Println(counter1())
  counter2()
  counter2()
  counter2()
  fmt.Println(counter2())

}
