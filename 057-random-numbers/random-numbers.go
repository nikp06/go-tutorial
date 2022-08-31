//            <p>Go&rsquo;s <code>math/rand</code> package provides
//<a href="https://en.wikipedia.org/wiki/Pseudorandom_number_generator">pseudorandom number</a>
//generation.</p>
//
//            

package main

//            

import (
"fmt"
"math/rand"
"time"
)

//            

func main () {

//            <p>For example, <code>rand.Intn</code> returns a random <code>int</code> n,
//<code>0 &lt;= n &lt; 100</code>.</p>
//


fmt . Print ( rand . Intn ( 100 ), "," )
fmt . Print ( rand . Intn ( 100 ))
fmt . Println ()

//            <p><code>rand.Float64</code> returns a <code>float64</code> <code>f</code>,
//<code>0.0 &lt;= f &lt; 1.0</code>.</p>
//


fmt . Println ( rand . Float64 ())

//            <p>This can be used to generate random floats in
//other ranges, for example <code>5.0 &lt;= f' &lt; 10.0</code>.</p>
//


fmt . Print (( rand . Float64 () * 5 ) + 5 , "," )
fmt . Print (( rand . Float64 () * 5 ) + 5 )
fmt . Println ()

//            <p>The default number generator is deterministic, so it&rsquo;ll
//produce the same sequence of numbers each time by default.
//To produce varying sequences, give it a seed that changes.
//Note that this is not safe to use for random numbers you
//intend to be secret, use <code>crypto/rand</code> for those.</p>
//


s1 := rand . NewSource ( time . Now (). UnixNano ())
r1 := rand . New ( s1 )

//            <p>Call the resulting <code>rand.Rand</code> just like the
//functions on the <code>rand</code> package.</p>
//


fmt . Print ( r1 . Intn ( 100 ), "," )
fmt . Print ( r1 . Intn ( 100 ))
fmt . Println ()

//            <p>If you seed a source with the same number, it
//produces the same sequence of random numbers.</p>
//


s2 := rand . NewSource ( 42 )
r2 := rand . New ( s2 )
fmt . Print ( r2 . Intn ( 100 ), "," )
fmt . Print ( r2 . Intn ( 100 ))
fmt . Println ()
s3 := rand . NewSource ( 42 )
r3 := rand . New ( s3 )
fmt . Print ( r3 . Intn ( 100 ), "," )
fmt . Print ( r3 . Intn ( 100 ))
}

