/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.20.6
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SignalDecisionReportResponse struct for SignalDecisionReportResponse
type SignalDecisionReportResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
}

// NewSignalDecisionReportResponse instantiates a new SignalDecisionReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalDecisionReportResponse(requestId string) *SignalDecisionReportResponse {
	this := SignalDecisionReportResponse{}
	this.RequestId = requestId
	return &this
}

// NewSignalDecisionReportResponseWithDefaults instantiates a new SignalDecisionReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalDecisionReportResponseWithDefaults() *SignalDecisionReportResponse {
	this := SignalDecisionReportResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SignalDecisionReportResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SignalDecisionReportResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SignalDecisionReportResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SignalDecisionReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullableSignalDecisionReportResponse struct {
	value *SignalDecisionReportResponse
	isSet bool
}

func (v NullableSignalDecisionReportResponse) Get() *SignalDecisionReportResponse {
	return v.value
}

func (v *NullableSignalDecisionReportResponse) Set(val *SignalDecisionReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalDecisionReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalDecisionReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalDecisionReportResponse(val *SignalDecisionReportResponse) *NullableSignalDecisionReportResponse {
	return &NullableSignalDecisionReportResponse{value: val, isSet: true}
}

func (v NullableSignalDecisionReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalDecisionReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


