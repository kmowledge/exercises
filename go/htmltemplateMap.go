package main

import (
	"html/template"
	"os"
)

type a struct {
	Numbers []string
	Letters []string
}

var data = &a{
	Numbers: []string{"one", "two"},
	Letters: []string{"a", "b", "b", "c"},
}

var tmplSrc = `start
{{$i := 0}}
{{range $number := $.Numbers}}
  {{range $letter := $.Letters}}
    {{if eq $letter "b"}}
      {{$i = add $i 1}}
        {{$i}}
     {{end}}
  {{end}}
{{end}}
fin
`

func main() {
	funcMap := template.FuncMap{
		"add": func(a int, b int) int {
			return a + b
		},
	}
	tmpl := template.Must(template.New("test").Funcs(funcMap).Parse(tmplSrc))
	tmpl.Execute(os.Stdout, data)
}
