// Dynamic content for showing cutomized content output to the user
// with the text/template package

// A sibling package named html/template has same API with additional security- but should be used for 
// generating HTML
package main

import (
	"os"
	"text/template"
)

func main(){
	
	// We can creat a template and parse its body froma string. 
	// Templates are a mix of static text and actions enclosed in {{...}} - dynamic text insertions
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")

	if err != nil{
		panic(err)
	}


	// Alternatively we can use the template. Must function to panic in case Parse returns an error.
	// Useful for templates initialized in global scope
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))




	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})



	// Creating a Helper Function
	// creates and parses string
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}


	// If the data is a struct we can use the {{.FieldName}} action to access the field
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct{
		Name string
	}{"Jane Doe" })


	// same applies to Maps- with maps there is no restriction on the case of key names
	t2.Execute(os.Stdout, map[string]string{"Name": "Mickey Daey" })

	// if/else provides conditional execution of templates.
	// a value is considered false if its the defualt value of a types
	// eg. 0, an empty string, a nil pointer etc
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "hello")
	t3.Execute(os.Stdout, "")



	// Range blocks allow for looping through slices, arrays, maps or channels.

	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

}