/**
 * @demo 变量类型转换
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32 = 0.0
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var s string = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", s, s)
}
