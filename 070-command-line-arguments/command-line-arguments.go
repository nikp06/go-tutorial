//            <p><a href="https://en.wikipedia.org/wiki/Command-line_interface#Arguments"><em>Command-line arguments</em></a>
//are a common way to parameterize execution of programs.
//For example, <code>go run hello.go</code> uses <code>run</code> and
//<code>hello.go</code> arguments to the <code>go</code> program.</p>
//
//            

package main

//            

import (
"fmt"
"os"
)

//            

func main () {

//            <p><code>os.Args</code> provides access to raw command-line
//arguments. Note that the first value in this slice
//is the path to the program, and <code>os.Args[1:]</code>
//holds the arguments to the program.</p>
//


argsWithProg := os . Args
argsWithoutProg := os . Args [ 1 :]

//            <p>You can get individual args with normal indexing.</p>
//


arg := os . Args [ 3 ]

//            

fmt . Println ( argsWithProg )
fmt . Println ( argsWithoutProg )
fmt . Println ( arg )
}

