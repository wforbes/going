// following along with youtube video
//
//	"Learn GO Fast: Full Tutorial"
//	https://www.youtube.com/watch?v=8uiZC0l4Ajw
//	This file includes content from 35:18 - 40:05
package main

import "fmt"

func main() {
	demoPointers()
	demoPointersWithFunctions()
}

func demoPointers() {
	// Pointer - a special type, they're variables
	//	that store memory locations instead of values
	//	like ints or floats do
	var p *int32 // the variable 'p' will hold the memory address of an int32 value
	// depending on your OS, the p variable will take up 32 bits or 64 bits
	//	for instance Windows 11 stores it as 64 bits or 8 bytes
	// 'p' starts as nil because we haven't initialized it
	//		(this pointer doesn't point to anything yet)
	//		(this pointer doesn't have an address assigned to it in which to store a int32 value yet)
	// We can give it an address (or memory location) with the new() func
	p = new(int32)
	// If we want to get the value stored at this memory location we can use the * symbol
	//	this is called 'dereferencing' the pointer
	fmt.Printf("The value p points to is: %v \n", *p)
	// we get a 0 because this is the default value of int32
	// Note: When you initialize a pointer with a memory location it
	//	zeroes out that memory location. It sets the value at that memory location
	//	to the zero value of that type. Like 0 for ints and floats, "" for strings, false for boolean

	// To change the value store at the memory location of a pointer we use * again
	*p = 10 // <-- (set the value at the memory location p is pointing to to 10)

	// Note: The * notation does double duty...
	// 		we use it to initialize the pointer
	//		we use it when we want to reference the value of the pointer
	//

	// a common headache is trying to set or get a nil pointer
	// if we dont initialize the pointer with a memory location...
	//-- var p2 *int32
	// It'll throw a run-time exception when trying to use it because we are trying
	// 	to work with a memory address that doesn't exist yet...
	//-- fmt.Printf("Trying to get the value stored at p2 memory location: %v \n", *p2)
	//-- *p2 = 10
	// So initialize it with a memory location instead..
	var p2 *int32 = new(int32)
	fmt.Printf("Trying to get the value stored at p2 memory location: %v \n", *p2)
	*p2 = 10
	fmt.Printf("Value of p2 after setting it: %v \n", *p2)

	// We can create a pointer from the address of another variable using
	//	the & symbol like this...
	var i int32
	fmt.Printf("The value of i is: %v \n", i)
	p2 = &i // the & symbol here means we want the memory address of the 'i' variable, not its value
	// Now both p2 and i reference the same int32 value in memory
	// What that means is...
	// if we use the * notation to change the value of p2
	*p2 = 1
	// the value of 'i' is also changed
	fmt.Printf("Since 1 was set on *p2 ... the value of 'i' is now also %v \n", i)

	// This is different than a regular variable
	var k int32 = 2
	i = k // assigning k into i will copy the value into the memory location of 'i'
	// so if we change either i or k value, the other will not change
	k = 24
	fmt.Printf("demonstrating that i is still %v, while k is now %v \n", i, k)

	// The main exception of this copy behavior of non-pointer variables
	//	is when working with slices
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice // lets copy a slice in a regular way without using a pointer
	// now lets try modifying the copy variable
	sliceCopy[2] = 4
	// and printing out both slices
	fmt.Println("The value of slice: ", slice)
	fmt.Println("The value of sliceCopy: ", sliceCopy)
	// You'll see that the value of the original slice HAS CHANGED
	// This is because under the hood, slices contain pointers to an underlying array
	// So with slices, we're just copying pointers when we do this
	//	So both variables refer to the same data
}

func demoPointersWithFunctions() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	// we have a function square() that takes an array of float64
	//	and squares all the values
	var result [5]float64 = square(thing1)
	fmt.Printf("The result is: %v \n", result)

	// When we pass 'thing1' into the function
	//	we're creating an exact copy of it inside the function to use as 'thing2'
	//	So this uses more memory than we need to...
	// 	Pointers can help optimize this
	//	We'll use function pointerSquare() that takes in a pointer and returns a pointer
	// pass it with & to say 'here is a pointer to the memory address of thing1'
	var newResult [5]float64 = pointerSquare(&thing1)
	fmt.Printf("The newResult is: %v \n", newResult)
}

func square(thing2 [5]float64) [5]float64 {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

func pointerSquare(thing2 *[5]float64) [5]float64 {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
