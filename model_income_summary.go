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

// IncomeSummary The verified fields from a paystub verification. All fields are provided as reported on the paystub.
type IncomeSummary struct {
	EmployerName EmployerIncomeSummaryFieldString `json:"employer_name"`
	EmployeeName EmployeeIncomeSummaryFieldString `json:"employee_name"`
	YtdGrossIncome YTDGrossIncomeSummaryFieldNumber `json:"ytd_gross_income"`
	YtdNetIncome YTDNetIncomeSummaryFieldNumber `json:"ytd_net_income"`
	PayFrequency NullablePayFrequency `json:"pay_frequency"`
	ProjectedWage ProjectedIncomeSummaryFieldNumber `json:"projected_wage"`
	VerifiedTransaction NullableTransactionData `json:"verified_transaction"`
	AdditionalProperties map[string]interface{}
}

type _IncomeSummary IncomeSummary

// NewIncomeSummary instantiates a new IncomeSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeSummary(employerName EmployerIncomeSummaryFieldString, employeeName EmployeeIncomeSummaryFieldString, ytdGrossIncome YTDGrossIncomeSummaryFieldNumber, ytdNetIncome YTDNetIncomeSummaryFieldNumber, payFrequency NullablePayFrequency, projectedWage ProjectedIncomeSummaryFieldNumber, verifiedTransaction NullableTransactionData) *IncomeSummary {
	this := IncomeSummary{}
	this.EmployerName = employerName
	this.EmployeeName = employeeName
	this.YtdGrossIncome = ytdGrossIncome
	this.YtdNetIncome = ytdNetIncome
	this.PayFrequency = payFrequency
	this.ProjectedWage = projectedWage
	this.VerifiedTransaction = verifiedTransaction
	return &this
}

// NewIncomeSummaryWithDefaults instantiates a new IncomeSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeSummaryWithDefaults() *IncomeSummary {
	this := IncomeSummary{}
	return &this
}

// GetEmployerName returns the EmployerName field value
func (o *IncomeSummary) GetEmployerName() EmployerIncomeSummaryFieldString {
	if o == nil {
		var ret EmployerIncomeSummaryFieldString
		return ret
	}

	return o.EmployerName
}

// GetEmployerNameOk returns a tuple with the EmployerName field value
// and a boolean to check if the value has been set.
func (o *IncomeSummary) GetEmployerNameOk() (*EmployerIncomeSummaryFieldString, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmployerName, true
}

// SetEmployerName sets field value
func (o *IncomeSummary) SetEmployerName(v EmployerIncomeSummaryFieldString) {
	o.EmployerName = v
}

// GetEmployeeName returns the EmployeeName field value
func (o *IncomeSummary) GetEmployeeName() EmployeeIncomeSummaryFieldString {
	if o == nil {
		var ret EmployeeIncomeSummaryFieldString
		return ret
	}

	return o.EmployeeName
}

// GetEmployeeNameOk returns a tuple with the EmployeeName field value
// and a boolean to check if the value has been set.
func (o *IncomeSummary) GetEmployeeNameOk() (*EmployeeIncomeSummaryFieldString, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmployeeName, true
}

// SetEmployeeName sets field value
func (o *IncomeSummary) SetEmployeeName(v EmployeeIncomeSummaryFieldString) {
	o.EmployeeName = v
}

// GetYtdGrossIncome returns the YtdGrossIncome field value
func (o *IncomeSummary) GetYtdGrossIncome() YTDGrossIncomeSummaryFieldNumber {
	if o == nil {
		var ret YTDGrossIncomeSummaryFieldNumber
		return ret
	}

	return o.YtdGrossIncome
}

// GetYtdGrossIncomeOk returns a tuple with the YtdGrossIncome field value
// and a boolean to check if the value has been set.
func (o *IncomeSummary) GetYtdGrossIncomeOk() (*YTDGrossIncomeSummaryFieldNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YtdGrossIncome, true
}

// SetYtdGrossIncome sets field value
func (o *IncomeSummary) SetYtdGrossIncome(v YTDGrossIncomeSummaryFieldNumber) {
	o.YtdGrossIncome = v
}

// GetYtdNetIncome returns the YtdNetIncome field value
func (o *IncomeSummary) GetYtdNetIncome() YTDNetIncomeSummaryFieldNumber {
	if o == nil {
		var ret YTDNetIncomeSummaryFieldNumber
		return ret
	}

	return o.YtdNetIncome
}

