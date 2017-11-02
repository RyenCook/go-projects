package main

import (
    "fmt"
)

func main() {
    var a, b = 1.0, 2.0

    /** Arithmetic Operators */
    fmt.Println(a + b)
    fmt.Println(a - b)
    fmt.Println(a * b)
    fmt.Println(b / a)
    fmt.Println(5 % 2)

    var h, i = true, false

    /** Relational Operators */
    fmt.Println(h == i)
    fmt.Println(h != i)
    fmt.Println(a > b)
    fmt.Println(a < b)
    fmt.Println(a >= b)
    fmt.Println(a <= b)

    /** Logical Operators */
    fmt.Println(h && i)
    fmt.Println(h || i)
    fmt.Println(!(h && i))

    x , y := 60, 35

    fmt.Println(x & y)
    fmt.Println(x | y)
    //fmt.Println(x^y)
    //fmt.Println(~x)
    fmt.Println(x << 2)
    fmt.Println(y >> 2)
}
