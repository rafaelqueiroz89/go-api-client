package accounts

import (
	"fmt"
	"github.com/rafaelqueiroz89/go-api-client/src/client"
	"net/http"
	"strconv"
)

const accountsBasePath = "v1/organisation/accounts"

//Constants for AccountClassification and Status
const (
	AccountClassificationPersonal string = "Personal"
	AccountClassificationBusiness        = "Business"
	PendingStatus                        = "pending"
	ConfirmedStatus                      = "confirmed"
	FailedStatus                         = "failed"
)

// AccountService Interface that exposes the methods to make operations in the Accounts resources
type AccountService interface {
	Fetch(id string) (*AccountDataResponse, *http.Response, error)
	Create(accountData *AccountDataRequest) (*AccountDataResponse, *http.Response, error)
	Delete(id string, version int64) (*http.Response, error)
	SetBaseUrl(id string)
}

// AccountServiceOperator The Account Service Operator to make the client work with the Accounts resources
// If no Base URL is given then the default ones will be used
type AccountServiceOperator struct {
	client client.Client
}

type AccountDataResponse struct {
	Data *AccountData `json:"data"`
}

type AccountDataRequest struct {
	Data *AccountData `json:"data"`
}

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

// SetBaseUrl Sets base url
func (a *AccountServiceOperator) SetBaseUrl(url string) {
	a.client.BaseUrl = url
}

// Fetch an Account based on a valid Id
func (a *AccountServiceOperator) Fetch(id string) (*AccountDataResponse, *http.Response, error) {
	accountData := &AccountDataResponse{}
	resp, err := a.client.NewClient().Request(accountData, http.MethodGet, fmt.Sprintf("%s", accountsBasePath+"/"+id), nil, nil)

	if err != nil {
		return nil, resp, err
	}

	return accountData, resp, nil
}

// Create an accounts based on a AccountDataRequest root object
func (a *AccountServiceOperator) Create(data *AccountDataRequest) (*AccountDataResponse, *http.Response, error) {
	accountData := &AccountDataResponse{}
	resp, err := a.client.NewClient().Request(accountData, http.MethodPost, fmt.Sprintf("%s", accountsBasePath), nil, data)

	if err != nil {
		return nil, resp, err
	}

	return accountData, resp, nil
}

// Delete an account based on Id and Version
func (a *AccountServiceOperator) Delete(id string, version int64) (*http.Response, error) {
	accountData := &AccountDataResponse{}
	query := make(map[string]string)
	query["version"] = strconv.FormatInt(version, 10)

	resp, err := a.client.NewClient().Request(accountData, http.MethodDelete, fmt.Sprintf("%s/%s", accountsBasePath, id), query, nil)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