// GetYtdNetIncomeOk returns a tuple with the YtdNetIncome field value
// and a boolean to check if the value has been set.
func (o *IncomeSummary) GetYtdNetIncomeOk() (*YTDNetIncomeSummaryFieldNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YtdNetIncome, true
}

// SetYtdNetIncome sets field value
func (o *IncomeSummary) SetYtdNetIncome(v YTDNetIncomeSummaryFieldNumber) {
	o.YtdNetIncome = v
}

// GetPayFrequency returns the PayFrequency field value
// If the value is explicit nil, the zero value for PayFrequency will be returned
func (o *IncomeSummary) GetPayFrequency() PayFrequency {
	if o == nil || o.PayFrequency.Get() == nil {
		var ret PayFrequency
		return ret
	}

	return *o.PayFrequency.Get()
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeSummary) GetPayFrequencyOk() (*PayFrequency, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayFrequency.Get(), o.PayFrequency.IsSet()
}

// SetPayFrequency sets field value
func (o *IncomeSummary) SetPayFrequency(v PayFrequency) {
	o.PayFrequency.Set(&v)
}

// GetProjectedWage returns the ProjectedWage field value
func (o *IncomeSummary) GetProjectedWage() ProjectedIncomeSummaryFieldNumber {
	if o == nil {
		var ret ProjectedIncomeSummaryFieldNumber
		return ret
	}

	return o.ProjectedWage
}

// GetProjectedWageOk returns a tuple with the ProjectedWage field value
// and a boolean to check if the value has been set.
func (o *IncomeSummary) GetProjectedWageOk() (*ProjectedIncomeSummaryFieldNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectedWage, true
}

// SetProjectedWage sets field value
func (o *IncomeSummary) SetProjectedWage(v ProjectedIncomeSummaryFieldNumber) {
	o.ProjectedWage = v
}

// GetVerifiedTransaction returns the VerifiedTransaction field value
// If the value is explicit nil, the zero value for TransactionData will be returned
func (o *IncomeSummary) GetVerifiedTransaction() TransactionData {
	if o == nil || o.VerifiedTransaction.Get() == nil {
		var ret TransactionData
		return ret
	}

	return *o.VerifiedTransaction.Get()
}

// GetVerifiedTransactionOk returns a tuple with the VerifiedTransaction field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeSummary) GetVerifiedTransactionOk() (*TransactionData, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VerifiedTransaction.Get(), o.VerifiedTransaction.IsSet()
}

// SetVerifiedTransaction sets field value
func (o *IncomeSummary) SetVerifiedTransaction(v TransactionData) {
	o.VerifiedTransaction.Set(&v)
}

func (o IncomeSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["employer_name"] = o.EmployerName
	}
	if true {
		toSerialize["employee_name"] = o.EmployeeName
	}
	if true {
		toSerialize["ytd_gross_income"] = o.YtdGrossIncome
	}
	if true {
		toSerialize["ytd_net_income"] = o.YtdNetIncome
	}
	if true {
		toSerialize["pay_frequency"] = o.PayFrequency.Get()
	}
	if true {
		toSerialize["projected_wage"] = o.ProjectedWage
	}
	if true {
		toSerialize["verified_transaction"] = o.VerifiedTransaction.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeSummary) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeSummary := _IncomeSummary{}

	if err = json.Unmarshal(bytes, &varIncomeSummary); err == nil {
		*o = IncomeSummary(varIncomeSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "employer_name")
		delete(additionalProperties, "employee_name")
		delete(additionalProperties, "ytd_gross_income")
		delete(additionalProperties, "ytd_net_income")
		delete(additionalProperties, "pay_frequency")
		delete(additionalProperties, "projected_wage")
		delete(additionalProperties, "verified_transaction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeSummary struct {
	value *IncomeSummary
	isSet bool
}

func (v NullableIncomeSummary) Get() *IncomeSummary {
	return v.value
}

func (v *NullableIncomeSummary) Set(val *IncomeSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeSummary(val *IncomeSummary) *NullableIncomeSummary {
	return &NullableIncomeSummary{value: val, isSet: true}
}

func (v NullableIncomeSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


