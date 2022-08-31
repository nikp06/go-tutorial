//            <p>In Go it&rsquo;s idiomatic to communicate errors via an
//explicit, separate return value. This contrasts with
//the exceptions used in languages like Java and Ruby and
//the overloaded single result / error value sometimes
//used in C. Go&rsquo;s approach makes it easy to see which
//functions return errors and to handle them using the
//same language constructs employed for any other,
//non-error tasks.</p>
//
//            

package main

//            

import (
"errors"
"fmt"
)

//            <p>By convention, errors are the last return value and
//have type <code>error</code>, a built-in interface.</p>
//


func f1 ( arg int ) ( int , error ) {
if arg == 42 {

//            <p><code>errors.New</code> constructs a basic <code>error</code> value
//with the given error message.</p>
//


return - 1 , errors . New ( "can't work with 42" )

//            

}

//            <p>A <code>nil</code> value in the error position indicates that
//there was no error.</p>
//


return arg + 3 , nil
}

//            <p>It&rsquo;s possible to use custom types as <code>error</code>s by
//implementing the <code>Error()</code> method on them. Here&rsquo;s a
//variant on the example above that uses a custom type
//to explicitly represent an argument error.</p>
//


type argError struct {
arg int
prob string
}

//            

func ( e * argError ) Error () string {
return fmt . Sprintf ( "%d - %s" , e . arg , e . prob )
}

//            

func f2 ( arg int ) ( int , error ) {
if arg == 42 {

//            <p>In this case we use <code>&amp;argError</code> syntax to build
//a new struct, supplying values for the two
//fields <code>arg</code> and <code>prob</code>.</p>
//


return - 1 , & argError { arg , "can't work with it" }
}
return arg + 3 , nil
}

//            

func main () {

//            <p>The two loops below test out each of our
//error-returning functions. Note that the use of an
//inline error check on the <code>if</code> line is a common
//idiom in Go code.</p>
//


for _ , i := range [] int { 7 , 42 } {
if r , e := f1 ( i ); e != nil {
fmt . Println ( "f1 failed:" , e )
} else {
fmt . Println ( "f1 worked:" , r )
}
}
for _ , i := range [] int { 7 , 42 } {
if r , e := f2 ( i ); e != nil {
fmt . Println ( "f2 failed:" , e )
} else {
fmt . Println ( "f2 worked:" , r )
}
}

//            <p>If you want to programmatically use the data in
//a custom error, you&rsquo;ll need to get the error as an
//instance of the custom error type via type
//assertion.</p>
//


_ , e := f2 ( 42 )
if ae , ok := e .( * argError ); ok {
fmt . Println ( ae . arg )
fmt . Println ( ae . prob )
}
}

