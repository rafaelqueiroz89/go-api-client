package accounts

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var accountService = AccountServiceOperator{}

//Test Fetch Accounts Tests
func TestFetch_With_Wrong_Id_Format(t *testing.T) {
	cases := []struct {
		name           string
		id             string
		expectedStatus int
	}{
		{
			name:           "Wrong Id",
			id:             "asfas",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty Id",
			id:             " ",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, resp, err := accountService.Fetch(tt.id)
			assert.NotNil(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}

func TestFetch_With_Unknown_Id(t *testing.T) {
	_, resp, err := accountService.Fetch(uuid.New().String())
	assert.NotNil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestFetch_With_Valid_AccountData(t *testing.T) {
	resC, respC, errC := accountService.Create(NewAccountDataRequestBuilder().Build())
	assert.Nil(t, errC)
	assert.Equal(t, http.StatusCreated, respC.StatusCode)
	assert.NotNil(t, respC)

	resF, respF, errF := accountService.Fetch(resC.Data.ID)
	assert.Nil(t, errF)
	assert.NotNil(t, resF)
	assert.Equal(t, http.StatusOK, respF.StatusCode)

	assert.Equal(t, resF.Data, resC.Data)
}

func TestFetch_With_Valid_AccountData_And_Valid_Attributes(t *testing.T) {
	resC, respC, errC := accountService.Create(
		NewAccountDataRequestBuilder().
			Build())
	assert.Nil(t, errC)
	assert.Equal(t, http.StatusCreated, respC.StatusCode)
	assert.NotNil(t, respC)

	resF, respF, errF := accountService.Fetch(resC.Data.ID)
	assert.Nil(t, errF)
	assert.NotNil(t, resF)
	assert.Equal(t, http.StatusOK, respF.StatusCode)

	assert.Equal(t, resF.Data, resC.Data)
}

//Test Create Accounts
func TestCreate_With_Valid_AccountData(t *testing.T) {
	account := NewAccountDataRequestBuilder().
		WithAccountClassification(AccountClassificationPersonal).
		Build()
	res, resp, err := accountService.Create(account)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.NotNil(t, res)
	assert.Equal(t, res.Data, account.Data)
}

func TestCreate_With_Invalid_AccountData(t *testing.T) {
	account := NewAccountDataRequestBuilder().Build()
	account.Data.ID = "xxx"
	res, resp, err := accountService.Create(account)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Nil(t, res)
}

func TestCreate_With_Valid_AccountData_And_Invalid_Attributes(t *testing.T) {
	cases := []struct {
		name           string
		request        *AccountDataRequest
		expectedResult *AccountDataResponse
		expectedError  bool
	}{
		{
			name:           "AccountData with wrong value for AccountClassification",
			request:        NewAccountDataRequestBuilder().WithAccountClassification("xpto").Build(),
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "AccountData with wrong value for Country",
			request:        NewAccountDataRequestBuilder().WithCountry("AAAAAAAF").Build(),
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "AccountData with wrong value for Status",
			request:        NewAccountDataRequestBuilder().WithStatus("AAAAAAAF").Build(),
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "AccountData with wrong value for Name",
			request:        NewAccountDataRequestBuilder().WithName(nil).Build(),
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "AccountData with bad formatted value for OrganisationId",
			request:        NewAccountDataRequestBuilder().WithOrganisationId("notaguid").Build(),
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "AccountData without any Attributes",
			request:        NewAccountDataRequestBuilder().WithAttributes(nil).Build(),
			expectedResult: nil,
			expectedError:  true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			res, resp, err := accountService.Create(tt.request)
			assert.Equal(t, res, tt.expectedResult)
			assert.NotNil(t, resp)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			assert.Equal(t, tt.expectedError, assert.NotNil(t, err))
		})
	}
}

func TestCreate_With_Valid_AccountData_And_Valid_Attributes(t *testing.T) {
	var account = NewAccountDataRequestBuilder().
		WithStatus(FailedStatus).
		Build()

	res, resp, err := accountService.Create(account)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.NotNil(t, res)
	assert.Equal(t, res.Data, account.Data)
}

//Test Delete Accounts
func TestDelete_With_Valid_AccountData(t *testing.T) {
	account := NewAccountDataRequestBuilder().Build()
	resp, err := accountService.Delete(account.Data.ID, *account.Data.Version)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDelete_With_Invalid_Version(t *testing.T) {
	account := NewAccountDataRequestBuilder().Build()
	version := int64(1)
	resp, err := accountService.Delete(account.Data.ID, version)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

//Test Builder
type AccountDataRequestBuilder struct {
	accountDataRequest *AccountDataRequest
}

func NewAccountDataRequestBuilder() *AccountDataRequestBuilder {
	version := int64(0)
	country, jointAccount, status := "UK", true, ConfirmedStatus
	switched, accMatchOutput, accountClassification := false, false, AccountClassificationBusiness

	return &AccountDataRequestBuilder{accountDataRequest: &AccountDataRequest{
		Data: &AccountData{
			Type:           "accounts",
			ID:             uuid.NewString(),
			OrganisationID: uuid.NewString(),
			Version:        &version,
			Attributes: &AccountAttributes{
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
				AccountClassification:   &accountClassification,
				AccountMatchingOptOut:   &accMatchOutput,
				AccountNumber:           "111231315",
				SecondaryIdentification: "Test1",
			},
		}},
	}
}

func (a *AccountDataRequestBuilder) Build() *AccountDataRequest {
	return a.accountDataRequest
}

func (a *AccountDataRequestBuilder) WithCountry(country string) *AccountDataRequestBuilder {
	a.accountDataRequest.Data.Attributes.AccountClassification = &country
	return a
}

func (a *AccountDataRequestBuilder) WithStatus(status string) *AccountDataRequestBuilder {
	a.accountDataRequest.Data.Attributes.Status = &status
	return a
}

func (a *AccountDataRequestBuilder) WithAccountClassification(accountClassification string) *AccountDataRequestBuilder {
	a.accountDataRequest.Data.Attributes.AccountClassification = &accountClassification
	return a
}

func (a *AccountDataRequestBuilder) WithName(name []string) *AccountDataRequestBuilder {
	a.accountDataRequest.Data.Attributes.Name = name
	return a
}

func (a *AccountDataRequestBuilder) WithOrganisationId(org string) *AccountDataRequestBuilder {
	a.accountDataRequest.Data.OrganisationID = org
	return a
}

func (a *AccountDataRequestBuilder) WithAttributes(attributes *AccountAttributes) *AccountDataRequestBuilder {
	a.accountDataRequest.Data.Attributes = attributes
	return a
}
