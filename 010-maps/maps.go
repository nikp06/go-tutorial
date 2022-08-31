//            <p><em>Maps</em> are Go&rsquo;s built-in <a href="https://en.wikipedia.org/wiki/Associative_array">associative data type</a>
//(sometimes called <em>hashes</em> or <em>dicts</em> in other languages).</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p>To create an empty map, use the builtin <code>make</code>:
//<code>make(map[key-type]val-type)</code>.</p>
//


m := make ( map [ string ] int )

//            <p>Set key/value pairs using typical <code>name[key] = val</code>
//syntax.</p>
//


m [ "k1" ] = 7
m [ "k2" ] = 13

//            <p>Printing a map with e.g. <code>fmt.Println</code> will show all of
//its key/value pairs.</p>
//


fmt . Println ( "map:" , m )

//            <p>Get a value for a key with <code>name[key]</code>.</p>
//


v1 := m [ "k1" ]
fmt . Println ( "v1: " , v1 )

//            <p>The builtin <code>len</code> returns the number of key/value
//pairs when called on a map.</p>
//


fmt . Println ( "len:" , len ( m ))

//            <p>The builtin <code>delete</code> removes key/value pairs from
//a map.</p>
//


delete ( m , "k2" )
fmt . Println ( "map:" , m )

//            <p>The optional second return value when getting a
//value from a map indicates if the key was present
//in the map. This can be used to disambiguate
//between missing keys and keys with zero values
//like <code>0</code> or <code>&quot;&quot;</code>. Here we didn&rsquo;t need the value
//itself, so we ignored it with the <em>blank identifier</em>
//<code>_</code>.</p>
//


_ , prs := m [ "k2" ]
fmt . Println ( "prs:" , prs )

//            <p>You can also declare and initialize a new map in
//the same line with this syntax.</p>
//


n := map [ string ] int { "foo" : 1 , "bar" : 2 }
fmt . Println ( "map:" , n )
}

