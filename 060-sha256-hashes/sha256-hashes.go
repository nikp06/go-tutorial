//            <p><a href="https://en.wikipedia.org/wiki/SHA-2"><em>SHA256 hashes</em></a> are
//frequently used to compute short identities for binary
//or text blobs. For example, TLS/SSL certificates use SHA256
//to compute a certificate&rsquo;s signature. Here&rsquo;s how to compute
//SHA256 hashes in Go.</p>
//
//            

package main

//            <p>Go implements several hash functions in various
//<code>crypto/*</code> packages.</p>
//


import (
"crypto/sha256"
"fmt"
)

//            

func main () {
s := "sha256 this string"

//            <p>Here we start with a new hash.</p>
//


h := sha256 . New ()

//            <p><code>Write</code> expects bytes. If you have a string <code>s</code>,
//use <code>[]byte(s)</code> to coerce it to bytes.</p>
//


h . Write ([] byte ( s ))

//            <p>This gets the finalized hash result as a byte
//slice. The argument to <code>Sum</code> can be used to append
//to an existing byte slice: it usually isn&rsquo;t needed.</p>
//


bs := h . Sum ( nil )

//            

fmt . Println ( s )
fmt . Printf ( "%x\n" , bs )
}

