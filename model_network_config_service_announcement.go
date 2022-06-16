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

// NetworkConfigServiceAnnouncement struct for NetworkConfigServiceAnnouncement
type NetworkConfigServiceAnnouncement struct {
	Netbios *bool `json:"netbios,omitempty"`
	Mdns *bool `json:"mdns,omitempty"`
	Wsd *bool `json:"wsd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkConfigServiceAnnouncement NetworkConfigServiceAnnouncement

// NewNetworkConfigServiceAnnouncement instantiates a new NetworkConfigServiceAnnouncement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkConfigServiceAnnouncement() *NetworkConfigServiceAnnouncement {
	this := NetworkConfigServiceAnnouncement{}
	return &this
}

// NewNetworkConfigServiceAnnouncementWithDefaults instantiates a new NetworkConfigServiceAnnouncement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkConfigServiceAnnouncementWithDefaults() *NetworkConfigServiceAnnouncement {
	this := NetworkConfigServiceAnnouncement{}
	return &this
}

// GetNetbios returns the Netbios field value if set, zero value otherwise.
func (o *NetworkConfigServiceAnnouncement) GetNetbios() bool {
	if o == nil || o.Netbios == nil {
		var ret bool
		return ret
	}
	return *o.Netbios
}

// GetNetbiosOk returns a tuple with the Netbios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfigServiceAnnouncement) GetNetbiosOk() (*bool, bool) {
	if o == nil || o.Netbios == nil {
		return nil, false
	}
	return o.Netbios, true
}

// HasNetbios returns a boolean if a field has been set.
func (o *NetworkConfigServiceAnnouncement) HasNetbios() bool {
	if o != nil && o.Netbios != nil {
		return true
	}

	return false
}

// SetNetbios gets a reference to the given bool and assigns it to the Netbios field.
func (o *NetworkConfigServiceAnnouncement) SetNetbios(v bool) {
	o.Netbios = &v
}

// GetMdns returns the Mdns field value if set, zero value otherwise.
func (o *NetworkConfigServiceAnnouncement) GetMdns() bool {
	if o == nil || o.Mdns == nil {
		var ret bool
		return ret
	}
	return *o.Mdns
}

// GetMdnsOk returns a tuple with the Mdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfigServiceAnnouncement) GetMdnsOk() (*bool, bool) {
	if o == nil || o.Mdns == nil {
		return nil, false
	}
	return o.Mdns, true
}

// HasMdns returns a boolean if a field has been set.
func (o *NetworkConfigServiceAnnouncement) HasMdns() bool {
	if o != nil && o.Mdns != nil {
		return true
	}

	return false
}

// SetMdns gets a reference to the given bool and assigns it to the Mdns field.
func (o *NetworkConfigServiceAnnouncement) SetMdns(v bool) {
	o.Mdns = &v
}

// GetWsd returns the Wsd field value if set, zero value otherwise.
func (o *NetworkConfigServiceAnnouncement) GetWsd() bool {
	if o == nil || o.Wsd == nil {
		var ret bool
		return ret
	}
	return *o.Wsd
}

// GetWsdOk returns a tuple with the Wsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfigServiceAnnouncement) GetWsdOk() (*bool, bool) {
	if o == nil || o.Wsd == nil {
		return nil, false
	}
	return o.Wsd, true
}

// HasWsd returns a boolean if a field has been set.
func (o *NetworkConfigServiceAnnouncement) HasWsd() bool {
	if o != nil && o.Wsd != nil {
		return true
	}

	return false
}

// SetWsd gets a reference to the given bool and assigns it to the Wsd field.
func (o *NetworkConfigServiceAnnouncement) SetWsd(v bool) {
	o.Wsd = &v
}

func (o NetworkConfigServiceAnnouncement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Netbios != nil {
		toSerialize["netbios"] = o.Netbios
	}
	if o.Mdns != nil {
		toSerialize["mdns"] = o.Mdns
	}
	if o.Wsd != nil {
		toSerialize["wsd"] = o.Wsd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkConfigServiceAnnouncement) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkConfigServiceAnnouncement := _NetworkConfigServiceAnnouncement{}

	if err = json.Unmarshal(bytes, &varNetworkConfigServiceAnnouncement); err == nil {
		*o = NetworkConfigServiceAnnouncement(varNetworkConfigServiceAnnouncement)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "netbios")
		delete(additionalProperties, "mdns")
		delete(additionalProperties, "wsd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkConfigServiceAnnouncement struct {
	value *NetworkConfigServiceAnnouncement
	isSet bool
}

func (v NullableNetworkConfigServiceAnnouncement) Get() *NetworkConfigServiceAnnouncement {
	return v.value
}

func (v *NullableNetworkConfigServiceAnnouncement) Set(val *NetworkConfigServiceAnnouncement) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkConfigServiceAnnouncement) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkConfigServiceAnnouncement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkConfigServiceAnnouncement(val *NetworkConfigServiceAnnouncement) *NullableNetworkConfigServiceAnnouncement {
	return &NullableNetworkConfigServiceAnnouncement{value: val, isSet: true}
}

func (v NullableNetworkConfigServiceAnnouncement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkConfigServiceAnnouncement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


