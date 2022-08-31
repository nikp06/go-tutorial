//            <p><em>Interfaces</em> are named collections of method
//signatures.</p>
//
//            

package main

//            

import (
"fmt"
"math"
)

//            <p>Here&rsquo;s a basic interface for geometric shapes.</p>
//


type geometry interface {
area () float64
perim () float64
}

//            <p>For our example we&rsquo;ll implement this interface on
//<code>rect</code> and <code>circle</code> types.</p>
//


type rect struct {
width , height float64
}
type circle struct {
radius float64
}

//            <p>To implement an interface in Go, we just need to
//implement all the methods in the interface. Here we
//implement <code>geometry</code> on <code>rect</code>s.</p>
//


func ( r rect ) area () float64 {
return r . width * r . height
}
func ( r rect ) perim () float64 {
return 2 * r . width + 2 * r . height
}

//            <p>The implementation for <code>circle</code>s.</p>
//


func ( c circle ) area () float64 {
return math . Pi * c . radius * c . radius
}
func ( c circle ) perim () float64 {
return 2 * math . Pi * c . radius
}

//            <p>If a variable has an interface type, then we can call
//methods that are in the named interface. Here&rsquo;s a
//generic <code>measure</code> function taking advantage of this
//to work on any <code>geometry</code>.</p>
//


func measure ( g geometry ) {
fmt . Println ( g )
fmt . Println ( g . area ())
fmt . Println ( g . perim ())
}

//            

func main () {
r := rect { width : 3 , height : 4 }
c := circle { radius : 5 }

//            <p>The <code>circle</code> and <code>rect</code> struct types both
//implement the <code>geometry</code> interface so we can use
//instances of
//these structs as arguments to <code>measure</code>.</p>
//


measure ( r )
measure ( c )
}

