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

// IncomeVerificationCreateRequest IncomeVerificationCreateRequest defines the request schema for `/income/verification/create`
type IncomeVerificationCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The URL endpoint to which Plaid should send webhooks related to the progress of the income verification process.
	Webhook string `json:"webhook"`
	// The ID of a precheck created with `/income/verification/precheck`. Will be used to improve conversion of the income verification flow.
	PrecheckId *string `json:"precheck_id,omitempty"`
	Options *IncomeVerificationCreateRequestOptions `json:"options,omitempty"`
}

// NewIncomeVerificationCreateRequest instantiates a new IncomeVerificationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationCreateRequest(webhook string) *IncomeVerificationCreateRequest {
	this := IncomeVerificationCreateRequest{}
	this.Webhook = webhook
	return &this
}

// NewIncomeVerificationCreateRequestWithDefaults instantiates a new IncomeVerificationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationCreateRequestWithDefaults() *IncomeVerificationCreateRequest {
	this := IncomeVerificationCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IncomeVerificationCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IncomeVerificationCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IncomeVerificationCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IncomeVerificationCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IncomeVerificationCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IncomeVerificationCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetWebhook returns the Webhook field value
func (o *IncomeVerificationCreateRequest) GetWebhook() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationCreateRequest) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Webhook, true
}

// SetWebhook sets field value
func (o *IncomeVerificationCreateRequest) SetWebhook(v string) {
	o.Webhook = v
}

// GetPrecheckId returns the PrecheckId field value if set, zero value otherwise.
func (o *IncomeVerificationCreateRequest) GetPrecheckId() string {
	if o == nil || o.PrecheckId == nil {
		var ret string
		return ret
	}
	return *o.PrecheckId
}

// GetPrecheckIdOk returns a tuple with the PrecheckId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationCreateRequest) GetPrecheckIdOk() (*string, bool) {
	if o == nil || o.PrecheckId == nil {
		return nil, false
	}
	return o.PrecheckId, true
}

// HasPrecheckId returns a boolean if a field has been set.
func (o *IncomeVerificationCreateRequest) HasPrecheckId() bool {
	if o != nil && o.PrecheckId != nil {
		return true
	}

	return false
}

// SetPrecheckId gets a reference to the given string and assigns it to the PrecheckId field.
func (o *IncomeVerificationCreateRequest) SetPrecheckId(v string) {
	o.PrecheckId = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *IncomeVerificationCreateRequest) GetOptions() IncomeVerificationCreateRequestOptions {
	if o == nil || o.Options == nil {
		var ret IncomeVerificationCreateRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationCreateRequest) GetOptionsOk() (*IncomeVerificationCreateRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *IncomeVerificationCreateRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given IncomeVerificationCreateRequestOptions and assigns it to the Options field.
func (o *IncomeVerificationCreateRequest) SetOptions(v IncomeVerificationCreateRequestOptions) {
	o.Options = &v
}

func (o IncomeVerificationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["webhook"] = o.Webhook
	}
	if o.PrecheckId != nil {
		toSerialize["precheck_id"] = o.PrecheckId
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationCreateRequest struct {
	value *IncomeVerificationCreateRequest
	isSet bool
}

func (v NullableIncomeVerificationCreateRequest) Get() *IncomeVerificationCreateRequest {
	return v.value
}

func (v *NullableIncomeVerificationCreateRequest) Set(val *IncomeVerificationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationCreateRequest(val *IncomeVerificationCreateRequest) *NullableIncomeVerificationCreateRequest {
	return &NullableIncomeVerificationCreateRequest{value: val, isSet: true}
}

func (v NullableIncomeVerificationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


