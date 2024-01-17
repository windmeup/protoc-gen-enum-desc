package shared

type enum struct {
	identifier string // enum 的名字可能一样,加上它的 parent 用来区分
}

func (e *enum) Identifier() string {
	return e.identifier
}
