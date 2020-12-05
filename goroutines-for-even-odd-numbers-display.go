package main

import (
	"sync"
	"fmt"
)

/*

During a coding interview - I was tasked to craft below synchronized output using thread :

Thread-1 :  1
Thread-2 :  2
Thread-1 :  3

<snip>

Thread-2 :  198
Thread-1 :  199
Thread-2 :  200

*/

var mutex *sync.Mutex // global mutex for goroutines/thread synchronization
var currentValue int // value to be displayed 
var max int = 200 // max value for testing purpose 

// thread 1 - display odd numbers
func printOddNumbers() {

	nextValue := 0
	for currentValue < (max - 1) {
		mutex.Lock()
		nextValue = currentValue + 1
		if nextValue % 2 != 0 {
			fmt.Println("Thread-1 : ", nextValue)
			currentValue = nextValue
		}

		mutex.Unlock()
	}
}

// thread 2 - display even numbers
func printEvenNumbers() {

	nextValue := 0
	for currentValue < max {
		mutex.Lock()
		nextValue = currentValue + 1
		if nextValue % 2 == 0 {
			fmt.Println("Thread-2 : ", nextValue)
			currentValue = nextValue
		}

		mutex.Unlock()
	}
}
	

func main() {

 	mutex = &sync.Mutex{}
 	go printOddNumbers() // start thread-1
	go printEvenNumbers() // start thread-2

	fmt.Scanln() // wait for user to type Enter to exit
}