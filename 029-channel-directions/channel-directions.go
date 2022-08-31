//            <p>When using channels as function parameters, you can
//specify if a channel is meant to only send or receive
//values. This specificity increases the type-safety of
//the program.</p>
//
//            

package main

//            

import "fmt"

//            <p>This <code>ping</code> function only accepts a channel for sending
//values. It would be a compile-time error to try to
//receive on this channel.</p>
//


func ping ( pings chan <- string , msg string ) {
pings <- msg
}

//            <p>The <code>pong</code> function accepts one channel for receives
//(<code>pings</code>) and a second for sends (<code>pongs</code>).</p>
//


func pong ( pings <- chan string , pongs chan <- string ) {
msg := <- pings
pongs <- msg
}

//            

func main () {
pings := make ( chan string , 1 )
pongs := make ( chan string , 1 )
ping ( pings , "passed message" )
pong ( pings , pongs )
fmt . Println ( <- pongs )
}

