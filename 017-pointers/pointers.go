//            <p>Go supports <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointers</a></em>,
//allowing you to pass references to values and records
//within your program.</p>
//
//            

package main

//            

import "fmt"

//            <p>We&rsquo;ll show how pointers work in contrast to values with
//2 functions: <code>zeroval</code> and <code>zeroptr</code>. <code>zeroval</code> has an
//<code>int</code> parameter, so arguments will be passed to it by
//value. <code>zeroval</code> will get a copy of <code>ival</code> distinct
//from the one in the calling function.</p>
//


func zeroval ( ival int ) {
ival = 0
}

//            <p><code>zeroptr</code> in contrast has an <code>*int</code> parameter, meaning
//that it takes an <code>int</code> pointer. The <code>*iptr</code> code in the
//function body then <em>dereferences</em> the pointer from its
//memory address to the current value at that address.
//Assigning a value to a dereferenced pointer changes the
//value at the referenced address.</p>
//


func zeroptr ( iptr * int ) {
* iptr = 0
}

//            

func main () {
i := 1
fmt . Println ( "initial:" , i )

//            

zeroval ( i )
fmt . Println ( "zeroval:" , i )

//            <p>The <code>&amp;i</code> syntax gives the memory address of <code>i</code>,
//i.e. a pointer to <code>i</code>.</p>
//


zeroptr ( & i )
fmt . Println ( "zeroptr:" , i )

//            <p>Pointers can be printed too.</p>
//


fmt . Println ( "pointer:" , & i )
}

