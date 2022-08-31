//            <p><a href="https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option"><em>Command-line flags</em></a>
//are a common way to specify options for command-line
//programs. For example, in <code>wc -l</code> the <code>-l</code> is a
//command-line flag.</p>
//
//            

package main

//            <p>Go provides a <code>flag</code> package supporting basic
//command-line flag parsing. We&rsquo;ll use this package to
//implement our example command-line program.</p>
//


import (
"flag"
"fmt"
)

//            

func main () {

//            <p>Basic flag declarations are available for string,
//integer, and boolean options. Here we declare a
//string flag <code>word</code> with a default value <code>&quot;foo&quot;</code>
//and a short description. This <code>flag.String</code> function
//returns a string pointer (not a string value);
//we&rsquo;ll see how to use this pointer below.</p>
//


wordPtr := flag . String ( "word" , "foo" , "a string" )

//            <p>This declares <code>numb</code> and <code>fork</code> flags, using a
//similar approach to the <code>word</code> flag.</p>
//


numbPtr := flag . Int ( "numb" , 42 , "an int" )
forkPtr := flag . Bool ( "fork" , false , "a bool" )

//            <p>It&rsquo;s also possible to declare an option that uses an
//existing var declared elsewhere in the program.
//Note that we need to pass in a pointer to the flag
//declaration function.</p>
//


var svar string
flag . StringVar ( & svar , "svar" , "bar" , "a string var" )

//            <p>Once all flags are declared, call <code>flag.Parse()</code>
//to execute the command-line parsing.</p>
//


flag . Parse ()

//            <p>Here we&rsquo;ll just dump out the parsed options and
//any trailing positional arguments. Note that we
//need to dereference the pointers with e.g. <code>*wordPtr</code>
//to get the actual option values.</p>
//


fmt . Println ( "word:" , * wordPtr )
fmt . Println ( "numb:" , * numbPtr )
fmt . Println ( "fork:" , * forkPtr )
fmt . Println ( "svar:" , svar )
fmt . Println ( "tail:" , flag . Args ())
}

