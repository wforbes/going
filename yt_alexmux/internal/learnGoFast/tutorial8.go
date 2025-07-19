// following along with youtube video
//
//	"Learn GO Fast: Full Tutorial"
//	https://www.youtube.com/watch?v=8uiZC0l4Ajw
//	This file includes content from 40:05 - 47:10
package learnGoFast

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var wg = sync.WaitGroup{}

var results = []string{}

func cleanUpResults() {
	results = []string{}
}

var m = sync.Mutex{}
var rwm = sync.RWMutex{}

func RunT8() {
	// Goroutine - A way to execute multiple functions
	//	that run concurrently
	demo1Sequential()
	demo2GoKeyword()
	demo3WriteProblem()
	demo4Mutex()
	demo5RWMutex()
	demo6PerfTest1()
	demo7PerfTest2()
}

// Here we're simulating a call to a database
// using the dbData array and a sleep time delay
// This is the baseline, Nothing concurrent is happening
func demo1Sequential() {
	fmt.Println("Baseline sequential execution")
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall1(i)
	}
	fmt.Printf("Total execution time: %v \n", time.Since(t0)) // ( 4.2827043s )
}
func dbCall1(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
}

// Go keyword and WaitGroups
//
//		Using the 'go' keyword in combination with a 'WaitGroup'
//		allows for concurrent execution of a function in a way that
//		feels similar to javascripts async/await promise resolution
//	 1. add 1 to wait group counter before each 'go' call
//	 2. using go keyword before a func call spawns a task in the background for it
//	    similar to javascript's 'await' ?
//	 3. calling done on wait group when each operation is done
//	    similar to javascript's Promise.resolve() ?
//	 4. calling Wait on the WaitGroup waits for all WG items to be done before continuing
//	    similar to javascript's Promise.all() ?
func demo2GoKeyword() {
	fmt.Println("Using Go/WaitGroup execution")
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)     // 1. add 1 to wait group counter before each 'go' call
		go dbCall2(i) // 2. go keyword spawns this task in the background
	}
	// 4. wait for all WaitGroup items to be done
	wg.Wait()
	fmt.Printf("Total execution time: %v \n", time.Since(t0)) // ( 1.6714907s )
}

func dbCall2(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	wg.Done() // 3. call done on wait group when each operation is done
}

// Problem with Writing to Shared Memory Location
//
//	A problem with multiple processes writing to
//	the same memory location at the same time is
//	demonstrated here with all the go routines trying
//	to write to the results slice simultaneously.
//	I didn't see any problems on my machine, but
//	this can lead to corrupt memory and other issues
//	and we shouldn't do it. Demo4 shows the fix
func demo3WriteProblem() {
	fmt.Println("Writing to shared array with Go/WaitGroup execution")
	defer cleanUpResults()
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)     // 1. add 1 to wait group counter before each 'go' call
		go dbCall3(i) // 2. go keyword spawns this task in the background
	}
	// 4. wait for all WaitGroup items to be done
	wg.Wait()
	fmt.Printf("Total execution time: %v \n", time.Since(t0)) // ( 2.0005831s )
	fmt.Printf("The results are %v \n", results)
}
func dbCall3(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	results = append(results, dbData[i])
	wg.Done() // 3. call done on wait group when each operation is done
}

// Mutex - short for 'mututal exclusion', contains
//
//	a 'Lock' and 'Unlock' method that causes a go routine
//	to check to see if another go routine has locked the mutex
//	if it has, then it will wait until it is unlocked before
//	continuing through the lock point. essentially causing
//	all the related go routines to take their turn through
//	that line of code, eliminating the Demo3 problem
//	 1. add 1 to wait group counter before each 'go' call
//	 2. using go keyword before a func call spawns a task in the background for it
//	 3. lock the mutex to ensure multiple goroutines dont modify the results slice at the same time
//	 4. unlock the mutex because this goroutine is done allowing others to obtain a lock
//	 5. calling done on wait group when each operation is done
//	 6. calling Wait on the WaitGroup waits for all WG items to be done before continuing
func demo4Mutex() {
	fmt.Println("Using Mutex with Go/WaitGroup execution")
	defer cleanUpResults()
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)     // 1. add 1 to wait group counter before each 'go' call
		go dbCall4(i) // 2. go keyword spawns this task in the background
	}
	// 6. wait for all WaitGroup items to be done
	wg.Wait()
	fmt.Printf("Total execution time: %v \n", time.Since(t0)) // ( 2.0003598s )
	fmt.Printf("The results are %v \n", results)
}
func dbCall4(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock() // 3. lock the mutex to ensure multiple goroutines dont modify
	//	the results slice at the same time
	results = append(results, dbData[i])
	m.Unlock() // 4. unlock the mutex because this goroutine is done
	// allowing others to obtain a lock
	wg.Done() // 5. call done on wait group when each operation is done
}

