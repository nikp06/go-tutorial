//            <p>Go&rsquo;s <code>sort</code> package implements sorting for builtins
//and user-defined types. We&rsquo;ll look at sorting for
//builtins first.</p>
//
//            

package main

//            

import (
"fmt"
"sort"
)

//            

func main () {

//            <p>Sort methods are specific to the builtin type;
//here&rsquo;s an example for strings. Note that sorting is
//in-place, so it changes the given slice and doesn&rsquo;t
//return a new one.</p>
//


strs := [] string { "c" , "a" , "b" }
sort . Strings ( strs )
fmt . Println ( "Strings:" , strs )

//            <p>An example of sorting <code>int</code>s.</p>
//


ints := [] int { 7 , 2 , 4 }
sort . Ints ( ints )
fmt . Println ( "Ints:   " , ints )

//            <p>We can also use <code>sort</code> to check if a slice is
//already in sorted order.</p>
//


s := sort . IntsAreSorted ( ints )
fmt . Println ( "Sorted: " , s )
}

