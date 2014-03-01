package main

import "fmt"

type phone struct {
  model       string
  awesomeness int
}

func (p *phone) promote() int {
  p.awesomeness += 5
  if p.awesomeness > 10 {
    p.awesomeness = 10
  }
  return p.awesomeness
}

func (p *phone) demote() int {
  p.awesomeness -= 5
  if p.awesomeness < 0 {
    p.awesomeness = 0
  }
  return p.awesomeness
}

func (p phone) introduce() string {
  if p.awesomeness < 2 {
    return "Throw the phone away! The " + p.model + " was a waste of money!"
  } else if p.awesomeness >=2 && p.awesomeness < 5 {
    return "The " + p.model + " is a good phone"
  } else if p.awesomeness >=5 && p.awesomeness <=7 {
    return "Brilliant phone, this! The " + p.model + " was a great buy"
  } else {
    return "Go buy a " + p.model + " like immediately!"
  }
}

func main() {

  xperiaSP := phone {model:"xperiaSP"}
  nexus4 := phone {"nexus4",7}
  motog := phone {"motog",9}

  fmt.Println("--Initial ratings")
  fmt.Println(motog.introduce())
  fmt.Println(nexus4.introduce())
  fmt.Println(xperiaSP.introduce())

  xperiaSP.promote()

  fmt.Println("--Promoting xperiaSP")
  fmt.Println(motog.introduce())
  fmt.Println(nexus4.introduce())
  fmt.Println(xperiaSP.introduce())

  motog.demote()

  fmt.Println("--Demoting motog")
  fmt.Println(motog.introduce())
  fmt.Println(nexus4.introduce())
  fmt.Println(xperiaSP.introduce())
}
