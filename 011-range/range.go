//            <p><em>range</em> iterates over elements in a variety of data
//structures. Let&rsquo;s see how to use <code>range</code> with some
//of the data structures we&rsquo;ve already learned.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p>Here we use <code>range</code> to sum the numbers in a slice.
//Arrays work like this too.</p>
//


nums := [] int { 2 , 3 , 4 }
sum := 0
for _ , num := range nums {
sum += num
}
fmt . Println ( "sum:" , sum )

//            <p><code>range</code> on arrays and slices provides both the
//index and value for each entry. Above we didn&rsquo;t
//need the index, so we ignored it with the
//blank identifier <code>_</code>. Sometimes we actually want
//the indexes though.</p>
//


for i , num := range nums {
if num == 3 {
fmt . Println ( "index:" , i )
}
}

//            <p><code>range</code> on map iterates over key/value pairs.</p>
//


kvs := map [ string ] string { "a" : "apple" , "b" : "banana" }
for k , v := range kvs {
fmt . Printf ( "%s -&gt; %s\n" , k , v )
}

//            <p><code>range</code> can also iterate over just the keys of a map.</p>
//


for k := range kvs {
fmt . Println ( "key:" , k )
}

//            <p><code>range</code> on strings iterates over Unicode code
//points. The first value is the starting byte index
//of the <code>rune</code> and the second the <code>rune</code> itself.
//See <a href="strings-and-runes">Strings and Runes</a> for more
//details.</p>
//


for i , c := range "go" {
fmt . Println ( i , c )
}
}

