package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/jamieomatthews/validation"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	//martini setup
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "public/templates",         // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		/*Funcs:      []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.*/
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
	}))

	m.Post("/test", binding.Bind(ContactRequest{}), func(req *http.Request, contactReq ContactRequest) string {
		return contactReq.FullName
	})

	m.Run()
}

// This method implements binding.Validator and is executed by the binding.Validate middleware
func (contactRequest ContactRequest) Validate(errors binding.Errors, req *http.Request) binding.Errors {
	log.Println("Validating User")
	// v := validation.Validation{Errors: *errors}

	v := validation.New(errors, req)

	// //run some validators
	v.Validate(contactRequest.FullName, "full_name").MaxLength(20)
	v.Validate(contactRequest.Email, "email").Default("Custom Email Validation Message").Classify("email-class").Email()
	v.Validate(contactRequest.Comments, "comments").TrimSpace().MinLength(10)

	return v.Errors
}

type ContactRequest struct {
	FullName string `form:"full_name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Subject  string `form:"subject"`
	Comments string `form:"comments"`
}

// func (v *Validation) MinLength(n int, minLength int, fieldName string) bool {
// 	return v.validate(MinLength{MinLength: minLength}, n, fieldName)
// }
