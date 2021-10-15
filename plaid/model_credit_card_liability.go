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

// CreditCardLiability An object representing a credit card account.
type CreditCardLiability struct {
	// The ID of the account that this liability belongs to.
	AccountId NullableString `json:"account_id"`
	// The various interest rates that apply to the account.
	Aprs []APR `json:"aprs"`
	// true if a payment is currently overdue. Availability for this field is limited.
	IsOverdue NullableBool `json:"is_overdue"`
	// The amount of the last payment.
	LastPaymentAmount float32 `json:"last_payment_amount"`
	// The date of the last payment. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD). Availability for this field is limited.
	LastPaymentDate string `json:"last_payment_date"`
	// The date of the last statement. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	LastStatementIssueDate string `json:"last_statement_issue_date"`
	// The total amount owed as of the last statement issued
	LastStatementBalance float32 `json:"last_statement_balance"`
	// The minimum payment due for the next billing cycle.
	MinimumPaymentAmount float32 `json:"minimum_payment_amount"`
	// The due date for the next payment. The due date is `null` if a payment is not expected. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD).
	NextPaymentDueDate NullableString `json:"next_payment_due_date"`
	AdditionalProperties map[string]interface{}
}

type _CreditCardLiability CreditCardLiability

// NewCreditCardLiability instantiates a new CreditCardLiability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCardLiability(accountId NullableString, aprs []APR, isOverdue NullableBool, lastPaymentAmount float32, lastPaymentDate string, lastStatementIssueDate string, lastStatementBalance float32, minimumPaymentAmount float32, nextPaymentDueDate NullableString) *CreditCardLiability {
	this := CreditCardLiability{}
	this.AccountId = accountId
	this.Aprs = aprs
	this.IsOverdue = isOverdue
	this.LastPaymentAmount = lastPaymentAmount
	this.LastPaymentDate = lastPaymentDate
	this.LastStatementIssueDate = lastStatementIssueDate
	this.LastStatementBalance = lastStatementBalance
	this.MinimumPaymentAmount = minimumPaymentAmount
	this.NextPaymentDueDate = nextPaymentDueDate
	return &this
}

// NewCreditCardLiabilityWithDefaults instantiates a new CreditCardLiability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCardLiabilityWithDefaults() *CreditCardLiability {
	this := CreditCardLiability{}
	return &this
}

// GetAccountId returns the AccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditCardLiability) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditCardLiability) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// SetAccountId sets field value
func (o *CreditCardLiability) SetAccountId(v string) {
	o.AccountId.Set(&v)
}

// GetAprs returns the Aprs field value
func (o *CreditCardLiability) GetAprs() []APR {
	if o == nil {
		var ret []APR
		return ret
	}

	return o.Aprs
}

// GetAprsOk returns a tuple with the Aprs field value
// and a boolean to check if the value has been set.
func (o *CreditCardLiability) GetAprsOk() (*[]APR, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Aprs, true
}

// SetAprs sets field value
func (o *CreditCardLiability) SetAprs(v []APR) {
	o.Aprs = v
}

// GetIsOverdue returns the IsOverdue field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *CreditCardLiability) GetIsOverdue() bool {
	if o == nil || o.IsOverdue.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsOverdue.Get()
}

// GetIsOverdueOk returns a tuple with the IsOverdue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditCardLiability) GetIsOverdueOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsOverdue.Get(), o.IsOverdue.IsSet()
}

// SetIsOverdue sets field value
func (o *CreditCardLiability) SetIsOverdue(v bool) {
	o.IsOverdue.Set(&v)
}

// GetLastPaymentAmount returns the LastPaymentAmount field value
func (o *CreditCardLiability) GetLastPaymentAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastPaymentAmount
}

// GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field value
// and a boolean to check if the value has been set.
func (o *CreditCardLiability) GetLastPaymentAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastPaymentAmount, true
}

// SetLastPaymentAmount sets field value
func (o *CreditCardLiability) SetLastPaymentAmount(v float32) {
	o.LastPaymentAmount = v
}

// GetLastPaymentDate returns the LastPaymentDate field value
func (o *CreditCardLiability) GetLastPaymentDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPaymentDate
}

// GetLastPaymentDateOk returns a tuple with the LastPaymentDate field value
// and a boolean to check if the value has been set.
func (o *CreditCardLiability) GetLastPaymentDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastPaymentDate, true
}

