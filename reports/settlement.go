package reports

import (
  "fmt"
	"github.com/tealeg/xlsx"
	"github.com/urfave/cli"
  "strconv"
	"os"
)

const (
  FIRSTNAME	= iota
  LASTNAME
  EMAIL
  SELLERNUMBER
  TOTALITEMSSOLD
  SUBTOTAL
  SELLERPERCENTAGE
  SELLERPERCENTAGEAMOUNT
  REGISTRATIONAMOUNT
  TOTAL
  CREATIONDATE
  REPORTVIEWED
  VIEWEDDATE
)

func parseSettlementReport(fn string) [][]string {
  var billsList [][]string

  srSlice, err := xlsx.FileToSlice(fn)
  if err != nil {
    fmt.Printf("Failed to read spreadsheet %s: %v\n", fn, err)
    return nil
  }

  for _, sellers := range srSlice {
    for _,  seller := range sellers {
      if seller[SELLERNUMBER] == "" {
        continue
      }

      if owed, err := strconv.ParseFloat(seller[TOTAL], 64); err == nil {
        if owed > 0.00 {
          billsList = append(billsList, seller)
        }
      }
    }
  }

  return billsList
}

func displayBills(c *cli.Context) error {
  if !c.IsSet("file") {
    cli.ShowCommandHelp(c, c.Command.Name)
    os.Exit(1)
  }

  fn := c.GlobalString("path") + c.String("file")

  bills := parseSettlementReport(fn)

  for _, bill := range bills{
    displayBill(bill)
  }
  fmt.Printf("Total checks to be writtent: %d\n", len(bills))

  return nil
}

func displayBill(bill []string) {
  fmt.Printf("%v\n", bill)
}
