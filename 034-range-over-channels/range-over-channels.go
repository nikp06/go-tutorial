//            <p>In a <a href="range">previous</a> example we saw how <code>for</code> and
//<code>range</code> provide iteration over basic data structures.
//We can also use this syntax to iterate over
//values received from a channel.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p>We&rsquo;ll iterate over 2 values in the <code>queue</code> channel.</p>
//


queue := make ( chan string , 2 )
queue <- "one"
queue <- "two"
close ( queue )

//            <p>This <code>range</code> iterates over each element as it&rsquo;s
//received from <code>queue</code>. Because we <code>close</code>d the
//channel above, the iteration terminates after
//receiving the 2 elements.</p>
//


for elem := range queue {
fmt . Println ( elem )
}
}

