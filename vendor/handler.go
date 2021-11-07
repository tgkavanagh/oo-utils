package vendor

import (
	"fmt"
	"github.com/tgkavanagh/oo-utils/utils"
	"github.com/urfave/cli"
	"sort"
)

func displayCurrentVendors(c *cli.Context) error {
	fn := utils.GetFilename(c, "file")

	vendors := utils.ParseQBxlsx(fn)

  keys := make([]int, 0, len(vendors))
	for k := range vendors {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
    utils.DisplayQBVendor(vendors[k])
	}

  fmt.Printf("# Current Vendors: %d\n", len(vendors))

	return nil
}

func displayContactsList(c *cli.Context) error {
	fn := utils.GetFilename(c, "file")

	contacts := utils.ParseContactList(fn)

	keys := make([]int, 0, len(contacts))
	for k := range contacts {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		utils.DisplayContact(contacts[k])
	}

	fmt.Printf("# Contacts: %d\n", len(contacts))
	return nil
}

func displayNewVendors(c *cli.Context) error {
	mfn := utils.GetFilename(c, "master")
	cfn := utils.GetFilename(c, "contacts")
  nfn := c.GlobalString("path") + "newvendors.xlsx"

  mvendors := utils.ParseQBxlsx(mfn)
  contacts := utils.ParseContactList(cfn)

  keys := make([]int, 0, len(contacts))
  for k := range contacts {
    keys = append(keys, k)
  }
  sort.Ints(keys)

  nvendors := make(map[int]utils.QBVendorInfo2)
	cvendors := make(map[int]utils.QBVendorInfo2)

  for _, k := range keys {
    // Check if the contact's seller number is in the mvendors list
		nv := utils.ConvertContactToVendor(contacts[k])

    if mv, ok := mvendors[k]; !ok {
      nvendors[k] = nv
      //utils.DisplayQBVendor(nv)
    } else {
			if mv != nv {
				fmt.Println("\nOld vendor information is out of date")
				fmt.Printf("Old: %v\n", mv)
				fmt.Printf("New: %v\n\n", nv)

				cvendors[k] = nv
			}
		}
  }

  fmt.Printf("# New Vendors: %d\n", len(nvendors))

  utils.WriteVendorSpreadsheet(nfn, nvendors)
	return nil
}
