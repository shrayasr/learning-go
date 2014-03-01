package main

import "fmt"

type phone struct {
  model       string
  awesomeness int
}

func makePhoneMoreAwesome(x *phone) {
  x.awesomeness += 5
}

func main() {

  motog := phone {"motog",7}
  nexus4 := phone {model:"nexus4", awesomeness:8}
  xperiaSP := phone {model:"xperiaSP"}

  fmt.Println(motog, nexus4, xperiaSP)
  makePhoneMoreAwesome(&xperiaSP)
  fmt.Println(xperiaSP)

}
