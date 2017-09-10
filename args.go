package args

import (
	"strings"
)

type alias struct {
	name  string
	value interface{}
}

type Args struct {
	commands []string
	options  map[string]interface{}
	aliases  map[string]alias
}

func New() *Args {
	return &Args{
		options: make(map[string]interface{}),
		aliases: make(map[string]alias),
	}
}

func (a *Args) Alias(long, short string, value interface{}) *Args {
	a.aliases[short] = alias{
		name:  long,
		value: value,
	}
	a.options[long] = value
	return a
}

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
	}
}
