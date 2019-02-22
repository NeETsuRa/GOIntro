package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
	fmt.Println("A number from 1-100", rand.Intn(100))
}

func add(x float64, y float64) float64 { //func add(x, y float64)
	return x + y
}

func combine(a, b string) (string, string) {
	return a, b
}

func main() {
	// Diffrent possible Types:
	//	 bool
	//	 string
	//	 int  int8  int16  int32  int64
	//	 uint uint8 uint16 uint32 uint64 uintptr
	//	 byte // alias for uint8
	//	 rune // alias for int32
	//	 float32 float64
	//	 complex64 complex128

	//How to declare variables:
	//var num1 float64 = 5.6
	//var num2 float64 = 9.5
	//var num1, num2 float64 = 5.6, 9.5
	num1, num2 := 5.6, 9.5
	//var a int = 62
	//var b float64 = float64(a)

	w1, w2 := combine("Hey", "there")

	foo()
	fmt.Println(add(num1, num2))
	fmt.Println(w1, w2)

	//Pointers:
	x := 15
	a := &x         // point to x  (memory address)
	fmt.Println(a)  // prints out the mem addr.
	fmt.Println(*a) // read a through the pointer, so this will print out a value (15 in this case)

	*a = 5          // sets the value pointed at to 5, which means x is modified (since x is stored at the mem addr)
	fmt.Println(x)  // see the new value of x
	*a = *a * *a    // what is the value of x now?
	fmt.Println(x)  // prints a value
	fmt.Println(*a) // prints a value

	fmt.Println(&x) // prints a memory address
	fmt.Println(a)  // prints a memory address
}
