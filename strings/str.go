package main

import "fmt"

func main() {
    var str1, str2 = "Hello", "World"
    fmt.Println(str1)
    fmt.Println(str2)
    // Concatenation by ',' automatically adds a space, for some reason.
    fmt.Println(str1,str2)

    /** Using golang's printing "verbs" one can customized printf's output:
     * %v  prints out the source-code that would produce said value
     * %v+ if the value is as truct, prints out the struct's field name
     * %T prints out a value's type
     * %t is for formatting booleans???
     * %d prints out in base10 formatting
     * %b prints a binary representation
     * %c prints the character corresponding to a given integer
     * %x prints a hexadecimal representation, ints and strings
     * %f decimal format for a float
     * %E scientific format for a float
     * %e scientific format for a float, with a lowercase e
     * %s basic string output
     * %q double-quotes strings in source
     * %p prints value's pointer
     * 
     */
    for i := 0; i < len(str1); i++ {
        fmt.Printf("%x ", str1[i])
    }

    fmt.Println()
}
