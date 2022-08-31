//            <p><code>//go:embed</code> is a <a href="https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives">compiler
//directive</a> that
//allows programs to include arbitrary files and folders in the Go binary at
//build time. Read more about the embed directive
//<a href="https://pkg.go.dev/embed">here</a>.</p>
//


package main

//            <p>Import the <code>embed</code> package; if you don&rsquo;t use any exported
//identifiers from this package, you can do a blank import with <code>_ &quot;embed&quot;</code>.</p>
//


import (
"embed"
)

//            <p><code>embed</code> directives accept paths relative to the directory containing the
//Go source file. This directive embeds the contents of the file into the
//<code>string</code> variable immediately following it.</p>
//


//go:embed folder/single_file.txt
 var fileString string

//            <p>Or embed the contents of the file into a <code>[]byte</code>.</p>
//


//go:embed folder/single_file.txt
 var fileByte [] byte

//            <p>We can also embed multiple files or even folders with wildcards. This uses
//a variable of the <a href="https://pkg.go.dev/embed#FS">embed.FS type</a>, which
//implements a simple virtual file system.</p>
//


//go:embed folder/single_file.txt
//go:embed folder/*.hash
 var folder embed . FS

//            

func main () {

//            <p>Print out the contents of <code>single_file.txt</code>.</p>
//


print ( fileString )
print ( string ( fileByte ))

//            <p>Retrieve some files from the embedded folder.</p>
//


content1 , _ := folder . ReadFile ( "folder/file1.hash" )
print ( string ( content1 ))

//            

content2 , _ := folder . ReadFile ( "folder/file2.hash" )
print ( string ( content2 ))
}

