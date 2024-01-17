package templates

import (
	"github.com/lyft/protoc-gen-star/v2"
	"github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/windmeup/protoc-gen-enum-desc/internal/templates/go"
	"text/template"
)

type (
	RegisterFunc func(tpl *template.Template, params pgs.Parameters)
	FilePathFunc func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath
)

func makeTemplate(ext string, f RegisterFunc, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	f(tpl, params)
	return tpl
}

func Templates(params pgs.Parameters) map[string][]*template.Template {
	return map[string][]*template.Template{
		"go": {makeTemplate("go", golang.Register, params)},
	}
}

func FilePathFor(tpl *template.Template) FilePathFunc {
	switch tpl.Name() {
	default: // golang
		return func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
			out := ctx.OutputPath(f)
			out = out.SetExt(".enumdesc." + tpl.Name())
			return &out
		}
	}
}
