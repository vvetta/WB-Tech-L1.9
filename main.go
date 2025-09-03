package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup	
	numbers := [4]int{1, 2, 3, 4}

	jobs := make(chan int, len(numbers))
	out := make(chan int, len(numbers))


	wg.Add(3)

	go CreateJob(&wg, numbers[:], jobs)
	go DoJob(&wg, jobs, out)
	go PrintOut(&wg, out)

	wg.Wait()
}

func PrintOut(wg *sync.WaitGroup, out <-chan int) {
	defer wg.Done()
	for number := range out {
		fmt.Println(number)
	}
}

func DoJob(wg *sync.WaitGroup, jobs <-chan int, out chan<- int) {
	defer wg.Done()
	defer close(out)
	for job := range jobs {
		out<-job*2
	}
}

func CreateJob(wg *sync.WaitGroup, numbers []int, jobs chan<- int) {
	defer wg.Done()
	defer close(jobs)
	for _, number := range numbers {
		jobs<-number
	}
}
