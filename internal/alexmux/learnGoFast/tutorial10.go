// following along with youtube video
//
//	"Learn GO Fast: Full Tutorial"
//	https://www.youtube.com/watch?v=8uiZC0l4Ajw
//	This file includes content from 52:50 - 55:42
package learnGoFast

import (
	"encoding/json"
	"fmt"
	"os"
)

func RunT10() {
	demoGenericFuncs()
	demoAnyType()
	demoUninferredGenerics()
	demoGenericStructs()
}

// Generics
//
//	allows you to define multiple different types
//	for the same parameter on a function, so that
//	you can reuse the same function with different
//	types of arguments
func demoGenericFuncs() {
	var intSlice = []int{1, 2, 3}
	// we can call sumSlice with a slice of ints
	fmt.Println(sumSlice(intSlice))

	var float32Slice = []float32{1, 2, 3}
	// and we can call it with a slice of float32s
	fmt.Println(sumSlice(float32Slice))
}

// using this bracket notation between the function name
// and the parameter parenthesis, everywhere T appears it
// represents whatever the type the parameter value was that
// got passed to the function at runtime
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Any Type
//
//	we can use this in some situations with Generics
//	to express that the type of the Generic could be anything
//	we can't use it in the last example because not every type
//	can work with the += operator
//	here's an example of a situation we can use Any
func demoAnyType() {
	var intSlice = []int{}
	fmt.Println(isEmpty(intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(isEmpty(float32Slice))
}

func isEmpty[T any](slice []T) bool {
	// ^ [T any] tells the compiler that this could be a slice of anything
	//	and since we're doing work on the slice (checking length) that's fine
	return len(slice) == 0
}

// Uninferred Generic calls
//
//	The previous two examples showed using Generics without defining what the
//	input argument type is when calling the function. This works sometimes
//	because the compiler can infer what the type is going to be when it gets
//	to the generic function.
//	Sometimes it can't do that, and we have to explicitly tell the function what
//	type is going into it. Here's an example:
//	We have two struct types and two .json files. We'll load the content of those
//	json files and unmarshal it into the struct types
type contactInfo struct {
	Name  string
	Email string
}
type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func demoUninferredGenerics() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	// ^ here we have to specify the type we need returned with brackets in the func call
	fmt.Printf("%+v \n", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("%+v \n", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, _ := os.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

// Generic Structs
//
// allows for defining generic types on struct fields
//
//	so that the struct can have a field with an interchangable type
func demoGenericStructs() {
	type gasEngine struct {
		gallons float32
		mpg     float32
	}
	type electricEngine struct {
		kwh   float32
		mpkwh float32
	}
	// we set generic T as either gasEngine or electricEngine
	type car[T gasEngine | electricEngine] struct {
		make   string
		model  string
		engine T // then the engine field can be either of those types
	}

	// then when using the car struct we just define what type of engine it needs
	// with the bracket notation, and theres no need for two different car structs
	var gasCar = car[gasEngine]{
		make:  "Honda",
		model: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}
	fmt.Printf("The %s %s holds $%.1f gallons of gas \n", gasCar.make, gasCar.model, gasCar.engine.gallons)

	var electricCar = car[electricEngine]{
		make:  "Tesla",
		model: "DumpsterTruck",
		engine: electricEngine{
			kwh:   57.2,
			mpkwh: 4.17,
		},
	}
	fmt.Printf("The %s %s holds $%.1f kilowatts of charge \n", electricCar.make, electricCar.model, electricCar.engine.kwh)

}