// Using RWMutex for specific Read or Write locks
//
//	The RWMutex expands Mutex by adding a RLock (read lock) and RUnlock (read unlock)
//	method. This helps ensure that a goroutine won't read from a memory location at the
//	same time as another goroutine is writing to it. Multiple goroutines may obtain a
//	a RLock at the same time, this is safe because reading is not altering the memory
//	location contents. a RLock will cause a Lock to wait for RUnlock, but a RLock wont
//	make another RLock wait
func demo5RWMutex() {
	fmt.Println("Using a Read/Write Mutex with Go/WaitGroup execution")
	defer cleanUpResults()
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)     // 1. add 1 to wait group counter before each 'go' call
		go dbCall5(i) // 2. go keyword spawns this task in the background
	}
	// 6. wait for all WaitGroup items to be done
	wg.Wait()
	fmt.Printf("Total execution time: %v \n", time.Since(t0)) // ( 2.0005823s )
	fmt.Printf("The results are %v \n", results)
}
func dbCall5(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log_t8()
	wg.Done() // 5. call done on wait group when each operation is done
}
func save(result string) {
	rwm.Lock() // 3. lock the mutex to ensure multiple goroutines dont modify
	//	the results slice at the same time
	results = append(results, result)
	rwm.Unlock() // 4. unlock the mutex because this goroutine is done
	// allowing others to obtain a lock
}
func log_t8() {
	rwm.RLock() // 5. lock the mutex to read from it, blocking Lock() calls
	fmt.Printf("The current results are: %v \n", results)
	rwm.RUnlock() // 6. unlock the mutex from reading, allowing Lock() calls again
}

// Performance Test 1
//
//	Simplifying the code to remove the 'dbCall' and 'results' related
//	lines, and setting 1000 iterations results in the execution time
//	still taking only as long as the sleepTime. This is because while
//	the process is waiting for the sleep to finish, it can start the next
//	goroutine(s). Going further, looping through increments of 1k to 10k,
//	10k to 100k, then 100k to 1million iterations results in roughly the
//	same execution time... tied to the sleepTime.
//	Note: 100k+ iterations started adding 0.10~ second, up to 2.6 sec.
//	my guess is this is the result of garbage collection?
func demo6PerfTest1() {
	fmt.Println("Testing performance of scaling empty slept goroutines from 1k to 1mil")
	perfTest_8(42, 1000, 2000)
	for i := range 10 {
		if i == 0 {
			continue
		}
		perfTest_8(i, (i+1)*1000, 2000)
	}
	for i := range 10 {
		if i == 0 {
			continue
		}
		perfTest_8(i, (i+1)*10000, 2000)
	}
	for i := range 10 {
		if i == 0 { // we already tested 100k in prev loop
			continue
		}
		perfTest_8(i, (i+1)*100000, 2000)
	}

}
func perfTest_8(testNum int, iterations int, sleepTime float32) {
	fmt.Printf("PF%v - Iterations: [ %v ] - Sleep Time: [ %v ]\n", testNum, iterations, sleepTime)
	t0 := time.Now()
	for i := 0; i < iterations; i++ {
		wg.Add(1)             // 1. add 1 to wait group counter before each 'go' call
		go dbCall6(sleepTime) // 2. go keyword spawns this task in the background
	}
	// 4. wait for all WaitGroup items to be done
	wg.Wait()
	fmt.Printf("PF%v - Total execution time: %v \n", testNum, time.Since(t0))
}
func dbCall6(sleepTime float32) {
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	wg.Done() // 3. call done on wait group when each operation is done
}

// Performance Test 2
//
//		In PerfTest1 there is no work being done in the goroutines, they just wait 2 seconds each.
//		Here in PerfTest2 each goroutine increments a variable 100 million times, which causes a
//	 linear increase in execution time as the number of times it is done increases.
func demo7PerfTest2() {
	fmt.Println("Testing performance of expensive goroutines, scaling quantity from 100 to 1k")
	for i := range 10 {
		perfTest2_8(i, (i+1)*100)
	}
}
func perfTest2_8(testNum int, iterations int) {
	fmt.Printf("PF%v - Iterations: [ %v ] \n", testNum, iterations)
	t0 := time.Now()
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("PF%v - Total execution time: %v \n", testNum, time.Since(t0))
}
func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res += 1
	}
	wg.Done()
}
