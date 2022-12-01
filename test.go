package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func processA(v map[int]string) {

	for k, v := range v {
		fmt.Printf("key: %s, value: %d\n", k, v)
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

func main() {

	m1 := make(map[int]string)
	m1[1] = "US"
	m1[3] = "JP"
	m2 := map[int]string{2: "John", 4: "Mana"}

	fmt.Println("Start - ProcessA")

	go processA(m1)
	go processA(m2)

	log.Println(runtime.NumGroutine()) // let us know numbers of goroutine

	fmt.Println("End - ProcessA")

	fmt.Println("Start - ProcessB")

	chB := make(chan string)
	go processB(chB, "test")

	msg := <-chB
	fmt.Println(msg)

	fmt.Println("End - ProcessB")

	// var ch1 chan int //
	// var ch2 <-chan int //
	// var ch3 chan<- int //

	fmt.Println("Start - ProcessB")

	chC := make(chan int, 20)
	go processC("go-routine 1", chC)
	go processC("go-routine 2", chC)
	go processC("go-routine 3", chC)

	i := 0
	for i < 100 {
		chC <- i
		i++
	}

	close(chC)
	fmt.Println(msg)

	fmt.Println("End - ProcessB")

}
