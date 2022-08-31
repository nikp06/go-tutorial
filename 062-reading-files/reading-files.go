//            <p>Reading and writing files are basic tasks needed for
//many Go programs. First we&rsquo;ll look at some examples of
//reading files.</p>
//
//            

package main

//            

import (
"bufio"
"fmt"
"io"
"os"
)

//            <p>Reading files requires checking most calls for errors.
//This helper will streamline our error checks below.</p>
//


func check ( e error ) {
if e != nil {
panic ( e )
}
}

//            

func main () {

//            <p>Perhaps the most basic file reading task is
//slurping a file&rsquo;s entire contents into memory.</p>
//


dat , err := os . ReadFile ( "/tmp/dat" )
check ( err )
fmt . Print ( string ( dat ))

//            <p>You&rsquo;ll often want more control over how and what
//parts of a file are read. For these tasks, start
//by <code>Open</code>ing a file to obtain an <code>os.File</code> value.</p>
//


f , err := os . Open ( "/tmp/dat" )
check ( err )

//            <p>Read some bytes from the beginning of the file.
//Allow up to 5 to be read but also note how many
//actually were read.</p>
//


b1 := make ([] byte , 5 )
n1 , err := f . Read ( b1 )
check ( err )
fmt . Printf ( "%d bytes: %s\n" , n1 , string ( b1 [: n1 ]))

//            <p>You can also <code>Seek</code> to a known location in the file
//and <code>Read</code> from there.</p>
//


o2 , err := f . Seek ( 6 , 0 )
check ( err )
b2 := make ([] byte , 2 )
n2 , err := f . Read ( b2 )
check ( err )
fmt . Printf ( "%d bytes @ %d: " , n2 , o2 )
fmt . Printf ( "%v\n" , string ( b2 [: n2 ]))

//            <p>The <code>io</code> package provides some functions that may
//be helpful for file reading. For example, reads
//like the ones above can be more robustly
//implemented with <code>ReadAtLeast</code>.</p>
//


o3 , err := f . Seek ( 6 , 0 )
check ( err )
b3 := make ([] byte , 2 )
n3 , err := io . ReadAtLeast ( f , b3 , 2 )
check ( err )
fmt . Printf ( "%d bytes @ %d: %s\n" , n3 , o3 , string ( b3 ))

//            <p>There is no built-in rewind, but <code>Seek(0, 0)</code>
//accomplishes this.</p>
//


_ , err = f . Seek ( 0 , 0 )
check ( err )

//            <p>The <code>bufio</code> package implements a buffered
//reader that may be useful both for its efficiency
//with many small reads and because of the additional
//reading methods it provides.</p>
//


r4 := bufio . NewReader ( f )
b4 , err := r4 . Peek ( 5 )
check ( err )
fmt . Printf ( "5 bytes: %s\n" , string ( b4 ))

//            <p>Close the file when you&rsquo;re done (usually this would
//be scheduled immediately after <code>Open</code>ing with
//<code>defer</code>).</p>
//


f . Close ()
}

