package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// start := time.Now()
	// for i := 1; i < 10; i++ {
	// 	go calculate_square(i)
	// }
	// fmt.Println(time.Since(start))
	// time.Sleep(2 * time.Second)
	fmt.Println("###########################################################")
	go start()
	time.Sleep(1 * time.Second)
	fmt.Println("########################### Wait Group ################################")
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 1; i < 11; i++ {
		go calculate_square(i, &wg)
	}
	wg.Wait()

	fmt.Println("########################### Channels ################################")
	ch := make(chan string)
	go sell(ch)
	go buy(ch)
	time.Sleep(1 * time.Second)
}
func calculate_square(i int, wg *sync.WaitGroup) {
	// time.Sleep(1 * time.Second)
	fmt.Println(i * i)
	wg.Done()
}

func start() {
	go Process()
	fmt.Println("In start")
}

func Process() {
	fmt.Println("In Process")
}

func sell(ch chan string) {
	ch <- "Fruit"
	fmt.Println("In Sell")
}
func buy(ch chan string) {
	fmt.Println("In buy")
	val := <-ch
	fmt.Println("Channel value : ", val)
}
