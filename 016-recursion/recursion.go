//            <p>Go supports
//<a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a>.
//Here&rsquo;s a classic example.</p>
//
//            

package main

//            

import "fmt"

//            <p>This <code>fact</code> function calls itself until it reaches the
//base case of <code>fact(0)</code>.</p>
//


func fact ( n int ) int {
if n == 0 {
return 1
}
return n * fact ( n - 1 )
}

//            

func main () {
fmt . Println ( fact ( 7 ))

//            <p>Closures can also be recursive, but this requires the
//closure to be declared with a typed <code>var</code> explicitly
//before it&rsquo;s defined.</p>
//


var fib func ( n int ) int

//            

fib = func ( n int ) int {
if n < 2 {
return n
}

//            <p>Since <code>fib</code> was previously declared in <code>main</code>, Go
//knows which function to call with <code>fib</code> here.</p>
//


return fib ( n - 1 ) + fib ( n - 2 )
}

//            

fmt . Println ( fib ( 7 ))
}

