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

// CreateISCSITargetExtentParams struct for CreateISCSITargetExtentParams
type CreateISCSITargetExtentParams struct {
	Target int32 `json:"target"`
	Lunid NullableInt32 `json:"lunid,omitempty"`
	Extent int32 `json:"extent"`
}

// NewCreateISCSITargetExtentParams instantiates a new CreateISCSITargetExtentParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateISCSITargetExtentParams(target int32, extent int32) *CreateISCSITargetExtentParams {
	this := CreateISCSITargetExtentParams{}
	this.Target = target
	this.Extent = extent
	return &this
}

// NewCreateISCSITargetExtentParamsWithDefaults instantiates a new CreateISCSITargetExtentParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateISCSITargetExtentParamsWithDefaults() *CreateISCSITargetExtentParams {
	this := CreateISCSITargetExtentParams{}
	return &this
}

// GetTarget returns the Target field value
func (o *CreateISCSITargetExtentParams) GetTarget() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *CreateISCSITargetExtentParams) GetTargetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *CreateISCSITargetExtentParams) SetTarget(v int32) {
	o.Target = v
}

// GetLunid returns the Lunid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateISCSITargetExtentParams) GetLunid() int32 {
	if o == nil || o.Lunid.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Lunid.Get()
}

// GetLunidOk returns a tuple with the Lunid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateISCSITargetExtentParams) GetLunidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lunid.Get(), o.Lunid.IsSet()
}

// HasLunid returns a boolean if a field has been set.
func (o *CreateISCSITargetExtentParams) HasLunid() bool {
	if o != nil && o.Lunid.IsSet() {
		return true
	}

	return false
}

// SetLunid gets a reference to the given NullableInt32 and assigns it to the Lunid field.
func (o *CreateISCSITargetExtentParams) SetLunid(v int32) {
	o.Lunid.Set(&v)
}
// SetLunidNil sets the value for Lunid to be an explicit nil
func (o *CreateISCSITargetExtentParams) SetLunidNil() {
	o.Lunid.Set(nil)
}

// UnsetLunid ensures that no value is present for Lunid, not even an explicit nil
func (o *CreateISCSITargetExtentParams) UnsetLunid() {
	o.Lunid.Unset()
}

// GetExtent returns the Extent field value
func (o *CreateISCSITargetExtentParams) GetExtent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Extent
}

// GetExtentOk returns a tuple with the Extent field value
// and a boolean to check if the value has been set.
func (o *CreateISCSITargetExtentParams) GetExtentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extent, true
}

// SetExtent sets field value
func (o *CreateISCSITargetExtentParams) SetExtent(v int32) {
	o.Extent = v
}

func (o CreateISCSITargetExtentParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["target"] = o.Target
	}
	if o.Lunid.IsSet() {
		toSerialize["lunid"] = o.Lunid.Get()
	}
	if true {
		toSerialize["extent"] = o.Extent
	}
	return json.Marshal(toSerialize)
}

type NullableCreateISCSITargetExtentParams struct {
	value *CreateISCSITargetExtentParams
	isSet bool
}

func (v NullableCreateISCSITargetExtentParams) Get() *CreateISCSITargetExtentParams {
	return v.value
}

func (v *NullableCreateISCSITargetExtentParams) Set(val *CreateISCSITargetExtentParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateISCSITargetExtentParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateISCSITargetExtentParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateISCSITargetExtentParams(val *CreateISCSITargetExtentParams) *NullableCreateISCSITargetExtentParams {
	return &NullableCreateISCSITargetExtentParams{value: val, isSet: true}
}

func (v NullableCreateISCSITargetExtentParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateISCSITargetExtentParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


