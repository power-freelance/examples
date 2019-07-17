package fserv

import (
	"github.com/labstack/echo"
	"github.com/urfave/cli"
)

func Start(c *cli.Context) error {

	// Parse command line flags
	flags := NewFlags(c)

	// Create static dir if not exists
	if err := checkDir(flags.StaticRoot, 0755); err != nil {
		return err
	}

	// Start echo http micro-framework
	e := echo.New()
	e.Debug = flags.Debug
	e.HideBanner = !flags.Debug
	e.Static("/", flags.StaticRoot)
	e.POST("/", Upload(flags))

	return e.Start(flags.ListenAddr)
}
