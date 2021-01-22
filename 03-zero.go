package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var sl []int
	var ptr *int
	var iface interface{}
	var arr [10]int
	fmt.Printf("%v %v %v %q %v %v %v\n", i, f, b, s, sl, ptr, iface)
	fmt.Printf("%v\n", arr)
}
