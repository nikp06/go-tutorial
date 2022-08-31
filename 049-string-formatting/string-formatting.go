//            <p>Go offers excellent support for string formatting in
//the <code>printf</code> tradition. Here are some examples of
//common string formatting tasks.</p>
//
//            

package main

//            

import (
"fmt"
"os"
)

//            

type point struct {
x , y int
}

//            

func main () {

//            <p>Go offers several printing &ldquo;verbs&rdquo; designed to
//format general Go values. For example, this prints
//an instance of our <code>point</code> struct.</p>
//


p := point { 1 , 2 }
fmt . Printf ( "struct1: %v\n" , p )

//            <p>If the value is a struct, the <code>%+v</code> variant will
//include the struct&rsquo;s field names.</p>
//


fmt . Printf ( "struct2: %+v\n" , p )

//            <p>The <code>%#v</code> variant prints a Go syntax representation
//of the value, i.e. the source code snippet that
//would produce that value.</p>
//


fmt . Printf ( "struct3: %#v\n" , p )

//            <p>To print the type of a value, use <code>%T</code>.</p>
//


fmt . Printf ( "type: %T\n" , p )

//            <p>Formatting booleans is straight-forward.</p>
//


fmt . Printf ( "bool: %t\n" , true )

//            <p>There are many options for formatting integers.
//Use <code>%d</code> for standard, base-10 formatting.</p>
//


fmt . Printf ( "int: %d\n" , 123 )

//            <p>This prints a binary representation.</p>
//


fmt . Printf ( "bin: %b\n" , 14 )

//            <p>This prints the character corresponding to the
//given integer.</p>
//


fmt . Printf ( "char: %c\n" , 33 )

//            <p><code>%x</code> provides hex encoding.</p>
//


fmt . Printf ( "hex: %x\n" , 456 )

//            <p>There are also several formatting options for
//floats. For basic decimal formatting use <code>%f</code>.</p>
//


fmt . Printf ( "float1: %f\n" , 78.9 )

//            <p><code>%e</code> and <code>%E</code> format the float in (slightly
//different versions of) scientific notation.</p>
//


fmt . Printf ( "float2: %e\n" , 123400000.0 )
fmt . Printf ( "float3: %E\n" , 123400000.0 )

//            <p>For basic string printing use <code>%s</code>.</p>
//


fmt . Printf ( "str1: %s\n" , "\"string\"" )

//            <p>To double-quote strings as in Go source, use <code>%q</code>.</p>
//


fmt . Printf ( "str2: %q\n" , "\"string\"" )

//            <p>As with integers seen earlier, <code>%x</code> renders
//the string in base-16, with two output characters
//per byte of input.</p>
//


fmt . Printf ( "str3: %x\n" , "hex this" )

//            <p>To print a representation of a pointer, use <code>%p</code>.</p>
//


fmt . Printf ( "pointer: %p\n" , & p )

//            <p>When formatting numbers you will often want to
//control the width and precision of the resulting
//figure. To specify the width of an integer, use a
//number after the <code>%</code> in the verb. By default the
//result will be right-justified and padded with
//spaces.</p>
//


fmt . Printf ( "width1: |%6d|%6d|\n" , 12 , 345 )

//            <p>You can also specify the width of printed floats,
//though usually you&rsquo;ll also want to restrict the
//decimal precision at the same time with the
//width.precision syntax.</p>
//


fmt . Printf ( "width2: |%6.2f|%6.2f|\n" , 1.2 , 3.45 )

//            <p>To left-justify, use the <code>-</code> flag.</p>
//


fmt . Printf ( "width3: |%-6.2f|%-6.2f|\n" , 1.2 , 3.45 )

//            <p>You may also want to control width when formatting
//strings, especially to ensure that they align in
//table-like output. For basic right-justified width.</p>
//


fmt . Printf ( "width4: |%6s|%6s|\n" , "foo" , "b" )

//            <p>To left-justify use the <code>-</code> flag as with numbers.</p>
//


fmt . Printf ( "width5: |%-6s|%-6s|\n" , "foo" , "b" )

//            <p>So far we&rsquo;ve seen <code>Printf</code>, which prints the
//formatted string to <code>os.Stdout</code>. <code>Sprintf</code> formats
//and returns a string without printing it anywhere.</p>
//


s := fmt . Sprintf ( "sprintf: a %s" , "string" )
fmt . Println ( s )

//            <p>You can format+print to <code>io.Writers</code> other than
//<code>os.Stdout</code> using <code>Fprintf</code>.</p>
//


fmt . Fprintf ( os . Stderr , "io: an %s\n" , "error" )
}

