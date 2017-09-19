package args

import (
	"strings"
)

// Private alias struct
type alias struct {
	// long name
	name string

	// actual value
	value interface{}
}

// Argument struct
type Args struct {
	// Parsed options
	options map[string]interface{}

	// Defined aliases
	aliases map[string]alias

	// Default values
	defaults map[string]interface{}
}

// Instantiate Args pointer
func New() *Args {
	return &Args{
		options:  make(map[string]interface{}),
		aliases:  make(map[string]alias),
		defaults: make(map[string]interface{}),
	}
}

// Define alias
func (a *Args) Alias(long, short string, value interface{}) *Args {
	// If short name is empty, accepts long options only
	if short != "" {
		a.aliases[short] = alias{
			name:  long,
			value: value,
		}
	}
	a.defaults[long] = value
	return a
}

// Parse args from supplied string slice
func (a *Args) Parse(args []string) *Context {
	commands := make([]string, 0)
	size := len(args)
	for i := 0; i < size; i++ {
		v := args[i]
		if v[0] != '-' {
			commands = append(commands, v)
			continue
		}
		if len(v) > 1 && v[1] == '-' {
			s := strings.Split(v, "=")
			if len(s) > 1 {
				a.options[s[0][2:]] = s[1]
			} else {
				a.options[s[0][2:]] = ""
			}
		} else if alias, ok := a.aliases[string(v[1])]; ok {
			if len(v) == 2 {
				if alias.value == nil {
					a.options[alias.name] = true
				} else if _, ok := alias.value.(bool); ok {
					a.options[alias.name] = true
				} else if i+1 < size {
					a.options[alias.name] = args[i+1]
					i++
				}
			} else {
				a.options[alias.name] = v[2:]
			}
		}
	}
	return &Context{
		commands: commands,
		options:  a.options,
		defaults: a.defaults,
	}
}