// SetLastPaymentDate sets field value
func (o *CreditCardLiability) SetLastPaymentDate(v string) {
	o.LastPaymentDate = v
}

// GetLastStatementIssueDate returns the LastStatementIssueDate field value
func (o *CreditCardLiability) GetLastStatementIssueDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastStatementIssueDate
}

// GetLastStatementIssueDateOk returns a tuple with the LastStatementIssueDate field value
// and a boolean to check if the value has been set.
func (o *CreditCardLiability) GetLastStatementIssueDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastStatementIssueDate, true
}

// SetLastStatementIssueDate sets field value
func (o *CreditCardLiability) SetLastStatementIssueDate(v string) {
	o.LastStatementIssueDate = v
}

// GetLastStatementBalance returns the LastStatementBalance field value
func (o *CreditCardLiability) GetLastStatementBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastStatementBalance
}

// GetLastStatementBalanceOk returns a tuple with the LastStatementBalance field value
// and a boolean to check if the value has been set.
func (o *CreditCardLiability) GetLastStatementBalanceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastStatementBalance, true
}

// SetLastStatementBalance sets field value
func (o *CreditCardLiability) SetLastStatementBalance(v float32) {
	o.LastStatementBalance = v
}

// GetMinimumPaymentAmount returns the MinimumPaymentAmount field value
func (o *CreditCardLiability) GetMinimumPaymentAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinimumPaymentAmount
}

// GetMinimumPaymentAmountOk returns a tuple with the MinimumPaymentAmount field value
// and a boolean to check if the value has been set.
func (o *CreditCardLiability) GetMinimumPaymentAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinimumPaymentAmount, true
}

// SetMinimumPaymentAmount sets field value
func (o *CreditCardLiability) SetMinimumPaymentAmount(v float32) {
	o.MinimumPaymentAmount = v
}

// GetNextPaymentDueDate returns the NextPaymentDueDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditCardLiability) GetNextPaymentDueDate() string {
	if o == nil || o.NextPaymentDueDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextPaymentDueDate.Get()
}

// GetNextPaymentDueDateOk returns a tuple with the NextPaymentDueDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditCardLiability) GetNextPaymentDueDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextPaymentDueDate.Get(), o.NextPaymentDueDate.IsSet()
}

// SetNextPaymentDueDate sets field value
func (o *CreditCardLiability) SetNextPaymentDueDate(v string) {
	o.NextPaymentDueDate.Set(&v)
}

func (o CreditCardLiability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if true {
		toSerialize["aprs"] = o.Aprs
	}
	if true {
		toSerialize["is_overdue"] = o.IsOverdue.Get()
	}
	if true {
		toSerialize["last_payment_amount"] = o.LastPaymentAmount
	}
	if true {
		toSerialize["last_payment_date"] = o.LastPaymentDate
	}
	if true {
		toSerialize["last_statement_issue_date"] = o.LastStatementIssueDate
	}
	if true {
		toSerialize["last_statement_balance"] = o.LastStatementBalance
	}
	if true {
		toSerialize["minimum_payment_amount"] = o.MinimumPaymentAmount
	}
	if true {
		toSerialize["next_payment_due_date"] = o.NextPaymentDueDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditCardLiability) UnmarshalJSON(bytes []byte) (err error) {
	varCreditCardLiability := _CreditCardLiability{}

	if err = json.Unmarshal(bytes, &varCreditCardLiability); err == nil {
		*o = CreditCardLiability(varCreditCardLiability)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "aprs")
		delete(additionalProperties, "is_overdue")
		delete(additionalProperties, "last_payment_amount")
		delete(additionalProperties, "last_payment_date")
		delete(additionalProperties, "last_statement_issue_date")
		delete(additionalProperties, "last_statement_balance")
		delete(additionalProperties, "minimum_payment_amount")
		delete(additionalProperties, "next_payment_due_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditCardLiability struct {
	value *CreditCardLiability
	isSet bool
}

func (v NullableCreditCardLiability) Get() *CreditCardLiability {
	return v.value
}

func (v *NullableCreditCardLiability) Set(val *CreditCardLiability) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCardLiability) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCardLiability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCardLiability(val *CreditCardLiability) *NullableCreditCardLiability {
	return &NullableCreditCardLiability{value: val, isSet: true}
}

func (v NullableCreditCardLiability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCardLiability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


