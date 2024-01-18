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
	defaultDescription := ""
	ddp := new(string)
	found, err := e.Extension(pged.E_DefaultDescription, ddp)
	if err != nil {
		return err
	}
	if found {
		defaultDescription = *ddp
	}
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
	if defaultDescription == "" && len(descriptions) == 0 {
		return nil
	}
	id := e.FullyQualifiedName()
	if pn := e.Package().ProtoName(); pn != "" {
		id = strings.TrimPrefix(id, "."+pn.String())
	}
	id = strings.ReplaceAll(strings.TrimPrefix(id, "."), ".", "_")
	c.enums[e] = &enum{
		identifier:         id,
		defaultDescription: defaultDescription,
		descriptions:       descriptions,
	}
	return nil
}

func (c *Context) EmptyDescription(f pgs.File) bool {
	for _, e := range f.Enums() {
		if c.enums[e] != nil {
			return false
		}
	}
	return true
}

// methods for template:

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

func (c *Context) defaultDescription(e pgs.Enum) string {
	if e, exists := c.enums[e]; exists {
		return e.defaultDescription
	}
	return ""
}
