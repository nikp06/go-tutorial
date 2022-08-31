//            <p>Throughout program execution, we often want to create
//data that isn&rsquo;t needed after the program exits.
//<em>Temporary files and directories</em> are useful for this
//purpose since they don&rsquo;t pollute the file system over
//time.</p>
//
//            

package main

//            

import (
"fmt"
"os"
"path/filepath"
)

//            

func check ( e error ) {
if e != nil {
panic ( e )
}
}

//            

func main () {

//            <p>The easiest way to create a temporary file is by
//calling <code>os.CreateTemp</code>. It creates a file <em>and</em>
//opens it for reading and writing. We provide <code>&quot;&quot;</code>
//as the first argument, so <code>os.CreateTemp</code> will
//create the file in the default location for our OS.</p>
//


f , err := os . CreateTemp ( "" , "sample" )
check ( err )

//            <p>Display the name of the temporary file. On
//Unix-based OSes the directory will likely be <code>/tmp</code>.
//The file name starts with the prefix given as the
//second argument to <code>os.CreateTemp</code> and the rest
//is chosen automatically to ensure that concurrent
//calls will always create different file names.</p>
//


fmt . Println ( "Temp file name:" , f . Name ())

//            <p>Clean up the file after we&rsquo;re done. The OS is
//likely to clean up temporary files by itself after
//some time, but it&rsquo;s good practice to do this
//explicitly.</p>
//


defer os . Remove ( f . Name ())

//            <p>We can write some data to the file.</p>
//


_ , err = f . Write ([] byte { 1 , 2 , 3 , 4 })
check ( err )

//            <p>If we intend to write many temporary files, we may
//prefer to create a temporary <em>directory</em>.
//<code>os.MkdirTemp</code>&rsquo;s arguments are the same as
//<code>CreateTemp</code>&rsquo;s, but it returns a directory <em>name</em>
//rather than an open file.</p>
//


dname , err := os . MkdirTemp ( "" , "sampledir" )
check ( err )
fmt . Println ( "Temp dir name:" , dname )

//            

defer os . RemoveAll ( dname )

//            <p>Now we can synthesize temporary file names by
//prefixing them with our temporary directory.</p>
//


fname := filepath . Join ( dname , "file1" )
err = os . WriteFile ( fname , [] byte { 1 , 2 }, 0666 )
check ( err )
}

