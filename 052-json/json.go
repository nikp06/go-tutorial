//            <p>Go offers built-in support for JSON encoding and
//decoding, including to and from built-in and custom
//data types.</p>
//
//            

package main

//            

import (
"encoding/json"
"fmt"
"os"
)

//            <p>We&rsquo;ll use these two structs to demonstrate encoding and
//decoding of custom types below.</p>
//


type response1 struct {
Page int
Fruits [] string
}

//            <p>Only exported fields will be encoded/decoded in JSON.
//Fields must start with capital letters to be exported.</p>
//


type response2 struct {
Page int `json:"page"`
Fruits [] string `json:"fruits"`
}

//            

func main () {

//            <p>First we&rsquo;ll look at encoding basic data types to
//JSON strings. Here are some examples for atomic
//values.</p>
//


bolB , _ := json . Marshal ( true )
fmt . Println ( string ( bolB ))

//            

intB , _ := json . Marshal ( 1 )
fmt . Println ( string ( intB ))

//            

fltB , _ := json . Marshal ( 2.34 )
fmt . Println ( string ( fltB ))

//            

strB , _ := json . Marshal ( "gopher" )
fmt . Println ( string ( strB ))

//            <p>And here are some for slices and maps, which encode
//to JSON arrays and objects as you&rsquo;d expect.</p>
//


slcD := [] string { "apple" , "peach" , "pear" }
slcB , _ := json . Marshal ( slcD )
fmt . Println ( string ( slcB ))

//            

mapD := map [ string ] int { "apple" : 5 , "lettuce" : 7 }
mapB , _ := json . Marshal ( mapD )
fmt . Println ( string ( mapB ))

//            <p>The JSON package can automatically encode your
//custom data types. It will only include exported
//fields in the encoded output and will by default
//use those names as the JSON keys.</p>
//


res1D := & response1 {
Page : 1 ,
Fruits : [] string { "apple" , "peach" , "pear" }}
res1B , _ := json . Marshal ( res1D )
fmt . Println ( string ( res1B ))

//            <p>You can use tags on struct field declarations
//to customize the encoded JSON key names. Check the
//definition of <code>response2</code> above to see an example
//of such tags.</p>
//


res2D := & response2 {
Page : 1 ,
Fruits : [] string { "apple" , "peach" , "pear" }}
res2B , _ := json . Marshal ( res2D )
fmt . Println ( string ( res2B ))

//            <p>Now let&rsquo;s look at decoding JSON data into Go
//values. Here&rsquo;s an example for a generic data
//structure.</p>
//


byt := [] byte ( `{"num":6.13,"strs":["a","b"]}` )

//            <p>We need to provide a variable where the JSON
//package can put the decoded data. This
//<code>map[string]interface{}</code> will hold a map of strings
//to arbitrary data types.</p>
//


var dat map [ string ] interface {}

//            <p>Here&rsquo;s the actual decoding, and a check for
//associated errors.</p>
//


if err := json . Unmarshal ( byt , & dat ); err != nil {
panic ( err )
}
fmt . Println ( dat )

//            <p>In order to use the values in the decoded map,
//we&rsquo;ll need to convert them to their appropriate type.
//For example here we convert the value in <code>num</code> to
//the expected <code>float64</code> type.</p>
//


num := dat [ "num" ].( float64 )
fmt . Println ( num )

//            <p>Accessing nested data requires a series of
//conversions.</p>
//


strs := dat [ "strs" ].([] interface {})
str1 := strs [ 0 ].( string )
fmt . Println ( str1 )

//            <p>We can also decode JSON into custom data types.
//This has the advantages of adding additional
//type-safety to our programs and eliminating the
//need for type assertions when accessing the decoded
//data.</p>
//


str := `{"page": 1, "fruits": ["apple", "peach"]}`
res := response2 {}
json . Unmarshal ([] byte ( str ), & res )
fmt . Println ( res )
fmt . Println ( res . Fruits [ 0 ])

//            <p>In the examples above we always used bytes and
//strings as intermediates between the data and
//JSON representation on standard out. We can also
//stream JSON encodings directly to <code>os.Writer</code>s like
//<code>os.Stdout</code> or even HTTP response bodies.</p>
//


enc := json . NewEncoder ( os . Stdout )
d := map [ string ] int { "apple" : 5 , "lettuce" : 7 }
enc . Encode ( d )
}

