package main

import "fmt"

func worker (jobs chan int) {

  for i:=0; i<10; i++ {
    jobs <- i
  }

  close(jobs)

}

func consumer (jobs chan int) {

  for job := range jobs {
    fmt.Println(job)
  }

}

func main() {

  jobs := make(chan int)


  go worker(jobs)
  consumer(jobs)

}
