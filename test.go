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

func main() {

	m1 := make(map[int]string)
	m1[1] = "US"
	m1[81] = "Japan"
	m2 := map[int]string{2: "John", 4: "aa"}

	fmt.Println("Start - ProcessA")

	go processA(m1)
	go processA(m2)

	log.Println(runtime.NumGroutine()) // let us know numbers of goroutine

	fmt.Println("End - ProcessA")

	fmt.Println("Start - ProcessB")

	ch := make(chan string)
	go processB(ch, "test")

	msg := <-ch
	fmt.Println(msg)

	fmt.Println("End - ProcessB")

}
