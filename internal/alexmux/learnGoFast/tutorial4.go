// following along with youtube video
//
//	"Learn GO Fast: Full Tutorial"
//	https://www.youtube.com/watch?v=8uiZC0l4Ajw
//	This file includes content from 19:30 - 26:35
package learnGoFast

import (
	"fmt"
	"time"
)

func RunT4() {
	demoArrays()
	demoSlices()
	demoMaps()
	demoLoops()

	perfTest()
}

func demoArrays() {
	fmt.Println("Arrays")
	// arrays
	//	Fixed length collection of data
	//	all of the same type that is indexable
	//	and stored in contiguous memory locations

	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	// int32 is 4 bytes of memory, array is 3 elements, size is 12 bytes (4 * 3)

	// we can print out the memory locations
	// of the array indicies with the & symbol
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// initialize the array with starting values
	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	intArr3 := [3]int32{4, 5, 6}
	intArr4 := [...]int32{7, 8, 42} // length inferred with ellipsis
	fmt.Println(intArr2)
	fmt.Println(intArr3)
	fmt.Println(intArr4)

	fmt.Println()
}

func demoSlices() {
	fmt.Println("Slices")
	// slices
	//	wrappers around arrays to give a more
	//	general, powerful, and convenient
	//	interface to sequences of data

	// omit the length digit to declare a slice
	var intSlice []int32 = []int32{100, 200, 300}
	fmt.Println(intSlice)

	// append function
	intSlice = append(intSlice, 400)
	// adds an item to the end of the slice, then returns the slice
	//	the system first checks the underlying array length
	//	then adds room enough for the item being added
	//	(with extra inaccessible elements at the end)
	fmt.Println(intSlice)

	// len() and cap()
	// we can check the length with len()
	//	and capacity of the slice with cap()
	intSliceBefore := []int32{11, 22, 33}
	fmt.Printf("The length of %v is %v with capacity %v\n", intSliceBefore, len(intSliceBefore), cap(intSliceBefore))
	fmt.Println("Appending one element to the slice...")
	intSliceAfter := append(intSliceBefore, 44)
	fmt.Printf("The length of %v is %v with capacity %v\n", intSliceAfter, len(intSliceAfter), cap(intSliceAfter))
	// although the array has 2 extra spaces of capacity
	// trying to access them with intSliceAfter[4] or intSliceAfter[5]
	//	will result in index out of bounds error

	// appending multiple items
	// we can append multiple new items with the ... (spread) operator
	newInts := []int32{55, 66}
	intSliceAfter2 := append(intSliceAfter, newInts...)
	fmt.Printf("Added two new items %v with one append call %v\n", newInts, intSliceAfter2)

	// make()
	// we can also use the make() function to create new slices
	intSlice3 := make([]int32, 3, 8) //<--- spotted an error in the video @ 22:50
	// we specify the length of the slice with the second arg
	//	and optionally specify the capacity of the slice with the third arg
	fmt.Println(intSlice3)

	fmt.Println()
}

func demoMaps() {
	fmt.Println("Maps")
	// map
	//	set of key/value pairs where you can look up a value by it's key
	var myMap map[string]uint8 = make(map[string]uint8)
	// ^ this means the keys are of type string and the values are of uint8
	fmt.Println(myMap)

	// you can also initialize a map with values immediately like this...
	var myMap2 = map[string]uint8{"Jimmy": 32, "Bob": 42}
	// so imagine this is a map of people with their age
	//	we can access someones age like this...
	fmt.Println(myMap2["Bob"])
	// if we try using a key that doesn't exist
	//	we get back the default value of the map's value type
	fmt.Println(myMap2["Potato"]) // uint8 default is 0

	// error checking
	//	since maps in Go always return a value
	//	its good to use the second optional error boolean return
	//	to check if the look up was successful
	var age, ok = myMap2["Jimmmmmmmmmy"]
	if ok {
		fmt.Printf("Jimmy's age is %d \n", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// delete() for deleting items
	delete(myMap2, "Jimmy")
	// this deletes by reference so no return value is given
	fmt.Println(myMap2)

	fmt.Println()
}

func demoLoops() {
	intArr := [...]int32{11, 22, 33, 44, 55, 66}
	myMap := map[string]uint8{"Adam": 23, "Sarah": 42}

	fmt.Println("For/Range loops")
	// for/range loop
	//	using the range keyword within a for loop
	//	allows initializing a variable for each iteration
	//	of the loop using the walrus operator
	fmt.Println("Iterating over a map")
	for name := range myMap {
		fmt.Printf("Name: %v \n", name)
	}
	// Note: when iterating over a map, order is not preserved
	//	so running this multiple times may result in the map items
	//	appearing in different orders

	fmt.Println("Iterating over a map, with the key and value")
	// we can also get the values of the map with a second optional for variable
	for name, age := range myMap {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	fmt.Println("Iterating over an array")
	// we can iterate over an array in a similar way
	//	where the first for variable is index
	//	and the second is the value at that index
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}
	fmt.Println()

	fmt.Println("Conditional for loops")
	// there is no while loop in Go
	// instead this can be achieved with conditional for loop
	var i int = 0
	for i < 10 {
		fmt.Println(i) // cant increment within print statement? lame.
		i++
	}

	fmt.Println("infinite for loop with break clause")
	// we can also define an infinite for loop (without condition)
	//	and place a break clause inside instead
	i = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	fmt.Println("traditional for loop")
	// of course go supports the traditional for loop
	//	with parts: initialization, condition, and post
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("For int range loop")
	// we can use range keyword to specify a number of iterations
	for range 10 {
		fmt.Println("iterating over range 10")
	}
	// and get the iteration number with a for var
	for i := range 10 {
		fmt.Printf("%d: iterating over range 10 \n", i)
	}
}

func perfTest() {
	// demonstration of the performance impact
	//	of using a slice with a predefined capacity
	//	compared to a slice without one
	var n int = 1000000
	var testSlice = []int{}            // no capacity set
	var testSlice2 = make([]int, 0, n) // capacity set exactly

	fmt.Printf("Total time without preallocation: %v \n", timeLoop(testSlice, n)) // 5.1756ms
	fmt.Printf("Total time with preallocation: %v \n", timeLoop(testSlice2, n))   // 1.0199ms
	// appears to be 5x faster
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
