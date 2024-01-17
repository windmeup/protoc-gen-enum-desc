package shared

type enum struct {
	identifier   string // enum 的名字可能一样,加上它的 parent 用来区分
	descriptions map[int32]string
}
