package main

import "fmt"

func main() {
	var num int
	var n1 int
	var suma int
	var negativo bool
	negativo = false
	fmt.Println("Cuantos numeros deseas agregar ")
	fmt.Scan(&num)
	if num > 0 {
		s := make([]int, 0, num)
		//  fmt.Println(len(s), cap(s))
		for i := 1; i <= num; i++ {
			fmt.Println("Ingresa el numero", i)
			fmt.Scan(&n1)
			if n1 > 0 {
				s = append(s, n1)
				suma = suma + n1
			} else {
				negativo = true
				fmt.Println("No son validos los numeros negativos")
				break
			}
		}

		if negativo == false {
			fmt.Println("La suma de los numeros es: ", suma)
		} else {
			fmt.Println("Intenta de nuevo.")
		}

	} else {
		fmt.Println("Intenta de nuevo.")
	}

}
