//            <p><em>Slices</em> are a key data type in Go, giving a more
//powerful interface to sequences than arrays.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p>Unlike arrays, slices are typed only by the
//elements they contain (not the number of elements).
//To create an empty slice with non-zero length, use
//the builtin <code>make</code>. Here we make a slice of
//<code>string</code>s of length <code>3</code> (initially zero-valued).</p>
//


s := make ([] string , 3 )
fmt . Println ( "emp:" , s )

//            <p>We can set and get just like with arrays.</p>
//


s [ 0 ] = "a"
s [ 1 ] = "b"
s [ 2 ] = "c"
fmt . Println ( "set:" , s )
fmt . Println ( "get:" , s [ 2 ])

//            <p><code>len</code> returns the length of the slice as expected.</p>
//


fmt . Println ( "len:" , len ( s ))

//            <p>In addition to these basic operations, slices
//support several more that make them richer than
//arrays. One is the builtin <code>append</code>, which
//returns a slice containing one or more new values.
//Note that we need to accept a return value from
//<code>append</code> as we may get a new slice value.</p>
//


s = append ( s , "d" )
s = append ( s , "e" , "f" )
fmt . Println ( "apd:" , s )

//            <p>Slices can also be <code>copy</code>&rsquo;d. Here we create an
//empty slice <code>c</code> of the same length as <code>s</code> and copy
//into <code>c</code> from <code>s</code>.</p>
//


c := make ([] string , len ( s ))
copy ( c , s )
fmt . Println ( "cpy:" , c )

//            <p>Slices support a &ldquo;slice&rdquo; operator with the syntax
//<code>slice[low:high]</code>. For example, this gets a slice
//of the elements <code>s[2]</code>, <code>s[3]</code>, and <code>s[4]</code>.</p>
//


l := s [ 2 : 5 ]
fmt . Println ( "sl1:" , l )

//            <p>This slices up to (but excluding) <code>s[5]</code>.</p>
//


l = s [: 5 ]
fmt . Println ( "sl2:" , l )

//            <p>And this slices up from (and including) <code>s[2]</code>.</p>
//


l = s [ 2 :]
fmt . Println ( "sl3:" , l )

//            <p>We can declare and initialize a variable for slice
//in a single line as well.</p>
//


t := [] string { "g" , "h" , "i" }
fmt . Println ( "dcl:" , t )

//            <p>Slices can be composed into multi-dimensional data
//structures. The length of the inner slices can
//vary, unlike with multi-dimensional arrays.</p>
//


twoD := make ([][] int , 3 )
for i := 0 ; i < 3 ; i ++ {
innerLen := i + 1
twoD [ i ] = make ([] int , innerLen )
for j := 0 ; j < innerLen ; j ++ {
twoD [ i ][ j ] = i + j
}
}
fmt . Println ( "2d: " , twoD )
}

