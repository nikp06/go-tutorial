//            <p>Sometimes we&rsquo;ll want to sort a collection by something
//other than its natural order. For example, suppose we
//wanted to sort strings by their length instead of
//alphabetically. Here&rsquo;s an example of custom sorts
//in Go.</p>
//
//            

package main

//            

import (
"fmt"
"sort"
)

//            <p>In order to sort by a custom function in Go, we need a
//corresponding type. Here we&rsquo;ve created a <code>byLength</code>
//type that is just an alias for the builtin <code>[]string</code>
//type.</p>
//


type byLength [] string

//            <p>We implement <code>sort.Interface</code> - <code>Len</code>, <code>Less</code>, and
//<code>Swap</code> - on our type so we can use the <code>sort</code> package&rsquo;s
//generic <code>Sort</code> function. <code>Len</code> and <code>Swap</code>
//will usually be similar across types and <code>Less</code> will
//hold the actual custom sorting logic. In our case we
//want to sort in order of increasing string length, so
//we use <code>len(s[i])</code> and <code>len(s[j])</code> here.</p>
//


func ( s byLength ) Len () int {
return len ( s )
}
func ( s byLength ) Swap ( i , j int ) {
s [ i ], s [ j ] = s [ j ], s [ i ]
}
func ( s byLength ) Less ( i , j int ) bool {
return len ( s [ i ]) < len ( s [ j ])
}

//            <p>With all of this in place, we can now implement our
//custom sort by converting the original <code>fruits</code> slice
//to <code>byLength</code>, and then use <code>sort.Sort</code> on that typed
//slice.</p>
//


func main () {
fruits := [] string { "peach" , "banana" , "kiwi" }
sort . Sort ( byLength ( fruits ))
fmt . Println ( fruits )
}

