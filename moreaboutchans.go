package main

import "fmt"

func worker(jobs chan int, done chan bool) {

  for i:=0; i<5; i++ {
    fmt.Println(<-jobs)
  }

  done <- true

}

func main() {

  jobsQueue := make(chan int)
  done := make(chan bool)

  go worker(jobsQueue, done)

  jobsQueue <- 1
  jobsQueue <- 2
  jobsQueue <- 3
  jobsQueue <- 4
  jobsQueue <- 5

  <-done

}
