package views

import (
	"fmt"
	"html/template"
)

// View struct for new layout files to be auto implemented in templates.
type View struct {
	Template *template.Template
	Layout   string
}

// NewView func to append layout files to parsed template files.
func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/navbar.gohtml",
		"views/layouts/materialize.gohtml",
		"views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Printf("Error parsing template files: %v", err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}
