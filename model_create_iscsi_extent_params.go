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

// CreateISCSIExtentParams struct for CreateISCSIExtentParams
type CreateISCSIExtentParams struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Disk NullableString `json:"disk,omitempty"`
	Serial NullableString `json:"serial,omitempty"`
	Path NullableString `json:"path,omitempty"`
	Filesize *int32 `json:"filesize,omitempty"`
	Blocksize *int32 `json:"blocksize,omitempty"`
	Pblocksize *bool `json:"pblocksize,omitempty"`
	AvailThreshold NullableInt32 `json:"avail_threshold,omitempty"`
	Comment *string `json:"comment,omitempty"`
	InsecureTpc *bool `json:"insecure_tpc,omitempty"`
	Xen *bool `json:"xen,omitempty"`
	Rpm *string `json:"rpm,omitempty"`
	Ro *bool `json:"ro,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCreateISCSIExtentParams instantiates a new CreateISCSIExtentParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateISCSIExtentParams() *CreateISCSIExtentParams {
	this := CreateISCSIExtentParams{}
	return &this
}

// NewCreateISCSIExtentParamsWithDefaults instantiates a new CreateISCSIExtentParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateISCSIExtentParamsWithDefaults() *CreateISCSIExtentParams {
	this := CreateISCSIExtentParams{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateISCSIExtentParams) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateISCSIExtentParams) SetType(v string) {
	o.Type = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateISCSIExtentParams) GetDisk() string {
	if o == nil || o.Disk.Get() == nil {
		var ret string
		return ret
	}
	return *o.Disk.Get()
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateISCSIExtentParams) GetDiskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disk.Get(), o.Disk.IsSet()
}

// HasDisk returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasDisk() bool {
	if o != nil && o.Disk.IsSet() {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NullableString and assigns it to the Disk field.
func (o *CreateISCSIExtentParams) SetDisk(v string) {
	o.Disk.Set(&v)
}
// SetDiskNil sets the value for Disk to be an explicit nil
func (o *CreateISCSIExtentParams) SetDiskNil() {
	o.Disk.Set(nil)
}

// UnsetDisk ensures that no value is present for Disk, not even an explicit nil
func (o *CreateISCSIExtentParams) UnsetDisk() {
	o.Disk.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateISCSIExtentParams) GetSerial() string {
	if o == nil || o.Serial.Get() == nil {
		var ret string
		return ret
	}
	return *o.Serial.Get()
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateISCSIExtentParams) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Serial.Get(), o.Serial.IsSet()
}

// HasSerial returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasSerial() bool {
	if o != nil && o.Serial.IsSet() {
		return true
	}

	return false
}

// SetSerial gets a reference to the given NullableString and assigns it to the Serial field.
func (o *CreateISCSIExtentParams) SetSerial(v string) {
	o.Serial.Set(&v)
}
// SetSerialNil sets the value for Serial to be an explicit nil
func (o *CreateISCSIExtentParams) SetSerialNil() {
	o.Serial.Set(nil)
}

// UnsetSerial ensures that no value is present for Serial, not even an explicit nil
func (o *CreateISCSIExtentParams) UnsetSerial() {
	o.Serial.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateISCSIExtentParams) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateISCSIExtentParams) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *CreateISCSIExtentParams) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *CreateISCSIExtentParams) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *CreateISCSIExtentParams) UnsetPath() {
	o.Path.Unset()
}

// GetFilesize returns the Filesize field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetFilesize() int32 {
	if o == nil || o.Filesize == nil {
		var ret int32
		return ret
	}
	return *o.Filesize
}

// GetFilesizeOk returns a tuple with the Filesize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetFilesizeOk() (*int32, bool) {
	if o == nil || o.Filesize == nil {
		return nil, false
	}
	return o.Filesize, true
}

// HasFilesize returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasFilesize() bool {
	if o != nil && o.Filesize != nil {
		return true
	}

	return false
}

// SetFilesize gets a reference to the given int32 and assigns it to the Filesize field.
func (o *CreateISCSIExtentParams) SetFilesize(v int32) {
	o.Filesize = &v
}

// GetBlocksize returns the Blocksize field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetBlocksize() int32 {
	if o == nil || o.Blocksize == nil {
		var ret int32
		return ret
	}
	return *o.Blocksize
}

// GetBlocksizeOk returns a tuple with the Blocksize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetBlocksizeOk() (*int32, bool) {
	if o == nil || o.Blocksize == nil {
		return nil, false
	}
	return o.Blocksize, true
}

// HasBlocksize returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasBlocksize() bool {
	if o != nil && o.Blocksize != nil {
		return true
	}

	return false
}

// SetBlocksize gets a reference to the given int32 and assigns it to the Blocksize field.
func (o *CreateISCSIExtentParams) SetBlocksize(v int32) {
	o.Blocksize = &v
}

// GetPblocksize returns the Pblocksize field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetPblocksize() bool {
	if o == nil || o.Pblocksize == nil {
		var ret bool
		return ret
	}
	return *o.Pblocksize
}

// GetPblocksizeOk returns a tuple with the Pblocksize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetPblocksizeOk() (*bool, bool) {
	if o == nil || o.Pblocksize == nil {
		return nil, false
	}
	return o.Pblocksize, true
}

// HasPblocksize returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasPblocksize() bool {
	if o != nil && o.Pblocksize != nil {
		return true
	}

	return false
}

// SetPblocksize gets a reference to the given bool and assigns it to the Pblocksize field.
func (o *CreateISCSIExtentParams) SetPblocksize(v bool) {
	o.Pblocksize = &v
}

// GetAvailThreshold returns the AvailThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateISCSIExtentParams) GetAvailThreshold() int32 {
	if o == nil || o.AvailThreshold.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AvailThreshold.Get()
}

// GetAvailThresholdOk returns a tuple with the AvailThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateISCSIExtentParams) GetAvailThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailThreshold.Get(), o.AvailThreshold.IsSet()
}

// HasAvailThreshold returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasAvailThreshold() bool {
	if o != nil && o.AvailThreshold.IsSet() {
		return true
	}

	return false
}

// SetAvailThreshold gets a reference to the given NullableInt32 and assigns it to the AvailThreshold field.
func (o *CreateISCSIExtentParams) SetAvailThreshold(v int32) {
	o.AvailThreshold.Set(&v)
}
// SetAvailThresholdNil sets the value for AvailThreshold to be an explicit nil
func (o *CreateISCSIExtentParams) SetAvailThresholdNil() {
	o.AvailThreshold.Set(nil)
}

// UnsetAvailThreshold ensures that no value is present for AvailThreshold, not even an explicit nil
func (o *CreateISCSIExtentParams) UnsetAvailThreshold() {
	o.AvailThreshold.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateISCSIExtentParams) SetComment(v string) {
	o.Comment = &v
}

// GetInsecureTpc returns the InsecureTpc field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetInsecureTpc() bool {
	if o == nil || o.InsecureTpc == nil {
		var ret bool
		return ret
	}
	return *o.InsecureTpc
}

// GetInsecureTpcOk returns a tuple with the InsecureTpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetInsecureTpcOk() (*bool, bool) {
	if o == nil || o.InsecureTpc == nil {
		return nil, false
	}
	return o.InsecureTpc, true
}

// HasInsecureTpc returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasInsecureTpc() bool {
	if o != nil && o.InsecureTpc != nil {
		return true
	}

	return false
}

// SetInsecureTpc gets a reference to the given bool and assigns it to the InsecureTpc field.
func (o *CreateISCSIExtentParams) SetInsecureTpc(v bool) {
	o.InsecureTpc = &v
}

// GetXen returns the Xen field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetXen() bool {
	if o == nil || o.Xen == nil {
		var ret bool
		return ret
	}
	return *o.Xen
}

// GetXenOk returns a tuple with the Xen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetXenOk() (*bool, bool) {
	if o == nil || o.Xen == nil {
		return nil, false
	}
	return o.Xen, true
}

// HasXen returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasXen() bool {
	if o != nil && o.Xen != nil {
		return true
	}

	return false
}

// SetXen gets a reference to the given bool and assigns it to the Xen field.
func (o *CreateISCSIExtentParams) SetXen(v bool) {
	o.Xen = &v
}

// GetRpm returns the Rpm field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetRpm() string {
	if o == nil || o.Rpm == nil {
		var ret string
		return ret
	}
	return *o.Rpm
}

// GetRpmOk returns a tuple with the Rpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetRpmOk() (*string, bool) {
	if o == nil || o.Rpm == nil {
		return nil, false
	}
	return o.Rpm, true
}

// HasRpm returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasRpm() bool {
	if o != nil && o.Rpm != nil {
		return true
	}

	return false
}

// SetRpm gets a reference to the given string and assigns it to the Rpm field.
func (o *CreateISCSIExtentParams) SetRpm(v string) {
	o.Rpm = &v
}

// GetRo returns the Ro field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetRo() bool {
	if o == nil || o.Ro == nil {
		var ret bool
		return ret
	}
	return *o.Ro
}

// GetRoOk returns a tuple with the Ro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetRoOk() (*bool, bool) {
	if o == nil || o.Ro == nil {
		return nil, false
	}
	return o.Ro, true
}

// HasRo returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasRo() bool {
	if o != nil && o.Ro != nil {
		return true
	}

	return false
}

// SetRo gets a reference to the given bool and assigns it to the Ro field.
func (o *CreateISCSIExtentParams) SetRo(v bool) {
	o.Ro = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateISCSIExtentParams) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIExtentParams) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateISCSIExtentParams) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateISCSIExtentParams) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CreateISCSIExtentParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Disk.IsSet() {
		toSerialize["disk"] = o.Disk.Get()
	}
	if o.Serial.IsSet() {
		toSerialize["serial"] = o.Serial.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Filesize != nil {
		toSerialize["filesize"] = o.Filesize
	}
	if o.Blocksize != nil {
		toSerialize["blocksize"] = o.Blocksize
	}
	if o.Pblocksize != nil {
		toSerialize["pblocksize"] = o.Pblocksize
	}
	if o.AvailThreshold.IsSet() {
		toSerialize["avail_threshold"] = o.AvailThreshold.Get()
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.InsecureTpc != nil {
		toSerialize["insecure_tpc"] = o.InsecureTpc
	}
	if o.Xen != nil {
		toSerialize["xen"] = o.Xen
	}
	if o.Rpm != nil {
		toSerialize["rpm"] = o.Rpm
	}
	if o.Ro != nil {
		toSerialize["ro"] = o.Ro
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableCreateISCSIExtentParams struct {
	value *CreateISCSIExtentParams
	isSet bool
}

func (v NullableCreateISCSIExtentParams) Get() *CreateISCSIExtentParams {
	return v.value
}

func (v *NullableCreateISCSIExtentParams) Set(val *CreateISCSIExtentParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateISCSIExtentParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateISCSIExtentParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateISCSIExtentParams(val *CreateISCSIExtentParams) *NullableCreateISCSIExtentParams {
	return &NullableCreateISCSIExtentParams{value: val, isSet: true}
}

func (v NullableCreateISCSIExtentParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateISCSIExtentParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


