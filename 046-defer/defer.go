//            <p><em>Defer</em> is used to ensure that a function call is
//performed later in a program&rsquo;s execution, usually for
//purposes of cleanup. <code>defer</code> is often used where e.g.
//<code>ensure</code> and <code>finally</code> would be used in other languages.</p>
//
//            

package main

//            

import (
"fmt"
"os"
)

//            <p>Suppose we wanted to create a file, write to it,
//and then close when we&rsquo;re done. Here&rsquo;s how we could
//do that with <code>defer</code>.</p>
//


func main () {

//            <p>Immediately after getting a file object with
//<code>createFile</code>, we defer the closing of that file
//with <code>closeFile</code>. This will be executed at the end
//of the enclosing function (<code>main</code>), after
//<code>writeFile</code> has finished.</p>
//


f := createFile ( "/tmp/defer.txt" )
defer closeFile ( f )
writeFile ( f )
}

//            

func createFile ( p string ) * os . File {
fmt . Println ( "creating" )
f , err := os . Create ( p )
if err != nil {
panic ( err )
}
return f
}

//            

func writeFile ( f * os . File ) {
fmt . Println ( "writing" )
fmt . Fprintln ( f , "data" )

//            

}

//            <p>It&rsquo;s important to check for errors when closing a
//file, even in a deferred function.</p>
//

func closeFile ( f * os . File ) {
fmt . Println ( "closing" )
err := f . Close ()

//            

if err != nil {
fmt . Fprintf ( os . Stderr , "error: %v\n" , err )
os . Exit ( 1 )
}
}

