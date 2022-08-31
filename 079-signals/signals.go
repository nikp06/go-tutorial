//            <p>Sometimes we&rsquo;d like our Go programs to intelligently
//handle <a href="https://en.wikipedia.org/wiki/Unix_signal">Unix signals</a>.
//For example, we might want a server to gracefully
//shutdown when it receives a <code>SIGTERM</code>, or a command-line
//tool to stop processing input if it receives a <code>SIGINT</code>.
//Here&rsquo;s how to handle signals in Go with channels.</p>
//
//            

package main

//            

import (
"fmt"
"os"
"os/signal"
"syscall"
)

//            

func main () {

//            <p>Go signal notification works by sending <code>os.Signal</code>
//values on a channel. We&rsquo;ll create a channel to
//receive these notifications. Note that this channel
//should be buffered.</p>
//


sigs := make ( chan os . Signal , 1 )

//            <p><code>signal.Notify</code> registers the given channel to
//receive notifications of the specified signals.</p>
//


signal . Notify ( sigs , syscall . SIGINT , syscall . SIGTERM )

//            <p>We could receive from <code>sigs</code> here in the main
//function, but let&rsquo;s see how this could also be
//done in a separate goroutine, to demonstrate
//a more realistic scenario of graceful shutdown.</p>
//


done := make ( chan bool , 1 )

//            <p>This goroutine executes a blocking receive for
//signals. When it gets one it&rsquo;ll print it out
//and then notify the program that it can finish.</p>
//

go func () {

//            

sig := <- sigs
fmt . Println ()
fmt . Println ( sig )
done <- true
}()

//            <p>The program will wait here until it gets the
//expected signal (as indicated by the goroutine
//above sending a value on <code>done</code>) and then exit.</p>
//


fmt . Println ( "awaiting signal" )
<- done
fmt . Println ( "exiting" )
}

