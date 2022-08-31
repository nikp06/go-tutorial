//            <p>A Go string is a read-only slice of bytes. The language
//and the standard library treat strings specially - as
//containers of text encoded in <a href="https://en.wikipedia.org/wiki/UTF-8">UTF-8</a>.
//In other languages, strings are made of &ldquo;characters&rdquo;.
//In Go, the concept of a character is called a <code>rune</code> - it&rsquo;s
//an integer that represents a Unicode code point.
//<a href="https://go.dev/blog/strings">This Go blog post</a> is a good
//introduction to the topic.</p>
//
//            

package main

//            

import (
"fmt"
"unicode/utf8"
)

//            

func main () {

//            <p><code>s</code> is a <code>string</code> assigned a literal value
//representing the word &ldquo;hello&rdquo; in the Thai
//language. Go string literals are UTF-8
//encoded text.</p>
//


const s = "สวัสดี"

//            <p>Since strings are equivalent to <code>[]byte</code>, this
//will produce the length of the raw bytes stored within.</p>
//


fmt . Println ( "Len:" , len ( s ))

//            <p>Indexing into a string produces the raw byte values at
//each index. This loop generates the hex values of all
//the bytes that constitute the code points in <code>s</code>.</p>
//


for i := 0 ; i < len ( s ); i ++ {
fmt . Printf ( "%x " , s [ i ])
}
fmt . Println ()

//            <p>To count how many <em>runes</em> are in a string, we can use
//the <code>utf8</code> package. Note that the run-time of
//<code>RuneCountInString</code> dependes on the size of the string,
//because it has to decode each UTF-8 rune sequentially.
//Some Thai characters are represented by multiple UTF-8
//code points, so the result of this count may be surprising.</p>
//


fmt . Println ( "Rune count:" , utf8 . RuneCountInString ( s ))

//            <p>A <code>range</code> loop handles strings specially and decodes
//each <code>rune</code> along with its offset in the string.</p>
//


for idx , runeValue := range s {
fmt . Printf ( "%#U starts at %d\n" , runeValue , idx )
}

//            <p>We can achieve the same iteration by using the
//<code>utf8.DecodeRuneInString</code> function explicitly.</p>
//


fmt . Println ( "\nUsing DecodeRuneInString" )
for i , w := 0 , 0 ; i < len ( s ); i += w {
runeValue , width := utf8 . DecodeRuneInString ( s [ i :])
fmt . Printf ( "%#U starts at %d\n" , runeValue , i )
w = width

//            <p>This demonstrates passing a <code>rune</code> value to a function.</p>
//


examineRune ( runeValue )
}
}

//            

func examineRune ( r rune ) {

//            <p>Values enclosed in single quotes are <em>rune literals</em>. We
//can compare a <code>rune</code> value to a rune literal directly.</p>
//


if r == 't' {
fmt . Println ( "found tee" )
} else if r == 'ส' {
fmt . Println ( "found so sua" )
}
}

