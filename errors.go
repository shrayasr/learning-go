package main

import (
  "fmt"
  "errors"
)

func builtInErrors(num int) (bool, error) {

  if num == 7 {
    return false, errors.New("7 is a bad number")
  } else {
    return true, nil
  }
}

type GopherError struct {
  errorMsg string
}

func (g *GopherError) Error() string {
  return "ERROR!: " + g.errorMsg
}

func gopherErrors(num int) (bool, error) {

  if num == 7 {
    return false, &GopherError{"gopheeerrr can't work with 7"}
  } else {
    return true, nil
  }
}

func main() {

  for i:=0; i<10; i++ {

    if ok, e := builtInErrors(i); e != nil {
      fmt.Printf("%t Failed %s\n", ok, e)
    } else {
      fmt.Printf("%t Success!\n", ok)
    }

    if ok, e := gopherErrors(i); e!= nil {
      fmt.Printf("%t Gopher Failed %s\n", ok, e)
    } else {
      fmt.Printf("%t Gopher yay\n", ok)
    }

  }

}
