package main

import "fmt"

func main() {

	num := 10
	var ptr *int

	ptr = &num;

	fmt.Println(ptr, *ptr);

	modifyPointerValue(ptr)
}

func modifyPointerValue(ptr *int) {
	(*ptr) *= (*ptr);
	fmt.Println(*ptr)
}
