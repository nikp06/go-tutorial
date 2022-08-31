//            <p>A <code>panic</code> typically means something went unexpectedly
//wrong. Mostly we use it to fail fast on errors that
//shouldn&rsquo;t occur during normal operation, or that we
//aren&rsquo;t prepared to handle gracefully.</p>
//
//            

package main

//            

import "os"

//            

func main () {

//            <p>We&rsquo;ll use panic throughout this site to check for
//unexpected errors. This is the only program on the
//site designed to panic.</p>
//


panic ( "a problem" )

//            <p>A common use of panic is to abort if a function
//returns an error value that we don&rsquo;t know how to
//(or want to) handle. Here&rsquo;s an example of
//<code>panic</code>king if we get an unexpected error when creating a new file.</p>
//


_ , err := os . Create ( "/tmp/file" )
if err != nil {
panic ( err )
}
}

