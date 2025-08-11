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
	fmt.Println("#############################################################################")
	// go start()
	// time.Sleep(1 * time.Second)
	fmt.Println("########################### Wait Group ######################################")
	// var wg sync.WaitGroup
	// wg.Add(10)
	// for i := 1; i < 11; i++ {
	// 	go calculate_square(i, &wg)
	// }
	// wg.Wait()

	fmt.Println("########################### Channels ########################################")
	// ch := make(chan string)
	// go sell(ch)
	// go buy(ch)
	// time.Sleep(2 * time.Second)
	fmt.Println("########################### buffered Channels ###############################")
	// var wg sync.WaitGroup
	// wg.Add(2)
	// ch:=make(chan int,3)
	// go sell2(ch,&wg)
	// wg.Wait()
	fmt.Println("########################### Select Channels #################################")
	// ch1:=make(chan int)
	// ch2:=make(chan int)
	// go goone(ch1)
	// go gotwo(ch2)
	// select {
	// 	case val1:=<-ch1 :
			// fmt.Println("Channel Value1 :",val1)
			// break
	// 	case val2:=<-ch2 :
	// 		fmt.Println("Channel Value2 :",val2)
	// 		break
	// 	// default:
	// 	// 	fmt.Println("default value")	
	// }
	// time.Sleep(2*time.Second)
	fmt.Println("########################### Go routine leak Channels ########################")
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go leak(&wg)
	// wg.Wait() 
	fmt.Println("########################### closure Channels ################################")
	// var wg sync.WaitGroup
	// wg.Add(10)
	// for i:=1;i<=10;i++{
	// 	go func(){
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()
	// } 
	// fmt.Println("ًwith repetiotion at some cases")
	// wg.Wait()
	
	// // for i:=1;i<=10;i++{
	// // 	go func (i int){
	// // 		fmt.Println(i)
	// // 		wg.Done()
	// // 	}(i)
	// // }
	// // fmt.Println("Without repetition")
	// // wg.Wait()
	fmt.Println("########################### After time Functions ############################")
	ch:=make(chan int)
	select {
		case <- time.After(1*time.Second):
			fmt.Println("Select this after time out")
		case val:= <- ch:
			fmt.Println("Channel Value1 :",val)
			break

	}
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
	val2:= <- ch
	fmt.Println("Channel value : ", val2)
}

func sell2(ch chan int,s *sync.WaitGroup){
	ch <- 10
	ch <- 11
	ch <- 12
	fmt.Println("all sent")
	s.Done()

	go buy2 (ch,s)
	// close(ch)
 }

func buy2(ch chan int,s *sync.WaitGroup){
	fmt.Println("waiting")
	
	fmt.Println("recieved",<-ch)
	
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	s.Done()
}

func goone(ch1 chan int){
	ch1 <- 1
	fmt.Println("goONE")
}
func gotwo(ch2 chan int){
	ch2 <- 2
	fmt.Println("goTWO")
}

func leak (s *sync.WaitGroup){
	ch :=make(chan int)
	go func(){
		val:=<-ch
		fmt.Println(val)
		s.Done()
	}()
	fmt.Println("exit with leak method")
	s.Done()
}

func test_time(ch chan int){
	time.Sleep(3*time.Second)
	ch <- 10
}