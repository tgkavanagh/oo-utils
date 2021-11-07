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

func ConvertContactToVendor(contact ContactInfo) QBVendorInfo2 {
  fullname := fmt.Sprintf("%s %s", contact.FirstName, contact.LastName)
	addr1 := fullname
	addr2 := fmt.Sprintf("%s %s", contact.Address1, contact.Address2)
	addr3 := fmt.Sprintf("%s, %s %s", contact.City, contact.State, contact.ZipCode)

fmt.Printf("Full name: %v\n", fullname)

	nv := QBVendorInfo2{
		Name:       contact.SellerNumber,
		Printas:    fullname,
		CompanyName: fullname,
		Firstname:  contact.FirstName,
		Lastname:   contact.LastName,
		Email:      contact.Email,
		Addr1:      addr1,
		Addr2:      addr2,
		Addr3:      addr3,
		Phone1:     contact.HomeTelephone,
		Phone2:     contact.CellTelephone,
	}

	return nv
}
