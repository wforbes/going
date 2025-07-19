// following along with youtube video
//
//	"Learn GO Fast: Full Tutorial"
//	https://www.youtube.com/watch?v=8uiZC0l4Ajw
//	This file includes content from 31:06 - 35:18
package learnGoFast

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func RunT6() {
	demoStructTypes()
	demoAnonymousStructs()
	demoStructMethods()
	demoInterfaces()
}

func demoStructTypes() {
	fmt.Println("Struct Types")

	// fields in a struct are initialized to their default values
	var engine1 gasEngine
	fmt.Printf("engine1 fields with default values: mpg=%v, gallons=%v \n", engine1.mpg, engine1.gallons)

	// we can initialize a struct with the struct literal syntax
	var engine2 gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Printf("engine2 fields with initial values set with struct literal syntax: mpg=%v, gallons=%v \n", engine2.mpg, engine2.gallons)

	var engine3 gasEngine = gasEngine{25, 15}
	// without field names the values are set in order of the type definition
	fmt.Printf("engine3 fields with initial values set with without field names: mpg=%v, gallons=%v \n", engine3.mpg, engine3.gallons)
	engine3.mpg = 20
	fmt.Printf("updated engine3 with direct field assignment (engine3.mpg = %v)\n", engine3.mpg)

	type owner struct {
		name string
	}
	// we can set fields in structs as anything, even other structs
	type gasEngineWithOwner struct {
		mpg       uint8
		gallons   uint8
		ownerInfo owner
	}

	var engineWithOwner gasEngineWithOwner = gasEngineWithOwner{25, 15, owner{"Bob"}}
	fmt.Printf(
		"engineWithOwner set with values: mpg=%v, gallons=%v, owner=%v \n",
		engineWithOwner.mpg, engineWithOwner.gallons, engineWithOwner.ownerInfo.name,
	)

	// we can also set the field with the struct as a direct ref to the child struct
	type gasEngineWithOwner2 struct {
		mpg     uint8
		gallons uint8
		owner
	}

	var engineWithOwner2 gasEngineWithOwner2 = gasEngineWithOwner2{25, 15, owner{"Bob"}}
	fmt.Printf(
		"engineWithOwner2 set with values: mpg=%v, gallons=%v, owner=%v \n",
		engineWithOwner2.mpg,
		engineWithOwner2.gallons,
		engineWithOwner2.owner, // this lets us just reference the field without dotting into it
		// (i guess since the owner struct has only one field, we just get "Bob" returned here)
	)

	// we can also set fields as just a primitive data type
	type structWithInt struct {
		int // struct has a field named 'int', with a data type of 'int'
	}
	var someNumber structWithInt = structWithInt{42}
	// but it looks like you can't dot into get its value like "someNumber.int"
	fmt.Printf("someNumber struct has an field named 'int' that is data type 'int'; someNumber.int : %v \n", someNumber.int)

	fmt.Println()
}

func demoAnonymousStructs() {
	fmt.Println("Anonymous Structs")
	// instead of defining a struct type we can
	//	define an anonymous struct into a variable
	var anonEngine = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Printf("anonEngine set as Anonymous Struct with values: (mpg=%v, gallons=%v)\n", anonEngine.mpg, anonEngine.gallons)

	// The main issue with this is that anon structs aren't as easily reusable as Struct Types

	fmt.Println()
}

func demoStructMethods() {
	// see the 'milesLeft()' method added to gasEngine struct below this 'demoStructMethods()' function
	var myEngine gasEngine = gasEngine{25, 15}
	// we can call the method on the struct like this...
	fmt.Printf("Total miles left in the tank: %v \n", myEngine.milesLeft())

	// (see the canMakeIt function)
	// we can pass this type into functions to use methods or fields on it
	fmt.Println("Q: do we have enough fuel for a 200 mile trip?")
	fmt.Printf("A: %v \n", canMakeIt(myEngine, 200))
	fmt.Println("Q: what about a 100 mile trip?")
	fmt.Printf("A: %v \n", canMakeIt(myEngine, 100))
}

//	Struct Methods - functions that are directly tied to a struct
//		and have access to the struct instance itself.
// 		Here the (e gasEngine) is giving the method access to the
// 		struct instance which we can use to manipulate field values
//		and return data based on them
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func canMakeIt(e gasEngine, miles uint8) string {
	if miles <= e.milesLeft() {
		return "You can make it there!"
	}
	return "Need to fuel up first!"
}

// Interfaces
//  - What if we want to make different types of engines?
//	we can make 'electricEngine' struct (see below)
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

//	and we can make a milesLeft() method that works
//	with our new 'electricEngine' struct
func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

//	- But we need an interface to allow 'canMakeIt()'
//	to work with both gasEngines and electricEngines
func canMakeIt2(e engine, miles uint8) string {
	if miles <= e.milesLeft() {
		return "You can make it there!"
	}
	return "Need to fuel up first!"
}

//	- So we create a 'engine' interface that requires a milesLeft2()
//	method and returns a uint8
type engine interface {
	milesLeft() uint8
}

func demoInterfaces() {
	fmt.Println("Interfaces")
	// see the above struct, method, func, and interface
	//	this pattern lets us use canMakeIt2 with multiple types of engines
	//	whether they're gas, electric, or something else
	var gasEngine gasEngine = gasEngine{25, 15}
	fmt.Printf("Can 'gasEngine' make the 200 mile trip? (%s) \n", canMakeIt2(gasEngine, 150))
	var electricEngine electricEngine = electricEngine{5, 82}
	fmt.Printf("Can 'electricEngine' make the 200 mile trip? (%s) \n", canMakeIt2(electricEngine, 150))

}
