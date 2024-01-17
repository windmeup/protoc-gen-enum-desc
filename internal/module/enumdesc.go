package module

import (
	"github.com/lyft/protoc-gen-star/v2"
	"github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/windmeup/protoc-gen-enum-desc/internal/templates"
	"strings"
)

const (
	moduleName    = "enumdesc"
	langParamName = "lang"
)

type Module struct {
	*pgs.ModuleBase
	ctx  pgsgo.Context
	lang string
}

func Enumdesc() pgs.Module {
	return &Module{
		ModuleBase: new(pgs.ModuleBase),
	}
}

func EnumdescForLanguage(lang string) pgs.Module {
	return &Module{
		ModuleBase: new(pgs.ModuleBase),
		lang:       lang,
	}
}

func (m *Module) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *Module) Name() string {
	return moduleName
}

func (m *Module) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
	lang, langParam := m.lang, m.Parameters().Str(langParamName)
	if lang == "" {
		lang = langParam
		m.Assert(lang != "", "`lang` parameter must be set")
	} else if langParam != "" {
		m.Fail("unknown `lang` parameter")
	}
	tpls := templates.Templates(m.Parameters())[lang]
	m.Assert(len(tpls) != 0, "could not find templates for `lang`: ", lang)
	for _, f := range targets {
		m.Push(f.Name().String())
		for _, tpl := range tpls {
			out := templates.FilePathFor(tpl)(f, m.ctx, tpl)
			if out == nil {
				continue
			}
			outPath := strings.TrimLeft(out.String(), "/")
			m.AddGeneratorTemplateFile(outPath, tpl, f)
		}
		m.Pop()
	}
	return m.Artifacts()
}
