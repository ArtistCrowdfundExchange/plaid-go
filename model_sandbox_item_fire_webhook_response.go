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

// SandboxItemFireWebhookResponse SandboxItemFireWebhookResponse defines the response schema for `/sandbox/item/fire_webhook`
type SandboxItemFireWebhookResponse struct {
	// Value is `true`  if the test` webhook_code`  was successfully fired.
	WebhookFired bool `json:"webhook_fired"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxItemFireWebhookResponse SandboxItemFireWebhookResponse

// NewSandboxItemFireWebhookResponse instantiates a new SandboxItemFireWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxItemFireWebhookResponse(webhookFired bool, requestId string) *SandboxItemFireWebhookResponse {
	this := SandboxItemFireWebhookResponse{}
	this.WebhookFired = webhookFired
	this.RequestId = requestId
	return &this
}

// NewSandboxItemFireWebhookResponseWithDefaults instantiates a new SandboxItemFireWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxItemFireWebhookResponseWithDefaults() *SandboxItemFireWebhookResponse {
	this := SandboxItemFireWebhookResponse{}
	return &this
}

// GetWebhookFired returns the WebhookFired field value
func (o *SandboxItemFireWebhookResponse) GetWebhookFired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WebhookFired
}

// GetWebhookFiredOk returns a tuple with the WebhookFired field value
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookResponse) GetWebhookFiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookFired, true
}

// SetWebhookFired sets field value
func (o *SandboxItemFireWebhookResponse) SetWebhookFired(v bool) {
	o.WebhookFired = v
}

// GetRequestId returns the RequestId field value
func (o *SandboxItemFireWebhookResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxItemFireWebhookResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxItemFireWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_fired"] = o.WebhookFired
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxItemFireWebhookResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxItemFireWebhookResponse := _SandboxItemFireWebhookResponse{}

	if err = json.Unmarshal(bytes, &varSandboxItemFireWebhookResponse); err == nil {
		*o = SandboxItemFireWebhookResponse(varSandboxItemFireWebhookResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_fired")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxItemFireWebhookResponse struct {
	value *SandboxItemFireWebhookResponse
	isSet bool
}

func (v NullableSandboxItemFireWebhookResponse) Get() *SandboxItemFireWebhookResponse {
	return v.value
}

func (v *NullableSandboxItemFireWebhookResponse) Set(val *SandboxItemFireWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxItemFireWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxItemFireWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxItemFireWebhookResponse(val *SandboxItemFireWebhookResponse) *NullableSandboxItemFireWebhookResponse {
	return &NullableSandboxItemFireWebhookResponse{value: val, isSet: true}
}

func (v NullableSandboxItemFireWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxItemFireWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


