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

// TransactionsSyncResponse TransactionsSyncResponse defines the response schema for `/transactions/sync`
type TransactionsSyncResponse struct {
	// Transactions that have been added to the item since `cursor` ordered by ascending last modified time.
	Added []Transaction `json:"added"`
	// Transactions that have been modified on the item since `cursor` ordered by ascending last modified time.
	Modified []Transaction `json:"modified"`
	// Transactions that have been removed from the item since `cursor` ordered by ascending last modified time.
	Removed []RemovedTransaction `json:"removed"`
	// Cursor used for fetching any future updates after the latest update provided in this response.
	NextCursor string `json:"next_cursor"`
	// Represents if more than requested count of transaction updates exist. If true, the additional updates can be fetched by making an additional request with `cursor` set to `next_cursor`.
	HasMore bool `json:"has_more"`
}

// NewTransactionsSyncResponse instantiates a new TransactionsSyncResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsSyncResponse(added []Transaction, modified []Transaction, removed []RemovedTransaction, nextCursor string, hasMore bool) *TransactionsSyncResponse {
	this := TransactionsSyncResponse{}
	this.Added = added
	this.Modified = modified
	this.Removed = removed
	this.NextCursor = nextCursor
	this.HasMore = hasMore
	return &this
}

// NewTransactionsSyncResponseWithDefaults instantiates a new TransactionsSyncResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsSyncResponseWithDefaults() *TransactionsSyncResponse {
	this := TransactionsSyncResponse{}
	return &this
}

// GetAdded returns the Added field value
func (o *TransactionsSyncResponse) GetAdded() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *TransactionsSyncResponse) GetAddedOk() (*[]Transaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value
func (o *TransactionsSyncResponse) SetAdded(v []Transaction) {
	o.Added = v
}

// GetModified returns the Modified field value
func (o *TransactionsSyncResponse) GetModified() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *TransactionsSyncResponse) GetModifiedOk() (*[]Transaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *TransactionsSyncResponse) SetModified(v []Transaction) {
	o.Modified = v
}

// GetRemoved returns the Removed field value
func (o *TransactionsSyncResponse) GetRemoved() []RemovedTransaction {
	if o == nil {
		var ret []RemovedTransaction
		return ret
	}

	return o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value
// and a boolean to check if the value has been set.
func (o *TransactionsSyncResponse) GetRemovedOk() (*[]RemovedTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Removed, true
}

// SetRemoved sets field value
func (o *TransactionsSyncResponse) SetRemoved(v []RemovedTransaction) {
	o.Removed = v
}

// GetNextCursor returns the NextCursor field value
func (o *TransactionsSyncResponse) GetNextCursor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
func (o *TransactionsSyncResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NextCursor, true
}

// SetNextCursor sets field value
func (o *TransactionsSyncResponse) SetNextCursor(v string) {
	o.NextCursor = v
}

// GetHasMore returns the HasMore field value
func (o *TransactionsSyncResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *TransactionsSyncResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *TransactionsSyncResponse) SetHasMore(v bool) {
	o.HasMore = v
}

func (o TransactionsSyncResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["added"] = o.Added
	}
	if true {
		toSerialize["modified"] = o.Modified
	}
	if true {
		toSerialize["removed"] = o.Removed
	}
	if true {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if true {
		toSerialize["has_more"] = o.HasMore
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsSyncResponse struct {
	value *TransactionsSyncResponse
	isSet bool
}

func (v NullableTransactionsSyncResponse) Get() *TransactionsSyncResponse {
	return v.value
}

func (v *NullableTransactionsSyncResponse) Set(val *TransactionsSyncResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsSyncResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsSyncResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsSyncResponse(val *TransactionsSyncResponse) *NullableTransactionsSyncResponse {
	return &NullableTransactionsSyncResponse{value: val, isSet: true}
}

func (v NullableTransactionsSyncResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsSyncResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


