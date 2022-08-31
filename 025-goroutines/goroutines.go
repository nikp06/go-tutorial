//            <p>A <em>goroutine</em> is a lightweight thread of execution.</p>
//
//            

package main

//            

import (
"fmt"
"time"
)

//            

func f ( from string ) {
for i := 0 ; i < 3 ; i ++ {
fmt . Println ( from , ":" , i )
}
}

//            

func main () {

//            <p>Suppose we have a function call <code>f(s)</code>. Here&rsquo;s how
//we&rsquo;d call that in the usual way, running it
//synchronously.</p>
//


f ( "direct" )

//            <p>To invoke this function in a goroutine, use
//<code>go f(s)</code>. This new goroutine will execute
//concurrently with the calling one.</p>
//


go f ( "goroutine" )

//            <p>You can also start a goroutine for an anonymous
//function call.</p>
//


go func ( msg string ) {
fmt . Println ( msg )
}( "going" )

//            <p>Our two function calls are running asynchronously in
//separate goroutines now. Wait for them to finish
//(for a more robust approach, use a <a href="waitgroups">WaitGroup</a>).</p>
//


time . Sleep ( time . Second )
fmt . Println ( "done" )
}

