/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// checks if the NetworkSummaryIpsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkSummaryIpsValue{}

// NetworkSummaryIpsValue struct for NetworkSummaryIpsValue
type NetworkSummaryIpsValue struct {
	IPV4                 []string `json:"IPV4,omitempty"`
	IPV6                 []string `json:"IPV6,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkSummaryIpsValue NetworkSummaryIpsValue

// NewNetworkSummaryIpsValue instantiates a new NetworkSummaryIpsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSummaryIpsValue() *NetworkSummaryIpsValue {
	this := NetworkSummaryIpsValue{}
	return &this
}

// NewNetworkSummaryIpsValueWithDefaults instantiates a new NetworkSummaryIpsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSummaryIpsValueWithDefaults() *NetworkSummaryIpsValue {
	this := NetworkSummaryIpsValue{}
	return &this
}

// GetIPV4 returns the IPV4 field value if set, zero value otherwise.
func (o *NetworkSummaryIpsValue) GetIPV4() []string {
	if o == nil || IsNil(o.IPV4) {
		var ret []string
		return ret
	}
	return o.IPV4
}

// GetIPV4Ok returns a tuple with the IPV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSummaryIpsValue) GetIPV4Ok() ([]string, bool) {
	if o == nil || IsNil(o.IPV4) {
		return nil, false
	}
	return o.IPV4, true
}

// HasIPV4 returns a boolean if a field has been set.
func (o *NetworkSummaryIpsValue) HasIPV4() bool {
	if o != nil && !IsNil(o.IPV4) {
		return true
	}

	return false
}

// SetIPV4 gets a reference to the given []string and assigns it to the IPV4 field.
func (o *NetworkSummaryIpsValue) SetIPV4(v []string) {
	o.IPV4 = v
}

// GetIPV6 returns the IPV6 field value if set, zero value otherwise.
func (o *NetworkSummaryIpsValue) GetIPV6() []string {
	if o == nil || IsNil(o.IPV6) {
		var ret []string
		return ret
	}
	return o.IPV6
}

// GetIPV6Ok returns a tuple with the IPV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSummaryIpsValue) GetIPV6Ok() ([]string, bool) {
	if o == nil || IsNil(o.IPV6) {
		return nil, false
	}
	return o.IPV6, true
}

// HasIPV6 returns a boolean if a field has been set.
func (o *NetworkSummaryIpsValue) HasIPV6() bool {
	if o != nil && !IsNil(o.IPV6) {
		return true
	}

	return false
}

// SetIPV6 gets a reference to the given []string and assigns it to the IPV6 field.
func (o *NetworkSummaryIpsValue) SetIPV6(v []string) {
	o.IPV6 = v
}

func (o NetworkSummaryIpsValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkSummaryIpsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IPV4) {
		toSerialize["IPV4"] = o.IPV4
	}
	if !IsNil(o.IPV6) {
		toSerialize["IPV6"] = o.IPV6
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkSummaryIpsValue) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkSummaryIpsValue := _NetworkSummaryIpsValue{}

	err = json.Unmarshal(bytes, &varNetworkSummaryIpsValue)

	if err != nil {
		return err
	}

	*o = NetworkSummaryIpsValue(varNetworkSummaryIpsValue)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "IPV4")
		delete(additionalProperties, "IPV6")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkSummaryIpsValue struct {
	value *NetworkSummaryIpsValue
	isSet bool
}

func (v NullableNetworkSummaryIpsValue) Get() *NetworkSummaryIpsValue {
	return v.value
}

func (v *NullableNetworkSummaryIpsValue) Set(val *NetworkSummaryIpsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSummaryIpsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSummaryIpsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSummaryIpsValue(val *NetworkSummaryIpsValue) *NullableNetworkSummaryIpsValue {
	return &NullableNetworkSummaryIpsValue{value: val, isSet: true}
}

func (v NullableNetworkSummaryIpsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSummaryIpsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
