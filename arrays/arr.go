package main

import (
    "fmt"
)

func main() {
    var arr [10]int;
    var i, j int
    /* initialize elements of array n to 0 */         
    for i = 0; i < 10; i++ {
      arr[i] = i + 100 /* set element at location i to i + 100 */
    }
    /* output each array element's value */
    for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, arr[j] )
    }
}
