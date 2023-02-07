/**
 * @demo 数组Array && 切片Slice
 */

package main

import (
	"fmt"
)

func main() {
	/*
		* @demo 数组Array
			NotePad:
				* 数组是值类型
				* 数组的长度是数组类型的一部分
	*/
	// 数组的实例化
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{4, 5, 6, 7, 8}
	var strTwoDimensional [4][5]string

	strTwoDimensional[2][3] = "hello"

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(strTwoDimensional, string(", length of strs is "), len(strTwoDimensional))
	// ------------------------------

	// 数组的遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("arr3[%d] = %d\t", i, arr3[i])
	}
	fmt.Println()

	for i := range arr3 {
		fmt.Printf("arr3[%d] = %d\t", i, arr3[i])
	}
	fmt.Println()

	for i, v := range arr3 {
		fmt.Printf("arr3[%d] = %d\t", i, v)
	}
	fmt.Println()
	// ------------------------------

	a := [...]int{1, 2, 3}
	b := a // NotePad: 数组是值类型，赋值和传参会复制整个数组

	b[1] = 100

	fmt.Printf("Array a: %v\n", a)
	fmt.Printf("Array b: %v\n", b)
	// ------------------------------

	/*
		* @demo 切片Slice
			NotePad:
				* 切片是数组的一个引用
				* 切片是引用类型
	*/
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))

	p := s // NotePad: 切片是引用类型，赋值和传参不会复制底层数组

	p[1] = 100

	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
	fmt.Printf("p = %v, len(p) = %d, cap(p) = %d\n", p, len(p), cap(p))

	// 切片的实例化
	s1 := s[:]   // slice of all elements
	s2 := s[3:]  // slice from 4th to end
	s3 := s[:6]  // slice first 6 elements
	s4 := s[3:6] // slice the 4th, 5th, and 6th elements

	s4[2] = 200

	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3 = %v, len(s3) = %d, cap(s3) = %d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4 = %v, len(s4) = %d, cap(s4) = %d\n", s4, len(s4), cap(s4))

	// 使用make函数实例化切片
	s5 := make([]int, 5, 10) // NotePad: make([]T, size, cap)

	fmt.Printf("s5 = %v, len(s5) = %d, cap(s5) = %d\n", s5, len(s5), cap(s5))

	// 使用 append() 函数向切片中追加元素
	s6 := []int{}

	fmt.Printf("s6 = %v\n", s6)
	fmt.Printf("Length: %v\n", len(s6))
	fmt.Printf("Capacity: %v\n", cap(s6))

	s6 = append(s6, 1)
	fmt.Printf("s6 = %v\n", s6)
	fmt.Printf("Length: %v\n", len(s6))
	fmt.Printf("Capacity: %v\n", cap(s6))

	s6 = append(s6, 2, 3, 4, 5)
	fmt.Printf("s6 = %v\n", s6)
	fmt.Printf("Length: %v\n", len(s6))
	fmt.Printf("Capacity: %v\n", cap(s6))

	s6 = append(s6, []int{6, 7, 8, 9, 10}...)
	fmt.Printf("s6 = %v\n", s6)
	fmt.Printf("Length: %v\n", len(s6))
	fmt.Printf("Capacity: %v\n", cap(s6))

	// 从数组中删除一个元素, 有破环性的 ！！！
	s7 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s7 = %v\n", s7)

	s8 := append(s7[:2], s7[3:]...)

	fmt.Printf("s8 = %v\n", s8)
	fmt.Printf("s7 = %v\n", s7)
	// ------------------------------
}
