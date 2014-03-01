package main

import "fmt"

func main() {

  // make a map
  myFirstMap := make(map[string]int)

  // print map itself and length
  fmt.Println(myFirstMap, len(myFirstMap))

  // set some vals
  myFirstMap["foo"] = 1234
  myFirstMap["bar"] = 4321
  myFirstMap["shit"] = -111

  // fetch values
  fmt.Println(myFirstMap["foo"])
  fmt.Println(myFirstMap["bar"])

  // print map itself and length
  fmt.Println(myFirstMap, len(myFirstMap))

  // remove shit from map
  delete(myFirstMap,"shit")

  // print map itself and length
  fmt.Println(myFirstMap, len(myFirstMap))

  // See if shit was still present
  _, shitIsThere := myFirstMap["shit"]
  fmt.Println(shitIsThere)

}
