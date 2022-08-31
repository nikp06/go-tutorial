//            <p>By default channels are <em>unbuffered</em>, meaning that they
//will only accept sends (<code>chan &lt;-</code>) if there is a
//corresponding receive (<code>&lt;- chan</code>) ready to receive the
//sent value. <em>Buffered channels</em> accept a limited
//number of  values without a corresponding receiver for
//those values.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {

//            <p>Here we <code>make</code> a channel of strings buffering up to
//2 values.</p>
//


messages := make ( chan string , 2 )

//            <p>Because this channel is buffered, we can send these
//values into the channel without a corresponding
//concurrent receive.</p>
//


messages <- "buffered"
messages <- "channel"

//            <p>Later we can receive these two values as usual.</p>
//


fmt . Println ( <- messages )
fmt . Println ( <- messages )
}

