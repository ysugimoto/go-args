package args

import (
	"strconv"
)

// Parsed context struct
type Context struct {
	// Sub command list
	commands []string

	// Parsed options from alias
	options map[string]interface{}
}

// Instantiate context pointer
func NewContext(commands []string, options map[string]interface{}) *Context {
	return &Context{
		commands: commands,
		options:  options,
	}
}

// Get option as string
func (c *Context) String(name string) (value string) {
	if v, ok := c.options[name]; ok {
		value = v.(string)
	}
	return
}

// Get option as int
func (c *Context) Int(name string) (value int) {
	if v, ok := c.options[name]; ok {
		if vv, err := strconv.Atoi(v.(string)); err == nil {
			value = vv
		}
	}
	return
}

// Get option as bool
func (c *Context) Bool(name string) (value bool) {
	if v, ok := c.options[name]; ok {
		value = v.(bool)
	}
	return
}

// Check option exintence
func (c *Context) Has(name string) (has bool) {
	if _, ok := c.options[name]; ok {
		has = true
	}
	return
}

// Alias for StringAt
func (c *Context) At(index int) string {
	if len(c.commands) < index {
		return ""
	}
	return c.commands[index]
}

// Get subcommand at index as string
func (c *Context) StringAt(index int) string {
	if len(c.commands) < index {
		return ""
	}
	return c.commands[index]
}

// Get subcommand at index as inttring
func (c *Context) IntAt(index int) int {
	if len(c.commands) < index {
		return 0
	}
	if i, err := strconv.Atoi(c.commands[index]); err == nil {
		return i
	}
	return 0
}

// Get subcommand size
func (c *Context) Len() int {
	return len(c.commands)
}
