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

// DepositSwitchAddressData The user's address.
type DepositSwitchAddressData struct {
	// The full city name
	City string `json:"city"`
	// The region or state Example: `\"NC\"`
	Region string `json:"region"`
	// The full street address Example: `\"564 Main Street, APT 15\"`
	Street string `json:"street"`
	// The postal code
	PostalCode string `json:"postal_code"`
	// The ISO 3166-1 alpha-2 country code
	Country string `json:"country"`
	AdditionalProperties map[string]interface{}
}

type _DepositSwitchAddressData DepositSwitchAddressData

// NewDepositSwitchAddressData instantiates a new DepositSwitchAddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchAddressData(city string, region string, street string, postalCode string, country string) *DepositSwitchAddressData {
	this := DepositSwitchAddressData{}
	this.City = city
	this.Region = region
	this.Street = street
	this.PostalCode = postalCode
	this.Country = country
	return &this
}

// NewDepositSwitchAddressDataWithDefaults instantiates a new DepositSwitchAddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchAddressDataWithDefaults() *DepositSwitchAddressData {
	this := DepositSwitchAddressData{}
	return &this
}

// GetCity returns the City field value
func (o *DepositSwitchAddressData) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchAddressData) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *DepositSwitchAddressData) SetCity(v string) {
	o.City = v
}

// GetRegion returns the Region field value
func (o *DepositSwitchAddressData) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchAddressData) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *DepositSwitchAddressData) SetRegion(v string) {
	o.Region = v
}

// GetStreet returns the Street field value
func (o *DepositSwitchAddressData) GetStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchAddressData) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *DepositSwitchAddressData) SetStreet(v string) {
	o.Street = v
}

// GetPostalCode returns the PostalCode field value
func (o *DepositSwitchAddressData) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchAddressData) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *DepositSwitchAddressData) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetCountry returns the Country field value
func (o *DepositSwitchAddressData) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchAddressData) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *DepositSwitchAddressData) SetCountry(v string) {
	o.Country = v
}

func (o DepositSwitchAddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["street"] = o.Street
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode
	}
	if true {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DepositSwitchAddressData) UnmarshalJSON(bytes []byte) (err error) {
	varDepositSwitchAddressData := _DepositSwitchAddressData{}

	if err = json.Unmarshal(bytes, &varDepositSwitchAddressData); err == nil {
		*o = DepositSwitchAddressData(varDepositSwitchAddressData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "street")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositSwitchAddressData struct {
	value *DepositSwitchAddressData
	isSet bool
}

func (v NullableDepositSwitchAddressData) Get() *DepositSwitchAddressData {
	return v.value
}

func (v *NullableDepositSwitchAddressData) Set(val *DepositSwitchAddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchAddressData(val *DepositSwitchAddressData) *NullableDepositSwitchAddressData {
	return &NullableDepositSwitchAddressData{value: val, isSet: true}
}

func (v NullableDepositSwitchAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


