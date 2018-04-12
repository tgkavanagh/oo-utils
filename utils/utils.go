package utils

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func GetStringTag(c *cli.Context, tag string) string {
	if !c.IsSet(tag) {
		fmt.Printf("Failed to get %s\n", tag)
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(1)
	}

	return c.String(tag)
}

func GetFilename(c *cli.Context, tag string) string {
	fn := GetStringTag(c, tag)
	path := c.GlobalString("path")

	return path + fn
}

func ConvertContactToVendor(contact ContactInfo) QBVendorInfo {
	nv := QBVendorInfo{
		Vend:       WBDefaultFields[QB_VEND],
		Name:       contact.SellerNumber,
		Addr1:      contact.Address1,
		Addr2:      contact.Address2,
		Addr3:      contact.City,
		Addr4:      contact.State,
		Addr5:      contact.ZipCode,
		Phone1:     contact.HomeTelephone,
		Phone2:     contact.CellTelephone,
		Faxnum:     contact.WorkTelephone,
		Email:      contact.Email,
		Salutation: contact.Suffix,
		Firstname:  contact.FirstName,
		Midinit:    contact.MiddleName,
		Lastname:   contact.LastName,
		T1099:      WBDefaultFields[QB_T1099],
		Hidden:     WBDefaultFields[QB_HIDDEN],
		Delcount:   0,
	}

	return nv
}
