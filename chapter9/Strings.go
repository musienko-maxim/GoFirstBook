package main

import "fmt"

type Title string

func main() {
	fmt.Println(Title("Alien") == Title("Alien"))
}
