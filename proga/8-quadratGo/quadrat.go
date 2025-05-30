package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C, D float64
	fmt.Println("Введите коэффициенты уравнения: ")
	fmt.Println("A =")
	fmt.Scan(&A)
	fmt.Println("B =")
	fmt.Scan(&B)
	fmt.Println("C =")
	fmt.Scan(&C)

	D = B*B - 4*A*C

	if D > 0 {
		X1 := (-B + math.Sqrt(D)) / (2 * A)
		X2 := (-B - math.Sqrt(D)) / (2 * A)
		fmt.Println("Уравнение имеет два действительных корня:")
		fmt.Println("X1 =", X1, "X2 =", X2)
	} else if D == 0 {
		X := -B / (2 * A)
		fmt.Println("Уравнение имеет один действительный корень (кратности 2):")
		fmt.Println("X =", X)
	} else {
		real := -B / (2 * A)
		imaginary := math.Sqrt(-D) / (2 * A)
		fmt.Println("Уравнение имеет два комплексных корня:")
		fmt.Printf("X1 = %.2f + %.2fi\n", real, imaginary)
		fmt.Printf("X2 = %.2f - %.2fi\n", real, imaginary)
	}
}
