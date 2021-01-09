package helpers

import (
	"bytes"
	"fmt"
	"neighborhood/config"
	"text/template"
)

var tmpl = `
{{.Name}}
--------------------------------------------
A very simple portknocker!!

USAGE
{{.Command}} <options> <host> <port:protocol> <port:protocol>

OPTIONS
{{range .Options}}
-{{.Flag}} {{.Description}}
{{end}}

EXAMPLE
{{.Command}} 192.168.10.10 3000:TCP 4000:UDP 5000
`

// Man : show the manual
func Man() {
	t, err := template.New("help").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var tBuffer bytes.Buffer
	t.Execute(&tBuffer, config.Config)

	fmt.Println(string(tBuffer.Bytes()))

}
