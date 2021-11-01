/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.44.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IncomeOverride Specify payroll data on the account.
type IncomeOverride struct {
	// A list of paystubs associated with the account.
	Paystubs *[]PaystubOverride `json:"paystubs,omitempty"`
}

// NewIncomeOverride instantiates a new IncomeOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeOverride() *IncomeOverride {
	this := IncomeOverride{}
	return &this
}

// NewIncomeOverrideWithDefaults instantiates a new IncomeOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeOverrideWithDefaults() *IncomeOverride {
	this := IncomeOverride{}
	return &this
}

// GetPaystubs returns the Paystubs field value if set, zero value otherwise.
func (o *IncomeOverride) GetPaystubs() []PaystubOverride {
	if o == nil || o.Paystubs == nil {
		var ret []PaystubOverride
		return ret
	}
	return *o.Paystubs
}

// GetPaystubsOk returns a tuple with the Paystubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeOverride) GetPaystubsOk() (*[]PaystubOverride, bool) {
	if o == nil || o.Paystubs == nil {
		return nil, false
	}
	return o.Paystubs, true
}

// HasPaystubs returns a boolean if a field has been set.
func (o *IncomeOverride) HasPaystubs() bool {
	if o != nil && o.Paystubs != nil {
		return true
	}

	return false
}

// SetPaystubs gets a reference to the given []PaystubOverride and assigns it to the Paystubs field.
func (o *IncomeOverride) SetPaystubs(v []PaystubOverride) {
	o.Paystubs = &v
}

func (o IncomeOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Paystubs != nil {
		toSerialize["paystubs"] = o.Paystubs
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeOverride struct {
	value *IncomeOverride
	isSet bool
}

func (v NullableIncomeOverride) Get() *IncomeOverride {
	return v.value
}

func (v *NullableIncomeOverride) Set(val *IncomeOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeOverride(val *IncomeOverride) *NullableIncomeOverride {
	return &NullableIncomeOverride{value: val, isSet: true}
}

func (v NullableIncomeOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


