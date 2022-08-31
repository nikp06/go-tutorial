//            <p>In the previous example we looked at
//<a href="spawning-processes">spawning external processes</a>. We
//do this when we need an external process accessible to
//a running Go process. Sometimes we just want to
//completely replace the current Go process with another
//(perhaps non-Go) one. To do this we&rsquo;ll use Go&rsquo;s
//implementation of the classic
//<a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
//function.</p>
//
//            

package main

//            

import (
"os"
"os/exec"
"syscall"
)

//            

func main () {

//            <p>For our example we&rsquo;ll exec <code>ls</code>. Go requires an
//absolute path to the binary we want to execute, so
//we&rsquo;ll use <code>exec.LookPath</code> to find it (probably
//<code>/bin/ls</code>).</p>
//


binary , lookErr := exec . LookPath ( "ls" )
if lookErr != nil {
panic ( lookErr )
}

//            <p><code>Exec</code> requires arguments in slice form (as
//opposed to one big string). We&rsquo;ll give <code>ls</code> a few
//common arguments. Note that the first argument should
//be the program name.</p>
//


args := [] string { "ls" , "-a" , "-l" , "-h" }

//            <p><code>Exec</code> also needs a set of <a href="environment-variables">environment variables</a>
//to use. Here we just provide our current
//environment.</p>
//


env := os . Environ ()

//            <p>Here&rsquo;s the actual <code>syscall.Exec</code> call. If this call is
//successful, the execution of our process will end
//here and be replaced by the <code>/bin/ls -a -l -h</code>
//process. If there is an error we&rsquo;ll get a return
//value.</p>
//


execErr := syscall . Exec ( binary , args , env )
if execErr != nil {
panic ( execErr )
}
}

