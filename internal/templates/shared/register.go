package shared

import "text/template"

func Register(tpl *template.Template, ctx *Context) {
	tpl.Funcs(map[string]any{
		"enumIdentifier":   ctx.enumIdentifier,
		"enumDescriptions": ctx.enumDescriptions,
	})
}
