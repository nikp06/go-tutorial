//            <p>Starting with version 1.18, Go has added support for
//<em>generics</em>, also known as <em>type parameters</em>.</p>
//
//            

package main

//            

import "fmt"

//            <p>As an example of a generic function, <code>MapKeys</code> takes
//a map of any type and returns a slice of its keys.
//This function has two type parameters - <code>K</code> and <code>V</code>;
//<code>K</code> has the <code>comparable</code> <em>constraint</em>, meaning that
//we can compare values of this type with the <code>==</code> and
//<code>!=</code> operators. This is required for map keys in Go.
//<code>V</code> has the <code>any</code> constraint, meaning that it&rsquo;s not
//restricted in any way (<code>any</code> is an alias for <code>interface{}</code>).</p>
//


func MapKeys [ K comparable , V any ]( m map [ K ] V ) [] K {
r := make ([] K , 0 , len ( m ))
for k := range m {
r = append ( r , k )
}
return r
}

//            <p>As an example of a generic type, <code>List</code> is a
//singly-linked list with values of any type.</p>
//


type List [ T any ] struct {
head , tail * element [ T ]
}

//            

type element [ T any ] struct {
next * element [ T ]
val T
}

//            <p>We can define methods on generic types just like we
//do on regular types, but we have to keep the type
//parameters in place. The type is <code>List[T]</code>, not <code>List</code>.</p>
//


func ( lst * List [ T ]) Push ( v T ) {
if lst . tail == nil {
lst . head = & element [ T ]{ val : v }
lst . tail = lst . head
} else {
lst . tail . next = & element [ T ]{ val : v }
lst . tail = lst . tail . next
}
}

//            

func ( lst * List [ T ]) GetAll () [] T {
var elems [] T
for e := lst . head ; e != nil ; e = e . next {
elems = append ( elems , e . val )
}
return elems
}

//            

func main () {
var m = map [ int ] string { 1 : "2" , 2 : "4" , 4 : "8" }

//            <p>When invoking generic functions, we can often rely
//on <em>type inference</em>. Note that we don&rsquo;t have to
//specify the types for <code>K</code> and <code>V</code> when
//calling <code>MapKeys</code> - the compiler infers them
//automatically.</p>
//


fmt . Println ( "keys:" , MapKeys ( m ))

//            <p>&hellip; though we could also specify them explicitly.</p>
//


_ = MapKeys [ int , string ]( m )

//            

lst := List [ int ]{}
lst . Push ( 10 )
lst . Push ( 13 )
lst . Push ( 23 )
fmt . Println ( "list:" , lst . GetAll ())
}

