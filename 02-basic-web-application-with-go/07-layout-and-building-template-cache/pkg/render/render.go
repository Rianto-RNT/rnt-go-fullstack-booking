package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap {

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// Create template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error ) {
	rntCache := map[string]*template.Template{}

	pages, err := filepath.Glob(".template/*.html")

	if err != nil {
		return rntCache, err
	}

	for _, page := range pages  {
		name := filepath.Base(page)
		
		// ts = template site
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return rntCache, err
		}
		
		matches, err := filepath.Glob("./template/*.html")
		if err != nil {
			return rntCache, err
		}
		
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./template/*.html")
			if err != nil {
				return rntCache, err
			}
		}

		rntCache[name] = ts 
	}

	return rntCache, nil
}
