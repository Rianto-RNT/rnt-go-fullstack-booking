package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap {

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	_, err :=  RenderTemplateBase(w)
	if err != nil {
		fmt.Println("Error getting template cache:", err)
	}
	
	parseTemplate, _ := template.ParseFiles("./template/" + tmpl)
	
	err = parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

func RenderTemplateBase(w http.ResponseWriter) (map[string]*template.Template, error ) {
	rntCache := map[string]*template.Template{}

	pages, err := filepath.Glob(".template/*.html")

	if err != nil {
		return rntCache, err
	}

	for _, page := range pages  {
		name := filepath.Base(page)

		fmt.Println("Page is currently", page)
		
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
