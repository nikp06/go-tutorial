//            <p>In the previous example we looked at setting up a simple
//<a href="http-servers">HTTP server</a>. HTTP servers are useful for
//demonstrating the usage of <code>context.Context</code> for
//controlling cancellation. A <code>Context</code> carries deadlines,
//cancellation signals, and other request-scoped values
//across API boundaries and goroutines.</p>
//


package main

//            

import (
"fmt"
"net/http"
"time"
)

//            

func hello ( w http . ResponseWriter , req * http . Request ) {

//            <p>A <code>context.Context</code> is created for each request by
//the <code>net/http</code> machinery, and is available with
//the <code>Context()</code> method.</p>
//


ctx := req . Context ()
fmt . Println ( "server: hello handler started" )
defer fmt . Println ( "server: hello handler ended" )

//            <p>Wait for a few seconds before sending a reply to the
//client. This could simulate some work the server is
//doing. While working, keep an eye on the context&rsquo;s
//<code>Done()</code> channel for a signal that we should cancel
//the work and return as soon as possible.</p>
//


select {
case <- time . After ( 10 * time . Second ):
fmt . Fprintf ( w , "hello\n" )
case <- ctx . Done ():

//            <p>The context&rsquo;s <code>Err()</code> method returns an error
//that explains why the <code>Done()</code> channel was
//closed.</p>
//


err := ctx . Err ()
fmt . Println ( "server:" , err )
internalError := http . StatusInternalServerError
http . Error ( w , err . Error (), internalError )
}
}

//            

func main () {

//            <p>As before, we register our handler on the &ldquo;/hello&rdquo;
//route, and start serving.</p>
//


http . HandleFunc ( "/hello" , hello )
http . ListenAndServe ( ":8090" , nil )
}

