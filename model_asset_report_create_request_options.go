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

// AssetReportCreateRequestOptions An optional object to filter `/asset_report/create` results. If provided, must be non-`null`. The optional `user` object is required for the report to be eligible for Fannie Mae's Day 1 Certainty program.
type AssetReportCreateRequestOptions struct {
	// Client-generated identifier, which can be used by lenders to track loan applications.
	ClientReportId *string `json:"client_report_id,omitempty"`
	// URL to which Plaid will send Assets webhooks, for example when the requested Asset Report is ready.
	Webhook *string `json:"webhook,omitempty"`
	User *AssetReportUser `json:"user,omitempty"`
}

// NewAssetReportCreateRequestOptions instantiates a new AssetReportCreateRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportCreateRequestOptions() *AssetReportCreateRequestOptions {
	this := AssetReportCreateRequestOptions{}
	return &this
}

// NewAssetReportCreateRequestOptionsWithDefaults instantiates a new AssetReportCreateRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportCreateRequestOptionsWithDefaults() *AssetReportCreateRequestOptions {
	this := AssetReportCreateRequestOptions{}
	return &this
}

// GetClientReportId returns the ClientReportId field value if set, zero value otherwise.
func (o *AssetReportCreateRequestOptions) GetClientReportId() string {
	if o == nil || o.ClientReportId == nil {
		var ret string
		return ret
	}
	return *o.ClientReportId
}

// GetClientReportIdOk returns a tuple with the ClientReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequestOptions) GetClientReportIdOk() (*string, bool) {
	if o == nil || o.ClientReportId == nil {
		return nil, false
	}
	return o.ClientReportId, true
}

// HasClientReportId returns a boolean if a field has been set.
func (o *AssetReportCreateRequestOptions) HasClientReportId() bool {
	if o != nil && o.ClientReportId != nil {
		return true
	}

	return false
}

// SetClientReportId gets a reference to the given string and assigns it to the ClientReportId field.
func (o *AssetReportCreateRequestOptions) SetClientReportId(v string) {
	o.ClientReportId = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *AssetReportCreateRequestOptions) GetWebhook() string {
	if o == nil || o.Webhook == nil {
		var ret string
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequestOptions) GetWebhookOk() (*string, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *AssetReportCreateRequestOptions) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given string and assigns it to the Webhook field.
func (o *AssetReportCreateRequestOptions) SetWebhook(v string) {
	o.Webhook = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AssetReportCreateRequestOptions) GetUser() AssetReportUser {
	if o == nil || o.User == nil {
		var ret AssetReportUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequestOptions) GetUserOk() (*AssetReportUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AssetReportCreateRequestOptions) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given AssetReportUser and assigns it to the User field.
func (o *AssetReportCreateRequestOptions) SetUser(v AssetReportUser) {
	o.User = &v
}

func (o AssetReportCreateRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientReportId != nil {
		toSerialize["client_report_id"] = o.ClientReportId
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportCreateRequestOptions struct {
	value *AssetReportCreateRequestOptions
	isSet bool
}

func (v NullableAssetReportCreateRequestOptions) Get() *AssetReportCreateRequestOptions {
	return v.value
}

func (v *NullableAssetReportCreateRequestOptions) Set(val *AssetReportCreateRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportCreateRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportCreateRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportCreateRequestOptions(val *AssetReportCreateRequestOptions) *NullableAssetReportCreateRequestOptions {
	return &NullableAssetReportCreateRequestOptions{value: val, isSet: true}
}

func (v NullableAssetReportCreateRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportCreateRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


