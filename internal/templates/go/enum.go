package golang

const enumTpl = `
var {{ enumIdentifier . }}_description = map[int32]string {
	{{ range $k, $v := enumDescriptions . -}}
		{{ $k }}: "{{ $v }}",
	{{ end -}}
}

func (x {{ enumIdentifier . }}) Description() string {
	return {{ enumIdentifier . }}_description[int32(protoreflect.EnumNumber(x))]
}
`
