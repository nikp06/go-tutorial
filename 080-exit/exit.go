//            <p>Use <code>os.Exit</code> to immediately exit with a given
//status.</p>
//
//            

package main

//            

import (
"fmt"
"os"
)

//            

func main () {

//            <p><code>defer</code>s will <em>not</em> be run when using <code>os.Exit</code>, so
//this <code>fmt.Println</code> will never be called.</p>
//


defer fmt . Println ( "!" )

//            <p>Exit with status 3.</p>
//


os . Exit ( 3 )
}

//            <p>Note that unlike e.g. C, Go does not use an integer
//return value from <code>main</code> to indicate exit status. If
//you&rsquo;d like to exit with a non-zero status you should
//use <code>os.Exit</code>.</p>
//
