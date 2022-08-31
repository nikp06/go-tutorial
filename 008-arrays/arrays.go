//            <p>In Go, an <em>array</em> is a numbered sequence of elements of a
//specific length.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p>Here we create an array <code>a</code> that will hold exactly
//5 <code>int</code>s. The type of elements and length are both
//part of the array&rsquo;s type. By default an array is
//zero-valued, which for <code>int</code>s means <code>0</code>s.</p>
//


var a [ 5 ] int
fmt . Println ( "emp:" , a )

//            <p>We can set a value at an index using the
//<code>array[index] = value</code> syntax, and get a value with
//<code>array[index]</code>.</p>
//


a [ 4 ] = 100
fmt . Println ( "set:" , a )
fmt . Println ( "get:" , a [ 4 ])

//            <p>The builtin <code>len</code> returns the length of an array.</p>
//


fmt . Println ( "len:" , len ( a ))

//            <p>Use this syntax to declare and initialize an array
//in one line.</p>
//


b := [ 5 ] int { 1 , 2 , 3 , 4 , 5 }
fmt . Println ( "dcl:" , b )

//            <p>Array types are one-dimensional, but you can
//compose types to build multi-dimensional data
//structures.</p>
//


var twoD [ 2 ][ 3 ] int
for i := 0 ; i < 2 ; i ++ {
for j := 0 ; j < 3 ; j ++ {
twoD [ i ][ j ] = i + j
}
}
fmt . Println ( "2d: " , twoD )
}

