package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type (
	cmdOption   func(c *baseCommand)
	baseCommand struct {
		isDebug bool
		cmdsMap map[cmdName]cmdable
		usage   func()
	}
)

func NewItems(cmds ...cmdable) []cmdable {
	return cmds
}

func NewCommandController(cmds []cmdable, opts ...cmdOption) *baseCommand {
	c := &baseCommand{
		cmdsMap: make(map[cmdName]cmdable),
	}

	for _, opt := range cmds {
		var name cmdName
		switch targetOpt := opt.(type) {
		case *boolCmd:
			targetOpt.ptr = flag.Bool(string(targetOpt.name), targetOpt.def, targetOpt.description)
			if targetOpt.alias != "" {
				flag.BoolVar(targetOpt.ptr, targetOpt.alias, targetOpt.def, "")
			}
			name = targetOpt.name
		case *stringCmd:
			targetOpt.ptr = flag.String(string(targetOpt.name), targetOpt.def, targetOpt.description)
			if targetOpt.alias != "" {
				flag.StringVar(targetOpt.ptr, targetOpt.alias, targetOpt.def, "")
			}
			name = targetOpt.name
		}
		c.cmdsMap[name] = opt
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
			c.PrintDefaults()
			os.Exit(0)
		}
	}
}

func (c *baseCommand) Usage() {
	c.usage()
}

func (c *baseCommand) Ready() {
	flag.Parse()

	for _, opt := range c.cmdsMap {
		switch targetOpt := opt.(type) {
		case *boolCmd:
			*c.cmdsMap[targetOpt.name].(*boolCmd).ptr = *targetOpt.ptr
		case *stringCmd:
			*c.cmdsMap[targetOpt.name].(*stringCmd).ptr = *targetOpt.ptr
		}
	}
}

func (c *baseCommand) GetBool(n cmdName) bool {
	v := *c.cmdsMap[n].(*boolCmd).ptr
	if c.isDebug {
		log.Printf("GetBool() name = '%s', value = %+v", n, v)
	}
	return v
}

func (c *baseCommand) GetString(n cmdName) string {
	v := *c.cmdsMap[n].(*stringCmd).ptr
	if c.isDebug {
		log.Printf("GetString() name = '%s', value = \"%+v\"", n, v)
	}
	return v
}

func (c *baseCommand) PrintDefaults() {
	keys := make([]string, 0, len(c.cmdsMap))
	for cn := range c.cmdsMap {
		keys = append(keys, string(cn))
	}
	sort.Strings(keys)

	for _, v := range keys {
		opt := c.cmdsMap[cmdName(v)]

		// Read default value from option
		var defStr string
		defaultStringReader, ok := opt.(defaultStringReader)
		if ok {
			defStr = defaultStringReader.DefaultString()
		}

		fmt.Printf(" %s\n", opt.Usage(defStr))
	}
}

type defaultStringReader interface {
	DefaultString() string
}

type cmdName string

type cmdable interface {
	Usage(defStr string) string
}

type Option struct {
	name        cmdName
	alias       string
	description string
}

func (c *Option) Name() cmdName {
	return c.name
}

func (c *Option) Usage(defStr string) string {
	if defStr == "" {
		defStr = "<empty>"
	}

	name := fmt.Sprintf("-%s", string(c.name))

	alias := ""
	if c.alias != "" {
		alias = fmt.Sprintf("-%s", c.alias)
	}

	return fmt.Sprintf(
		"%s (default: %s)\n    %s",
		strings.Join([]string{name, alias}, ", "),
		defStr,
		c.description,
	)
}

type boolCmd struct {
	Option
	def bool
	ptr *bool
}

func (c *boolCmd) DefaultString() string {
	if !c.def {
		return "false"
	}

	return "true"
}

func NewBoolCommand(name cmdName, alias string, def bool, desc string) cmdable {
	return &boolCmd{
		Option: Option{
			name:        name,
			alias:       alias,
			description: desc,
		},
		def: def,
		ptr: new(bool),
	}
}

type stringCmd struct {
	Option
	def string
	ptr *string
}

func (c *stringCmd) DefaultString() string {
	return c.def
}

func NewStringCommand(name cmdName, alias string, def string, desc string) cmdable {
	return &stringCmd{
		Option: Option{
			name:        name,
			alias:       alias,
			description: desc,
		},
		def: def,
		ptr: new(string),
	}
}
