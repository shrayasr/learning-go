package main

import "fmt"

func worker(jobs chan int, done chan bool) {
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
}

func main() {

  jobs := make(chan int, 5)
  done := make(chan bool)

  go worker(jobs, done)

  for j:=0; j<3; j++ {
    jobs <- j
    fmt.Println("sent job")
  }

  close(jobs)

  <-done
}
