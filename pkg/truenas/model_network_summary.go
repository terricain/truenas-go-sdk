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

// checks if the NetworkSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkSummary{}

// NetworkSummary struct for NetworkSummary
type NetworkSummary struct {
	Ips                  *map[string]NetworkSummaryIpsValue `json:"ips,omitempty"`
	DefaultRoutes        []string                           `json:"default_routes,omitempty"`
	Nameservers          []string                           `json:"nameservers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkSummary NetworkSummary

// NewNetworkSummary instantiates a new NetworkSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSummary() *NetworkSummary {
	this := NetworkSummary{}
	return &this
}

// NewNetworkSummaryWithDefaults instantiates a new NetworkSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSummaryWithDefaults() *NetworkSummary {
	this := NetworkSummary{}
	return &this
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *NetworkSummary) GetIps() map[string]NetworkSummaryIpsValue {
	if o == nil || IsNil(o.Ips) {
		var ret map[string]NetworkSummaryIpsValue
		return ret
	}
	return *o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSummary) GetIpsOk() (*map[string]NetworkSummaryIpsValue, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *NetworkSummary) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given map[string]NetworkSummaryIpsValue and assigns it to the Ips field.
func (o *NetworkSummary) SetIps(v map[string]NetworkSummaryIpsValue) {
	o.Ips = &v
}

// GetDefaultRoutes returns the DefaultRoutes field value if set, zero value otherwise.
func (o *NetworkSummary) GetDefaultRoutes() []string {
	if o == nil || IsNil(o.DefaultRoutes) {
		var ret []string
		return ret
	}
	return o.DefaultRoutes
}

// GetDefaultRoutesOk returns a tuple with the DefaultRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSummary) GetDefaultRoutesOk() ([]string, bool) {
	if o == nil || IsNil(o.DefaultRoutes) {
		return nil, false
	}
	return o.DefaultRoutes, true
}

// HasDefaultRoutes returns a boolean if a field has been set.
func (o *NetworkSummary) HasDefaultRoutes() bool {
	if o != nil && !IsNil(o.DefaultRoutes) {
		return true
	}

	return false
}

// SetDefaultRoutes gets a reference to the given []string and assigns it to the DefaultRoutes field.
func (o *NetworkSummary) SetDefaultRoutes(v []string) {
	o.DefaultRoutes = v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *NetworkSummary) GetNameservers() []string {
	if o == nil || IsNil(o.Nameservers) {
		var ret []string
		return ret
	}
	return o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSummary) GetNameserversOk() ([]string, bool) {
	if o == nil || IsNil(o.Nameservers) {
		return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *NetworkSummary) HasNameservers() bool {
	if o != nil && !IsNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given []string and assigns it to the Nameservers field.
func (o *NetworkSummary) SetNameservers(v []string) {
	o.Nameservers = v
}

func (o NetworkSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.DefaultRoutes) {
		toSerialize["default_routes"] = o.DefaultRoutes
	}
	if !IsNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkSummary) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkSummary := _NetworkSummary{}

	err = json.Unmarshal(bytes, &varNetworkSummary)

	if err != nil {
		return err
	}

	*o = NetworkSummary(varNetworkSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ips")
		delete(additionalProperties, "default_routes")
		delete(additionalProperties, "nameservers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkSummary struct {
	value *NetworkSummary
	isSet bool
}

func (v NullableNetworkSummary) Get() *NetworkSummary {
	return v.value
}

func (v *NullableNetworkSummary) Set(val *NetworkSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSummary(val *NetworkSummary) *NullableNetworkSummary {
	return &NullableNetworkSummary{value: val, isSet: true}
}

func (v NullableNetworkSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
