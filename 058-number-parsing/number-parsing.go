//            <p>Parsing numbers from strings is a basic but common task
//in many programs; here&rsquo;s how to do it in Go.</p>
//
//            

package main

//            <p>The built-in package <code>strconv</code> provides the number
//parsing.</p>
//


import (
"fmt"
"strconv"
)

//            

func main () {

//            <p>With <code>ParseFloat</code>, this <code>64</code> tells how many bits of
//precision to parse.</p>
//


f , _ := strconv . ParseFloat ( "1.234" , 64 )
fmt . Println ( f )

//            <p>For <code>ParseInt</code>, the <code>0</code> means infer the base from
//the string. <code>64</code> requires that the result fit in 64
//bits.</p>
//


i , _ := strconv . ParseInt ( "123" , 0 , 64 )
fmt . Println ( i )

//            <p><code>ParseInt</code> will recognize hex-formatted numbers.</p>
//


d , _ := strconv . ParseInt ( "0x1c8" , 0 , 64 )
fmt . Println ( d )

//            <p>A <code>ParseUint</code> is also available.</p>
//


u , _ := strconv . ParseUint ( "789" , 0 , 64 )
fmt . Println ( u )

//            <p><code>Atoi</code> is a convenience function for basic base-10
//<code>int</code> parsing.</p>
//


k , _ := strconv . Atoi ( "135" )
fmt . Println ( k )

//            <p>Parse functions return an error on bad input.</p>
//


_ , e := strconv . Atoi ( "wat" )
fmt . Println ( e )
}

