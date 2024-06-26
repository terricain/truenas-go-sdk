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

// checks if the VMDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMDevicesInner{}

// VMDevicesInner struct for VMDevicesInner
type VMDevicesInner struct {
	Id                   int32                  `json:"id"`
	Dtype                string                 `json:"dtype"`
	Order                *int32                 `json:"order,omitempty"`
	Vm                   *int32                 `json:"vm,omitempty"`
	Attributes           map[string]interface{} `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VMDevicesInner VMDevicesInner

// NewVMDevicesInner instantiates a new VMDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMDevicesInner(id int32, dtype string) *VMDevicesInner {
	this := VMDevicesInner{}
	this.Id = id
	this.Dtype = dtype
	return &this
}

// NewVMDevicesInnerWithDefaults instantiates a new VMDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMDevicesInnerWithDefaults() *VMDevicesInner {
	this := VMDevicesInner{}
	return &this
}

// GetId returns the Id field value
func (o *VMDevicesInner) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VMDevicesInner) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VMDevicesInner) SetId(v int32) {
	o.Id = v
}

// GetDtype returns the Dtype field value
func (o *VMDevicesInner) GetDtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dtype
}

// GetDtypeOk returns a tuple with the Dtype field value
// and a boolean to check if the value has been set.
func (o *VMDevicesInner) GetDtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dtype, true
}

// SetDtype sets field value
func (o *VMDevicesInner) SetDtype(v string) {
	o.Dtype = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VMDevicesInner) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMDevicesInner) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VMDevicesInner) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *VMDevicesInner) SetOrder(v int32) {
	o.Order = &v
}

// GetVm returns the Vm field value if set, zero value otherwise.
func (o *VMDevicesInner) GetVm() int32 {
	if o == nil || IsNil(o.Vm) {
		var ret int32
		return ret
	}
	return *o.Vm
}

// GetVmOk returns a tuple with the Vm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMDevicesInner) GetVmOk() (*int32, bool) {
	if o == nil || IsNil(o.Vm) {
		return nil, false
	}
	return o.Vm, true
}

// HasVm returns a boolean if a field has been set.
func (o *VMDevicesInner) HasVm() bool {
	if o != nil && !IsNil(o.Vm) {
		return true
	}

	return false
}

// SetVm gets a reference to the given int32 and assigns it to the Vm field.
func (o *VMDevicesInner) SetVm(v int32) {
	o.Vm = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *VMDevicesInner) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMDevicesInner) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *VMDevicesInner) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *VMDevicesInner) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o VMDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["dtype"] = o.Dtype
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Vm) {
		toSerialize["vm"] = o.Vm
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VMDevicesInner) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"dtype",
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

	varVMDevicesInner := _VMDevicesInner{}

	err = json.Unmarshal(bytes, &varVMDevicesInner)

	if err != nil {
		return err
	}

	*o = VMDevicesInner(varVMDevicesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "dtype")
		delete(additionalProperties, "order")
		delete(additionalProperties, "vm")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVMDevicesInner struct {
	value *VMDevicesInner
	isSet bool
}

func (v NullableVMDevicesInner) Get() *VMDevicesInner {
	return v.value
}

func (v *NullableVMDevicesInner) Set(val *VMDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVMDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVMDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMDevicesInner(val *VMDevicesInner) *NullableVMDevicesInner {
	return &NullableVMDevicesInner{value: val, isSet: true}
}

func (v NullableVMDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
