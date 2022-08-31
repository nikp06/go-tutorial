//            <p>Sometimes our Go programs need to spawn other, non-Go
//processes.</p>
//
//            

package main

//            

import (
"fmt"
"io"
"os/exec"
)

//            

func main () {

//            <p>We&rsquo;ll start with a simple command that takes no
//arguments or input and just prints something to
//stdout. The <code>exec.Command</code> helper creates an object
//to represent this external process.</p>
//


dateCmd := exec . Command ( "date" )

//            <p>The <code>Output</code> method runs the command, waits for it
//to finish and collects its standard output.
// If there were no errors, <code>dateOut</code> will hold bytes
//with the date info.</p>
//


dateOut , err := dateCmd . Output ()
if err != nil {
panic ( err )
}
fmt . Println ( "&gt; date" )
fmt . Println ( string ( dateOut ))

//            <p><code>Output</code> and other methods of <code>Command</code> will return
//<code>*exec.Error</code> if there was a problem executing the
//command (e.g. wrong path), and <code>*exec.ExitError</code>
//if the command ran but exited with a non-zero return
//code.</p>
//


_ , err = exec . Command ( "date" , "-x" ). Output ()
if err != nil {
switch e := err .( type ) {
case * exec . Error :
fmt . Println ( "failed executing:" , err )
case * exec . ExitError :
fmt . Println ( "command exit rc =" , e . ExitCode ())
default :
panic ( err )
}
}

//            <p>Next we&rsquo;ll look at a slightly more involved case
//where we pipe data to the external process on its
//<code>stdin</code> and collect the results from its <code>stdout</code>.</p>
//


grepCmd := exec . Command ( "grep" , "hello" )

//            <p>Here we explicitly grab input/output pipes, start
//the process, write some input to it, read the
//resulting output, and finally wait for the process
//to exit.</p>
//


grepIn , _ := grepCmd . StdinPipe ()
grepOut , _ := grepCmd . StdoutPipe ()
grepCmd . Start ()
grepIn . Write ([] byte ( "hello grep\ngoodbye grep" ))
grepIn . Close ()
grepBytes , _ := io . ReadAll ( grepOut )
grepCmd . Wait ()

//            <p>We omitted error checks in the above example, but
//you could use the usual <code>if err != nil</code> pattern for
//all of them. We also only collect the <code>StdoutPipe</code>
//results, but you could collect the <code>StderrPipe</code> in
//exactly the same way.</p>
//


fmt . Println ( "&gt; grep hello" )
fmt . Println ( string ( grepBytes ))

//            <p>Note that when spawning commands we need to
//provide an explicitly delineated command and
//argument array, vs. being able to just pass in one
//command-line string. If you want to spawn a full
//command with a string, you can use <code>bash</code>&rsquo;s <code>-c</code>
//option:</p>
//


lsCmd := exec . Command ( "bash" , "-c" , "ls -a -l -h" )
lsOut , err := lsCmd . Output ()
if err != nil {
panic ( err )
}
fmt . Println ( "&gt; ls -a -l -h" )
fmt . Println ( string ( lsOut ))
}

