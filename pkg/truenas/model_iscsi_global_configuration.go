/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
	"fmt"
)

// checks if the ISCSIGlobalConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ISCSIGlobalConfiguration{}

// ISCSIGlobalConfiguration struct for ISCSIGlobalConfiguration
type ISCSIGlobalConfiguration struct {
	Id                 int32         `json:"id"`
	Basename           string        `json:"basename"`
	IsnsServers        []string      `json:"isns_servers"`
	PoolAvailThreshold NullableInt32 `json:"pool_avail_threshold"`
	Alua               bool          `json:"alua"`
}

type _ISCSIGlobalConfiguration ISCSIGlobalConfiguration

// NewISCSIGlobalConfiguration instantiates a new ISCSIGlobalConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewISCSIGlobalConfiguration(id int32, basename string, isnsServers []string, poolAvailThreshold NullableInt32, alua bool) *ISCSIGlobalConfiguration {
	this := ISCSIGlobalConfiguration{}
	this.Id = id
	this.Basename = basename
	this.IsnsServers = isnsServers
	this.PoolAvailThreshold = poolAvailThreshold
	this.Alua = alua
	return &this
}

// NewISCSIGlobalConfigurationWithDefaults instantiates a new ISCSIGlobalConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewISCSIGlobalConfigurationWithDefaults() *ISCSIGlobalConfiguration {
	this := ISCSIGlobalConfiguration{}
	return &this
}

// GetId returns the Id field value
func (o *ISCSIGlobalConfiguration) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ISCSIGlobalConfiguration) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ISCSIGlobalConfiguration) SetId(v int32) {
	o.Id = v
}

// GetBasename returns the Basename field value
func (o *ISCSIGlobalConfiguration) GetBasename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Basename
}

// GetBasenameOk returns a tuple with the Basename field value
// and a boolean to check if the value has been set.
func (o *ISCSIGlobalConfiguration) GetBasenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Basename, true
}

// SetBasename sets field value
func (o *ISCSIGlobalConfiguration) SetBasename(v string) {
	o.Basename = v
}

// GetIsnsServers returns the IsnsServers field value
func (o *ISCSIGlobalConfiguration) GetIsnsServers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IsnsServers
}

// GetIsnsServersOk returns a tuple with the IsnsServers field value
// and a boolean to check if the value has been set.
func (o *ISCSIGlobalConfiguration) GetIsnsServersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsnsServers, true
}

// SetIsnsServers sets field value
func (o *ISCSIGlobalConfiguration) SetIsnsServers(v []string) {
	o.IsnsServers = v
}

// GetPoolAvailThreshold returns the PoolAvailThreshold field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ISCSIGlobalConfiguration) GetPoolAvailThreshold() int32 {
	if o == nil || o.PoolAvailThreshold.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PoolAvailThreshold.Get()
}

// GetPoolAvailThresholdOk returns a tuple with the PoolAvailThreshold field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISCSIGlobalConfiguration) GetPoolAvailThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PoolAvailThreshold.Get(), o.PoolAvailThreshold.IsSet()
}

// SetPoolAvailThreshold sets field value
func (o *ISCSIGlobalConfiguration) SetPoolAvailThreshold(v int32) {
	o.PoolAvailThreshold.Set(&v)
}

// GetAlua returns the Alua field value
func (o *ISCSIGlobalConfiguration) GetAlua() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Alua
}

// GetAluaOk returns a tuple with the Alua field value
// and a boolean to check if the value has been set.
func (o *ISCSIGlobalConfiguration) GetAluaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alua, true
}

// SetAlua sets field value
func (o *ISCSIGlobalConfiguration) SetAlua(v bool) {
	o.Alua = v
}

func (o ISCSIGlobalConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ISCSIGlobalConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["basename"] = o.Basename
	toSerialize["isns_servers"] = o.IsnsServers
	toSerialize["pool_avail_threshold"] = o.PoolAvailThreshold.Get()
	toSerialize["alua"] = o.Alua
	return toSerialize, nil
}

func (o *ISCSIGlobalConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"basename",
		"isns_servers",
		"pool_avail_threshold",
		"alua",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varISCSIGlobalConfiguration := _ISCSIGlobalConfiguration{}

	err = json.Unmarshal(bytes, &varISCSIGlobalConfiguration)

	if err != nil {
		return err
	}

	*o = ISCSIGlobalConfiguration(varISCSIGlobalConfiguration)

	return err
}

type NullableISCSIGlobalConfiguration struct {
	value *ISCSIGlobalConfiguration
	isSet bool
}

func (v NullableISCSIGlobalConfiguration) Get() *ISCSIGlobalConfiguration {
	return v.value
}

func (v *NullableISCSIGlobalConfiguration) Set(val *ISCSIGlobalConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableISCSIGlobalConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableISCSIGlobalConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableISCSIGlobalConfiguration(val *ISCSIGlobalConfiguration) *NullableISCSIGlobalConfiguration {
	return &NullableISCSIGlobalConfiguration{value: val, isSet: true}
}

func (v NullableISCSIGlobalConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableISCSIGlobalConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
