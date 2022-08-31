//            <p>We often want to execute Go code at some point in the
//future, or repeatedly at some interval. Go&rsquo;s built-in
//<em>timer</em> and <em>ticker</em> features make both of these tasks
//easy. We&rsquo;ll look first at timers and then
//at <a href="tickers">tickers</a>.</p>
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

//            <p>Timers represent a single event in the future. You
//tell the timer how long you want to wait, and it
//provides a channel that will be notified at that
//time. This timer will wait 2 seconds.</p>
//


timer1 := time . NewTimer ( 2 * time . Second )

//            <p>The <code>&lt;-timer1.C</code> blocks on the timer&rsquo;s channel <code>C</code>
//until it sends a value indicating that the timer
//fired.</p>
//


<- timer1 . C
fmt . Println ( "Timer 1 fired" )

//            <p>If you just wanted to wait, you could have used
//<code>time.Sleep</code>. One reason a timer may be useful is
//that you can cancel the timer before it fires.
//Here&rsquo;s an example of that.</p>
//


timer2 := time . NewTimer ( time . Second )
go func () {
<- timer2 . C
fmt . Println ( "Timer 2 fired" )
}()
stop2 := timer2 . Stop ()
if stop2 {
fmt . Println ( "Timer 2 stopped" )
}

//            <p>Give the <code>timer2</code> enough time to fire, if it ever
//was going to, to show it is in fact stopped.</p>
//


time . Sleep ( 2 * time . Second )
}

