//            <p>A common requirement in programs is getting the number
//of seconds, milliseconds, or nanoseconds since the
//<a href="https://en.wikipedia.org/wiki/Unix_time">Unix epoch</a>.
//Here&rsquo;s how to do it in Go.</p>
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

//            <p>Use <code>time.Now</code> with <code>Unix</code>, <code>UnixMilli</code> or <code>UnixNano</code>
//to get elapsed time since the Unix epoch in seconds,
//milliseconds or nanoseconds, respectively.</p>
//


now := time . Now ()
fmt . Println ( now )

//            

fmt . Println ( now . Unix ())
fmt . Println ( now . UnixMilli ())
fmt . Println ( now . UnixNano ())

//            <p>You can also convert integer seconds or nanoseconds
//since the epoch into the corresponding <code>time</code>.</p>
//


fmt . Println ( time . Unix ( now . Unix (), 0 ))
fmt . Println ( time . Unix ( 0 , now . UnixNano ()))
}

