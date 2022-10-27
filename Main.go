/**
 * @demo 字符串
 */

package main

import (
	"fmt"
)

func main() {
	s := "this is a string"

	fmt.Printf("%v, %T\n", s, s)

	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	// ! cannot assign to s[2]
	// s[2] = "u"

	s1 := "this is also a string"

	fmt.Printf("%v, %T\n", s+s1, s+s1)

	b := []byte(s)

	fmt.Printf("%v, %T\n", b, b)

	var r rune = 'a'

	fmt.Printf("%v, %T\n", r, r)
}
