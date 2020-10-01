package main

import "fmt"

func main() {
	var valor int
	var resultado float64

	fmt.Scanf("%d\n", &valor)

	if valor >= 0 {
		resultado = calcularNumeroE(valor)
		fmt.Print(resultado)
	} else {
		fmt.Print("Ingrese un valor mayor o igual a 0")
	}
}

func calcularNumeroE(valor int) float64 {
	var numeroE float64

	numeroE = 0

	for i := 0; i <= valor; i++ {
		numeroE += float64(1) / float64(calcularFactorial(i))
	}

	return numeroE
}

func calcularFactorial(num int) int {
	var factorial int

	if num == 1 || num == 0 {
		factorial = 1
	} else {
		factorial = 1

		for i := 1; i <= num; i++ {
			factorial *= i
		}
	}

	return factorial
}