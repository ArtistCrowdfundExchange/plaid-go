/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.39.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AccountFilter Enumerates the account subtypes that the application wishes for the user to be able to select from. For more details refer to Plaid documentation on account filters.
type AccountFilter struct {
	// A list of account subtypes to be filtered.
	Depository *[]string `json:"depository,omitempty"`
	// A list of account subtypes to be filtered.
	Credit *[]string `json:"credit,omitempty"`
	// A list of account subtypes to be filtered.
	Loan *[]string `json:"loan,omitempty"`
	// A list of account subtypes to be filtered.
	Investment *[]string `json:"investment,omitempty"`
}

// NewAccountFilter instantiates a new AccountFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountFilter() *AccountFilter {
	this := AccountFilter{}
	return &this
}

// NewAccountFilterWithDefaults instantiates a new AccountFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountFilterWithDefaults() *AccountFilter {
	this := AccountFilter{}
	return &this
}

// GetDepository returns the Depository field value if set, zero value otherwise.
func (o *AccountFilter) GetDepository() []string {
	if o == nil || o.Depository == nil {
		var ret []string
		return ret
	}
	return *o.Depository
}

// GetDepositoryOk returns a tuple with the Depository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFilter) GetDepositoryOk() (*[]string, bool) {
	if o == nil || o.Depository == nil {
		return nil, false
	}
	return o.Depository, true
}

// HasDepository returns a boolean if a field has been set.
func (o *AccountFilter) HasDepository() bool {
	if o != nil && o.Depository != nil {
		return true
	}

	return false
}

// SetDepository gets a reference to the given []string and assigns it to the Depository field.
func (o *AccountFilter) SetDepository(v []string) {
	o.Depository = &v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *AccountFilter) GetCredit() []string {
	if o == nil || o.Credit == nil {
		var ret []string
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFilter) GetCreditOk() (*[]string, bool) {
	if o == nil || o.Credit == nil {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *AccountFilter) HasCredit() bool {
	if o != nil && o.Credit != nil {
		return true
	}

	return false
}

// SetCredit gets a reference to the given []string and assigns it to the Credit field.
func (o *AccountFilter) SetCredit(v []string) {
	o.Credit = &v
}

// GetLoan returns the Loan field value if set, zero value otherwise.
func (o *AccountFilter) GetLoan() []string {
	if o == nil || o.Loan == nil {
		var ret []string
		return ret
	}
	return *o.Loan
}

// GetLoanOk returns a tuple with the Loan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFilter) GetLoanOk() (*[]string, bool) {
	if o == nil || o.Loan == nil {
		return nil, false
	}
	return o.Loan, true
}

// HasLoan returns a boolean if a field has been set.
func (o *AccountFilter) HasLoan() bool {
	if o != nil && o.Loan != nil {
		return true
	}

	return false
}

// SetLoan gets a reference to the given []string and assigns it to the Loan field.
func (o *AccountFilter) SetLoan(v []string) {
	o.Loan = &v
}

// GetInvestment returns the Investment field value if set, zero value otherwise.
func (o *AccountFilter) GetInvestment() []string {
	if o == nil || o.Investment == nil {
		var ret []string
		return ret
	}
	return *o.Investment
}

// GetInvestmentOk returns a tuple with the Investment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFilter) GetInvestmentOk() (*[]string, bool) {
	if o == nil || o.Investment == nil {
		return nil, false
	}
	return o.Investment, true
}

// HasInvestment returns a boolean if a field has been set.
func (o *AccountFilter) HasInvestment() bool {
	if o != nil && o.Investment != nil {
		return true
	}

	return false
}

// SetInvestment gets a reference to the given []string and assigns it to the Investment field.
func (o *AccountFilter) SetInvestment(v []string) {
	o.Investment = &v
}

func (o AccountFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Depository != nil {
		toSerialize["depository"] = o.Depository
	}
	if o.Credit != nil {
		toSerialize["credit"] = o.Credit
	}
	if o.Loan != nil {
		toSerialize["loan"] = o.Loan
	}
	if o.Investment != nil {
		toSerialize["investment"] = o.Investment
	}
	return json.Marshal(toSerialize)
}

type NullableAccountFilter struct {
	value *AccountFilter
	isSet bool
}

func (v NullableAccountFilter) Get() *AccountFilter {
	return v.value
}

func (v *NullableAccountFilter) Set(val *AccountFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountFilter(val *AccountFilter) *NullableAccountFilter {
	return &NullableAccountFilter{value: val, isSet: true}
}

func (v NullableAccountFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


