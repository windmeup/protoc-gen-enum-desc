package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/windmeup/protoc-gen-enum-desc/proto/pged"
	"strings"
)

type Context struct {
	enums map[pgs.Enum]*enum
}

func NewContext() *Context {
	return &Context{
		enums: make(map[pgs.Enum]*enum),
	}
}

func (c *Context) AddEnum(e pgs.Enum) error {
	descriptions := make(map[int32]string)
	for _, v := range e.Values() {
		dp := new(string)
		found, err := v.Extension(pged.E_Description, dp)
		if err != nil {
			return err
		}
		if !found {
			continue
		}
		if desc := *dp; desc != "" {
			descriptions[v.Value()] = desc
		}
	}
	if len(descriptions) == 0 {
		return nil
	}
	id := e.FullyQualifiedName()
	if pn := e.Package().ProtoName(); pn != "" {
		id = strings.TrimPrefix(id, "."+pn.String())
	}
	id = strings.ReplaceAll(strings.TrimPrefix(id, "."), ".", "_")
	c.enums[e] = &enum{
		identifier:   id,
		descriptions: descriptions,
	}
	return nil
}

func (c *Context) enumIdentifier(e pgs.Enum) string {
	if e, exists := c.enums[e]; exists {
		return e.identifier
	}
	return ""
}

func (c *Context) enumDescriptions(e pgs.Enum) map[int32]string {
	if e, exists := c.enums[e]; exists {
		return e.descriptions
	}
	return nil
}
