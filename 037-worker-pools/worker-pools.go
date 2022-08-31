//            <p>In this example we&rsquo;ll look at how to implement
//a <em>worker pool</em> using goroutines and channels.</p>
//
//            

package main

//            

import (
"fmt"
"time"
)

//            <p>Here&rsquo;s the worker, of which we&rsquo;ll run several
//concurrent instances. These workers will receive
//work on the <code>jobs</code> channel and send the corresponding
//results on <code>results</code>. We&rsquo;ll sleep a second per job to
//simulate an expensive task.</p>
//


func worker ( id int , jobs <- chan int , results chan <- int ) {
for j := range jobs {
fmt . Println ( "worker" , id , "started  job" , j )
time . Sleep ( time . Second )
fmt . Println ( "worker" , id , "finished job" , j )
results <- j * 2
}
}

//            

func main () {

//            <p>In order to use our pool of workers we need to send
//them work and collect their results. We make 2
//channels for this.</p>
//


const numJobs = 5
jobs := make ( chan int , numJobs )
results := make ( chan int , numJobs )

//            <p>This starts up 3 workers, initially blocked
//because there are no jobs yet.</p>
//


for w := 1 ; w <= 3 ; w ++ {
go worker ( w , jobs , results )
}

//            <p>Here we send 5 <code>jobs</code> and then <code>close</code> that
//channel to indicate that&rsquo;s all the work we have.</p>
//


for j := 1 ; j <= numJobs ; j ++ {
jobs <- j
}
close ( jobs )

//            <p>Finally we collect all the results of the work.
//This also ensures that the worker goroutines have
//finished. An alternative way to wait for multiple
//goroutines is to use a <a href="waitgroups">WaitGroup</a>.</p>
//


for a := 1 ; a <= numJobs ; a ++ {
<- results
}
}

