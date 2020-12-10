package utils

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"sort"
	"strconv"
	"strings"
)

const (
	QB_VEND = iota
	QB_NAME
	QB_REFNUM
	QB_TIMESTAMP
	QB_PRINTAS
	QB_ADDR1
	QB_ADDR2
	QB_ADDR3
	QB_ADDR4
	QB_ADDR5
	QB_VTYPE
	QB_CONT1
	QB_CONT2
	QB_PHONE1
	QB_PHONE2
	QB_FAXNUM
	QB_EMAIL
	QB_NOTE
	QB_TAXID
	QB_LIMIT
	QB_TERMS
	QB_NOTEPAD
	QB_SALUTATION
	QB_COMPANYNAME
	QB_FIRSTNAME
	QB_MIDINIT
	QB_LASTNAME
	QB_CUSTFLD1
	QB_CUSTFLD2
	QB_CUSTFLD3
	QB_CUSTFLD4
	QB_CUSTFLD5
	QB_CUSTFLD6
	QB_CUSTFLD7
	QB_CUSTFLD8
	QB_CUSTFLD9
	QB_CUSTFLD10
	QB_CUSTFLD11
	QB_CUSTFLD12
	QB_CUSTFLD13
	QB_CUSTFLD14
	QB_CUSTFLD15
	QB_T1099
	QB_HIDDEN
	QB_DELCOUNT
)

var QBFields = []string{
	QB_VEND:        "!VEND",
	QB_NAME:        "NAME",
	QB_REFNUM:      "REFNUM",
	QB_TIMESTAMP:   "TIMESTAMP",
	QB_PRINTAS:     "PRINTAS",
	QB_ADDR1:       "ADDR1",
	QB_ADDR2:       "ADDR2",
	QB_ADDR3:       "ADDR3",
	QB_ADDR4:       "ADDR4",
	QB_ADDR5:       "ADDR5",
	QB_VTYPE:       "VTYPE",
	QB_CONT1:       "CONT1",
	QB_CONT2:       "CONT2",
	QB_PHONE1:      "PHONE1",
	QB_PHONE2:      "PHONE2",
	QB_FAXNUM:      "FAXNUM",
	QB_EMAIL:       "EMAIL",
	QB_NOTE:        "NOTE",
	QB_TAXID:       "TAXID",
	QB_LIMIT:       "LIMIT",
	QB_TERMS:       "TERMS",
	QB_NOTEPAD:     "NOTEPAD",
	QB_SALUTATION:  "SALUTATION",
	QB_COMPANYNAME: "COMPANYNAME",
	QB_FIRSTNAME:   "FIRSTNAME",
	QB_MIDINIT:     "MIDINIT",
	QB_LASTNAME:    "LASTNAME",
	QB_CUSTFLD1:    "CUSTFLD1",
	QB_CUSTFLD2:    "CUSTFLD2",
	QB_CUSTFLD3:    "CUSTFLD3",
	QB_CUSTFLD4:    "CUSTFLD4",
	QB_CUSTFLD5:    "CUSTFLD5",
	QB_CUSTFLD6:    "CUSTFLD6",
	QB_CUSTFLD7:    "CUSTFLD7",
	QB_CUSTFLD8:    "CUSTFLD8",
	QB_CUSTFLD9:    "CUSTFLD9",
	QB_CUSTFLD10:   "CUSTFLD10",
	QB_CUSTFLD11:   "CUSTFLD11",
	QB_CUSTFLD12:   "CUSTFLD12",
	QB_CUSTFLD13:   "CUSTFLD13",
	QB_CUSTFLD14:   "CUSTFLD14",
	QB_CUSTFLD15:   "CUSTFLD15",
	QB_T1099:       "1099",
	QB_HIDDEN:      "HIDDEN",
	QB_DELCOUNT:    "DELCOUNT",
}

type QBVendorInfo struct {
	Vend        string
	Name        int
	Refnum      string
	Timestamp   string
	Printas     string
	Addr1       string
	Addr2       string
	Addr3       string
	Addr4       string
	Addr5       string
	VType       string
	Cont1       string
	Cont2       string
	Phone1      string
	Phone2      string
	Faxnum      string
	Email       string
	Note        string
	Taxid       string
	Limit       string
	Terms       string
	Notepad     string
	Salutation  string
	CompanyName string
	Firstname   string
	Midinit     string
	Lastname    string
	Custfld1    string
	Custfld2    string
	Custfld3    string
	Custfld4    string
	Custfld5    string
	Custfld6    string
	Custfld7    string
	Custfld8    string
	Custfld9    string
	Custfld10   string
	Custfld11   string
	Custfld12   string
	Custfld13   string
	Custfld14   string
	Custfld15   string
	T1099       string
	Hidden      string
	Delcount    int
}

