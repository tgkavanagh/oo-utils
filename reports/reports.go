package reports

import (
//  "fmt"
  "github.com/urfave/cli"
)

var Commands cli.Command = cli.Command{
	Name:  "reports",
	Usage: "Report Commands",

	Subcommands: []cli.Command{
		{
			Name:   "display",
			Usage:  "Display Sellers owed money",
			Action: displayBills,
      Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "Settlement report file",
				},
			},
		},
	},
}
