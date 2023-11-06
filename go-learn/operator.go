package main

import "fmt"

func operator(){
	var a,b = 10,20;
	//sum
	sum:=a+b;
	fmt.Println(sum);
	//product
	product:=a*b;
	fmt.Println(product);
	//division
	division:=a/b;
	fmt.Println(division);
	//modulus
	modulus:=a%b;
	fmt.Println(modulus);
	//comparison
	comp:=a>b;
	fmt.Println(comp);
	//logic
	log:=a>b && b<a;
	fmt.Println(log);

	
}