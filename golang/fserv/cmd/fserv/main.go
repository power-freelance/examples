package main

import (
	"log"
	"os"

	"github.com/power-freelance/examples/golang/fserv"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "FServ"
	app.Usage = "Simple static server for uploading and serving files."
	app.Version = "0.0.1"
	app.Action = fserv.Start
	app.Flags = fserv.FlagDefinitions

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
