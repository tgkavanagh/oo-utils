package main

import (
  "fmt"
  "github.com/urfave/cli"
  "github.com/tgkavanagh/oo-utils/reports"
  "github.com/tgkavanagh/oo-utils/vendor"
  "math/rand"
  "os"
  "time"
)

var majorVersion = 1
var minorVersion = 0
var maintenanceVersion = 0

func main() {
  // Seed random so apps can use it to generate opaque topic identifiers.
	rand.Seed(time.Now().UTC().UnixNano())

	appVersion := fmt.Sprintf("%d.%d.%d", majorVersion, minorVersion, maintenanceVersion)

	app := cli.NewApp()
	app.Name = "oo-utils"
	app.Version = appVersion
	app.Usage = "OO Utils"
	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}

	// Global Option Flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "path, p",
			EnvVar: "OO_PATH",
      Value: "./",
			Usage:  "Path to OO files",
		},
	}

	// Commands
	app.Commands = []cli.Command{
		reports.Commands,
		vendor.Commands,
	}

	app.Run(os.Args)
}
