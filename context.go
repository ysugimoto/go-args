package args

import (
	"strconv"
)

type Context struct {
	commands []string
	options  map[string]interface{}
}

func NewContext(commands []string, options map[string]interface{}) *Context {
	return &Context{
		commands: commands,
		options:  options,
	}
}

func (c *Context) String(name string) (value string) {
	if v, ok := c.options[name]; ok {
		value = v.(string)
	}
	return
}

func (c *Context) Int(name string) (value int) {
	if v, ok := c.options[name]; ok {
		if vv, err := strconv.Atoi(v.(string)); err == nil {
			value = vv
		}
	}
	return
}

func (c *Context) Bool(name string) (value bool) {
	if v, ok := c.options[name]; ok {
		value = v.(bool)
	}
	return
}

func (c *Context) Has(name string) (has bool) {
	if _, ok := c.options[name]; ok {
		has = true
	}
	return
}

func (c *Context) At(index int) string {
	if len(c.commands) < index {
		return ""
	}
	return c.commands[index]
}

func (c *Context) Len() int {
	return len(c.commands)
}
