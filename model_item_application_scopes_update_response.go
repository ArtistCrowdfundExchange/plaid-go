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

// ItemApplicationScopesUpdateResponse ItemApplicationScopesUpdateResponse defines the response schema for `/item/application/scopes/update`
type ItemApplicationScopesUpdateResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
}

// NewItemApplicationScopesUpdateResponse instantiates a new ItemApplicationScopesUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemApplicationScopesUpdateResponse(requestId string) *ItemApplicationScopesUpdateResponse {
	this := ItemApplicationScopesUpdateResponse{}
	this.RequestId = requestId
	return &this
}

// NewItemApplicationScopesUpdateResponseWithDefaults instantiates a new ItemApplicationScopesUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemApplicationScopesUpdateResponseWithDefaults() *ItemApplicationScopesUpdateResponse {
	this := ItemApplicationScopesUpdateResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ItemApplicationScopesUpdateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ItemApplicationScopesUpdateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ItemApplicationScopesUpdateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ItemApplicationScopesUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullableItemApplicationScopesUpdateResponse struct {
	value *ItemApplicationScopesUpdateResponse
	isSet bool
}

func (v NullableItemApplicationScopesUpdateResponse) Get() *ItemApplicationScopesUpdateResponse {
	return v.value
}

func (v *NullableItemApplicationScopesUpdateResponse) Set(val *ItemApplicationScopesUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableItemApplicationScopesUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableItemApplicationScopesUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemApplicationScopesUpdateResponse(val *ItemApplicationScopesUpdateResponse) *NullableItemApplicationScopesUpdateResponse {
	return &NullableItemApplicationScopesUpdateResponse{value: val, isSet: true}
}

func (v NullableItemApplicationScopesUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemApplicationScopesUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


