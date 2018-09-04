//Mutex:For more complex state we can use a mutex to safely access data across multiple goroutines.
//In computer science, mutual exclusion is a property of concurrency control, which is instituted for the purpose of preventing race conditions;
//It is the requirement that one thread of execution never enter its critical section at the same time that another concurrent thread of execution enters its own critical section
//A race condition or race hazard is the behavior of an electronics, software, or other system where the output is dependent on the sequence or timing of other uncontrollable events.
package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
	"sync/atomic"
)

func main(){
	var  state  = make(map[int]int)  //the state will be a map.
	var mutex = &sync.Mutex{}        //This mutex will synchronize access to state.
	var readOps uint64               //We’ll keep track of how many read and write operations we do.
	var writeOps uint64
	for r := 0; r < 100; r++ {        //Here we start 100 goroutines to execute repeated reads against the state, once per millisecond in each goroutine.
	// For each read we pick a key to access, Lock() the mutex to ensure exclusive access to the state, read the value at the chosen key, Unlock() the mutex, and increment the readOps count.
	go func() {
			total := 0
			for {key := rand.Intn(5)
				mutex.Lock()
				total += state[key] //
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond) //Wait a bit between reads.
			}
		}()
	}
	for w := 0; w < 10; w++ {    //We’ll also start 10 goroutines to simulate writes, using the same pattern we did for reads
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

time.Sleep(time.Second)  //Let the 10 goroutines work on the state and mutex for a second.


	readOpsFinal := atomic.LoadUint64(&readOps)  //Take and report final operation counts.
fmt.Println("readOps:", readOpsFinal)
writeOpsFinal := atomic.LoadUint64(&writeOps)
fmt.Println("writeOps:", writeOpsFinal)
	mutex.Lock() //With a final lock of state, show how it ended up.
	fmt.Println("state:", state)
	mutex.Unlock()
}