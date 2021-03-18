//EJEMPLOS DE COMO FUNCIONAN LOS SLICES Y ARREGLOS EN GO

package main

import "fmt"

func main() {
	//x := [5]int{1, 2, 3, 4, 5}
	//x[4] = 100
	//x[2] = 50
	//fmt.Println(len(x), cap(x))

	//for i := 0; i < len(x); i++ {
	//		fmt.Println(x[i])
	//}

	/*	for i, v := range x { //devuelve el indice y luego el valor
			fmt.Println(v, i)

		}
	*/
	//fmt.Println(x)

	//////////////////////////////////SLICE

	/*	arr := [5]float64{1, 2, 3, 4, 5}
			x := arr[0:3]
			fmt.Println(x)
		//
		x := [...]int{1, 2, 3, 4, 5, 6, 7}
		s := x[0:4]
		fmt.Println(len(s), cap(s))

	*/

	var s []int
	//	s = append(s, 1)
	fmt.Println(len(s), cap(s))
	s = append(s, 1, 4, 5, 6, 9)
	fmt.Println(len(s), cap(s))
	for _, v := range s {
		fmt.Println(v)
	}

}
