//            <p><em>Switch statements</em> express conditionals across many
//branches.</p>
//
//            

package main

//            

import (
"fmt"
"time"
)

//            

func main () {

//            <p>Here&rsquo;s a basic <code>switch</code>.</p>
//


i := 2
fmt . Print ( "Write " , i , " as " )
switch i {
case 1 :
fmt . Println ( "one" )
case 2 :
fmt . Println ( "two" )
case 3 :
fmt . Println ( "three" )
}

//            <p>You can use commas to separate multiple expressions
//in the same <code>case</code> statement. We use the optional
//<code>default</code> case in this example as well.</p>
//


switch time . Now (). Weekday () {
case time . Saturday , time . Sunday :
fmt . Println ( "It's the weekend" )
default :
fmt . Println ( "It's a weekday" )
}

//            <p><code>switch</code> without an expression is an alternate way
//to express if/else logic. Here we also show how the
//<code>case</code> expressions can be non-constants.</p>
//


t := time . Now ()
switch {
case t . Hour () < 12 :
fmt . Println ( "It's before noon" )
default :
fmt . Println ( "It's after noon" )
}

//            <p>A type <code>switch</code> compares types instead of values.  You
//can use this to discover the type of an interface
//value.  In this example, the variable <code>t</code> will have the
//type corresponding to its clause.</p>
//


whatAmI := func ( i interface {}) {
switch t := i .( type ) {
case bool :
fmt . Println ( "I'm a bool" )
case int :
fmt . Println ( "I'm an int" )
default :
fmt . Printf ( "Don't know type %T\n" , t )
}
}
whatAmI ( true )
whatAmI ( 1 )
whatAmI ( "hey" )
}

