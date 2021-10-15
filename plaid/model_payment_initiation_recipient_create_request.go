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

// PaymentInitiationRecipientCreateRequest PaymentInitiationRecipientCreateRequest defines the request schema for `/payment_initiation/recipient/create`
type PaymentInitiationRecipientCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The name of the recipient
	Name string `json:"name"`
	// The International Bank Account Number (IBAN) for the recipient. If BACS data is not provided, an IBAN is required.
	Iban NullableString `json:"iban,omitempty"`
	Bacs NullableRecipientBACSNullable `json:"bacs,omitempty"`
	Address NullablePaymentInitiationAddress `json:"address,omitempty"`
}

// NewPaymentInitiationRecipientCreateRequest instantiates a new PaymentInitiationRecipientCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationRecipientCreateRequest(name string) *PaymentInitiationRecipientCreateRequest {
	this := PaymentInitiationRecipientCreateRequest{}
	this.Name = name
	return &this
}

// NewPaymentInitiationRecipientCreateRequestWithDefaults instantiates a new PaymentInitiationRecipientCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationRecipientCreateRequestWithDefaults() *PaymentInitiationRecipientCreateRequest {
	this := PaymentInitiationRecipientCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationRecipientCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationRecipientCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationRecipientCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationRecipientCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetName returns the Name field value
func (o *PaymentInitiationRecipientCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientCreateRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaymentInitiationRecipientCreateRequest) SetName(v string) {
	o.Name = v
}

// GetIban returns the Iban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationRecipientCreateRequest) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}
	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationRecipientCreateRequest) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// HasIban returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientCreateRequest) HasIban() bool {
	if o != nil && o.Iban.IsSet() {
		return true
	}

	return false
}

// SetIban gets a reference to the given NullableString and assigns it to the Iban field.
func (o *PaymentInitiationRecipientCreateRequest) SetIban(v string) {
	o.Iban.Set(&v)
}
// SetIbanNil sets the value for Iban to be an explicit nil
func (o *PaymentInitiationRecipientCreateRequest) SetIbanNil() {
	o.Iban.Set(nil)
}

// UnsetIban ensures that no value is present for Iban, not even an explicit nil
func (o *PaymentInitiationRecipientCreateRequest) UnsetIban() {
	o.Iban.Unset()
}

// GetBacs returns the Bacs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationRecipientCreateRequest) GetBacs() RecipientBACSNullable {
	if o == nil || o.Bacs.Get() == nil {
		var ret RecipientBACSNullable
		return ret
	}
	return *o.Bacs.Get()
}

// GetBacsOk returns a tuple with the Bacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationRecipientCreateRequest) GetBacsOk() (*RecipientBACSNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bacs.Get(), o.Bacs.IsSet()
}

// HasBacs returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientCreateRequest) HasBacs() bool {
	if o != nil && o.Bacs.IsSet() {
		return true
	}

	return false
}

// SetBacs gets a reference to the given NullableRecipientBACSNullable and assigns it to the Bacs field.
func (o *PaymentInitiationRecipientCreateRequest) SetBacs(v RecipientBACSNullable) {
	o.Bacs.Set(&v)
}
// SetBacsNil sets the value for Bacs to be an explicit nil
func (o *PaymentInitiationRecipientCreateRequest) SetBacsNil() {
	o.Bacs.Set(nil)
}

// UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
func (o *PaymentInitiationRecipientCreateRequest) UnsetBacs() {
	o.Bacs.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationRecipientCreateRequest) GetAddress() PaymentInitiationAddress {
	if o == nil || o.Address.Get() == nil {
		var ret PaymentInitiationAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationRecipientCreateRequest) GetAddressOk() (*PaymentInitiationAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientCreateRequest) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullablePaymentInitiationAddress and assigns it to the Address field.
func (o *PaymentInitiationRecipientCreateRequest) SetAddress(v PaymentInitiationAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *PaymentInitiationRecipientCreateRequest) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *PaymentInitiationRecipientCreateRequest) UnsetAddress() {
	o.Address.Unset()
}

func (o PaymentInitiationRecipientCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Iban.IsSet() {
		toSerialize["iban"] = o.Iban.Get()
	}
	if o.Bacs.IsSet() {
		toSerialize["bacs"] = o.Bacs.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationRecipientCreateRequest struct {
	value *PaymentInitiationRecipientCreateRequest
	isSet bool
}

func (v NullablePaymentInitiationRecipientCreateRequest) Get() *PaymentInitiationRecipientCreateRequest {
	return v.value
}

func (v *NullablePaymentInitiationRecipientCreateRequest) Set(val *PaymentInitiationRecipientCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRecipientCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRecipientCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRecipientCreateRequest(val *PaymentInitiationRecipientCreateRequest) *NullablePaymentInitiationRecipientCreateRequest {
	return &NullablePaymentInitiationRecipientCreateRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationRecipientCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRecipientCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


