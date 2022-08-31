//            <p><em>Closing</em> a channel indicates that no more values
//will be sent on it. This can be useful to communicate
//completion to the channel&rsquo;s receivers.</p>
//
//            

package main

//            

import "fmt"

//            <p>In this example we&rsquo;ll use a <code>jobs</code> channel to
//communicate work to be done from the <code>main()</code> goroutine
//to a worker goroutine. When we have no more jobs for
//the worker we&rsquo;ll <code>close</code> the <code>jobs</code> channel.</p>
//


func main () {
jobs := make ( chan int , 5 )
done := make ( chan bool )

//            <p>Here&rsquo;s the worker goroutine. It repeatedly receives
//from <code>jobs</code> with <code>j, more := &lt;-jobs</code>. In this
//special 2-value form of receive, the <code>more</code> value
//will be <code>false</code> if <code>jobs</code> has been <code>close</code>d and all
//values in the channel have already been received.
//We use this to notify on <code>done</code> when we&rsquo;ve worked
//all our jobs.</p>
//


go func () {
for {
j , more := <- jobs
if more {
fmt . Println ( "received job" , j )
} else {
fmt . Println ( "received all jobs" )
done <- true
return
}
}
}()

//            <p>This sends 3 jobs to the worker over the <code>jobs</code>
//channel, then closes it.</p>
//


for j := 1 ; j <= 3 ; j ++ {
jobs <- j
fmt . Println ( "sent job" , j )
}
close ( jobs )
fmt . Println ( "sent all jobs" )

//            <p>We await the worker using the
//<a href="channel-synchronization">synchronization</a> approach
//we saw earlier.</p>
//


<- done
}