type QBVendorInfo2 struct {
	Name        int
	Printas     string
  CompanyName string
	Firstname   string
	Lastname    string
  Email       string
	Addr1       string
	Addr2       string
	Addr3       string
	Addr4       string
	Addr5       string
	Phone1      string
	Phone2      string
}

var WBDefaultFields = []string{
	QB_VEND:   "VEND",
	QB_T1099:  "N",
	QB_HIDDEN: "N",
}

func convertQBVendorToSlice(vendor QBVendorInfo2) []string {
	vslice := []string{
		QB_NAME:        fmt.Sprintf("%d", vendor.Name),
		QB_PRINTAS:     vendor.Printas,
    QB_COMPANYNAME: vendor.CompanyName,
		QB_FIRSTNAME:   vendor.Firstname,
		QB_LASTNAME:    vendor.Lastname,
    QB_EMAIL:       vendor.Email,
		QB_ADDR1:       vendor.Addr1,
		QB_ADDR2:       vendor.Addr2,
		QB_ADDR3:       vendor.Addr3,
		QB_ADDR4:       vendor.Addr4,
		QB_ADDR5:       vendor.Addr5,
		QB_PHONE1:      vendor.Phone1,
		QB_PHONE2:      vendor.Phone2,
	}

	return vslice
}

func convertSliceToQBVendor(vn int, vendor []string) QBVendorInfo2 {
	nv := QBVendorInfo2{
		Name:        vn,
		Printas:     vendor[QB_PRINTAS],
    CompanyName: vendor[QB_COMPANYNAME],
		Firstname:   vendor[QB_FIRSTNAME],
		Lastname:    vendor[QB_LASTNAME],
    Email:       vendor[QB_EMAIL],
		Addr1:       vendor[QB_ADDR1],
		Addr2:       vendor[QB_ADDR2],
		Addr3:       vendor[QB_ADDR3],
		Addr4:       vendor[QB_ADDR4],
		Addr5:       vendor[QB_ADDR5],
		Phone1:      vendor[QB_PHONE1],
		Phone2:      vendor[QB_PHONE2],
	}

	return nv
}

func ParseQBxlsx(fn string) map[int]QBVendorInfo2 {
	vendorList := make(map[int]QBVendorInfo2)

	qbvSlice, err := xlsx.FileToSlice(fn)
	if err != nil {
		fmt.Printf("Failed to read spreadsheet %s: %v\n", fn, err)
		return nil
	}

	for _, vendors := range qbvSlice {
		for _, vendor := range vendors {
			if vendor[QB_VEND] != "VEND" || len(vendor) != 45 {
				continue
			}

			if vn, err := strconv.Atoi(vendor[QB_NAME]); err == nil {
				if _, ok := vendorList[vn]; !ok {
					nv := convertSliceToQBVendor(vn, vendor)
					vendorList[vn] = nv
				}
			}
		}
	}

	return vendorList
}

func DisplayQBVendor(vendor QBVendorInfo2) {
	fmt.Printf("%d ", vendor.Name)
	fmt.Printf("%s ", strings.Title(vendor.Firstname))

	fmt.Printf("%s ", strings.Title(vendor.Lastname))
	fmt.Printf("%s ", vendor.Email)

	if vendor.Addr1 != "" {
		fmt.Printf("%s ", strings.Title(vendor.Addr1))
	}

	if vendor.Addr2 != "" {
		fmt.Printf("%s ", strings.Title(vendor.Addr2))
	}

	if vendor.Addr3 != "" {
		fmt.Printf("%s ", strings.Title(vendor.Addr3))
	}

	if vendor.Addr4 != "" {
		fmt.Printf("%s ", strings.Title(vendor.Addr4))
	}

	if vendor.Addr5 != "" {
		fmt.Printf("%s ", strings.Title(vendor.Addr5))
	}

	fmt.Printf("\n")
}

func WriteVendorSpreadsheet(fn string, vendors map[int]QBVendorInfo2) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	keys := make([]int, 0, len(vendors))
	for k := range vendors {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i, k := range keys {
		row := sheet.AddRow()
		v := vendors[k]
    fmt.Printf("\rWriting new vendor: %d", i+1)
		row.WriteStruct(&v, 45)

		err = file.Save(fn)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

  fmt.Printf("\nFinished writing %s\n", fn)
}
