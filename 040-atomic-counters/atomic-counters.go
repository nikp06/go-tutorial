//            <p>The primary mechanism for managing state in Go is
//communication over channels. We saw this for example
//with <a href="worker-pools">worker pools</a>. There are a few other
//options for managing state though. Here we&rsquo;ll
//look at using the <code>sync/atomic</code> package for <em>atomic
//counters</em> accessed by multiple goroutines.</p>
//
//            

package main

//            

import (
"fmt"
"sync"
"sync/atomic"
)

//            

func main () {

//            <p>We&rsquo;ll use an unsigned integer to represent our
//(always-positive) counter.</p>
//


var ops uint64

//            <p>A WaitGroup will help us wait for all goroutines
//to finish their work.</p>
//


var wg sync . WaitGroup

//            <p>We&rsquo;ll start 50 goroutines that each increment the
//counter exactly 1000 times.</p>
//


for i := 0 ; i < 50 ; i ++ {
wg . Add ( 1 )

//            <p>To atomically increment the counter we
//use <code>AddUint64</code>, giving it the memory
//address of our <code>ops</code> counter with the
//<code>&amp;</code> syntax.</p>
//

go func () {
for c := 0 ; c < 1000 ; c ++ {

//            

atomic . AddUint64 ( & ops , 1 )
}
wg . Done ()
}()
}

//            <p>Wait until all the goroutines are done.</p>
//


wg . Wait ()

//            <p>It&rsquo;s safe to access <code>ops</code> now because we know
//no other goroutine is writing to it. Reading
//atomics safely while they are being updated is
//also possible, using functions like
//<code>atomic.LoadUint64</code>.</p>
//


fmt . Println ( "ops:" , ops )
}

