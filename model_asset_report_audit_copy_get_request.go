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

// AssetReportAuditCopyGetRequest AssetReportAuditCopyGetRequest defines the request schema for `/asset_report/audit_copy/get`
type AssetReportAuditCopyGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The `audit_copy_token` granting access to the Audit Copy you would like to get.
	AuditCopyToken string `json:"audit_copy_token"`
}

// NewAssetReportAuditCopyGetRequest instantiates a new AssetReportAuditCopyGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportAuditCopyGetRequest(auditCopyToken string) *AssetReportAuditCopyGetRequest {
	this := AssetReportAuditCopyGetRequest{}
	this.AuditCopyToken = auditCopyToken
	return &this
}

// NewAssetReportAuditCopyGetRequestWithDefaults instantiates a new AssetReportAuditCopyGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportAuditCopyGetRequestWithDefaults() *AssetReportAuditCopyGetRequest {
	this := AssetReportAuditCopyGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportAuditCopyGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportAuditCopyGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportAuditCopyGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportAuditCopyGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportAuditCopyGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportAuditCopyGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAuditCopyToken returns the AuditCopyToken field value
func (o *AssetReportAuditCopyGetRequest) GetAuditCopyToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditCopyToken
}

// GetAuditCopyTokenOk returns a tuple with the AuditCopyToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportAuditCopyGetRequest) GetAuditCopyTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditCopyToken, true
}

// SetAuditCopyToken sets field value
func (o *AssetReportAuditCopyGetRequest) SetAuditCopyToken(v string) {
	o.AuditCopyToken = v
}

func (o AssetReportAuditCopyGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["audit_copy_token"] = o.AuditCopyToken
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportAuditCopyGetRequest struct {
	value *AssetReportAuditCopyGetRequest
	isSet bool
}

func (v NullableAssetReportAuditCopyGetRequest) Get() *AssetReportAuditCopyGetRequest {
	return v.value
}

func (v *NullableAssetReportAuditCopyGetRequest) Set(val *AssetReportAuditCopyGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportAuditCopyGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportAuditCopyGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportAuditCopyGetRequest(val *AssetReportAuditCopyGetRequest) *NullableAssetReportAuditCopyGetRequest {
	return &NullableAssetReportAuditCopyGetRequest{value: val, isSet: true}
}

func (v NullableAssetReportAuditCopyGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportAuditCopyGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


