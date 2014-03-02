package main

import "fmt"

type drivable interface {
  whoami() string
  start() string
  cruise() string
  stop() string
}

type vehicle struct {
  typeof string
  wheels int
  base_sound string
}

func (v vehicle) whoami() string {
  return "---" + v.typeof
}

func (v vehicle) start() string {
  return "Starting engine.. " + v.base_sound[:3]
}

func (v vehicle) cruise() string {

  var x string
  for i:=0; i<v.wheels; i++ {
    x += v.base_sound
  }
  return x
}

func (v vehicle) stop() string {
  return "screeeech"
}

func drive (d drivable) {
  fmt.Println(d.whoami())
  fmt.Println(d.start())
  fmt.Println(d.cruise())
  fmt.Println(d.cruise())
  fmt.Println(d.cruise())
  fmt.Println(d.stop())
}

func main() {

  drivableVehicles := []drivable{&vehicle{"car",4,"vroom"}, &vehicle{"bullet",2,"budu"}}

  for _,drivableVehicle := range drivableVehicles {
    drive(drivableVehicle)
  }

}
