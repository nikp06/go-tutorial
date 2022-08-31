//            <p>Go offers built-in support for <a href="https://en.wikipedia.org/wiki/Regular_expression">regular expressions</a>.
//Here are some examples of  common regexp-related tasks
//in Go.</p>
//
//            

package main

//            

import (
"bytes"
"fmt"
"regexp"
)

//            

func main () {

//            <p>This tests whether a pattern matches a string.</p>
//


match , _ := regexp . MatchString ( "p([a-z]+)ch" , "peach" )
fmt . Println ( match )

//            <p>Above we used a string pattern directly, but for
//other regexp tasks you&rsquo;ll need to <code>Compile</code> an
//optimized <code>Regexp</code> struct.</p>
//


r , _ := regexp . Compile ( "p([a-z]+)ch" )

//            <p>Many methods are available on these structs. Here&rsquo;s
//a match test like we saw earlier.</p>
//


fmt . Println ( r . MatchString ( "peach" ))

//            <p>This finds the match for the regexp.</p>
//


fmt . Println ( r . FindString ( "peach punch" ))

//            <p>This also finds the first match but returns the
//start and end indexes for the match instead of the
//matching text.</p>
//


fmt . Println ( "idx:" , r . FindStringIndex ( "peach punch" ))

//            <p>The <code>Submatch</code> variants include information about
//both the whole-pattern matches and the submatches
//within those matches. For example this will return
//information for both <code>p([a-z]+)ch</code> and <code>([a-z]+)</code>.</p>
//


fmt . Println ( r . FindStringSubmatch ( "peach punch" ))

//            <p>Similarly this will return information about the
//indexes of matches and submatches.</p>
//


fmt . Println ( r . FindStringSubmatchIndex ( "peach punch" ))

//            <p>The <code>All</code> variants of these functions apply to all
//matches in the input, not just the first. For
//example to find all matches for a regexp.</p>
//


fmt . Println ( r . FindAllString ( "peach punch pinch" , - 1 ))

//            <p>These <code>All</code> variants are available for the other
//functions we saw above as well.</p>
//


fmt . Println ( "all:" , r . FindAllStringSubmatchIndex (
"peach punch pinch" , - 1 ))

//            <p>Providing a non-negative integer as the second
//argument to these functions will limit the number
//of matches.</p>
//


fmt . Println ( r . FindAllString ( "peach punch pinch" , 2 ))

//            <p>Our examples above had string arguments and used
//names like <code>MatchString</code>. We can also provide
//<code>[]byte</code> arguments and drop <code>String</code> from the
//function name.</p>
//


fmt . Println ( r . Match ([] byte ( "peach" )))

//            <p>When creating global variables with regular
//expressions you can use the <code>MustCompile</code> variation
//of <code>Compile</code>. <code>MustCompile</code> panics instead of
//returning an error, which makes it safer to use for
//global variables.</p>
//


r = regexp . MustCompile ( "p([a-z]+)ch" )
fmt . Println ( "regexp:" , r )

//            <p>The <code>regexp</code> package can also be used to replace
//subsets of strings with other values.</p>
//


fmt . Println ( r . ReplaceAllString ( "a peach" , "<fruit&gt;" ))

//            <p>The <code>Func</code> variant allows you to transform matched
//text with a given function.</p>
//


in := [] byte ( "a peach" )
out := r . ReplaceAllFunc ( in , bytes . ToUpper )
fmt . Println ( string ( out ))
}

