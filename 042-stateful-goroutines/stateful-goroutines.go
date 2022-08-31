//            <p>In the previous example we used explicit locking with
//<a href="mutexes">mutexes</a> to synchronize access to shared state
//across multiple goroutines. Another option is to use the
//built-in synchronization features of  goroutines and
//channels to achieve the same result. This channel-based
//approach aligns with Go&rsquo;s ideas of sharing memory by
//communicating and having each piece of data owned
//by exactly 1 goroutine.</p>
//
//            

package main

//            

import (
"fmt"
"math/rand"
"sync/atomic"
"time"
)

//            <p>In this example our state will be owned by a single
//goroutine. This will guarantee that the data is never
//corrupted with concurrent access. In order to read or
//write that state, other goroutines will send messages
//to the owning goroutine and receive corresponding
//replies. These <code>readOp</code> and <code>writeOp</code> <code>struct</code>s
//encapsulate those requests and a way for the owning
//goroutine to respond.</p>
//


type readOp struct {
key int
resp chan int
}
type writeOp struct {
key int
val int
resp chan bool
}

//            

func main () {

//            <p>As before we&rsquo;ll count how many operations we perform.</p>
//


var readOps uint64
var writeOps uint64

//            <p>The <code>reads</code> and <code>writes</code> channels will be used by
//other goroutines to issue read and write requests,
//respectively.</p>
//


reads := make ( chan readOp )
writes := make ( chan writeOp )

//            <p>Here is the goroutine that owns the <code>state</code>, which
//is a map as in the previous example but now private
//to the stateful goroutine. This goroutine repeatedly
//selects on the <code>reads</code> and <code>writes</code> channels,
//responding to requests as they arrive. A response
//is executed by first performing the requested
//operation and then sending a value on the response
//channel <code>resp</code> to indicate success (and the desired
//value in the case of <code>reads</code>).</p>
//


go func () {
var state = make ( map [ int ] int )
for {
select {
case read := <- reads :
read . resp <- state [ read . key ]
case write := <- writes :
state [ write . key ] = write . val
write . resp <- true
}
}
}()

//            <p>This starts 100 goroutines to issue reads to the
//state-owning goroutine via the <code>reads</code> channel.
//Each read requires constructing a <code>readOp</code>, sending
//it over the <code>reads</code> channel, and then receiving the
//result over the provided <code>resp</code> channel.</p>
//


for r := 0 ; r < 100 ; r ++ {
go func () {
for {
read := readOp {
key : rand . Intn ( 5 ),
resp : make ( chan int )}
reads <- read
<- read . resp
atomic . AddUint64 ( & readOps , 1 )
time . Sleep ( time . Millisecond )
}
}()
}

//            <p>We start 10 writes as well, using a similar
//approach.</p>
//


for w := 0 ; w < 10 ; w ++ {
go func () {
for {
write := writeOp {
key : rand . Intn ( 5 ),
val : rand . Intn ( 100 ),
resp : make ( chan bool )}
writes <- write
<- write . resp
atomic . AddUint64 ( & writeOps , 1 )
time . Sleep ( time . Millisecond )
}
}()
}

//            <p>Let the goroutines work for a second.</p>
//


time . Sleep ( time . Second )

//            <p>Finally, capture and report the op counts.</p>
//


readOpsFinal := atomic . LoadUint64 ( & readOps )
fmt . Println ( "readOps:" , readOpsFinal )
writeOpsFinal := atomic . LoadUint64 ( & writeOps )
fmt . Println ( "writeOps:" , writeOpsFinal )
}

