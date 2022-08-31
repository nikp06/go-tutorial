//            <p>Go has several useful functions for working with
//<em>directories</em> in the file system.</p>
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

//            <p>Create a new sub-directory in the current working
//directory.</p>
//


err := os . Mkdir ( "subdir" , 0755 )
check ( err )

//            <p>When creating temporary directories, it&rsquo;s good
//practice to <code>defer</code> their removal. <code>os.RemoveAll</code>
//will delete a whole directory tree (similarly to
//<code>rm -rf</code>).</p>
//


defer os . RemoveAll ( "subdir" )

//            <p>Helper function to create a new empty file.</p>
//


createEmptyFile := func ( name string ) {
d := [] byte ( "" )
check ( os . WriteFile ( name , d , 0644 ))
}

//            

createEmptyFile ( "subdir/file1" )

//            <p>We can create a hierarchy of directories, including
//parents with <code>MkdirAll</code>. This is similar to the
//command-line <code>mkdir -p</code>.</p>
//


err = os . MkdirAll ( "subdir/parent/child" , 0755 )
check ( err )

//            

createEmptyFile ( "subdir/parent/file2" )
createEmptyFile ( "subdir/parent/file3" )
createEmptyFile ( "subdir/parent/child/file4" )

//            <p><code>ReadDir</code> lists directory contents, returning a
//slice of <code>os.DirEntry</code> objects.</p>
//


c , err := os . ReadDir ( "subdir/parent" )
check ( err )

//            

fmt . Println ( "Listing subdir/parent" )
for _ , entry := range c {
fmt . Println ( " " , entry . Name (), entry . IsDir ())
}

//            <p><code>Chdir</code> lets us change the current working directory,
//similarly to <code>cd</code>.</p>
//


err = os . Chdir ( "subdir/parent/child" )
check ( err )

//            <p>Now we&rsquo;ll see the contents of <code>subdir/parent/child</code>
//when listing the <em>current</em> directory.</p>
//


c , err = os . ReadDir ( "." )
check ( err )

//            

fmt . Println ( "Listing subdir/parent/child" )
for _ , entry := range c {
fmt . Println ( " " , entry . Name (), entry . IsDir ())
}

//            <p><code>cd</code> back to where we started.</p>
//


err = os . Chdir ( "../../.." )
check ( err )

//            <p>We can also visit a directory <em>recursively</em>,
//including all its sub-directories. <code>Walk</code> accepts
//a callback function to handle every file or
//directory visited.</p>
//


fmt . Println ( "Visiting subdir" )
err = filepath . Walk ( "subdir" , visit )
}

//            <p><code>visit</code> is called for every file or directory found
//recursively by <code>filepath.Walk</code>.</p>
//


func visit ( p string , info os . FileInfo , err error ) error {
if err != nil {
return err
}
fmt . Println ( " " , p , info . IsDir ())
return nil
}

