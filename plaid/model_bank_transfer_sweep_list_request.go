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
	"time"
)

// BankTransferSweepListRequest BankTransferSweepListRequest defines the request schema for `/bank_transfer/sweep/list`
type BankTransferSweepListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// If multiple origination accounts are available, `origination_account_id` must be used to specify the account that the sweeps belong to.
	OriginationAccountId NullableString `json:"origination_account_id,omitempty"`
	// Starting ID of sweeps to return.
	StartId NullableInt32 `json:"start_id,omitempty"`
	// The start datetime of sweeps to return (RFC 3339 format).
	StartTime NullableTime `json:"start_time,omitempty"`
	// The end datetime of sweeps to return (RFC 3339 format).
	EndTime NullableTime `json:"end_time,omitempty"`
	// The maximum number of sweeps to return.
	Count NullableInt32 `json:"count,omitempty"`
}

// NewBankTransferSweepListRequest instantiates a new BankTransferSweepListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferSweepListRequest() *BankTransferSweepListRequest {
	this := BankTransferSweepListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	return &this
}

// NewBankTransferSweepListRequestWithDefaults instantiates a new BankTransferSweepListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferSweepListRequestWithDefaults() *BankTransferSweepListRequest {
	this := BankTransferSweepListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BankTransferSweepListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferSweepListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BankTransferSweepListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BankTransferSweepListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *BankTransferSweepListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferSweepListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *BankTransferSweepListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *BankTransferSweepListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetOriginationAccountId returns the OriginationAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferSweepListRequest) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferSweepListRequest) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// HasOriginationAccountId returns a boolean if a field has been set.
func (o *BankTransferSweepListRequest) HasOriginationAccountId() bool {
	if o != nil && o.OriginationAccountId.IsSet() {
		return true
	}

	return false
}

// SetOriginationAccountId gets a reference to the given NullableString and assigns it to the OriginationAccountId field.
func (o *BankTransferSweepListRequest) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}
// SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil
func (o *BankTransferSweepListRequest) SetOriginationAccountIdNil() {
	o.OriginationAccountId.Set(nil)
}

// UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
func (o *BankTransferSweepListRequest) UnsetOriginationAccountId() {
	o.OriginationAccountId.Unset()
}

// GetStartId returns the StartId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferSweepListRequest) GetStartId() int32 {
	if o == nil || o.StartId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.StartId.Get()
}

// GetStartIdOk returns a tuple with the StartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferSweepListRequest) GetStartIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartId.Get(), o.StartId.IsSet()
}

// HasStartId returns a boolean if a field has been set.
func (o *BankTransferSweepListRequest) HasStartId() bool {
	if o != nil && o.StartId.IsSet() {
		return true
	}

	return false
}

// SetStartId gets a reference to the given NullableInt32 and assigns it to the StartId field.
func (o *BankTransferSweepListRequest) SetStartId(v int32) {
	o.StartId.Set(&v)
}
// SetStartIdNil sets the value for StartId to be an explicit nil
func (o *BankTransferSweepListRequest) SetStartIdNil() {
	o.StartId.Set(nil)
}

// UnsetStartId ensures that no value is present for StartId, not even an explicit nil
func (o *BankTransferSweepListRequest) UnsetStartId() {
	o.StartId.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferSweepListRequest) GetStartTime() time.Time {
	if o == nil || o.StartTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferSweepListRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *BankTransferSweepListRequest) HasStartTime() bool {
	if o != nil && o.StartTime.IsSet() {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given NullableTime and assigns it to the StartTime field.
func (o *BankTransferSweepListRequest) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}
// SetStartTimeNil sets the value for StartTime to be an explicit nil
func (o *BankTransferSweepListRequest) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
func (o *BankTransferSweepListRequest) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferSweepListRequest) GetEndTime() time.Time {
	if o == nil || o.EndTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferSweepListRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// HasEndTime returns a boolean if a field has been set.
func (o *BankTransferSweepListRequest) HasEndTime() bool {
	if o != nil && o.EndTime.IsSet() {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given NullableTime and assigns it to the EndTime field.
func (o *BankTransferSweepListRequest) SetEndTime(v time.Time) {
	o.EndTime.Set(&v)
}
// SetEndTimeNil sets the value for EndTime to be an explicit nil
func (o *BankTransferSweepListRequest) SetEndTimeNil() {
	o.EndTime.Set(nil)
}

// UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
func (o *BankTransferSweepListRequest) UnsetEndTime() {
	o.EndTime.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferSweepListRequest) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferSweepListRequest) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *BankTransferSweepListRequest) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *BankTransferSweepListRequest) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *BankTransferSweepListRequest) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *BankTransferSweepListRequest) UnsetCount() {
	o.Count.Unset()
}

func (o BankTransferSweepListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.OriginationAccountId.IsSet() {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	if o.StartId.IsSet() {
		toSerialize["start_id"] = o.StartId.Get()
	}
	if o.StartTime.IsSet() {
		toSerialize["start_time"] = o.StartTime.Get()
	}
	if o.EndTime.IsSet() {
		toSerialize["end_time"] = o.EndTime.Get()
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBankTransferSweepListRequest struct {
	value *BankTransferSweepListRequest
	isSet bool
}

func (v NullableBankTransferSweepListRequest) Get() *BankTransferSweepListRequest {
	return v.value
}

func (v *NullableBankTransferSweepListRequest) Set(val *BankTransferSweepListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferSweepListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferSweepListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferSweepListRequest(val *BankTransferSweepListRequest) *NullableBankTransferSweepListRequest {
	return &NullableBankTransferSweepListRequest{value: val, isSet: true}
}

func (v NullableBankTransferSweepListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferSweepListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


