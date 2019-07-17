package fserv

import (
	"github.com/labstack/echo"
	"github.com/urfave/cli"
)

func Start(c *cli.Context) error {
	flags := NewFlags(c)

	e := echo.New()
	e.Static("/", flags.StaticRoot)

	return e.Start(flags.ListenAddr)
}
