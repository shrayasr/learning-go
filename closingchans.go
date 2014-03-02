package main

import "fmt"

func main() {

  jobs := make(chan int, 5)
  done := make(chan bool)

  go func(){

    for {
      j, more := <-jobs
      if more {
        fmt.Println("Job received, ",j)
      } else {
        fmt.Println("No more jobs")
        done <- true
        return
      }
    }

  }()

  for j:=0; j<3; j++ {
    jobs <- j
    fmt.Println("sent job")
  }

  close(jobs)

  <-done
}
