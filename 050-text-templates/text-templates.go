//            <p>Go offers built-in support for creating dynamic content or showing customized
//output to the user with the <code>text/template</code> package. A sibling package
//named <code>html/template</code> provides the same API but has additional security
//features and should be used for generating HTML.</p>
//
//            

package main

//            

import (
"os"
"text/template"
)

//            

func main () {

//            <p>We can create a new template and parse its body from
//a string.
//Templates are a mix of static text and &ldquo;actions&rdquo; enclosed in
//<code>{{...}}</code> that are used to dynamically insert content.</p>
//


t1 := template . New ( "t1" )
t1 , err := t1 . Parse ( "Value is {{.}}\n" )
if err != nil {
panic ( err )
}

//            <p>Alternatively, we can use the <code>template.Must</code> function to
//panic in case <code>Parse</code> returns an error. This is especially
//useful for templates initialized in the global scope.</p>
//


t1 = template . Must ( t1 . Parse ( "Value: {{.}}\n" ))

//            <p>By &ldquo;executing&rdquo; the template we generate its text with
//specific values for its actions. The <code>{{.}}</code> action is
//replaced by the value passed as a parameter to <code>Execute</code>.</p>
//


t1 . Execute ( os . Stdout , "some text" )
t1 . Execute ( os . Stdout , 5 )
t1 . Execute ( os . Stdout , [] string {
"Go" ,
"Rust" ,
"C++" ,
"C#" ,
})

//            <p>Helper function we&rsquo;ll use below.</p>
//


Create := func ( name , t string ) * template . Template {
return template . Must ( template . New ( name ). Parse ( t ))
}

//            <p>If the data is a struct we can use the <code>{{.FieldName}}</code> action to access
//its fields. The fields should be exported to be accessible when a
//template is executing.</p>
//


t2 := Create ( "t2" , "Name: {{.Name}}\n" )

//            

t2 . Execute ( os . Stdout , struct {
Name string
}{ "Jane Doe" })

//            <p>The same applies to maps; with maps there is no restriction on the
//case of key names.</p>
//


t2 . Execute ( os . Stdout , map [ string ] string {
"Name" : "Mickey Mouse" ,
})

//            <p>if/else provide conditional execution for templates. A value is considered
//false if it&rsquo;s the default value of a type, such as 0, an empty string,
//nil pointer, etc.
//This sample demonstrates another
//feature of templates: using <code>-</code> in actions to trim whitespace.</p>
//


t3 := Create ( "t3" ,
"{{if . -}} yes {{else -}} no {{end}}\n" )
t3 . Execute ( os . Stdout , "not empty" )
t3 . Execute ( os . Stdout , "" )

//            <p>range blocks let us loop through slices, arrays, maps or channels. Inside
//the range block <code>{{.}}</code> is set to the current item of the iteration.</p>
//


t4 := Create ( "t4" ,
"Range: {{range .}}{{.}} {{end}}\n" )
t4 . Execute ( os . Stdout ,
[] string {
"Go" ,
"Rust" ,
"C++" ,
"C#" ,
})
}

