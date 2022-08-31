//            <p><em>Functions</em> are central in Go. We&rsquo;ll learn about
//functions with a few different examples.</p>
//
//            

package main

//            

import "fmt"

//            <p>Here&rsquo;s a function that takes two <code>int</code>s and returns
//their sum as an <code>int</code>.</p>
//


func plus ( a int , b int ) int {

//            <p>Go requires explicit returns, i.e. it won&rsquo;t
//automatically return the value of the last
//expression.</p>
//


return a + b
}

//            <p>When you have multiple consecutive parameters of
//the same type, you may omit the type name for the
//like-typed parameters up to the final parameter that
//declares the type.</p>
//


func plusPlus ( a , b , c int ) int {
return a + b + c
}

//            

func main () {

//            <p>Call a function just as you&rsquo;d expect, with
//<code>name(args)</code>.</p>
//


res := plus ( 1 , 2 )
fmt . Println ( "1+2 =" , res )

//            

res = plusPlus ( 1 , 2 , 3 )
fmt . Println ( "1+2+3 =" , res )
}

