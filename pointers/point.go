package main

import "fmt"

func main() {
    var a = 20
    var b *int

    b = &a

    fmt.Printf("Address of a variable: %x\n", &a  )

    /* address stored in pointer variable */
    fmt.Printf("Address stored in b variable: %x\n", b )

    /* access the value using the pointer */
    fmt.Printf("Value of *b variable: %d\n", *b )

    /* Null Pointer, contains no value */
    var c *int
    fmt.Printf("Value of c variable: %d\n", c)

    if(c != nil) {
        fmt.Printf("c not nil\n")
    }
    if(c == nil) {
        fmt.Printf("c def nil\n")
    }
}

