//            <p>Writing a basic HTTP server is easy using the
//<code>net/http</code> package.</p>
//


package main

//            

import (
"fmt"
"net/http"
)

//            <p>A fundamental concept in <code>net/http</code> servers is
//<em>handlers</em>. A handler is an object implementing the
//<code>http.Handler</code> interface. A common way to write
//a handler is by using the <code>http.HandlerFunc</code> adapter
//on functions with the appropriate signature.</p>
//


func hello ( w http . ResponseWriter , req * http . Request ) {

//            <p>Functions serving as handlers take a
//<code>http.ResponseWriter</code> and a <code>http.Request</code> as
//arguments. The response writer is used to fill in the
//HTTP response. Here our simple response is just
//&ldquo;hello\n&rdquo;.</p>
//


fmt . Fprintf ( w , "hello\n" )
}

//            

func headers ( w http . ResponseWriter , req * http . Request ) {

//            <p>This handler does something a little more
//sophisticated by reading all the HTTP request
//headers and echoing them into the response body.</p>
//


for name , headers := range req . Header {
for _ , h := range headers {
fmt . Fprintf ( w , "%v: %v\n" , name , h )
}
}
}

//            

func main () {

//            <p>We register our handlers on server routes using the
//<code>http.HandleFunc</code> convenience function. It sets up
//the <em>default router</em> in the <code>net/http</code> package and
//takes a function as an argument.</p>
//


http . HandleFunc ( "/hello" , hello )
http . HandleFunc ( "/headers" , headers )

//            <p>Finally, we call the <code>ListenAndServe</code> with the port
//and a handler. <code>nil</code> tells it to use the default
//router we&rsquo;ve just set up.</p>
//


http . ListenAndServe ( ":8090" , nil )
}

