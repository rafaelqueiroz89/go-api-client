package main

import (
	"fmt"
	"github.com/rafaelqueiroz89/form3-client-api/src/accounts"
)

// The first step is to create an Account Service and assign it to its operator
// This will give all the operations in the API for the Accounts resource
var accountService = accounts.AccountServiceOperator{}

// Example function for Create, Fetch and Delete
func main() {
	country, accClassification, status := "GB", accounts.AccountClassificationBusiness, accounts.PendingStatus
	jointAccount, switched, accMatchOutput := false, false, false
	version := int64(0)

	a, b, c := accountService.Create(
		&accounts.AccountDataRequest{
			Data: &accounts.AccountData{
				Type:           "accounts",
				ID:             "4c54ff77-8067-43a7-807f-da216d598ad4",
				OrganisationID: "8acbc689-0d73-447d-96a1-c071b3e6ba5f",
				Version:        &version,
				Attributes: &accounts.AccountAttributes{
					Country:                 &country,
					BaseCurrency:            "GBP",
					BankID:                  "400300",
					BankIDCode:              "MASCX",
					Bic:                     "BICABM12",
					Iban:                    "GB94BARC10201530093459",
					JointAccount:            &jointAccount,
					Status:                  &status,
					Switched:                &switched,
					Name:                    []string{"Rafael Queiroz"},
					AlternativeNames:        []string{"My alternative name"},
					AccountClassification:   &accClassification,
					AccountMatchingOptOut:   &accMatchOutput,
					AccountNumber:           "111231315",
					SecondaryIdentification: "Test1",
				},
			},
		})

	fmt.Println(a)
	fmt.Print(b)
	fmt.Print(c)
	_, _, _ = accountService.Fetch("4c54ff77-8067-43a7-807f-da216d598ad4")

	_, _ = accountService.Delete("4c54ff77-8067-43a7-807f-da216d598ad4", 0)
}
