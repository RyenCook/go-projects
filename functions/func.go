package main

import "fmt"

func main() {
   var a int = 100
   var b int = 200
   var ret = max(a, b)

   fmt.Printf("Max value is : %d\n", ret)
   fmt.Println(a, b)

   a, b = swap(a,b)

   fmt.Println("After swap:",a, b)
}

/* Unlike most C-fam langs, Go can actually use functions declared AFTER main */
func max(num1, num2 int) int {
   /* local variable declaration */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}

/* Also unlike most C-fam langs Go can RETURN MULTIPLE VALUES */
func swap(x, y int) (int, int) {
   return y, x
}
