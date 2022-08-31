//            <p>Go&rsquo;s <em>select</em> lets you wait on multiple channel
//operations. Combining goroutines and channels with
//select is a powerful feature of Go.</p>
//
//            

package main

//            

import (
"fmt"
"time"
)

//            

func main () {

//            <p>For our example we&rsquo;ll select across two channels.</p>
//


c1 := make ( chan string )
c2 := make ( chan string )

//            <p>Each channel will receive a value after some amount
//of time, to simulate e.g. blocking RPC operations
//executing in concurrent goroutines.</p>
//


go func () {
time . Sleep ( 1 * time . Second )
c1 <- "one"
}()
go func () {
time . Sleep ( 2 * time . Second )
c2 <- "two"
}()

//            <p>We&rsquo;ll use <code>select</code> to await both of these values
//simultaneously, printing each one as it arrives.</p>
//


for i := 0 ; i < 2 ; i ++ {
select {
case msg1 := <- c1 :
fmt . Println ( "received" , msg1 )
case msg2 := <- c2 :
fmt . Println ( "received" , msg2 )
}
}
}

