//            <p>Go has built-in support for <em>multiple return values</em>.
//This feature is used often in idiomatic Go, for example
//to return both result and error values from a function.</p>
//
//            

package main

//            

import "fmt"

//            <p>The <code>(int, int)</code> in this function signature shows that
//the function returns 2 <code>int</code>s.</p>
//


func vals () ( int , int ) {
return 3 , 7
}

//            

func main () {

//            <p>Here we use the 2 different return values from the
//call with <em>multiple assignment</em>.</p>
//


a , b := vals ()
fmt . Println ( a )
fmt . Println ( b )

//            <p>If you only want a subset of the returned values,
//use the blank identifier <code>_</code>.</p>
//


_ , c := vals ()
fmt . Println ( c )
}

