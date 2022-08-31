//            <p>The Go standard library comes with excellent support
//for HTTP clients and servers in the <code>net/http</code>
//package. In this example we&rsquo;ll use it to issue simple
//HTTP requests.</p>
//


package main

//            

import (
"bufio"
"fmt"
"net/http"
)

//            

func main () {

//            <p>Issue an HTTP GET request to a server. <code>http.Get</code> is a
//convenient shortcut around creating an <code>http.Client</code>
//object and calling its <code>Get</code> method; it uses the
//<code>http.DefaultClient</code> object which has useful default
//settings.</p>
//


resp , err := http . Get ( "https://gobyexample.com" )
if err != nil {
panic ( err )
}
defer resp . Body . Close ()

//            <p>Print the HTTP response status.</p>
//


fmt . Println ( "Response status:" , resp . Status )

//            <p>Print the first 5 lines of the response body.</p>
//


scanner := bufio . NewScanner ( resp . Body )
for i := 0 ; scanner . Scan () && i < 5 ; i ++ {
fmt . Println ( scanner . Text ())
}

//            

if err := scanner . Err (); err != nil {
panic ( err )
}
}

