/**
 * @demo 在包一层声明变量
 */

/*
 Tips:
      * 无法用 varName := value 的方式声明变量
*/

package main

import "fmt"

var (
	name   string = "name"
	number int    = 3
	season int    = 11
)

func main() {
	fmt.Println(name)
	fmt.Println(number)
	fmt.Println(season)
}
