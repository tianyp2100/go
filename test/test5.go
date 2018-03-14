package main

import "fmt"

func main(){
 var a int = 21
 var b int = 10
 var c int

 c = a + b
 fmt.Printf("a + b = %d\n", c)
 c = a - b
 fmt.Printf("a - b = %d\n", c)
 c = a * b
 fmt.Printf("a * b = %d\n", c)
 c = a / b 
 fmt.Printf("a / b = %d\n", c)
 c = a % b
 fmt.Printf("a求余b = %d\n", c)
 a++
 fmt.Printf("a++ %d\n", a)
 a = 21
 a--
 fmt.Printf("a-- %d\n", a)
}
