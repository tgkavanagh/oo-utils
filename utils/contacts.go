package utils

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
)

const (
	CN_SUFFIX = iota
	CN_FIRSTNAME
	CN_MIDDLENAME
	CN_LASTNAME
	CN_HOMETELEPHONE
	CN_WORKTELEPHONE
	CN_CELL
	CN_EMAIL
	CN_ADDR1
	CN_ADDR2
	CN_CITY
	CN_STATE
	CN_ZIPCODE
	CN_ISSELLER
	CN_SELLERNUMBER
	CN_ISVOLUNTEER
	CN_ISBUYER
	CN_ISVENDOR
	CN_ISNEWMOM
	CN_DATEADDED
)

type ContactInfo struct {
	Suffix        string
	FirstName     string
	MiddleName    string
	LastName      string
	HomeTelephone string
	WorkTelephone string
	CellTelephone string
	Email         string
	Address1      string
	Address2      string
	City          string
	State         string
	ZipCode       string
	IsSeller      string
	SellerNumber  int
	IsVolunteer   string
	IsBuyer       string
	IsVendor      string
	IsNewMom      string
	DateAdded     string
}

func convertSliceToContact(sn int, contact []string) ContactInfo {
	nc := ContactInfo{
		Suffix:        contact[CN_SUFFIX],
		FirstName:     contact[CN_FIRSTNAME],
		MiddleName:    contact[CN_MIDDLENAME],
		LastName:      contact[CN_LASTNAME],
		HomeTelephone: contact[CN_HOMETELEPHONE],
		WorkTelephone: contact[CN_WORKTELEPHONE],
		CellTelephone: contact[CN_CELL],
		Email:         contact[CN_EMAIL],
		Address1:      contact[CN_ADDR1],
		Address2:      contact[CN_ADDR2],
		City:          contact[CN_CITY],
		State:         contact[CN_STATE],
		ZipCode:       contact[CN_ZIPCODE],
		IsSeller:      contact[CN_ISSELLER],
		SellerNumber:  sn,
		IsVolunteer:   contact[CN_ISVOLUNTEER],
		IsBuyer:       contact[CN_ISBUYER],
		IsVendor:      contact[CN_ISVENDOR],
		IsNewMom:      contact[CN_ISNEWMOM],
		DateAdded:     contact[CN_DATEADDED],
	}

	return nc
}

func ParseContactList(fn string) map[int]ContactInfo {
	contactsList := make(map[int]ContactInfo)

	mlSlice, err := xlsx.FileToSlice(fn)
	if err != nil {
		fmt.Printf("Failed to read spreadsheet %s: %v\n", fn, err)
		return nil
	}

	for _, contacts := range mlSlice {
		for _, contact := range contacts {
			if contact[CN_ISSELLER] != "YES" {
				continue
			}

			if sn, err := strconv.Atoi(contact[CN_SELLERNUMBER]); err == nil {
				if _, ok := contactsList[sn]; !ok {
					nc := convertSliceToContact(sn, contact)
					contactsList[sn] = nc
				}
			}
		}
	}

	return contactsList
}

func DisplayContact(contact ContactInfo) {
	fmt.Printf("%d ", contact.SellerNumber)
	fmt.Printf("%s ", strings.Title(contact.FirstName))

	if contact.MiddleName != "" {
		fmt.Printf("%s ", strings.Title(contact.MiddleName))
	}

	fmt.Printf("%s ", strings.Title(contact.LastName))
	fmt.Printf("%s ", contact.Email)

	if contact.Address1 != "" {
		fmt.Printf("%s ", strings.Title(contact.Address1))
	}

	if contact.Address2 != "" {
		fmt.Printf("%s ", strings.Title(contact.Address2))
	}

	if contact.City != "" {
		fmt.Printf("%s ", strings.Title(contact.City))
	}

	if contact.State != "" {
		fmt.Printf("%s ", strings.Title(contact.State))
	}

	if contact.ZipCode != "" {
		fmt.Printf("%s ", strings.Title(contact.ZipCode))
	}

	fmt.Printf("\n")
}
