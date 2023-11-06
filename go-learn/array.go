package main

import "fmt"

var array_num = [...]int{1,2,3,4,5,6}
var array_names =[10]string{"lemon","lemon","lemon"}

func array() {
	// Add code for the array function here if needed.
	nameit_array:=[...]int{1,3,4,5,6,6}
	fmt.Println("pass it")
	fmt.Println(array_num);
	fmt.Println(array_names);
	fmt.Println(nameit_array);
	//change array elements
	nameit_array[2]=90;
	fmt.Println(nameit_array);
	//find array length
	fmt.Println(len(array_names))

}
