//            <p>Go supports time formatting and parsing via
//pattern-based layouts.</p>
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

//            <p>Here&rsquo;s a basic example of formatting a time
//according to RFC3339, using the corresponding layout
//constant.</p>
//


t := time . Now ()
p ( t . Format ( time . RFC3339 ))

//            <p>Time parsing uses the same layout values as <code>Format</code>.</p>
//


t1 , e := time . Parse (
time . RFC3339 ,
"2012-11-01T22:08:41+00:00" )
p ( t1 )

//            <p><code>Format</code> and <code>Parse</code> use example-based layouts. Usually
//you&rsquo;ll use a constant from <code>time</code> for these layouts, but
//you can also supply custom layouts. Layouts must use the
//reference time <code>Mon Jan 2 15:04:05 MST 2006</code> to show the
//pattern with which to format/parse a given time/string.
//The example time must be exactly as shown: the year 2006,
//15 for the hour, Monday for the day of the week, etc.</p>
//


p ( t . Format ( "3:04PM" ))
p ( t . Format ( "Mon Jan _2 15:04:05 2006" ))
p ( t . Format ( "2006-01-02T15:04:05.999999-07:00" ))
form := "3 04 PM"
t2 , e := time . Parse ( form , "8 41 PM" )
p ( t2 )

//            <p>For purely numeric representations you can also
//use standard string formatting with the extracted
//components of the time value.</p>
//


fmt . Printf ( "%d-%02d-%02dT%02d:%02d:%02d-00:00\n" ,
t . Year (), t . Month (), t . Day (),
t . Hour (), t . Minute (), t . Second ())

//            <p><code>Parse</code> will return an error on malformed input
//explaining the parsing problem.</p>
//


ansic := "Mon Jan _2 15:04:05 2006"
_ , e = time . Parse ( ansic , "8:41PM" )
p ( e )
}

