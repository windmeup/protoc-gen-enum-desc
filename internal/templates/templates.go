package templates

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	golang "github.com/windmeup/protoc-gen-enum-desc/internal/templates/go"
	"github.com/windmeup/protoc-gen-enum-desc/internal/templates/shared"
	"text/template"
)

type (
	RegisterFunc func(tpl *template.Template, params pgs.Parameters)
	FilePathFunc func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath
)

func makeTemplate(ctx *shared.Context, ext string, f RegisterFunc, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	shared.Register(tpl, ctx)
	f(tpl, params)
	return tpl
}

func Templates(ctx *shared.Context, params pgs.Parameters) map[string][]*template.Template {
	return map[string][]*template.Template{
		"go": {makeTemplate(ctx, "go", golang.Register, params)},
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
