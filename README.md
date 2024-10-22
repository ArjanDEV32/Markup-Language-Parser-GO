# A General Markup Language parser in Golang

This module is able to parse any type of Markup language,
wether it's html or xml 
and is able to turn it from a string to a data structure 
or from a data structure to a string.

### This module consist of 3 Structs:

* prop struct:
```golang
type prop struct{Name string; Value string}
```
* tag struct:
```golang
type Tag struct{
	Name,InnerText string
	Props []prop
	Children []Tag
	Type byte 
} 
```
* Markup Language data struct:
```golang
type MULDS struct{Data []Tag}
```

### and 2 functions:
* parse function:
```golang
func Parse(src *string) []Tag
```
* stringify function:
```golang
func Stringify(data *MULDS) string
```

## Example:

```golang
type MULDS = mul.MULDS

func main() {
	var MarkUp string = `
<!DOCTYPE glossary PUBLIC "-//OASIS//DTD DocBook V3.1//EN">
<glossary><title>example glossary</title>
<GlossDiv><title>S</title>
 <GlossList>
	<GlossEntry ID="SGML" SortAs="SGML">
	 <GlossTerm>Standard Generalized Markup Language</GlossTerm>
	 <Acronym>SGML</Acronym>
	 <Abbrev>ISO 8879:1986</Abbrev>
	 <GlossDef>
		<para>A meta-markup language, used to create markup
languages such as DocBook.</para>
		<GlossSeeAlso OtherTerm="GML">
		<GlossSeeAlso OtherTerm="XML">
	 </GlossDef>
	 <GlossSee OtherTerm="markup">
	</GlossEntry>
 </GlossList>
</GlossDiv>
</glossary>
`
	
  var ds MULDS
  ds.Data = mulp.Parse(&MarkUp)
  var MarkUp2 string = mulp.Stringify(&ds.Data)

  fmt.Println(ds.Data)
  fmt.Println(MarkUp2)
}
```

