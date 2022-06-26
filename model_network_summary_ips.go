/*
 * TrueNAS RESTful API
 *
 * Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// NetworkSummaryIps struct for NetworkSummaryIps
type NetworkSummaryIps struct {
	IPV4                 *[]string `json:"IPV4,omitempty"`
	IPV6                 *[]string `json:"IPV6,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkSummaryIps NetworkSummaryIps

// NewNetworkSummaryIps instantiates a new NetworkSummaryIps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSummaryIps() *NetworkSummaryIps {
	this := NetworkSummaryIps{}
	return &this
}

// NewNetworkSummaryIpsWithDefaults instantiates a new NetworkSummaryIps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSummaryIpsWithDefaults() *NetworkSummaryIps {
	this := NetworkSummaryIps{}
	return &this
}

// GetIPV4 returns the IPV4 field value if set, zero value otherwise.
func (o *NetworkSummaryIps) GetIPV4() []string {
	if o == nil || o.IPV4 == nil {
		var ret []string
		return ret
	}
	return *o.IPV4
}

// GetIPV4Ok returns a tuple with the IPV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSummaryIps) GetIPV4Ok() (*[]string, bool) {
	if o == nil || o.IPV4 == nil {
		return nil, false
	}
	return o.IPV4, true
}

// HasIPV4 returns a boolean if a field has been set.
func (o *NetworkSummaryIps) HasIPV4() bool {
	if o != nil && o.IPV4 != nil {
		return true
	}

	return false
}

// SetIPV4 gets a reference to the given []string and assigns it to the IPV4 field.
func (o *NetworkSummaryIps) SetIPV4(v []string) {
	o.IPV4 = &v
}

// GetIPV6 returns the IPV6 field value if set, zero value otherwise.
func (o *NetworkSummaryIps) GetIPV6() []string {
	if o == nil || o.IPV6 == nil {
		var ret []string
		return ret
	}
	return *o.IPV6
}

// GetIPV6Ok returns a tuple with the IPV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSummaryIps) GetIPV6Ok() (*[]string, bool) {
	if o == nil || o.IPV6 == nil {
		return nil, false
	}
	return o.IPV6, true
}

// HasIPV6 returns a boolean if a field has been set.
func (o *NetworkSummaryIps) HasIPV6() bool {
	if o != nil && o.IPV6 != nil {
		return true
	}

	return false
}

// SetIPV6 gets a reference to the given []string and assigns it to the IPV6 field.
func (o *NetworkSummaryIps) SetIPV6(v []string) {
	o.IPV6 = &v
}

func (o NetworkSummaryIps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IPV4 != nil {
		toSerialize["IPV4"] = o.IPV4
	}
	if o.IPV6 != nil {
		toSerialize["IPV6"] = o.IPV6
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkSummaryIps) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkSummaryIps := _NetworkSummaryIps{}

	if err = json.Unmarshal(bytes, &varNetworkSummaryIps); err == nil {
		*o = NetworkSummaryIps(varNetworkSummaryIps)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "IPV4")
		delete(additionalProperties, "IPV6")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkSummaryIps struct {
	value *NetworkSummaryIps
	isSet bool
}

func (v NullableNetworkSummaryIps) Get() *NetworkSummaryIps {
	return v.value
}

func (v *NullableNetworkSummaryIps) Set(val *NetworkSummaryIps) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSummaryIps) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSummaryIps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSummaryIps(val *NetworkSummaryIps) *NullableNetworkSummaryIps {
	return &NullableNetworkSummaryIps{value: val, isSet: true}
}

func (v NullableNetworkSummaryIps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSummaryIps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}