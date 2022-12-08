/**
 * @demo 常量 --  const 关键字
 */

package main

import (
	"fmt"
)

/* ****************************************
枚举常量
***************************************** */
const (
	a0 = iota
	b1 = iota
	c2 = iota
)

const (
	_ = iota // 丢弃 枚举中的 0 值
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

/* ****************************************
枚举常量 -- 文件大小的计量单位
NotePad: 位运算参与枚举常量取值
***************************************** */
const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // == 1024
	MB                    // == 1048576
	GB                    // == 1073741824
	TB                    // == 1099511627776
	PB                    // == 1125899906842624
	EB                    // == 1152921504606846976
	ZB                    // == 1180591620717411303424
	YB                    // == 1208925819614629174706176
)

/* ****************************************
枚举常量 -- 比特位 表 布尔值-用户权限
***************************************** */
const (
	isAdmin            = 1 << iota // == 1
	isHeadquarters                 // == 2
	canSeeFinancials               // == 4
	canSeeAfrica                   // == 8
	canSeeAsia                     // == 16
	canSeeEurope                   // == 32
	canSeeNorthAmerica             // == 64
	canSeeSouthAmerica             // == 128
)

func main() {
	const myConst int = 42

	fmt.Printf("%v, %T\n", myConst, myConst)

	const a = 42
	var b int16 = 27

	fmt.Printf("%v, %T\n", a+b, a+b) // const 隐式类型转换
	// ----------------------------------------

	// 枚举常量 demo
	fmt.Printf("%v, %T\n", a0, a0)
	fmt.Printf("%v, %T\n", b1, b1)
	fmt.Printf("%v, %T\n", c2, c2)

	var specialistType_1 int
	var specialistType_2 int = catSpecialist

	fmt.Printf("%v\n", specialistType_1 == catSpecialist)
	fmt.Printf("%v\n", specialistType_2 == catSpecialist)
	// ----------------------------------------

	// 文件大小转换
	fileSize := 4000000000.

	fmt.Printf("%.2fGB\n", fileSize/GB)
	// ----------------------------------------

	// 位或运算实现比特位的布尔值
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope

	fmt.Printf("%b\n", roles)
	// 位与运算取回特定比特位的值
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
	// ----------------------------------------
}
