package main

import (
	"flag"
	"fmt"
	"os"
)

type (
	cmdOption   func(c *baseCommand)
	baseCommand struct {
		cmds    cmdableSlice
		cmdsMap map[cmdName]interface{}
		usage   func()
	}
)

type cmdableSlice []cmdable

func NewItems(cmds ...cmdable) []cmdable {
	return cmds
}

func NewCommandController(cmds []cmdable, opts ...cmdOption) *baseCommand {
	c := &baseCommand{cmds: cmds, cmdsMap: make(map[cmdName]interface{})}

	for i, opt := range cmds {
		switch targetOpt := opt.(type) {
		case *boolCmd:
			targetOpt.ptr = flag.Bool(string(targetOpt.name), targetOpt.value, targetOpt.description)
			if targetOpt.alias != "" {
				flag.BoolVar(targetOpt.ptr, targetOpt.alias, *targetOpt.ptr, "")
			}
		case *stringCmd:
			targetOpt.ptr = flag.String(string(targetOpt.name), targetOpt.value, targetOpt.description)
			if targetOpt.alias != "" {
				flag.StringVar(targetOpt.ptr, targetOpt.alias, *targetOpt.ptr, "")
			}
		}
		c.cmds[i] = opt
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithUsageHeader(header string) cmdOption {
	return func(c *baseCommand) {
		c.usage = func() {
			fmt.Fprintln(os.Stderr, header)
			c.cmds.PrintDefaults()
			os.Exit(0)
		}
	}
}

func (c *baseCommand) Usage() {
	c.usage()
}

func (c *baseCommand) Ready() {
	flag.Parse()

	for _, opt := range c.cmds {
		switch targetOpt := opt.(type) {
		case *boolCmd:
			c.cmdsMap[opt.Name()] = *targetOpt.ptr
		case *stringCmd:
			c.cmdsMap[opt.Name()] = *targetOpt.ptr
		}
	}
}

func (c *baseCommand) GetBool(n cmdName) bool {
	return c.cmdsMap[n].(bool)
}

func (c *baseCommand) GetString(n cmdName) string {
	return c.cmdsMap[n].(string)
}

func (c cmdableSlice) GetBool(n cmdName) bool {
	return false
}

func (c cmdableSlice) GetString(n cmdName) string {
	return ""
}

func (s cmdableSlice) PrintDefaults() cmdableSlice {
	for _, opt := range s {
		fmt.Printf(" %s\n", opt.Usage())
	}

	return s
}

func (s cmdableSlice) AndExit(code int) {
	os.Exit(code)
}

type cmdName string

type cmdable interface {
	Name() cmdName
	Usage() string
}

type Option struct {
	name        cmdName
	alias       string
	def         string
	description string
}

func (c *Option) Name() cmdName {
	return c.name
}

func (c *Option) Alias() string {
	return c.alias
}

func (c *Option) Usage() string {
	def := c.def
	if def == "" {
		def = "<empty>"
	}
	switch {
	case c.alias != "":
		return fmt.Sprintf("-%s or -%s\n    %s (default: %s)", c.alias, c.name, c.description, def)
	default:
		return fmt.Sprintf("-%s\n     %s (default: %s)", c.name, c.description, def)
	}
}

func (c *Option) Description() string {
	return c.description
}

type boolCmd struct {
	Option
	value bool
	ptr   *bool
}

func NewBoolCommand(name cmdName, alias, usage, desc string) cmdable {
	return &boolCmd{
		Option: Option{name, alias, usage, desc},
	}
}

type stringCmd struct {
	Option
	value string
	ptr   *string
}

func NewStringCommand(name cmdName, alias, usage, desc string) cmdable {
	return &stringCmd{
		Option: Option{name, alias, usage, desc},
	}
}
