package golang

const enumTpl = `
{{ if enumDescriptions . }}
var {{ enumIdentifier . }}_description = map[int32]string {
	{{ range $k, $v := enumDescriptions . -}}
		{{ $k }}: "{{ $v }}",
	{{ end -}}
}
{{ end }}

func (x {{ enumIdentifier . }}) Description() string {
	{{- if enumDescriptions . }}
		{{- if $dds := defaultDescription . }}
			if d := {{ enumIdentifier . }}_description[int32(protoreflect.EnumNumber(x))]; d != "" {
				return d
			}
			return "{{ $dds }}"
		{{- else }}
			return {{ enumIdentifier . }}_description[int32(protoreflect.EnumNumber(x))]
		{{- end }}
	{{- else }}
		return "{{ defaultDescription . }}"
	{{- end }}
}
`
