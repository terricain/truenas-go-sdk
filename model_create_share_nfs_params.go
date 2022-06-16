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

// CreateShareNFSParams struct for CreateShareNFSParams
type CreateShareNFSParams struct {
	Paths []string `json:"paths"`
	Comment *string `json:"comment,omitempty"`
	Networks []string `json:"networks,omitempty"`
	Hosts []string `json:"hosts,omitempty"`
	Alldirs *bool `json:"alldirs,omitempty"`
	Ro *bool `json:"ro,omitempty"`
	Quiet *bool `json:"quiet,omitempty"`
	MaprootUser *string `json:"maproot_user,omitempty"`
	MaprootGroup *string `json:"maproot_group,omitempty"`
	MapallUser *string `json:"mapall_user,omitempty"`
	MapallGroup *string `json:"mapall_group,omitempty"`
	Security []string `json:"security,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateShareNFSParams CreateShareNFSParams

// NewCreateShareNFSParams instantiates a new CreateShareNFSParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShareNFSParams(paths []string) *CreateShareNFSParams {
	this := CreateShareNFSParams{}
	this.Paths = paths
	return &this
}

// NewCreateShareNFSParamsWithDefaults instantiates a new CreateShareNFSParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShareNFSParamsWithDefaults() *CreateShareNFSParams {
	this := CreateShareNFSParams{}
	return &this
}

// GetPaths returns the Paths field value
func (o *CreateShareNFSParams) GetPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetPathsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Paths, true
}

// SetPaths sets field value
func (o *CreateShareNFSParams) SetPaths(v []string) {
	o.Paths = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateShareNFSParams) SetComment(v string) {
	o.Comment = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetNetworks() []string {
	if o == nil || o.Networks == nil {
		var ret []string
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetNetworksOk() ([]string, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *CreateShareNFSParams) SetNetworks(v []string) {
	o.Networks = v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetHosts() []string {
	if o == nil || o.Hosts == nil {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetHostsOk() ([]string, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *CreateShareNFSParams) SetHosts(v []string) {
	o.Hosts = v
}

// GetAlldirs returns the Alldirs field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetAlldirs() bool {
	if o == nil || o.Alldirs == nil {
		var ret bool
		return ret
	}
	return *o.Alldirs
}

// GetAlldirsOk returns a tuple with the Alldirs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetAlldirsOk() (*bool, bool) {
	if o == nil || o.Alldirs == nil {
		return nil, false
	}
	return o.Alldirs, true
}

// HasAlldirs returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasAlldirs() bool {
	if o != nil && o.Alldirs != nil {
		return true
	}

	return false
}

// SetAlldirs gets a reference to the given bool and assigns it to the Alldirs field.
func (o *CreateShareNFSParams) SetAlldirs(v bool) {
	o.Alldirs = &v
}

// GetRo returns the Ro field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetRo() bool {
	if o == nil || o.Ro == nil {
		var ret bool
		return ret
	}
	return *o.Ro
}

// GetRoOk returns a tuple with the Ro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetRoOk() (*bool, bool) {
	if o == nil || o.Ro == nil {
		return nil, false
	}
	return o.Ro, true
}

// HasRo returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasRo() bool {
	if o != nil && o.Ro != nil {
		return true
	}

	return false
}

// SetRo gets a reference to the given bool and assigns it to the Ro field.
func (o *CreateShareNFSParams) SetRo(v bool) {
	o.Ro = &v
}

// GetQuiet returns the Quiet field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetQuiet() bool {
	if o == nil || o.Quiet == nil {
		var ret bool
		return ret
	}
	return *o.Quiet
}

// GetQuietOk returns a tuple with the Quiet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetQuietOk() (*bool, bool) {
	if o == nil || o.Quiet == nil {
		return nil, false
	}
	return o.Quiet, true
}

// HasQuiet returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasQuiet() bool {
	if o != nil && o.Quiet != nil {
		return true
	}

	return false
}

// SetQuiet gets a reference to the given bool and assigns it to the Quiet field.
func (o *CreateShareNFSParams) SetQuiet(v bool) {
	o.Quiet = &v
}

// GetMaprootUser returns the MaprootUser field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetMaprootUser() string {
	if o == nil || o.MaprootUser == nil {
		var ret string
		return ret
	}
	return *o.MaprootUser
}

// GetMaprootUserOk returns a tuple with the MaprootUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetMaprootUserOk() (*string, bool) {
	if o == nil || o.MaprootUser == nil {
		return nil, false
	}
	return o.MaprootUser, true
}

// HasMaprootUser returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasMaprootUser() bool {
	if o != nil && o.MaprootUser != nil {
		return true
	}

	return false
}

// SetMaprootUser gets a reference to the given string and assigns it to the MaprootUser field.
func (o *CreateShareNFSParams) SetMaprootUser(v string) {
	o.MaprootUser = &v
}

// GetMaprootGroup returns the MaprootGroup field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetMaprootGroup() string {
	if o == nil || o.MaprootGroup == nil {
		var ret string
		return ret
	}
	return *o.MaprootGroup
}

// GetMaprootGroupOk returns a tuple with the MaprootGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetMaprootGroupOk() (*string, bool) {
	if o == nil || o.MaprootGroup == nil {
		return nil, false
	}
	return o.MaprootGroup, true
}

// HasMaprootGroup returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasMaprootGroup() bool {
	if o != nil && o.MaprootGroup != nil {
		return true
	}

	return false
}

// SetMaprootGroup gets a reference to the given string and assigns it to the MaprootGroup field.
func (o *CreateShareNFSParams) SetMaprootGroup(v string) {
	o.MaprootGroup = &v
}

// GetMapallUser returns the MapallUser field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetMapallUser() string {
	if o == nil || o.MapallUser == nil {
		var ret string
		return ret
	}
	return *o.MapallUser
}

// GetMapallUserOk returns a tuple with the MapallUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetMapallUserOk() (*string, bool) {
	if o == nil || o.MapallUser == nil {
		return nil, false
	}
	return o.MapallUser, true
}

// HasMapallUser returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasMapallUser() bool {
	if o != nil && o.MapallUser != nil {
		return true
	}

	return false
}

// SetMapallUser gets a reference to the given string and assigns it to the MapallUser field.
func (o *CreateShareNFSParams) SetMapallUser(v string) {
	o.MapallUser = &v
}

// GetMapallGroup returns the MapallGroup field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetMapallGroup() string {
	if o == nil || o.MapallGroup == nil {
		var ret string
		return ret
	}
	return *o.MapallGroup
}

// GetMapallGroupOk returns a tuple with the MapallGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetMapallGroupOk() (*string, bool) {
	if o == nil || o.MapallGroup == nil {
		return nil, false
	}
	return o.MapallGroup, true
}

// HasMapallGroup returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasMapallGroup() bool {
	if o != nil && o.MapallGroup != nil {
		return true
	}

	return false
}

// SetMapallGroup gets a reference to the given string and assigns it to the MapallGroup field.
func (o *CreateShareNFSParams) SetMapallGroup(v string) {
	o.MapallGroup = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetSecurity() []string {
	if o == nil || o.Security == nil {
		var ret []string
		return ret
	}
	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetSecurityOk() ([]string, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given []string and assigns it to the Security field.
func (o *CreateShareNFSParams) SetSecurity(v []string) {
	o.Security = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateShareNFSParams) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareNFSParams) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateShareNFSParams) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateShareNFSParams) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CreateShareNFSParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["paths"] = o.Paths
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Networks != nil {
		toSerialize["networks"] = o.Networks
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.Alldirs != nil {
		toSerialize["alldirs"] = o.Alldirs
	}
	if o.Ro != nil {
		toSerialize["ro"] = o.Ro
	}
	if o.Quiet != nil {
		toSerialize["quiet"] = o.Quiet
	}
	if o.MaprootUser != nil {
		toSerialize["maproot_user"] = o.MaprootUser
	}
	if o.MaprootGroup != nil {
		toSerialize["maproot_group"] = o.MaprootGroup
	}
	if o.MapallUser != nil {
		toSerialize["mapall_user"] = o.MapallUser
	}
	if o.MapallGroup != nil {
		toSerialize["mapall_group"] = o.MapallGroup
	}
	if o.Security != nil {
		toSerialize["security"] = o.Security
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateShareNFSParams) UnmarshalJSON(bytes []byte) (err error) {
	varCreateShareNFSParams := _CreateShareNFSParams{}

	if err = json.Unmarshal(bytes, &varCreateShareNFSParams); err == nil {
		*o = CreateShareNFSParams(varCreateShareNFSParams)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "paths")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "networks")
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "alldirs")
		delete(additionalProperties, "ro")
		delete(additionalProperties, "quiet")
		delete(additionalProperties, "maproot_user")
		delete(additionalProperties, "maproot_group")
		delete(additionalProperties, "mapall_user")
		delete(additionalProperties, "mapall_group")
		delete(additionalProperties, "security")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateShareNFSParams struct {
	value *CreateShareNFSParams
	isSet bool
}

func (v NullableCreateShareNFSParams) Get() *CreateShareNFSParams {
	return v.value
}

func (v *NullableCreateShareNFSParams) Set(val *CreateShareNFSParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShareNFSParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShareNFSParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShareNFSParams(val *CreateShareNFSParams) *NullableCreateShareNFSParams {
	return &NullableCreateShareNFSParams{value: val, isSet: true}
}

func (v NullableCreateShareNFSParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShareNFSParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


