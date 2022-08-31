//            <p>Go&rsquo;s <em>structs</em> are typed collections of fields.
//They&rsquo;re useful for grouping data together to form
//records.</p>
//
//            

package main

//            

import "fmt"

//            <p>This <code>person</code> struct type has <code>name</code> and <code>age</code> fields.</p>
//


type person struct {
name string
age int
}

//            <p><code>newPerson</code> constructs a new person struct with the given name.</p>
//


func newPerson ( name string ) * person {

//            <p>You can safely return a pointer to local variable
//as a local variable will survive the scope of the function.</p>
//


p := person { name : name }
p . age = 42
return & p
}

//            

func main () {

//            <p>This syntax creates a new struct.</p>
//


fmt . Println ( person { "Bob" , 20 })

//            <p>You can name the fields when initializing a struct.</p>
//


fmt . Println ( person { name : "Alice" , age : 30 })

//            <p>Omitted fields will be zero-valued.</p>
//


fmt . Println ( person { name : "Fred" })

//            <p>An <code>&amp;</code> prefix yields a pointer to the struct.</p>
//


fmt . Println ( & person { name : "Ann" , age : 40 })

//            <p>It&rsquo;s idiomatic to encapsulate new struct creation in constructor functions</p>
//


fmt . Println ( newPerson ( "Jon" ))

//            <p>Access struct fields with a dot.</p>
//


s := person { name : "Sean" , age : 50 }
fmt . Println ( s . name )

//            <p>You can also use dots with struct pointers - the
//pointers are automatically dereferenced.</p>
//


sp := & s
fmt . Println ( sp . age )

//            <p>Structs are mutable.</p>
//


sp . age = 51
fmt . Println ( sp . age )
}

