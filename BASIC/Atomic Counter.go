//for managing state though. Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
package main
//Unsigned can hold a larger positive value, and no negative value.
// to manage simple counter state using atomic operations.
import (
	"fmt"
	"time"
	"sync/atomic"
)
func main() {
	var ops uint64 //We’ll use an unsigned integer to represent our (always-positive) counter.
	for i := 0; i <= 50; i++ { //To simulate concurrent updates, we’ll start 50 goroutines that each increment the counter about once a millisecond.
		go func() {
			for {
				atomic.AddUint64(&ops, 1) //To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax.
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second) //Wait a second to allow some ops to accumulate.


	opsFinal := atomic.LoadUint64(&ops) //In order to safely use the counter while it’s still being updated by other goroutines,
	// we extract a copy of the current value into opsFinal via LoadUint64.
	// As above we need to give this function the memory address &ops from which to fetch the value
	fmt.Println("ops:", opsFinal)
}