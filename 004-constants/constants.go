//            <p>Go supports <em>constants</em> of character, string, boolean,
//and numeric values.</p>
//
//            

package main

//            

import (
"fmt"
"math"
)

//            <p><code>const</code> declares a constant value.</p>
//


const s string = "constant"

//            

func main () {
fmt . Println ( s )

//            <p>A <code>const</code> statement can appear anywhere a <code>var</code>
//statement can.</p>
//


const n = 500000000

//            <p>Constant expressions perform arithmetic with
//arbitrary precision.</p>
//


const d = 3e20 / n
fmt . Println ( d )

//            <p>A numeric constant has no type until it&rsquo;s given
//one, such as by an explicit conversion.</p>
//


fmt . Println ( int64 ( d ))

//            <p>A number can be given a type by using it in a
//context that requires one, such as a variable
//assignment or function call. For example, here
//<code>math.Sin</code> expects a <code>float64</code>.</p>
//


fmt . Println ( math . Sin ( n ))
}

