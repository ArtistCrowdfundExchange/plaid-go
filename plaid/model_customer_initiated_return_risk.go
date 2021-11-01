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

// CustomerInitiatedReturnRisk The object contains a risk score and a risk tier that evaluate the transaction return risk of an unauthorized debit. Common return codes in this category include: “R05”, \"R07\", \"R10\", \"R11\", \"R29\". These returns typically have a return time frame of up to 60 calendar days. During this period, the customer of financial institutions can dispute a transaction as unauthorized.
type CustomerInitiatedReturnRisk struct {
	// A score from 0-99 that indicates the transaction return risk: a higher risk score suggests a higher return likelihood.
	Score int32 `json:"score"`
	// A tier corresponding to the projected likelihood that the transaction, if initiated, will be subject to a return.  In the `customer_initiated_return_risk` object, there are five risk tiers corresponding to the scores:   1: Predicted customer-initiated return incidence rate between 0.00% - 0.02%   2: Predicted customer-initiated return incidence rate between 0.02% - 0.05%   3: Predicted customer-initiated return incidence rate between 0.05% - 0.1%   4: Predicted customer-initiated return incidence rate between 0.1% - 0.5%   5: Predicted customer-initiated return incidence rate greater than 0.5% 
	RiskTier int32 `json:"risk_tier"`
}

// NewCustomerInitiatedReturnRisk instantiates a new CustomerInitiatedReturnRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerInitiatedReturnRisk(score int32, riskTier int32) *CustomerInitiatedReturnRisk {
	this := CustomerInitiatedReturnRisk{}
	this.Score = score
	this.RiskTier = riskTier
	return &this
}

// NewCustomerInitiatedReturnRiskWithDefaults instantiates a new CustomerInitiatedReturnRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerInitiatedReturnRiskWithDefaults() *CustomerInitiatedReturnRisk {
	this := CustomerInitiatedReturnRisk{}
	return &this
}

// GetScore returns the Score field value
func (o *CustomerInitiatedReturnRisk) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *CustomerInitiatedReturnRisk) GetScoreOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *CustomerInitiatedReturnRisk) SetScore(v int32) {
	o.Score = v
}

// GetRiskTier returns the RiskTier field value
func (o *CustomerInitiatedReturnRisk) GetRiskTier() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RiskTier
}

// GetRiskTierOk returns a tuple with the RiskTier field value
// and a boolean to check if the value has been set.
func (o *CustomerInitiatedReturnRisk) GetRiskTierOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RiskTier, true
}

// SetRiskTier sets field value
func (o *CustomerInitiatedReturnRisk) SetRiskTier(v int32) {
	o.RiskTier = v
}

func (o CustomerInitiatedReturnRisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["score"] = o.Score
	}
	if true {
		toSerialize["risk_tier"] = o.RiskTier
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerInitiatedReturnRisk struct {
	value *CustomerInitiatedReturnRisk
	isSet bool
}

func (v NullableCustomerInitiatedReturnRisk) Get() *CustomerInitiatedReturnRisk {
	return v.value
}

func (v *NullableCustomerInitiatedReturnRisk) Set(val *CustomerInitiatedReturnRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerInitiatedReturnRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerInitiatedReturnRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerInitiatedReturnRisk(val *CustomerInitiatedReturnRisk) *NullableCustomerInitiatedReturnRisk {
	return &NullableCustomerInitiatedReturnRisk{value: val, isSet: true}
}

func (v NullableCustomerInitiatedReturnRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerInitiatedReturnRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


