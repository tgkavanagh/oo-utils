package reports

import (
	"fmt"
	"github.com/tgkavanagh/oo-utils/utils"
	"github.com/urfave/cli"
	"sort"
)
func displaySettlementReport(c *cli.Context) error {
  fn := utils.GetFilename(c, "file")
	nfn := c.GlobalString("path") + "newsettlements.xlsx"

	settlements := utils.ParseSettlementReport(fn)

  keys := make([]int, 0, len(settlements))
	for k := range settlements {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
    utils.DisplaySettlementInfo(settlements[k])
	}

	fmt.Printf("Total checks to be written: %d\n", len(settlements))

	utils.WriteSettlementSpreadsheet(nfn, settlements)

	return nil
}
