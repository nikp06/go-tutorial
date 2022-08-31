//            <p><a href="timers">Timers</a> are for when you want to do
//something once in the future - <em>tickers</em> are for when
//you want to do something repeatedly at regular
//intervals. Here&rsquo;s an example of a ticker that ticks
//periodically until we stop it.</p>
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

//            <p>Tickers use a similar mechanism to timers: a
//channel that is sent values. Here we&rsquo;ll use the
//<code>select</code> builtin on the channel to await the
//values as they arrive every 500ms.</p>
//


ticker := time . NewTicker ( 500 * time . Millisecond )
done := make ( chan bool )

//            

go func () {
for {
select {
case <- done :
return
case t := <- ticker . C :
fmt . Println ( "Tick at" , t )
}
}
}()

//            <p>Tickers can be stopped like timers. Once a ticker
//is stopped it won&rsquo;t receive any more values on its
//channel. We&rsquo;ll stop ours after 1600ms.</p>
//


time . Sleep ( 1600 * time . Millisecond )
ticker . Stop ()
done <- true
fmt . Println ( "Ticker stopped" )
}

