//            <p>Go provides built-in support for <a href="https://en.wikipedia.org/wiki/Base64">base64
//encoding/decoding</a>.</p>
//
//            

package main

//            <p>This syntax imports the <code>encoding/base64</code> package with
//the <code>b64</code> name instead of the default <code>base64</code>. It&rsquo;ll
//save us some space below.</p>
//


import (
b64 "encoding/base64"
"fmt"
)

//            

func main () {

//            <p>Here&rsquo;s the <code>string</code> we&rsquo;ll encode/decode.</p>
//


data := "abc123!?$*&()'-=@~"

//            <p>Go supports both standard and URL-compatible
//base64. Here&rsquo;s how to encode using the standard
//encoder. The encoder requires a <code>[]byte</code> so we
//convert our <code>string</code> to that type.</p>
//


sEnc := b64 . StdEncoding . EncodeToString ([] byte ( data ))
fmt . Println ( sEnc )

//            <p>Decoding may return an error, which you can check
//if you don&rsquo;t already know the input to be
//well-formed.</p>
//


sDec , _ := b64 . StdEncoding . DecodeString ( sEnc )
fmt . Println ( string ( sDec ))
fmt . Println ()

//            <p>This encodes/decodes using a URL-compatible base64
//format.</p>
//


uEnc := b64 . URLEncoding . EncodeToString ([] byte ( data ))
fmt . Println ( uEnc )
uDec , _ := b64 . URLEncoding . DecodeString ( uEnc )
fmt . Println ( string ( uDec ))
}

