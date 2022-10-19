package main

import "fmt"

func main() {
      fmt.Println("Hello, World!")
      fmt.Println()
      
      /* *********************************
       * Variables - 3种声明变量的方式
      ********************************** */

      var a int
      a = 42
      a = 27

      var b float32 = 27

      c := 27.0

      fmt.Println(a, b, c)

      fmt.Printf("%v, %T", a, a)
      fmt.Println()
      fmt.Printf("%v, %T", b, b)
      fmt.Println()
      fmt.Printf("%v, %T", c, c)
      fmt.Println()

      /* *********************************
      // End of Variables ****************
      ********************************** */

      fmt.Print("\n\nEnd this program!\n")
}
