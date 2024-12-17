package main

import "fmt"

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these as both binary and hexadecimal
	a,b,c,d,e,f := 0,1,2,3,4,5
	fmt.Printf("a-f as binary, %b %b %b %b %b %b \n", a,b,c,d,e,f)
	fmt.Printf("a-f as hexadecimal, %x %x %x %x %x %x\n", a,b,c,d,e,f)
}