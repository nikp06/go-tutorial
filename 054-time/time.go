//            <p>Go offers extensive support for times and durations;
//here are some examples.</p>
//
//            

package main

//            

import (
"fmt"
"time"
)

//            

func main () {
p := fmt . Println

//            <p>We&rsquo;ll start by getting the current time.</p>
//


now := time . Now ()
p ( now )

//            <p>You can build a <code>time</code> struct by providing the
//year, month, day, etc. Times are always associated
//with a <code>Location</code>, i.e. time zone.</p>
//


then := time . Date (
2009 , 11 , 17 , 20 , 34 , 58 , 651387237 , time . UTC )
p ( then )

//            <p>You can extract the various components of the time
//value as expected.</p>
//


p ( then . Year ())
p ( then . Month ())
p ( then . Day ())
p ( then . Hour ())
p ( then . Minute ())
p ( then . Second ())
p ( then . Nanosecond ())
p ( then . Location ())

//            <p>The Monday-Sunday <code>Weekday</code> is also available.</p>
//


p ( then . Weekday ())

//            <p>These methods compare two times, testing if the
//first occurs before, after, or at the same time
//as the second, respectively.</p>
//


p ( then . Before ( now ))
p ( then . After ( now ))
p ( then . Equal ( now ))

//            <p>The <code>Sub</code> methods returns a <code>Duration</code> representing
//the interval between two times.</p>
//


diff := now . Sub ( then )
p ( diff )

//            <p>We can compute the length of the duration in
//various units.</p>
//


p ( diff . Hours ())
p ( diff . Minutes ())
p ( diff . Seconds ())
p ( diff . Nanoseconds ())

//            <p>You can use <code>Add</code> to advance a time by a given
//duration, or with a <code>-</code> to move backwards by a
//duration.</p>
//


p ( then . Add ( diff ))
p ( then . Add ( - diff ))
}

