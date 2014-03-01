package main

import "fmt"

func main() {

  namesOfFrands := []string{"foo","bar","baz"}

  for _, frand := range namesOfFrands {
    fmt.Println(frand)
  }

  for idNo, frand := range namesOfFrands {
    fmt.Println(idNo, frand)
  }

  numsOfFrands := make(map[string]int)
  numsOfFrands["foo"] = 1234
  numsOfFrands["bar"] = 5678
  numsOfFrands["baz"] = 9012

  for k,v := range numsOfFrands {
    fmt.Println(k, v)
  }

}
