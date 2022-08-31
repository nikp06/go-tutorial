//            <p><em>Timeouts</em> are important for programs that connect to
//external resources or that otherwise need to bound
//execution time. Implementing timeouts in Go is easy and
//elegant thanks to channels and <code>select</code>.</p>
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

//            <p>For our example, suppose we&rsquo;re executing an external
//call that returns its result on a channel <code>c1</code>
//after 2s. Note that the channel is buffered, so the
//send in the goroutine is nonblocking. This is a
//common pattern to prevent goroutine leaks in case the
//channel is never read.</p>
//


c1 := make ( chan string , 1 )
go func () {
time . Sleep ( 2 * time . Second )
c1 <- "result 1"
}()

//            <p>Here&rsquo;s the <code>select</code> implementing a timeout.
//<code>res := &lt;-c1</code> awaits the result and <code>&lt;-time.After</code>
//awaits a value to be sent after the timeout of
//1s. Since <code>select</code> proceeds with the first
//receive that&rsquo;s ready, we&rsquo;ll take the timeout case
//if the operation takes more than the allowed 1s.</p>
//


select {
case res := <- c1 :
fmt . Println ( res )
case <- time . After ( 1 * time . Second ):
fmt . Println ( "timeout 1" )
}

//            <p>If we allow a longer timeout of 3s, then the receive
//from <code>c2</code> will succeed and we&rsquo;ll print the result.</p>
//


c2 := make ( chan string , 1 )
go func () {
time . Sleep ( 2 * time . Second )
c2 <- "result 2"
}()
select {
case res := <- c2 :
fmt . Println ( res )
case <- time . After ( 3 * time . Second ):
fmt . Println ( "timeout 2" )
}
}

