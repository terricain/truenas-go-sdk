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

// checks if the ISCSIExtent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ISCSIExtent{}

// ISCSIExtent struct for ISCSIExtent
type ISCSIExtent struct {
	Id             int32          `json:"id"`
	Name           string         `json:"name"`
	Type           string         `json:"type"`
	Disk           NullableString `json:"disk,omitempty"`
	Serial         NullableString `json:"serial,omitempty"`
	Path           NullableString `json:"path,omitempty"`
	Filesize       *string        `json:"filesize,omitempty"`
	Blocksize      *int32         `json:"blocksize,omitempty"`
	Pblocksize     *bool          `json:"pblocksize,omitempty"`
	AvailThreshold NullableInt32  `json:"avail_threshold,omitempty"`
	Comment        *string        `json:"comment,omitempty"`
	InsecureTpc    *bool          `json:"insecure_tpc,omitempty"`
	Xen            *bool          `json:"xen,omitempty"`
	Rpm            *string        `json:"rpm,omitempty"`
	Ro             *bool          `json:"ro,omitempty"`
	Enabled        *bool          `json:"enabled,omitempty"`
}

type _ISCSIExtent ISCSIExtent

// NewISCSIExtent instantiates a new ISCSIExtent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewISCSIExtent(id int32, name string, type_ string) *ISCSIExtent {
	this := ISCSIExtent{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewISCSIExtentWithDefaults instantiates a new ISCSIExtent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewISCSIExtentWithDefaults() *ISCSIExtent {
	this := ISCSIExtent{}
	return &this
}

// GetId returns the Id field value
func (o *ISCSIExtent) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ISCSIExtent) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ISCSIExtent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ISCSIExtent) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ISCSIExtent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ISCSIExtent) SetType(v string) {
	o.Type = v
}

// GetDisk returns the Disk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISCSIExtent) GetDisk() string {
	if o == nil || IsNil(o.Disk.Get()) {
		var ret string
		return ret
	}
	return *o.Disk.Get()
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISCSIExtent) GetDiskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disk.Get(), o.Disk.IsSet()
}

// HasDisk returns a boolean if a field has been set.
func (o *ISCSIExtent) HasDisk() bool {
	if o != nil && o.Disk.IsSet() {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NullableString and assigns it to the Disk field.
func (o *ISCSIExtent) SetDisk(v string) {
	o.Disk.Set(&v)
}

// SetDiskNil sets the value for Disk to be an explicit nil
func (o *ISCSIExtent) SetDiskNil() {
	o.Disk.Set(nil)
}

// UnsetDisk ensures that no value is present for Disk, not even an explicit nil
func (o *ISCSIExtent) UnsetDisk() {
	o.Disk.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISCSIExtent) GetSerial() string {
	if o == nil || IsNil(o.Serial.Get()) {
		var ret string
		return ret
	}
	return *o.Serial.Get()
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISCSIExtent) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Serial.Get(), o.Serial.IsSet()
}

// HasSerial returns a boolean if a field has been set.
func (o *ISCSIExtent) HasSerial() bool {
	if o != nil && o.Serial.IsSet() {
		return true
	}

	return false
}

// SetSerial gets a reference to the given NullableString and assigns it to the Serial field.
func (o *ISCSIExtent) SetSerial(v string) {
	o.Serial.Set(&v)
}

// SetSerialNil sets the value for Serial to be an explicit nil
func (o *ISCSIExtent) SetSerialNil() {
	o.Serial.Set(nil)
}

// UnsetSerial ensures that no value is present for Serial, not even an explicit nil
func (o *ISCSIExtent) UnsetSerial() {
	o.Serial.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISCSIExtent) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISCSIExtent) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *ISCSIExtent) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *ISCSIExtent) SetPath(v string) {
	o.Path.Set(&v)
}

// SetPathNil sets the value for Path to be an explicit nil
func (o *ISCSIExtent) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *ISCSIExtent) UnsetPath() {
	o.Path.Unset()
}

// GetFilesize returns the Filesize field value if set, zero value otherwise.
func (o *ISCSIExtent) GetFilesize() string {
	if o == nil || IsNil(o.Filesize) {
		var ret string
		return ret
	}
	return *o.Filesize
}

// GetFilesizeOk returns a tuple with the Filesize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetFilesizeOk() (*string, bool) {
	if o == nil || IsNil(o.Filesize) {
		return nil, false
	}
	return o.Filesize, true
}

