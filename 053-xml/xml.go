//            <p>Go offers built-in support for XML and XML-like
//formats with the <code>encoding.xml</code> package.</p>
//
//            

package main

//            

import (
"encoding/xml"
"fmt"
)

//            <p>Plant will be mapped to XML. Similarly to the
//JSON examples, field tags contain directives for the
//encoder and decoder. Here we use some special features
//of the XML package: the <code>XMLName</code> field name dictates
//the name of the XML element representing this struct;
//<code>id,attr</code> means that the <code>Id</code> field is an XML
//<em>attribute</em> rather than a nested element.</p>
//


type Plant struct {
XMLName xml . Name `xml:"plant"`
Id int `xml:"id,attr"`
Name string `xml:"name"`
Origin [] string `xml:"origin"`
}

//            

func ( p Plant ) String () string {
return fmt . Sprintf ( "Plant id=%v, name=%v, origin=%v" ,
p . Id , p . Name , p . Origin )
}

//            

func main () {
coffee := & Plant { Id : 27 , Name : "Coffee" }
coffee . Origin = [] string { "Ethiopia" , "Brazil" }

//            <p>Emit XML representing our plant; using
//<code>MarshalIndent</code> to produce a more
//human-readable output.</p>
//


out , _ := xml . MarshalIndent ( coffee , " " , "  " )
fmt . Println ( string ( out ))

//            <p>To add a generic XML header to the output, append
//it explicitly.</p>
//


fmt . Println ( xml . Header + string ( out ))

//            <p>Use <code>Unmarshal</code> to parse a stream of bytes with XML
//into a data structure. If the XML is malformed or
//cannot be mapped onto Plant, a descriptive error
//will be returned.</p>
//


var p Plant
if err := xml . Unmarshal ( out , & p ); err != nil {
panic ( err )
}
fmt . Println ( p )

//            

tomato := & Plant { Id : 81 , Name : "Tomato" }
tomato . Origin = [] string { "Mexico" , "California" }

//            <p>The <code>parent&gt;child&gt;plant</code> field tag tells the encoder
//to nest all <code>plant</code>s under <code>&lt;parent&gt;&lt;child&gt;...</code></p>
//


type Nesting struct {
XMLName xml . Name `xml:"nesting"`
Plants [] * Plant `xml:"parent&gt;child&gt;plant"`
}

//            

nesting := & Nesting {}
nesting . Plants = [] * Plant { coffee , tomato }

//            

out , _ = xml . MarshalIndent ( nesting , " " , "  " )
fmt . Println ( string ( out ))
}

