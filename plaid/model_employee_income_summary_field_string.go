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

// EmployeeIncomeSummaryFieldString struct for EmployeeIncomeSummaryFieldString
type EmployeeIncomeSummaryFieldString struct {
	// The value of the field.
	Value string `json:"value"`
	VerificationStatus VerificationStatus `json:"verification_status"`
}

// NewEmployeeIncomeSummaryFieldString instantiates a new EmployeeIncomeSummaryFieldString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeIncomeSummaryFieldString(value string, verificationStatus VerificationStatus) *EmployeeIncomeSummaryFieldString {
	this := EmployeeIncomeSummaryFieldString{}
	this.Value = value
	this.VerificationStatus = verificationStatus
	return &this
}

// NewEmployeeIncomeSummaryFieldStringWithDefaults instantiates a new EmployeeIncomeSummaryFieldString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeIncomeSummaryFieldStringWithDefaults() *EmployeeIncomeSummaryFieldString {
	this := EmployeeIncomeSummaryFieldString{}
	return &this
}

// GetValue returns the Value field value
func (o *EmployeeIncomeSummaryFieldString) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EmployeeIncomeSummaryFieldString) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EmployeeIncomeSummaryFieldString) SetValue(v string) {
	o.Value = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *EmployeeIncomeSummaryFieldString) GetVerificationStatus() VerificationStatus {
	if o == nil {
		var ret VerificationStatus
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *EmployeeIncomeSummaryFieldString) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *EmployeeIncomeSummaryFieldString) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = v
}

func (o EmployeeIncomeSummaryFieldString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableEmployeeIncomeSummaryFieldString struct {
	value *EmployeeIncomeSummaryFieldString
	isSet bool
}

func (v NullableEmployeeIncomeSummaryFieldString) Get() *EmployeeIncomeSummaryFieldString {
	return v.value
}

func (v *NullableEmployeeIncomeSummaryFieldString) Set(val *EmployeeIncomeSummaryFieldString) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeIncomeSummaryFieldString) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeIncomeSummaryFieldString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeIncomeSummaryFieldString(val *EmployeeIncomeSummaryFieldString) *NullableEmployeeIncomeSummaryFieldString {
	return &NullableEmployeeIncomeSummaryFieldString{value: val, isSet: true}
}

func (v NullableEmployeeIncomeSummaryFieldString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeIncomeSummaryFieldString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


