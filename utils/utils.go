package utils

import "fmt"

func Add(a, b, c int) (int, int) {
	fmt.Println(a + b + c)
	result := a + b + c
	first := a;
	return result, first;
}
