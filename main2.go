package main

import "fmt"

//EJEMPLO DE COMO FUNCIONA UN SLICE
func main() {
	s := make([]int, 0, 100)
	fmt.Println(len(s), cap(s))
	s = append(s, 1, 4, 5, 6, 9)
	fmt.Println(len(s), cap(s))
	for _, v := range s {
		fmt.Println(v)
	}

}
