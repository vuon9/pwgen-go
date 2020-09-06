package main

import (
	"flag"
	"fmt"
	"os"
)

type cmdOption func(c *baseCommand)
type baseCommand struct {
	cmds    cmdableSlice
	cmdsMap map[string]interface{}
	usage   func()
}

type cmdableSlice []cmdable

func NewItems(cmds ...cmdable) []cmdable {
	return cmds
}

func NewCommandController(cmds []cmdable, opts ...cmdOption) *baseCommand {
	c := &baseCommand{cmds: cmds, cmdsMap: make(map[string]interface{})}

	for i, opt := range cmds {
		switch targetOpt := opt.(type) {
		case *boolCmd:
			targetOpt.ptr = flag.Bool(targetOpt.name, targetOpt.value, targetOpt.description)
			if targetOpt.alias != "" {
				flag.BoolVar(targetOpt.ptr, targetOpt.alias, *targetOpt.ptr, "")
			}
		case *stringCmd:
			targetOpt.ptr = flag.String(targetOpt.name, targetOpt.value, targetOpt.description)
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

func (c *baseCommand) GetBool(n string) bool {
	return c.cmdsMap[n].(bool)
}

func (c *baseCommand) GetString(n string) string {
	return c.cmdsMap[n].(string)
}

func (c cmdableSlice) GetBool(n string) bool {
	return false
}

func (c cmdableSlice) GetString(n string) string {
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

type cmdable interface {
	Name() string
	Usage() string
}

type Option struct {
	name        string
	alias       string
	usage       string
	description string
}

func (c *Option) Name() string {
	return c.name
}

func (c *Option) Alias() string {
	return c.alias
}

func (c *Option) Usage() string {
	switch {
	case c.usage != "":
		return fmt.Sprintf("%s\n     %s", c.usage, c.description)
	case c.alias != "":
		return fmt.Sprintf("-%s or -%s\n    %s", c.alias, c.name, c.description)
	default:
		return fmt.Sprintf("-%s\n     %s", c.name, c.description)
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

func NewBoolCommand(opt Option) cmdable {
	return &boolCmd{
		Option: opt,
	}
}

type stringCmd struct {
	Option
	value string
	ptr   *string
}

func NewStringCommand(opt Option) cmdable {
	return &stringCmd{
		Option: opt,
	}
}
