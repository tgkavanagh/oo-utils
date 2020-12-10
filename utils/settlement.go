package utils

import (
	"fmt"
	"github.com/tealeg/xlsx"
  "sort"
	"strconv"
  "strings"
)

const (
	FIRSTNAME = iota
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

var SettlementFields = []string{
	FIRSTNAME:              "First Name",
	LASTNAME:               "Last Name",
	EMAIL:                  "Email",
	SELLERNUMBER:           "Seller Number",
	TOTALITEMSSOLD:         "Total Items Sold",
	SUBTOTAL:               "SubTotal",
	SELLERPERCENTAGE:       "Seller %",
	SELLERPERCENTAGEAMOUNT: "Seller % Amount",
	REGISTRATIONAMOUNT:     "Registration Amount",
	TOTAL:                  "Total",
	CREATIONDATE:           "Creation Date",
	REPORTVIEWED:           "Report Viewed",
	VIEWEDDATE:             "Viewed Date",
}

/*
type SettlementInfo struct {
	FirstName              string
	LastName               string
	Email                  string
	SellerNumber           int
	TotalItemsSold         int
	SubTotal               float64
	SellerPercentage       int
	SellerPercentageAmount float64
	RegistrationAmount     float64
	Total                  float64
	CreationDate           string
	ReportViewed           string
	ViewedDate             string
}


func convertSliceToSettlementInfo(sn int, ss []string) SettlementInfo {
	nsi := SettlementInfo{
		FirstName:    ss[FIRSTNAME],
		LastName:     ss[LASTNAME],
		Email:        ss[EMAIL],
		CreationDate: ss[CREATIONDATE],
		ReportViewed: ss[REPORTVIEWED],
		ViewedDate:   ss[VIEWEDDATE],
	}

	nsi.SellerNumber = sn
	nsi.TotalItemsSold, _ = strconv.Atoi(ss[TOTALITEMSSOLD])

	nsi.SubTotal, _ = strconv.ParseFloat(ss[SUBTOTAL], 64)
	nsi.SellerPercentage, _ = strconv.Atoi(ss[SELLERPERCENTAGE])
	nsi.SellerPercentageAmount, _ = strconv.ParseFloat(ss[SELLERPERCENTAGEAMOUNT], 64)
	nsi.RegistrationAmount, _ = strconv.ParseFloat(ss[REGISTRATIONAMOUNT], 64)
	nsi.Total, _ = strconv.ParseFloat(ss[TOTAL], 64)

	return nsi
}
*/

type SettlementInfo struct {
  SellerNumber           int
	FirstName              string
	LastName               string
	SubTotal							 float64
	SellerPercentage			 string
	SellerPercentageAmount float64
	RegistrationAmount     float64
	Total                  float64
}


func convertSliceToSettlementInfo(sn int, ss []string) SettlementInfo {
	nsi := SettlementInfo{
		FirstName:    ss[FIRSTNAME],
		LastName:     ss[LASTNAME],
	}

	nsi.SellerNumber = sn
	nsi.SubTotal, _ = strconv.ParseFloat(ss[SUBTOTAL], 64)
	nsi.SellerPercentage = ss[SELLERPERCENTAGE]
	nsi.SellerPercentageAmount, _ = strconv.ParseFloat(ss[SELLERPERCENTAGEAMOUNT], 64)
	nsi.RegistrationAmount, _ = strconv.ParseFloat(ss[REGISTRATIONAMOUNT], 64)
	nsi.Total, _ = strconv.ParseFloat(ss[TOTAL], 64)

	return nsi
}

func ParseSettlementReport(fn string) map[int]SettlementInfo {
	siList := make(map[int]SettlementInfo)

	srSlice, err := xlsx.FileToSlice(fn)
	if err != nil {
		fmt.Printf("Failed to read spreadsheet %s: %v\n", fn, err)
		return nil
	}

	for _, settlements := range srSlice {
		for _, settlement := range settlements {
      if sn, err := strconv.Atoi(settlement[SELLERNUMBER]); err == nil {
        if _, ok := siList[sn]; !ok {
          si := convertSliceToSettlementInfo(sn, settlement)
          if owed, err := strconv.ParseFloat(settlement[TOTAL], 64); err == nil && owed > 0.00 {
            siList[sn] = si
    			}
        }
			}
		}
	}

	return siList
}

func DisplaySettlementInfo(si SettlementInfo) {
  fmt.Printf("%d ", si.SellerNumber)
  fmt.Printf("%s ", strings.Title(si.FirstName))
  fmt.Printf("%s ", strings.Title(si.LastName))
	fmt.Printf("%s ", si.SellerPercentage)
  fmt.Printf("%.2f ", si.SellerPercentageAmount)
  fmt.Printf("%.2f ", si.RegistrationAmount)
  fmt.Printf("%.2f ", si.Total)
  fmt.Printf("\n")
}

func WriteSettlementSpreadsheet(fn string, settlements map[int]SettlementInfo) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	keys := make([]int, 0, len(settlements))
	for k := range settlements {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i, k := range keys {
		row := sheet.AddRow()
		v := settlements[k]
    fmt.Printf("\rWriting new settlement: %d", i+1)
		row.WriteStruct(&v, 45)

		err = file.Save(fn)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

  fmt.Printf("\nFinished writing %s\n", fn)
}
