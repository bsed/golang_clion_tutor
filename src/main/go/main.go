package main

import "fmt"

func main() {

	var a string = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	_ = "breakpoint"
	fmt.Println(e)

	f := "short"
	_ = "breakpoint"
	fmt.Println(f)

	fmt.Printf("Hello, world.\n")

}