/**
 * @demo 变量命名规范及作用域
 */

/*
 NotePad:
      * 作用域 -- 3钟可见级别
	  *. package level : 放在代码最顶部,用此可在整个package 中使用此变量
	    *. 小写: 只能在包内部使用,但是在包内任何地方均可访问
	    *. 大写: 暴露该变量(全局),可以被本包外的其他包访问,被调用
	  *. funciton level : 可见范围在当前函数中
	  *. block level : 代码块  在 {} 直接影响 {} 内的 变量
*/

package main

import "fmt"

var I int = 42

var i int = 42

func main() {
	var j int = 21

	fmt.Println(i, j)
}
