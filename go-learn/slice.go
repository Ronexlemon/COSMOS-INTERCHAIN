package main

import "fmt"

func slice(){
	slice_name:=[]int{};
	
	fmt.Println(slice_name);
	//create slice from array
	myArray:=[10]int{1,2,3,43,55,46,4,6,46,4}
slice_name=myArray[0:5];
fmt.Println(slice_name);
//create a slice with the make function
slice_make:= make([]string,5,10)
fmt.Println(slice_make);

}