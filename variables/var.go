package main

import (
	"fmt"
	"reflect"
)

/**
 * Data-types in Go
 * ---------------
 * Go has four data-type categories:
 * 1. Boolean (bool) - consists of true and false
 * 2. Numeric - arithmetic types (aka numbers)
 * 3. String (string) - string of text
 * 4. Derived - includes pointers, arrays, structures, union, function, slice, interface, map and channel types
 */

/**
 * Numeric types
 * -------------
 * 1. Integer
 *     - unit8: unsigned 8-bit int (0 to 255)
 *     - unit16: unsigned 16-bit int (0 to 65535)
 *     - unit32: unsigned 32-bit int (0 to 4294967295)
 *     - unit64: unsigned 64-bit int (0 to 18446744073709551615)
 *     - int8: signed 8-bit int (-128 to 127)
 *     - int16: signed 16-bit int (-32768 to 32767)
 *     - int32: signed 32-bit int (-2147483648 to 2147483647)
 *     - int64: signed 64-bit int (-9223372036854775808 to 9223372036854775807)
 * 2. Floating
 *     - float32: 32-bit floating-point number
 *     - float64: 64-bit floating-point number
 *     - complex64: Complex numbers, with float32 real and imaginary parts
 *     - complex128: Complex numbers, with float64 real and imaginary parts
 * 3. Other
 *     - byte: same as unit8
 *     - rune: same as int32
 *     - unit: natural-sized unsigned int (32-bit for x86, 64-bit for x64)
 *     - int: natural-sized signed int
 *     - unitptr: an unsigned int to store the uninterpreted bits of a pointer value
 */
func main() {
	// Variables are defined by the following syntax: var identifier(s) optional_data_type
    var a, b , c = 1, .99, "foo" // Type is automatically determined during initialization
    
    // The type of each of these is dynamically set based upon the given input, kinda like JS
    fmt.Println("The value of a is" , a, "with a type of", reflect.TypeOf(a))
    fmt.Println("The value of b is" , b, "with a type of", reflect.TypeOf(b))
    fmt.Println("The value of c is" , c, "with a type of", reflect.TypeOf(c))

    var d, e int = 1, 2 // Forced type "int" on d and e
    fmt.Println("The value of d is" , d);
    fmt.Println("The value of e is" , e);

    var f, g float64 = .99, .5 // float64 == double
    fmt.Println("The value of f is", f);
    fmt.Println("The value of g is", g);

    var h, i bool = true, false
    fmt.Println("The value of h is", h);
    fmt.Println("the value of i is", i);

    var j, k string = "Hello", "World"
    l := " " // Shorthand for var l = " "
    fmt.Println(j+l+k);
}
