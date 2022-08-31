//            <p>The <code>filepath</code> package provides functions to parse
//and construct <em>file paths</em> in a way that is portable
//between operating systems; <code>dir/file</code> on Linux vs.
//<code>dir\file</code> on Windows, for example.</p>
//


package main

//            

import (
"fmt"
"path/filepath"
"strings"
)

//            

func main () {

//            <p><code>Join</code> should be used to construct paths in a
//portable way. It takes any number of arguments
//and constructs a hierarchical path from them.</p>
//


p := filepath . Join ( "dir1" , "dir2" , "filename" )
fmt . Println ( "p:" , p )

//            <p>You should always use <code>Join</code> instead of
//concatenating <code>/</code>s or <code>\</code>s manually. In addition
//to providing portability, <code>Join</code> will also
//normalize paths by removing superfluous separators
//and directory changes.</p>
//


fmt . Println ( filepath . Join ( "dir1//" , "filename" ))
fmt . Println ( filepath . Join ( "dir1/../dir1" , "filename" ))

//            <p><code>Dir</code> and <code>Base</code> can be used to split a path to the
//directory and the file. Alternatively, <code>Split</code> will
//return both in the same call.</p>
//


fmt . Println ( "Dir(p):" , filepath . Dir ( p ))
fmt . Println ( "Base(p):" , filepath . Base ( p ))

//            <p>We can check whether a path is absolute.</p>
//


fmt . Println ( filepath . IsAbs ( "dir/file" ))
fmt . Println ( filepath . IsAbs ( "/dir/file" ))

//            

filename := "config.json"

//            <p>Some file names have extensions following a dot. We
//can split the extension out of such names with <code>Ext</code>.</p>
//


ext := filepath . Ext ( filename )
fmt . Println ( ext )

//            <p>To find the file&rsquo;s name with the extension removed,
//use <code>strings.TrimSuffix</code>.</p>
//


fmt . Println ( strings . TrimSuffix ( filename , ext ))

//            <p><code>Rel</code> finds a relative path between a <em>base</em> and a
//<em>target</em>. It returns an error if the target cannot
//be made relative to base.</p>
//


rel , err := filepath . Rel ( "a/b" , "a/b/t/file" )
if err != nil {
panic ( err )
}
fmt . Println ( rel )

//            

rel , err = filepath . Rel ( "a/b" , "a/c/t/file" )
if err != nil {
panic ( err )
}
fmt . Println ( rel )
}

