package golang

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"text/template"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	ctx := pgsgo.InitContext(params)
	tpl.Funcs(map[string]any{
		"pkg": ctx.PackageName,
	})
	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("enum").Parse(enumTpl))
}
