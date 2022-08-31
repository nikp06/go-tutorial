//            <p><a href="https://en.wikipedia.org/wiki/Variadic_function"><em>Variadic functions</em></a>
//can be called with any number of trailing arguments.
//For example, <code>fmt.Println</code> is a common variadic
//function.</p>
//
//            

package main

//            

import "fmt"

//            <p>Here&rsquo;s a function that will take an arbitrary number
//of <code>int</code>s as arguments.</p>
//


func sum ( nums ... int ) {
fmt . Print ( nums , " " )
total := 0

//            <p>Within the function, the type of <code>nums</code> is
//equivalent to <code>[]int</code>. We can call <code>len(nums)</code>,
//iterate over it with <code>range</code>, etc.</p>
//


for _ , num := range nums {
total += num
}
fmt . Println ( total )
}

//            

func main () {

//            <p>Variadic functions can be called in the usual way
//with individual arguments.</p>
//


sum ( 1 , 2 )
sum ( 1 , 2 , 3 )

//            <p>If you already have multiple args in a slice,
//apply them to a variadic function using
//<code>func(slice...)</code> like this.</p>
//


nums := [] int { 1 , 2 , 3 , 4 }
sum ( nums ... )
}

