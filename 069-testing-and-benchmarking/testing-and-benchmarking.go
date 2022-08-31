//            <p>Unit testing is an important part of writing
//principled Go programs. The <code>testing</code> package
//provides the tools we need to write unit tests
//and the <code>go test</code> command runs tests.</p>
//
//            <p>For the sake of demonstration, this code is in package
//<code>main</code>, but it could be any package. Testing code
//typically lives in the same package as the code it tests.</p>
//


package main

//            

import (
"fmt"
"testing"
)

//            <p>We&rsquo;ll be testing this simple implementation of an
//integer minimum. Typically, the code we&rsquo;re testing
//would be in a source file named something like
//<code>intutils.go</code>, and the test file for it would then
//be named <code>intutils_test.go</code>.</p>
//


func IntMin ( a , b int ) int {
if a < b {
return a
}
return b
}

//            <p>A test is created by writing a function with a name
//beginning with <code>Test</code>.</p>
//


func TestIntMinBasic ( t * testing . T ) {
ans := IntMin ( 2 , - 2 )
if ans != - 2 {

//            <p><code>t.Error*</code> will report test failures but continue
//executing the test. <code>t.Fatal*</code> will report test
//failures and stop the test immediately.</p>
//


t . Errorf ( "IntMin(2, -2) = %d; want -2" , ans )
}
}

//            <p>Writing tests can be repetitive, so it&rsquo;s idiomatic to
//use a <em>table-driven style</em>, where test inputs and
//expected outputs are listed in a table and a single loop
//walks over them and performs the test logic.</p>
//


func TestIntMinTableDriven ( t * testing . T ) {
var tests = [] struct {
a , b int
want int
}{
{ 0 , 1 , 0 },
{ 1 , 0 , 0 },
{ 2 , - 2 , - 2 },
{ 0 , - 1 , - 1 },
{ - 1 , 0 , - 1 },
}

//            <p>t.Run enables running &ldquo;subtests&rdquo;, one for each
//table entry. These are shown separately
//when executing <code>go test -v</code>.</p>
//

for _ , tt := range tests {

//            

testname := fmt . Sprintf ( "%d,%d" , tt . a , tt . b )
t . Run ( testname , func ( t * testing . T ) {
ans := IntMin ( tt . a , tt . b )
if ans != tt . want {
t . Errorf ( "got %d, want %d" , ans , tt . want )
}
})
}
}

//            <p>Benchmark tests typically go in <code>_test.go</code> files and are
//named beginning with <code>Benchmark</code>. The <code>testing</code> runner
//executes each benchmark function several times, increasing
//<code>b.N</code> on each run until it collects a precise measurement.</p>
//


func BenchmarkIntMin ( b * testing . B ) {

//            <p>Typically the benchmark runs a function we&rsquo;re
//benchmarking in a loop <code>b.N</code> times.</p>
//


for i := 0 ; i < b . N ; i ++ {
IntMin ( 1 , 2 )
}
}

