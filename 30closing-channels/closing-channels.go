package main

import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job:", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job")
	}

	close(jobs)

	fmt.Println("stop all job")

	<-done
}