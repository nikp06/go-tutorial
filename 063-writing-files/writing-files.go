//            <p>Writing files in Go follows similar patterns to the
//ones we saw earlier for reading.</p>
//
//            

package main

//            

import (
"bufio"
"fmt"
"os"
)

//            

func check ( e error ) {
if e != nil {
panic ( e )
}
}

//            

func main () {

//            <p>To start, here&rsquo;s how to dump a string (or just
//bytes) into a file.</p>
//


d1 := [] byte ( "hello\ngo\n" )
err := os . WriteFile ( "/tmp/dat1" , d1 , 0644 )
check ( err )

//            <p>For more granular writes, open a file for writing.</p>
//


f , err := os . Create ( "/tmp/dat2" )
check ( err )

//            <p>It&rsquo;s idiomatic to defer a <code>Close</code> immediately
//after opening a file.</p>
//


defer f . Close ()

//            <p>You can <code>Write</code> byte slices as you&rsquo;d expect.</p>
//


d2 := [] byte { 115 , 111 , 109 , 101 , 10 }
n2 , err := f . Write ( d2 )
check ( err )
fmt . Printf ( "wrote %d bytes\n" , n2 )

//            <p>A <code>WriteString</code> is also available.</p>
//


n3 , err := f . WriteString ( "writes\n" )
check ( err )
fmt . Printf ( "wrote %d bytes\n" , n3 )

//            <p>Issue a <code>Sync</code> to flush writes to stable storage.</p>
//


f . Sync ()

//            <p><code>bufio</code> provides buffered writers in addition
//to the buffered readers we saw earlier.</p>
//


w := bufio . NewWriter ( f )
n4 , err := w . WriteString ( "buffered\n" )
check ( err )
fmt . Printf ( "wrote %d bytes\n" , n4 )

//            <p>Use <code>Flush</code> to ensure all buffered operations have
//been applied to the underlying writer.</p>
//


w . Flush ()

//            

}

