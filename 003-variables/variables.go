//            <p>In Go, <em>variables</em> are explicitly declared and used by
//the compiler to e.g. check type-correctness of function
//calls.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p><code>var</code> declares 1 or more variables.</p>
//


var a = "initial"
fmt . Println ( a )

//            <p>You can declare multiple variables at once.</p>
//


var b , c int = 1 , 2
fmt . Println ( b , c )

//            <p>Go will infer the type of initialized variables.</p>
//


var d = true
fmt . Println ( d )

//            <p>Variables declared without a corresponding
//initialization are <em>zero-valued</em>. For example, the
//zero value for an <code>int</code> is <code>0</code>.</p>
//


var e int
fmt . Println ( e )

//            <p>The <code>:=</code> syntax is shorthand for declaring and
//initializing a variable, e.g. for
//<code>var f string = &quot;apple&quot;</code> in this case.</p>
//


f := "apple"
fmt . Println ( f )
}

