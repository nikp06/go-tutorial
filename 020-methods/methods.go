//            <p>Go supports <em>methods</em> defined on struct types.</p>
//
//            

package main

//            

import "fmt"

//            

type rect struct {
width , height int
}

//            <p>This <code>area</code> method has a <em>receiver type</em> of <code>*rect</code>.</p>
//


func ( r * rect ) area () int {
return r . width * r . height
}

//            <p>Methods can be defined for either pointer or value
//receiver types. Here&rsquo;s an example of a value receiver.</p>
//


func ( r rect ) perim () int {
return 2 * r . width + 2 * r . height
}

//            

func main () {
r := rect { width : 10 , height : 5 }

//            <p>Here we call the 2 methods defined for our struct.</p>
//


fmt . Println ( "area: " , r . area ())
fmt . Println ( "perim:" , r . perim ())

//            <p>Go automatically handles conversion between values
//and pointers for method calls. You may want to use
//a pointer receiver type to avoid copying on method
//calls or to allow the method to mutate the
//receiving struct.</p>
//


rp := & r
fmt . Println ( "area: " , rp . area ())
fmt . Println ( "perim:" , rp . perim ())
}

