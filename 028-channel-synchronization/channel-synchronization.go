//            <p>We can use channels to synchronize execution
//across goroutines. Here&rsquo;s an example of using a
//blocking receive to wait for a goroutine to finish.
//When waiting for multiple goroutines to finish,
//you may prefer to use a <a href="waitgroups">WaitGroup</a>.</p>
//
//            

package main

//            

import (
"fmt"
"time"
)

//            <p>This is the function we&rsquo;ll run in a goroutine. The
//<code>done</code> channel will be used to notify another
//goroutine that this function&rsquo;s work is done.</p>
//


func worker ( done chan bool ) {
fmt . Print ( "working..." )
time . Sleep ( time . Second )
fmt . Println ( "done" )

//            <p>Send a value to notify that we&rsquo;re done.</p>
//


done <- true
}

//            

func main () {

//            <p>Start a worker goroutine, giving it the channel to
//notify on.</p>
//


done := make ( chan bool , 1 )
go worker ( done )

//            <p>Block until we receive a notification from the
//worker on the channel.</p>
//


<- done
}

