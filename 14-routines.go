package main

/* Googol is 1 multiplied by 10 in goroutines stopping when it's Googol.

   Primary takeaways:

   - Goroutines cannot call the function they are contained within. No
     recursion then.
   - Parallel execution is hella fast, especially for simple functions
     (like increaseGoogol).
   - Global variables are not guarded in any way, thus mere goroutines
     are not safe by themselves.
   - Goroutines are better used with channels (this I did not do yet).
 */

import "fmt"

var exit bool = false
var googol float64 = 1.0

func increaseGoogol () {
	if googol >= 1.000000e+100 {
		exit = true
		return
	}
	googol *= 10
}

func main() {
	for !exit {
		go increaseGoogol()
	}
	fmt.Printf("%e\n", googol)
}
