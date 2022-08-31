//            <p>Some command-line tools, like the <code>go</code> tool or <code>git</code>
//have many <em>subcommands</em>, each with its own set of
//flags. For example, <code>go build</code> and <code>go get</code> are two
//different subcommands of the <code>go</code> tool.
//The <code>flag</code> package lets us easily define simple
//subcommands that have their own flags.</p>
//
//            

package main

//            

import (
"flag"
"fmt"
"os"
)

//            

func main () {

//            <p>We declare a subcommand using the <code>NewFlagSet</code>
//function, and proceed to define new flags specific
//for this subcommand.</p>
//


fooCmd := flag . NewFlagSet ( "foo" , flag . ExitOnError )
fooEnable := fooCmd . Bool ( "enable" , false , "enable" )
fooName := fooCmd . String ( "name" , "" , "name" )

//            <p>For a different subcommand we can define different
//supported flags.</p>
//


barCmd := flag . NewFlagSet ( "bar" , flag . ExitOnError )
barLevel := barCmd . Int ( "level" , 0 , "level" )

//            <p>The subcommand is expected as the first argument
//to the program.</p>
//


if len ( os . Args ) < 2 {
fmt . Println ( "expected 'foo' or 'bar' subcommands" )
os . Exit ( 1 )
}

//            <p>Check which subcommand is invoked.</p>
//


switch os . Args [ 1 ] {

//            <p>For every subcommand, we parse its own flags and
//have access to trailing positional arguments.</p>
//


case "foo" :
fooCmd . Parse ( os . Args [ 2 :])
fmt . Println ( "subcommand 'foo'" )
fmt . Println ( "  enable:" , * fooEnable )
fmt . Println ( "  name:" , * fooName )
fmt . Println ( "  tail:" , fooCmd . Args ())
case "bar" :
barCmd . Parse ( os . Args [ 2 :])
fmt . Println ( "subcommand 'bar'" )
fmt . Println ( "  level:" , * barLevel )
fmt . Println ( "  tail:" , barCmd . Args ())
default :
fmt . Println ( "expected 'foo' or 'bar' subcommands" )
os . Exit ( 1 )
}
}

