package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChanel chan string) {
	reportNap("sending routine", 2)
	fmt.Println("***sending value***")
	myChanel <- "a"
	fmt.Println("***sending value***")
	myChanel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
