/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.19.12
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// BankTransferType The type of bank transfer. This will be either `debit` or `credit`.  A `debit` indicates a transfer of money into your origination account; a `credit` indicates a transfer of money out of your origination account.
type BankTransferType string

// List of BankTransferType
const (
	BANKTRANSFERTYPE_DEBIT BankTransferType = "debit"
	BANKTRANSFERTYPE_CREDIT BankTransferType = "credit"
)

func (v *BankTransferType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BankTransferType(value)
	for _, existing := range []BankTransferType{ "debit", "credit",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BankTransferType", value)
}

// Ptr returns reference to BankTransferType value
func (v BankTransferType) Ptr() *BankTransferType {
	return &v
}

type NullableBankTransferType struct {
	value *BankTransferType
	isSet bool
}

func (v NullableBankTransferType) Get() *BankTransferType {
	return v.value
}

func (v *NullableBankTransferType) Set(val *BankTransferType) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferType) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferType(val *BankTransferType) *NullableBankTransferType {
	return &NullableBankTransferType{value: val, isSet: true}
}

func (v NullableBankTransferType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

