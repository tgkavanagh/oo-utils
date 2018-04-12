package vendor

import (
	//  "fmt"
	"github.com/urfave/cli"
)

var Commands cli.Command = cli.Command{
	Name:  "vendor",
	Usage: "Vendor List Commands",

	Subcommands: []cli.Command{
		{
			Name:   "current",
			Usage:  "List all vendors",
			Action: displayCurrentVendors,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "Quickbooks xlsx file",
				},
			},
		},
		{
			Name:   "contacts",
			Usage:  "List all contacts",
			Action: displayContactsList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "Mailing List xlsx file from MyCM",
				},
			},
		},
		{
			Name:   "new",
			Usage:  "List new vendors",
			Action: displayNewVendors,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "master, m",
					Usage: "master xlsx file (from Quickbooks)",
				},
				cli.StringFlag{
					Name:  "contacts, c",
					Usage: "contacts xlsx file (from MyCM)",
				},
			},
		},
	},
}
