//            <p><code>for</code> is Go&rsquo;s only looping construct. Here are
//some basic types of <code>for</code> loops.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p>The most basic type, with a single condition.</p>
//


i := 1
for i <= 3 {
fmt . Println ( i )
i = i + 1
}

//            <p>A classic initial/condition/after <code>for</code> loop.</p>
//


for j := 7 ; j <= 9 ; j ++ {
fmt . Println ( j )
}

//            <p><code>for</code> without a condition will loop repeatedly
//until you <code>break</code> out of the loop or <code>return</code> from
//the enclosing function.</p>
//


for {
fmt . Println ( "loop" )
break
}

//            <p>You can also <code>continue</code> to the next iteration of
//the loop.</p>
//


for n := 0 ; n <= 5 ; n ++ {
if n % 2 == 0 {
continue
}
fmt . Println ( n )
}
}

