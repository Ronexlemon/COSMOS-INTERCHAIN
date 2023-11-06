package main

import (
	"fmt"
)

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum

}
func main() {
	array()
	slice()
	operator()
	forLoop()
	fmt.Println(summation(3,4));
	var num int = 2
	x := "yoolow"
	//varibale without initial value
	var name string
	var b bool
	var d float32
	//value assignment after declaration
	var interchain string
	interchain = "COSMOS"
	//Multiple Variable Declaration
	var aa, bb, c, dd int = 1, 2, 3, 4
	// constants
	const pie int16 = 60
	fmt.Println("Hellow Interchain")
	fmt.Println(num)
	fmt.Println(x)
	fmt.Println(name)
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(interchain)
	fmt.Println(aa, bb, c, dd)
	fmt.Println(pie)
	fmt.Println(add(3, 4))
}
