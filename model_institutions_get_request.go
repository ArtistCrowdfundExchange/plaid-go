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
)

// InstitutionsGetRequest InstitutionsGetRequest defines the request schema for `/institutions/get`
type InstitutionsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The total number of Institutions to return.
	Count int32 `json:"count"`
	// The number of Institutions to skip.
	Offset int32 `json:"offset"`
	// Specify an array of Plaid-supported country codes this institution supports, using the ISO-3166-1 alpha-2 country code standard. 
	CountryCodes []CountryCode `json:"country_codes"`
	Options *InstitutionsGetRequestOptions `json:"options,omitempty"`
}

// NewInstitutionsGetRequest instantiates a new InstitutionsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsGetRequest(count int32, offset int32, countryCodes []CountryCode) *InstitutionsGetRequest {
	this := InstitutionsGetRequest{}
	this.Count = count
	this.Offset = offset
	this.CountryCodes = countryCodes
	return &this
}

// NewInstitutionsGetRequestWithDefaults instantiates a new InstitutionsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsGetRequestWithDefaults() *InstitutionsGetRequest {
	this := InstitutionsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InstitutionsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InstitutionsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InstitutionsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *InstitutionsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *InstitutionsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *InstitutionsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetCount returns the Count field value
func (o *InstitutionsGetRequest) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *InstitutionsGetRequest) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *InstitutionsGetRequest) SetCount(v int32) {
	o.Count = v
}

// GetOffset returns the Offset field value
func (o *InstitutionsGetRequest) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *InstitutionsGetRequest) GetOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *InstitutionsGetRequest) SetOffset(v int32) {
	o.Offset = v
}

// GetCountryCodes returns the CountryCodes field value
func (o *InstitutionsGetRequest) GetCountryCodes() []CountryCode {
	if o == nil {
		var ret []CountryCode
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *InstitutionsGetRequest) GetCountryCodesOk() (*[]CountryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *InstitutionsGetRequest) SetCountryCodes(v []CountryCode) {
	o.CountryCodes = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InstitutionsGetRequest) GetOptions() InstitutionsGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret InstitutionsGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsGetRequest) GetOptionsOk() (*InstitutionsGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InstitutionsGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given InstitutionsGetRequestOptions and assigns it to the Options field.
func (o *InstitutionsGetRequest) SetOptions(v InstitutionsGetRequestOptions) {
	o.Options = &v
}

func (o InstitutionsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["country_codes"] = o.CountryCodes
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionsGetRequest struct {
	value *InstitutionsGetRequest
	isSet bool
}

func (v NullableInstitutionsGetRequest) Get() *InstitutionsGetRequest {
	return v.value
}

func (v *NullableInstitutionsGetRequest) Set(val *InstitutionsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsGetRequest(val *InstitutionsGetRequest) *NullableInstitutionsGetRequest {
	return &NullableInstitutionsGetRequest{value: val, isSet: true}
}

func (v NullableInstitutionsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


