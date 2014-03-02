package main

import "fmt"

type Vehicle interface {
  howManyWheels() int
  milage() float32
}

type FourWheeledVehicles struct {
}

func (f *FourWheeledVehicles) howManyWheels() int {
  return 4
}

func (f *FourWheeledVehicles) milage() float32 {
  return 11.3
}

type TwoWheeledVehicles struct {
}

func (t TwoWheeledVehicles) howManyWheels() int {
  return 2
}

func (t TwoWheeledVehicles) milage() float32 {
  return 40.2
}

func getInfo(v Vehicle) {
  fmt.Println(v.howManyWheels())
  fmt.Println(v.milage())
}


func main() {

  vehicles := []Vehicle{new(FourWheeledVehicles), &TwoWheeledVehicles{}}

  for _, vehicle := range vehicles {
    getInfo(vehicle)
  }

}
