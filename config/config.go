package config

import (
	"flag"
	"io/ioutil"

	"ginTemplate/utils"
	"github.com/urfave/cli"
)


var (
	config       Config
)

type Config interface {
	IsSet(name string) bool
	Bool(name string) bool
	Int(name string) int
	IntSlice(name string) []int
	Int64(name string) int64
	Int64Slice(name string) []int64
	String(name string) string
	StringSlice(name string) []string
	Uint(name string) uint
	Uint64(name string) uint64
	Set(name, value string) error
}

func IsSet(name string) bool           { return config.IsSet(name) }
func Bool(name string) bool            { return config.Bool(name) }
func Int(name string) int              { return config.Int(name) }
func IntSlice(name string) []int       { return config.IntSlice(name) }
func Int64(name string) int64          { return config.Int64(name) }
func Int64Slice(name string) []int64   { return config.Int64Slice(name) }
func String(name string) string        { return config.String(name) }
func StringSlice(name string) []string { return config.StringSlice(name) }
func Uint(name string) uint            { return config.Uint(name) }
func Uint64(name string) uint64        { return config.Uint64(name) }
func Set(name, value string) error     { return config.Set(name, value) }

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "host",
		Value: utils.GetInternalIPv4Address(),
		Usage: "service listen address",
	},
	cli.UintFlag{
		Name:  "port,p",
		Value: 8033,
		Usage: "service port",
	},
	cli.BoolFlag{
		Name:  "internet",
		Usage: "is run for internet",
	},
	cli.StringFlag{
		Name:  "mode,m",
		Value: "debug",
		Usage: "run mode",
	},
	cli.StringFlag{
		Name:  "config,c",
		Value: "",
		Usage: "configure file",
	},
	cli.BoolFlag{
		Name:  "weblogin",
		Usage: "use web login",
	},
}

// Initialize initialize process configure
func Initialize(c Config) {
	if c == nil {
		app := cli.NewApp()
		app.Name = "gin_template"
		app.Usage = "gin框架web server项目脚手架"
		app.Flags = Flags

		set := flag.NewFlagSet(app.Name, flag.ContinueOnError)

		for _, f := range Flags {
			f.Apply(set)
		}
		set.SetOutput(ioutil.Discard)

		c = cli.NewContext(app, set, nil)
	}
	config = c
	initializeDatabase(c)
}
