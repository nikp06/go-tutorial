//            <p>A <em>line filter</em> is a common type of program that reads
//input on stdin, processes it, and then prints some
//derived result to stdout. <code>grep</code> and <code>sed</code> are common
//line filters.</p>
//
//            <p>Here&rsquo;s an example line filter in Go that writes a
//capitalized version of all input text. You can use this
//pattern to write your own Go line filters.</p>
//


package main

//            

import (
"bufio"
"fmt"
"os"
"strings"
)

//            

func main () {

//            <p>Wrapping the unbuffered <code>os.Stdin</code> with a buffered
//scanner gives us a convenient <code>Scan</code> method that
//advances the scanner to the next token; which is
//the next line in the default scanner.</p>
//


scanner := bufio . NewScanner ( os . Stdin )

//            <p><code>Text</code> returns the current token, here the next line,
//from the input.</p>
//

for scanner . Scan () {

//            

ucl := strings . ToUpper ( scanner . Text ())

//            <p>Write out the uppercased line.</p>
//


fmt . Println ( ucl )
}

//            <p>Check for errors during <code>Scan</code>. End of file is
//expected and not reported by <code>Scan</code> as an error.</p>
//


if err := scanner . Err (); err != nil {
fmt . Fprintln ( os . Stderr , "error:" , err )
os . Exit ( 1 )
}
}