// HasFilesize returns a boolean if a field has been set.
func (o *ISCSIExtent) HasFilesize() bool {
	if o != nil && !IsNil(o.Filesize) {
		return true
	}

	return false
}

// SetFilesize gets a reference to the given string and assigns it to the Filesize field.
func (o *ISCSIExtent) SetFilesize(v string) {
	o.Filesize = &v
}

// GetBlocksize returns the Blocksize field value if set, zero value otherwise.
func (o *ISCSIExtent) GetBlocksize() int32 {
	if o == nil || IsNil(o.Blocksize) {
		var ret int32
		return ret
	}
	return *o.Blocksize
}

// GetBlocksizeOk returns a tuple with the Blocksize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetBlocksizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Blocksize) {
		return nil, false
	}
	return o.Blocksize, true
}

// HasBlocksize returns a boolean if a field has been set.
func (o *ISCSIExtent) HasBlocksize() bool {
	if o != nil && !IsNil(o.Blocksize) {
		return true
	}

	return false
}

// SetBlocksize gets a reference to the given int32 and assigns it to the Blocksize field.
func (o *ISCSIExtent) SetBlocksize(v int32) {
	o.Blocksize = &v
}

// GetPblocksize returns the Pblocksize field value if set, zero value otherwise.
func (o *ISCSIExtent) GetPblocksize() bool {
	if o == nil || IsNil(o.Pblocksize) {
		var ret bool
		return ret
	}
	return *o.Pblocksize
}

// GetPblocksizeOk returns a tuple with the Pblocksize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetPblocksizeOk() (*bool, bool) {
	if o == nil || IsNil(o.Pblocksize) {
		return nil, false
	}
	return o.Pblocksize, true
}

// HasPblocksize returns a boolean if a field has been set.
func (o *ISCSIExtent) HasPblocksize() bool {
	if o != nil && !IsNil(o.Pblocksize) {
		return true
	}

	return false
}

// SetPblocksize gets a reference to the given bool and assigns it to the Pblocksize field.
func (o *ISCSIExtent) SetPblocksize(v bool) {
	o.Pblocksize = &v
}

// GetAvailThreshold returns the AvailThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ISCSIExtent) GetAvailThreshold() int32 {
	if o == nil || IsNil(o.AvailThreshold.Get()) {
		var ret int32
		return ret
	}
	return *o.AvailThreshold.Get()
}

// GetAvailThresholdOk returns a tuple with the AvailThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ISCSIExtent) GetAvailThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailThreshold.Get(), o.AvailThreshold.IsSet()
}

// HasAvailThreshold returns a boolean if a field has been set.
func (o *ISCSIExtent) HasAvailThreshold() bool {
	if o != nil && o.AvailThreshold.IsSet() {
		return true
	}

	return false
}

// SetAvailThreshold gets a reference to the given NullableInt32 and assigns it to the AvailThreshold field.
func (o *ISCSIExtent) SetAvailThreshold(v int32) {
	o.AvailThreshold.Set(&v)
}

// SetAvailThresholdNil sets the value for AvailThreshold to be an explicit nil
func (o *ISCSIExtent) SetAvailThresholdNil() {
	o.AvailThreshold.Set(nil)
}

// UnsetAvailThreshold ensures that no value is present for AvailThreshold, not even an explicit nil
func (o *ISCSIExtent) UnsetAvailThreshold() {
	o.AvailThreshold.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ISCSIExtent) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ISCSIExtent) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ISCSIExtent) SetComment(v string) {
	o.Comment = &v
}

// GetInsecureTpc returns the InsecureTpc field value if set, zero value otherwise.
func (o *ISCSIExtent) GetInsecureTpc() bool {
	if o == nil || IsNil(o.InsecureTpc) {
		var ret bool
		return ret
	}
	return *o.InsecureTpc
}

// GetInsecureTpcOk returns a tuple with the InsecureTpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetInsecureTpcOk() (*bool, bool) {
	if o == nil || IsNil(o.InsecureTpc) {
		return nil, false
	}
	return o.InsecureTpc, true
}

// HasInsecureTpc returns a boolean if a field has been set.
func (o *ISCSIExtent) HasInsecureTpc() bool {
	if o != nil && !IsNil(o.InsecureTpc) {
		return true
	}

	return false
}

// SetInsecureTpc gets a reference to the given bool and assigns it to the InsecureTpc field.
func (o *ISCSIExtent) SetInsecureTpc(v bool) {
	o.InsecureTpc = &v
}

