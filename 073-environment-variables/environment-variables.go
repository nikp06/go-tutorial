//            <p><a href="https://en.wikipedia.org/wiki/Environment_variable">Environment variables</a>
//are a universal mechanism for <a href="https://www.12factor.net/config">conveying configuration
//information to Unix programs</a>.
//Let&rsquo;s look at how to set, get, and list environment variables.</p>
//
//            

package main

//            

import (
"fmt"
"os"
"strings"
)

//            

func main () {

//            <p>To set a key/value pair, use <code>os.Setenv</code>. To get a
//value for a key, use <code>os.Getenv</code>. This will return
//an empty string if the key isn&rsquo;t present in the
//environment.</p>
//


os . Setenv ( "FOO" , "1" )
fmt . Println ( "FOO:" , os . Getenv ( "FOO" ))
fmt . Println ( "BAR:" , os . Getenv ( "BAR" ))

//            <p>Use <code>os.Environ</code> to list all key/value pairs in the
//environment. This returns a slice of strings in the
//form <code>KEY=value</code>. You can <code>strings.SplitN</code> them to
//get the key and value. Here we print all the keys.</p>
//


fmt . Println ()
for _ , e := range os . Environ () {
pair := strings . SplitN ( e , "=" , 2 )
fmt . Println ( pair [ 0 ])
}
}

