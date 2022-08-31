//            <p>Go has various value types including strings,
//integers, floats, booleans, etc. Here are a few
//basic examples.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p>Strings, which can be added together with <code>+</code>.</p>
//


fmt . Println ( "go" + "lang" )

//            <p>Integers and floats.</p>
//


fmt . Println ( "1+1 =" , 1 + 1 )
fmt . Println ( "7.0/3.0 =" , 7.0 / 3.0 )

//            <p>Booleans, with boolean operators as you&rsquo;d expect.</p>
//


fmt . Println ( true && false )
fmt . Println ( true || false )
fmt . Println (! true )
}