// GetXen returns the Xen field value if set, zero value otherwise.
func (o *ISCSIExtent) GetXen() bool {
	if o == nil || IsNil(o.Xen) {
		var ret bool
		return ret
	}
	return *o.Xen
}

// GetXenOk returns a tuple with the Xen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetXenOk() (*bool, bool) {
	if o == nil || IsNil(o.Xen) {
		return nil, false
	}
	return o.Xen, true
}

// HasXen returns a boolean if a field has been set.
func (o *ISCSIExtent) HasXen() bool {
	if o != nil && !IsNil(o.Xen) {
		return true
	}

	return false
}

// SetXen gets a reference to the given bool and assigns it to the Xen field.
func (o *ISCSIExtent) SetXen(v bool) {
	o.Xen = &v
}

// GetRpm returns the Rpm field value if set, zero value otherwise.
func (o *ISCSIExtent) GetRpm() string {
	if o == nil || IsNil(o.Rpm) {
		var ret string
		return ret
	}
	return *o.Rpm
}

// GetRpmOk returns a tuple with the Rpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetRpmOk() (*string, bool) {
	if o == nil || IsNil(o.Rpm) {
		return nil, false
	}
	return o.Rpm, true
}

// HasRpm returns a boolean if a field has been set.
func (o *ISCSIExtent) HasRpm() bool {
	if o != nil && !IsNil(o.Rpm) {
		return true
	}

	return false
}

// SetRpm gets a reference to the given string and assigns it to the Rpm field.
func (o *ISCSIExtent) SetRpm(v string) {
	o.Rpm = &v
}

// GetRo returns the Ro field value if set, zero value otherwise.
func (o *ISCSIExtent) GetRo() bool {
	if o == nil || IsNil(o.Ro) {
		var ret bool
		return ret
	}
	return *o.Ro
}

// GetRoOk returns a tuple with the Ro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetRoOk() (*bool, bool) {
	if o == nil || IsNil(o.Ro) {
		return nil, false
	}
	return o.Ro, true
}

// HasRo returns a boolean if a field has been set.
func (o *ISCSIExtent) HasRo() bool {
	if o != nil && !IsNil(o.Ro) {
		return true
	}

	return false
}

// SetRo gets a reference to the given bool and assigns it to the Ro field.
func (o *ISCSIExtent) SetRo(v bool) {
	o.Ro = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ISCSIExtent) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ISCSIExtent) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ISCSIExtent) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ISCSIExtent) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ISCSIExtent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ISCSIExtent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.Disk.IsSet() {
		toSerialize["disk"] = o.Disk.Get()
	}
	if o.Serial.IsSet() {
		toSerialize["serial"] = o.Serial.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if !IsNil(o.Filesize) {
		toSerialize["filesize"] = o.Filesize
	}
	if !IsNil(o.Blocksize) {
		toSerialize["blocksize"] = o.Blocksize
	}
	if !IsNil(o.Pblocksize) {
		toSerialize["pblocksize"] = o.Pblocksize
	}
	if o.AvailThreshold.IsSet() {
		toSerialize["avail_threshold"] = o.AvailThreshold.Get()
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.InsecureTpc) {
		toSerialize["insecure_tpc"] = o.InsecureTpc
	}
	if !IsNil(o.Xen) {
		toSerialize["xen"] = o.Xen
	}
	if !IsNil(o.Rpm) {
		toSerialize["rpm"] = o.Rpm
	}
	if !IsNil(o.Ro) {
		toSerialize["ro"] = o.Ro
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

func (o *ISCSIExtent) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
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

	varISCSIExtent := _ISCSIExtent{}

	err = json.Unmarshal(bytes, &varISCSIExtent)

	if err != nil {
		return err
	}

	*o = ISCSIExtent(varISCSIExtent)

	return err
}

type NullableISCSIExtent struct {
	value *ISCSIExtent
	isSet bool
}

func (v NullableISCSIExtent) Get() *ISCSIExtent {
	return v.value
}

func (v *NullableISCSIExtent) Set(val *ISCSIExtent) {
	v.value = val
	v.isSet = true
}

func (v NullableISCSIExtent) IsSet() bool {
	return v.isSet
}

func (v *NullableISCSIExtent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableISCSIExtent(val *ISCSIExtent) *NullableISCSIExtent {
	return &NullableISCSIExtent{value: val, isSet: true}
}

func (v NullableISCSIExtent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableISCSIExtent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
