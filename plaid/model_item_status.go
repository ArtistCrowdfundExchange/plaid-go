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

// ItemStatus An object with information about the status of the Item.
type ItemStatus struct {
	Investments NullableItemStatusInvestments `json:"investments,omitempty"`
	Transactions NullableItemStatusTransactions `json:"transactions,omitempty"`
	LastWebhook NullableItemStatusLastWebhook `json:"last_webhook,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ItemStatus ItemStatus

// NewItemStatus instantiates a new ItemStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemStatus() *ItemStatus {
	this := ItemStatus{}
	return &this
}

// NewItemStatusWithDefaults instantiates a new ItemStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemStatusWithDefaults() *ItemStatus {
	this := ItemStatus{}
	return &this
}

// GetInvestments returns the Investments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatus) GetInvestments() ItemStatusInvestments {
	if o == nil || o.Investments.Get() == nil {
		var ret ItemStatusInvestments
		return ret
	}
	return *o.Investments.Get()
}

// GetInvestmentsOk returns a tuple with the Investments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatus) GetInvestmentsOk() (*ItemStatusInvestments, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Investments.Get(), o.Investments.IsSet()
}

// HasInvestments returns a boolean if a field has been set.
func (o *ItemStatus) HasInvestments() bool {
	if o != nil && o.Investments.IsSet() {
		return true
	}

	return false
}

// SetInvestments gets a reference to the given NullableItemStatusInvestments and assigns it to the Investments field.
func (o *ItemStatus) SetInvestments(v ItemStatusInvestments) {
	o.Investments.Set(&v)
}
// SetInvestmentsNil sets the value for Investments to be an explicit nil
func (o *ItemStatus) SetInvestmentsNil() {
	o.Investments.Set(nil)
}

// UnsetInvestments ensures that no value is present for Investments, not even an explicit nil
func (o *ItemStatus) UnsetInvestments() {
	o.Investments.Unset()
}

// GetTransactions returns the Transactions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatus) GetTransactions() ItemStatusTransactions {
	if o == nil || o.Transactions.Get() == nil {
		var ret ItemStatusTransactions
		return ret
	}
	return *o.Transactions.Get()
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatus) GetTransactionsOk() (*ItemStatusTransactions, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Transactions.Get(), o.Transactions.IsSet()
}

// HasTransactions returns a boolean if a field has been set.
func (o *ItemStatus) HasTransactions() bool {
	if o != nil && o.Transactions.IsSet() {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given NullableItemStatusTransactions and assigns it to the Transactions field.
func (o *ItemStatus) SetTransactions(v ItemStatusTransactions) {
	o.Transactions.Set(&v)
}
// SetTransactionsNil sets the value for Transactions to be an explicit nil
func (o *ItemStatus) SetTransactionsNil() {
	o.Transactions.Set(nil)
}

// UnsetTransactions ensures that no value is present for Transactions, not even an explicit nil
func (o *ItemStatus) UnsetTransactions() {
	o.Transactions.Unset()
}

// GetLastWebhook returns the LastWebhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatus) GetLastWebhook() ItemStatusLastWebhook {
	if o == nil || o.LastWebhook.Get() == nil {
		var ret ItemStatusLastWebhook
		return ret
	}
	return *o.LastWebhook.Get()
}

// GetLastWebhookOk returns a tuple with the LastWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatus) GetLastWebhookOk() (*ItemStatusLastWebhook, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastWebhook.Get(), o.LastWebhook.IsSet()
}

// HasLastWebhook returns a boolean if a field has been set.
func (o *ItemStatus) HasLastWebhook() bool {
	if o != nil && o.LastWebhook.IsSet() {
		return true
	}

	return false
}

// SetLastWebhook gets a reference to the given NullableItemStatusLastWebhook and assigns it to the LastWebhook field.
func (o *ItemStatus) SetLastWebhook(v ItemStatusLastWebhook) {
	o.LastWebhook.Set(&v)
}
// SetLastWebhookNil sets the value for LastWebhook to be an explicit nil
func (o *ItemStatus) SetLastWebhookNil() {
	o.LastWebhook.Set(nil)
}

// UnsetLastWebhook ensures that no value is present for LastWebhook, not even an explicit nil
func (o *ItemStatus) UnsetLastWebhook() {
	o.LastWebhook.Unset()
}

func (o ItemStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Investments.IsSet() {
		toSerialize["investments"] = o.Investments.Get()
	}
	if o.Transactions.IsSet() {
		toSerialize["transactions"] = o.Transactions.Get()
	}
	if o.LastWebhook.IsSet() {
		toSerialize["last_webhook"] = o.LastWebhook.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemStatus) UnmarshalJSON(bytes []byte) (err error) {
	varItemStatus := _ItemStatus{}

	if err = json.Unmarshal(bytes, &varItemStatus); err == nil {
		*o = ItemStatus(varItemStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "investments")
		delete(additionalProperties, "transactions")
		delete(additionalProperties, "last_webhook")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemStatus struct {
	value *ItemStatus
	isSet bool
}

func (v NullableItemStatus) Get() *ItemStatus {
	return v.value
}

func (v *NullableItemStatus) Set(val *ItemStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableItemStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableItemStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemStatus(val *ItemStatus) *NullableItemStatus {
	return &NullableItemStatus{value: val, isSet: true}
}

func (v NullableItemStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


