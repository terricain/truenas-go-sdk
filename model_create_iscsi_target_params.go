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

// CreateISCSITargetParams struct for CreateISCSITargetParams
type CreateISCSITargetParams struct {
	Name *string `json:"name,omitempty"`
	Alias NullableString `json:"alias,omitempty"`
	Mode *string `json:"mode,omitempty"`
	Groups []CreateISCSITargetParamsGroupsInner `json:"groups,omitempty"`
}

// NewCreateISCSITargetParams instantiates a new CreateISCSITargetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateISCSITargetParams() *CreateISCSITargetParams {
	this := CreateISCSITargetParams{}
	return &this
}

// NewCreateISCSITargetParamsWithDefaults instantiates a new CreateISCSITargetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateISCSITargetParamsWithDefaults() *CreateISCSITargetParams {
	this := CreateISCSITargetParams{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateISCSITargetParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSITargetParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateISCSITargetParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateISCSITargetParams) SetName(v string) {
	o.Name = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateISCSITargetParams) GetAlias() string {
	if o == nil || o.Alias.Get() == nil {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateISCSITargetParams) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *CreateISCSITargetParams) HasAlias() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *CreateISCSITargetParams) SetAlias(v string) {
	o.Alias.Set(&v)
}
// SetAliasNil sets the value for Alias to be an explicit nil
func (o *CreateISCSITargetParams) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *CreateISCSITargetParams) UnsetAlias() {
	o.Alias.Unset()
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *CreateISCSITargetParams) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSITargetParams) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *CreateISCSITargetParams) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *CreateISCSITargetParams) SetMode(v string) {
	o.Mode = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *CreateISCSITargetParams) GetGroups() []CreateISCSITargetParamsGroupsInner {
	if o == nil || o.Groups == nil {
		var ret []CreateISCSITargetParamsGroupsInner
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSITargetParams) GetGroupsOk() ([]CreateISCSITargetParamsGroupsInner, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *CreateISCSITargetParams) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []CreateISCSITargetParamsGroupsInner and assigns it to the Groups field.
func (o *CreateISCSITargetParams) SetGroups(v []CreateISCSITargetParamsGroupsInner) {
	o.Groups = v
}

func (o CreateISCSITargetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableCreateISCSITargetParams struct {
	value *CreateISCSITargetParams
	isSet bool
}

func (v NullableCreateISCSITargetParams) Get() *CreateISCSITargetParams {
	return v.value
}

func (v *NullableCreateISCSITargetParams) Set(val *CreateISCSITargetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateISCSITargetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateISCSITargetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateISCSITargetParams(val *CreateISCSITargetParams) *NullableCreateISCSITargetParams {
	return &NullableCreateISCSITargetParams{value: val, isSet: true}
}

func (v NullableCreateISCSITargetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateISCSITargetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

