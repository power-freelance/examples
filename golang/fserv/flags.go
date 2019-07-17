package fserv

import "github.com/urfave/cli"

var FlagDefinitions = []cli.Flag{
	cli.BoolFlag{
		Name:   "debug",
		Usage:  "Debug mode",
		EnvVar: "DEBUG",
	},
	cli.StringFlag{
		Name:   "listenAddr",
		Usage:  "The address with the port on which the server will be launched.",
		Value:  ":9000",
		EnvVar: "LISTEN_ADDR",
	},
	cli.StringFlag{
		Name:   "staticRoot",
		Usage:  "Directory in which files will be saved.",
		Value:  "static",
		EnvVar: "STATIC_ROOT",
	},
}

func NewFlags(c *cli.Context) *Flags {
	return &Flags{
		Debug:      c.Bool("debug"),
		ListenAddr: c.String("listenAddr"),
		StaticRoot: c.String("staticRoot"),
	}
}

type Flags struct {
	Debug      bool
	ListenAddr string
	StaticRoot string
}
