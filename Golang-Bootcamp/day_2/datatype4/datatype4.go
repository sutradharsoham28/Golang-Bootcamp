package main

import "fmt"

func main() {

	var a complex128 = complex(6, 2)
	var b complex64 = complex(9, 2)

	fmt.Printf("The type of a is %T and the type of b is %T", a, b)
	fmt.Println("\na=", a)
	fmt.Println("b=", b)

	a_real := real(a)
	a_imag := imag(a)

	fmt.Println("\nThe real part of a is: ", a_real)
	fmt.Println("\nThe imaginary part of a is: ", a_imag)

}
