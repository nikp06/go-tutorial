//            <p>Go supports <em>embedding</em> of structs and interfaces
//to express a more seamless <em>composition</em> of types.
//This is not to be confused with <code>//go:embed</code> which is
//a go directive introduced in Go version 1.16+ to embed
//files and folders into the application binary.</p>
//
//            

package main

//            

import "fmt"

//            

type base struct {
num int
}

//            

func ( b base ) describe () string {
return fmt . Sprintf ( "base with num=%v" , b . num )
}

//            <p>A <code>container</code> <em>embeds</em> a <code>base</code>. An embedding looks
//like a field without a name.</p>
//


type container struct {
base
str string
}

//            

func main () {

//            <p>When creating structs with literals, we have to
//initialize the embedding explicitly; here the
//embedded type serves as the field name.</p>
//


co := container {
base : base {
num : 1 ,
},
str : "some name" ,
}

//            <p>We can access the base&rsquo;s fields directly on <code>co</code>,
//e.g. <code>co.num</code>.</p>
//


fmt . Printf ( "co={num: %v, str: %v}\n" , co . num , co . str )

//            <p>Alternatively, we can spell out the full path using
//the embedded type name.</p>
//


fmt . Println ( "also num:" , co . base . num )

//            <p>Since <code>container</code> embeds <code>base</code>, the methods of
//<code>base</code> also become methods of a <code>container</code>. Here
//we invoke a method that was embedded from <code>base</code>
//directly on <code>co</code>.</p>
//


fmt . Println ( "describe:" , co . describe ())

//            

type describer interface {
describe () string
}

//            <p>Embedding structs with methods may be used to bestow
//interface implementations onto other structs. Here
//we see that a <code>container</code> now implements the
//<code>describer</code> interface because it embeds <code>base</code>.</p>
//


var d describer = co
fmt . Println ( "describer:" , d . describe ())
}

