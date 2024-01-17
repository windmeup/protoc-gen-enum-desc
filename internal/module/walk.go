package module

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/windmeup/protoc-gen-enum-desc/internal/templates/shared"
)

type visitor struct {
	pgs.Visitor
	pgs.DebuggerCommon
	ctx *shared.Context
}

func newVisitor(d pgs.DebuggerCommon) *visitor {
	v := &visitor{
		DebuggerCommon: d,
		ctx:            shared.NewContext(),
	}
	v.Visitor = pgs.PassThroughVisitor(v)
	return v
}

func (v *visitor) VisitEnum(e pgs.Enum) (pgs.Visitor, error) {
	if err := v.ctx.AddEnum(e); err != nil {
		return nil, err
	}
	v.Debugf("add enum: %v", e)
	return v, nil
}
