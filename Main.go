/**
 * @demo 复数类型
 */

package main

import (
	"fmt"
)

func main() {
	var n complex64 = 1 + 2i

	fmt.Printf("%v, %T\n", n, n)

	n = 2i

	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))

	var c complex128 = 1 + 2i

	fmt.Printf("%v, %T\n", real(c), real(c))
	fmt.Printf("%v, %T\n", imag(c), imag(c))

	a := 1 + 2i
	b := 2 + 5.2i

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	var x complex128 = complex(5, 12)

	fmt.Printf("%v, %T\n", x, x)
}
