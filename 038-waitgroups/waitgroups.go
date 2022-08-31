//            <p>To wait for multiple goroutines to finish, we can
//use a <em>wait group</em>.</p>
//
//            

package main

//            

import (
"fmt"
"sync"
"time"
)

//            <p>This is the function we&rsquo;ll run in every goroutine.</p>
//


func worker ( id int ) {
fmt . Printf ( "Worker %d starting\n" , id )

//            <p>Sleep to simulate an expensive task.</p>
//


time . Sleep ( time . Second )
fmt . Printf ( "Worker %d done\n" , id )
}

//            

func main () {

//            <p>This WaitGroup is used to wait for all the
//goroutines launched here to finish. Note: if a WaitGroup is
//explicitly passed into functions, it should be done <em>by pointer</em>.</p>
//


var wg sync . WaitGroup

//            <p>Launch several goroutines and increment the WaitGroup
//counter for each.</p>
//


for i := 1 ; i <= 5 ; i ++ {
wg . Add ( 1 )

//            <p>Avoid re-use of the same <code>i</code> value in each goroutine closure.
//See <a href="https://golang.org/doc/faq#closures_and_goroutines">the FAQ</a>
//for more details.</p>
//


i := i

//            <p>Wrap the worker call in a closure that makes sure to tell
//the WaitGroup that this worker is done. This way the worker
//itself does not have to be aware of the concurrency primitives
//involved in its execution.</p>
//


go func () {
defer wg . Done ()
worker ( i )
}()
}

//            <p>Block until the WaitGroup counter goes back to 0;
//all the workers notified they&rsquo;re done.</p>
//


wg . Wait ()

//            <p>Note that this approach has no straightforward way
//to propagate errors from workers. For more
//advanced use cases, consider using the
//<a href="https://pkg.go.dev/golang.org/x/sync/errgroup">errgroup package</a>.</p>
//


}

