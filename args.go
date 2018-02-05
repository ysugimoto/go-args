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
		fp.WriteString(v + "\n")
		if v[0] != '-' {
			commands = append(commands, v)
			continue
		}
		if len(v) > 1 && v[1] == '-' {
			s := strings.Split(v, "=")
			name := s[0][2:]
			d, ok := a.defaults[name]
			if !ok {
				continue
			}
			if len(s) > 1 {
				a.options[name] = s[1]
				continue
			}
			if d == nil {
				a.options[name] = ""
				continue
			}
			if i+1 < size {
				a.options[name] = args[i+1]
				i++
			} else {
				a.options[name] = ""
			}
			continue
		}
		alias, ok := a.aliases[string(v[1])]
		if !ok {
			continue
		}
		if len(v) == 2 {
			if alias.value == nil {
				a.options[alias.name] = true
			} else if _, ok := alias.value.(bool); ok {
				a.options[alias.name] = true
			} else if i+1 < size {
				a.options[alias.name] = args[i+1]
				i++
			}
			continue
		}
		a.options[alias.name] = v[2:]
	}
	return &Context{
		commands: commands,
		options:  a.options,
		defaults: a.defaults,
	}
}
