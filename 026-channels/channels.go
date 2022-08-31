//            <p><em>Channels</em> are the pipes that connect concurrent
//goroutines. You can send values into channels from one
//goroutine and receive those values into another
//goroutine.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p>Create a new channel with <code>make(chan val-type)</code>.
//Channels are typed by the values they convey.</p>
//


messages := make ( chan string )

//            <p><em>Send</em> a value into a channel using the <code>channel &lt;-</code>
//syntax. Here we send <code>&quot;ping&quot;</code>  to the <code>messages</code>
//channel we made above, from a new goroutine.</p>
//


go func () { messages <- "ping" }()

//            <p>The <code>&lt;-channel</code> syntax <em>receives</em> a value from the
//channel. Here we&rsquo;ll receive the <code>&quot;ping&quot;</code> message
//we sent above and print it out.</p>
//


msg := <- messages
fmt . Println ( msg )
}

