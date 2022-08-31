//            <p>Basic sends and receives on channels are blocking.
//However, we can use <code>select</code> with a <code>default</code> clause to
//implement <em>non-blocking</em> sends, receives, and even
//non-blocking multi-way <code>select</code>s.</p>
//
//            

package main

//            

import "fmt"

//            

func main () {
messages := make ( chan string )
signals := make ( chan bool )

//            <p>Here&rsquo;s a non-blocking receive. If a value is
//available on <code>messages</code> then <code>select</code> will take
//the <code>&lt;-messages</code> <code>case</code> with that value. If not
//it will immediately take the <code>default</code> case.</p>
//


select {
case msg := <- messages :
fmt . Println ( "received message" , msg )
default :
fmt . Println ( "no message received" )
}

//            <p>A non-blocking send works similarly. Here <code>msg</code>
//cannot be sent to the <code>messages</code> channel, because
//the channel has no buffer and there is no receiver.
//Therefore the <code>default</code> case is selected.</p>
//


msg := "hi"
select {
case messages <- msg :
fmt . Println ( "sent message" , msg )
default :
fmt . Println ( "no message sent" )
}

//            <p>We can use multiple <code>case</code>s above the <code>default</code>
//clause to implement a multi-way non-blocking
//select. Here we attempt non-blocking receives
//on both <code>messages</code> and <code>signals</code>.</p>
//


select {
case msg := <- messages :
fmt . Println ( "received message" , msg )
case sig := <- signals :
fmt . Println ( "received signal" , sig )
default :
fmt . Println ( "no activity" )
}
}

