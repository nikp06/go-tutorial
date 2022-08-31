//            <p>In the previous example we saw how to manage simple
//counter state using <a href="atomic-counters">atomic operations</a>.
//For more complex state we can use a <a href="https://en.wikipedia.org/wiki/Mutual_exclusion"><em>mutex</em></a>
//to safely access data across multiple goroutines.</p>
//
//            

package main

//            

import (
"fmt"
"sync"
)

//            <p>Container holds a map of counters; since we want to
//update it concurrently from multiple goroutines, we
//add a <code>Mutex</code> to synchronize access.
//Note that mutexes must not be copied, so if this
//<code>struct</code> is passed around, it should be done by
//pointer.</p>
//


type Container struct {
mu sync . Mutex
counters map [ string ] int
}

//            <p>Lock the mutex before accessing <code>counters</code>; unlock
//it at the end of the function using a <a href="defer">defer</a>
//statement.</p>
//

func ( c * Container ) inc ( name string ) {

//            

c . mu . Lock ()
defer c . mu . Unlock ()
c . counters [ name ] ++
}

//            <p>Note that the zero value of a mutex is usable as-is, so no
//initialization is required here.</p>
//

func main () {
c := Container {

//            

counters : map [ string ] int { "a" : 0 , "b" : 0 },
}

//            

var wg sync . WaitGroup

//            <p>This function increments a named counter
//in a loop.</p>
//


doIncrement := func ( name string , n int ) {
for i := 0 ; i < n ; i ++ {
c . inc ( name )
}
wg . Done ()
}

//            <p>Run several goroutines concurrently; note
//that they all access the same <code>Container</code>,
//and two of them access the same counter.</p>
//


wg . Add ( 3 )
go doIncrement ( "a" , 10000 )
go doIncrement ( "a" , 10000 )
go doIncrement ( "b" , 10000 )

//            <p>Wait a for the goroutines to finish</p>
//


wg . Wait ()
fmt . Println ( c . counters )
}

