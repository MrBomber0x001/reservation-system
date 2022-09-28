package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{} // a map of function, I can use in templates
// this allows us to create our own functions to work with templates, because the templating languages may not give us specific functionalites

//* RenderTemplaet renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get the template cache from app config

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err) // don't go any further
	}

	t, exists := tc[tmpl]
	if !exists {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer) // hold the information [parsedTemplate] into bytes

	_ = t.Execute(buf, nil) // take the bytes which contain the template and execute [parse] this template

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(`Error writing template to browser`, err)
	}
	parsedTemplte, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplte.Execute(w, nil)
	if err != nil {
		fmt.Print("Error parsing template", err)
		return
	}
}

//* CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{} //empty map of templates

	pages, err := filepath.Glob("./templates/*.template.htm") // returning the names of file matching a pattern or error, and only error is ErrBadPattern(), it ignores the I/O errors such as reading directories
	log.Println(pages)
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page) // return the last element of a path /a/b/c {c}
		log.Println("Page is currently", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page) // allocate new HTML template
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.htm")
		log.Println("Matches", matches)
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.htm")
			log.Println("ts", ts)
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
		log.Println("Cache", myCache)
	}
	return myCache, nil
}

// CreateTemplateCache is not effient because every single time we visit a route, we build the cache by traversing across
// all templates, we can avoid this by doing a wide app configurations to setup caching for once the server is up and running and set the cache inside app config and check this configuration here
