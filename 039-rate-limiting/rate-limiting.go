//            <p><a href="https://en.wikipedia.org/wiki/Rate_limiting"><em>Rate limiting</em></a>
//is an important mechanism for controlling resource
//utilization and maintaining quality of service. Go
//elegantly supports rate limiting with goroutines,
//channels, and <a href="tickers">tickers</a>.</p>
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

//            <p>First we&rsquo;ll look at basic rate limiting. Suppose
//we want to limit our handling of incoming requests.
//We&rsquo;ll serve these requests off a channel of the
//same name.</p>
//


requests := make ( chan int , 5 )
for i := 1 ; i <= 5 ; i ++ {
requests <- i
}
close ( requests )

//            <p>This <code>limiter</code> channel will receive a value
//every 200 milliseconds. This is the regulator in
//our rate limiting scheme.</p>
//


limiter := time . Tick ( 200 * time . Millisecond )

//            <p>By blocking on a receive from the <code>limiter</code> channel
//before serving each request, we limit ourselves to
//1 request every 200 milliseconds.</p>
//


for req := range requests {
<- limiter
fmt . Println ( "request" , req , time . Now ())
}

//            <p>We may want to allow short bursts of requests in
//our rate limiting scheme while preserving the
//overall rate limit. We can accomplish this by
//buffering our limiter channel. This <code>burstyLimiter</code>
//channel will allow bursts of up to 3 events.</p>
//


burstyLimiter := make ( chan time . Time , 3 )

//            <p>Fill up the channel to represent allowed bursting.</p>
//


for i := 0 ; i < 3 ; i ++ {
burstyLimiter <- time . Now ()
}

//            <p>Every 200 milliseconds we&rsquo;ll try to add a new
//value to <code>burstyLimiter</code>, up to its limit of 3.</p>
//


go func () {
for t := range time . Tick ( 200 * time . Millisecond ) {
burstyLimiter <- t
}
}()

//            <p>Now simulate 5 more incoming requests. The first
//3 of these will benefit from the burst capability
//of <code>burstyLimiter</code>.</p>
//


burstyRequests := make ( chan int , 5 )
for i := 1 ; i <= 5 ; i ++ {
burstyRequests <- i
}
close ( burstyRequests )
for req := range burstyRequests {
<- burstyLimiter
fmt . Println ( "request" , req , time . Now ())
}
}

