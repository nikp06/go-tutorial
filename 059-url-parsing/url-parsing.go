//            <p>URLs provide a <a href="https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/">uniform way to locate resources</a>.
//Here&rsquo;s how to parse URLs in Go.</p>
//
//            

package main

//            

import (
"fmt"
"net"
"net/url"
)

//            

func main () {

//            <p>We&rsquo;ll parse this example URL, which includes a
//scheme, authentication info, host, port, path,
//query params, and query fragment.</p>
//


s := "postgres://user:pass@host.com:5432/path?k=v#f"

//            <p>Parse the URL and ensure there are no errors.</p>
//


u , err := url . Parse ( s )
if err != nil {
panic ( err )
}

//            <p>Accessing the scheme is straightforward.</p>
//


fmt . Println ( u . Scheme )

//            <p><code>User</code> contains all authentication info; call
//<code>Username</code> and <code>Password</code> on this for individual
//values.</p>
//


fmt . Println ( u . User )
fmt . Println ( u . User . Username ())
p , _ := u . User . Password ()
fmt . Println ( p )

//            <p>The <code>Host</code> contains both the hostname and the port,
//if present. Use <code>SplitHostPort</code> to extract them.</p>
//


fmt . Println ( u . Host )
host , port , _ := net . SplitHostPort ( u . Host )
fmt . Println ( host )
fmt . Println ( port )

//            <p>Here we extract the <code>path</code> and the fragment after
//the <code>#</code>.</p>
//


fmt . Println ( u . Path )
fmt . Println ( u . Fragment )

//            <p>To get query params in a string of <code>k=v</code> format,
//use <code>RawQuery</code>. You can also parse query params
//into a map. The parsed query param maps are from
//strings to slices of strings, so index into <code>[0]</code>
//if you only want the first value.</p>
//


fmt . Println ( u . RawQuery )
m , _ := url . ParseQuery ( u . RawQuery )
fmt . Println ( m )
fmt . Println ( m [ "k" ][ 0 ])
}

