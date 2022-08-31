//            <p>Go makes it possible to <em>recover</em> from a panic, by
//using the <code>recover</code> built-in function. A <code>recover</code> can
//stop a <code>panic</code> from aborting the program and let it
//continue with execution instead.</p>
//
//            <p>An example of where this can be useful: a server
//wouldn&rsquo;t want to crash if one of the client connections
//exhibits a critical error. Instead, the server would
//want to close that connection and continue serving
//other clients. In fact, this is what Go&rsquo;s <code>net/http</code>
//does by default for HTTP servers.</p>
//
//            

package main

//            

import "fmt"

//            <p>This function panics.</p>
//


func mayPanic () {
panic ( "a problem" )
}

//            <p><code>recover</code> must be called within a deferred function.
//When the enclosing function panics, the defer will
//activate and a <code>recover</code> call within it will catch
//the panic.</p>
//

func main () {

//            <p>The return value of <code>recover</code> is the error raised in
//the call to <code>panic</code>.</p>
//

defer func () {
if r := recover (); r != nil {

//            

fmt . Println ( "Recovered. Error:\n" , r )
}
}()

//            

mayPanic ()

//            <p>This code will not run, because <code>mayPanic</code> panics.
//The execution of <code>main</code> stops at the point of the
//panic and resumes in the deferred closure.</p>
//


fmt . Println ( "After mayPanic()" )
}

