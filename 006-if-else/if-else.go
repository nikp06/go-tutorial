//            <p>Branching with <code>if</code> and <code>else</code> in Go is
//straight-forward.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p>Here&rsquo;s a basic example.</p>
//


if 7 % 2 == 0 {
fmt . Println ( "7 is even" )
} else {
fmt . Println ( "7 is odd" )
}

//            <p>You can have an <code>if</code> statement without an else.</p>
//


if 8 % 4 == 0 {
fmt . Println ( "8 is divisible by 4" )
}

//            <p>A statement can precede conditionals; any variables
//declared in this statement are available in all
//branches.</p>
//


if num := 9 ; num < 0 {
fmt . Println ( num , "is negative" )
} else if num < 10 {
fmt . Println ( num , "has 1 digit" )
} else {
fmt . Println ( num , "has multiple digits" )
}
}

//            <p>Note that you don&rsquo;t need parentheses around conditions
//in Go, but that the braces are required.</p>
//
