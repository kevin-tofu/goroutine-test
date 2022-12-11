package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func processA(v map[int]string) {

	for k, v := range v {
		fmt.Printf("key: %d, value: %s\n", k, v)
		time.Sleep(1 * time.Second)
		// fmt.Println(i, str)
	}

}

func processB(ch chan string, v string) {
	ch <- v
}

func processC(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func processD(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func main() {

	m1 := make(map[int]string)
	m1[1] = "US"
	m1[3] = "JP"
	m2 := map[int]string{2: "John", 4: "Mana"}

	fmt.Println("Start - ProcessA")

	go processA(m1)
	go processA(m2)

	log.Println(runtime.NumGoroutine()) // let us know numbers of goroutine
	time.Sleep(5 * time.Second)

	fmt.Println("End - ProcessA")

	fmt.Println("Start - ProcessB")

	chB := make(chan string)
	go processB(chB, "test")

	msg := <-chB
	fmt.Println(msg)
	time.Sleep(5 * time.Second)
	fmt.Println("End - ProcessB")

	fmt.Println("Start - ProcessC")

	chC := make(chan int, 20)
	// var ch1 chan int //
	// var ch2 <-chan int //
	// var ch3 chan<- int //

	go processC("go-routine 1", chC)
	go processC("go-routine 2", chC)
	go processC("go-routine 3", chC)

	i := 0
	for i < 100 {
		chC <- i
		i++
	}

	close(chC)
	time.Sleep(5 * time.Second)
	fmt.Println("End - ProcessC")

	fmt.Println("Start - processD")
	var wg sync.WaitGroup
	wg.Add(1)
	go processD("Hello", &wg)
	wg.Wait()
	fmt.Println("Start - processD")

}
