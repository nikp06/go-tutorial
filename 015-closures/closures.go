//            <p>Go supports <a href="https://en.wikipedia.org/wiki/Anonymous_function"><em>anonymous functions</em></a>,
//which can form <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
//Anonymous functions are useful when you want to define
//a function inline without having to name it.</p>
//
//            

package main

//            

import "fmt"

//            <p>This function <code>intSeq</code> returns another function, which
//we define anonymously in the body of <code>intSeq</code>. The
//returned function <em>closes over</em> the variable <code>i</code> to
//form a closure.</p>
//


func intSeq () func () int {
i := 0
return func () int {
i ++
return i
}
}

//            

func main () {

//            <p>We call <code>intSeq</code>, assigning the result (a function)
//to <code>nextInt</code>. This function value captures its
//own <code>i</code> value, which will be updated each time
//we call <code>nextInt</code>.</p>
//


nextInt := intSeq ()

//            <p>See the effect of the closure by calling <code>nextInt</code>
//a few times.</p>
//


fmt . Println ( nextInt ())
fmt . Println ( nextInt ())
fmt . Println ( nextInt ())

//            <p>To confirm that the state is unique to that
//particular function, create and test a new one.</p>
//


newInts := intSeq ()
fmt . Println ( newInts ())
}

