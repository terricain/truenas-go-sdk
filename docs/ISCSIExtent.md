# ISCSIExtent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Disk** | Pointer to **NullableString** |  | [optional] 
**Serial** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Filesize** | Pointer to **int32** |  | [optional] 
**Blocksize** | Pointer to **int32** |  | [optional] 
**Pblocksize** | Pointer to **bool** |  | [optional] 
**AvailThreshold** | Pointer to **NullableInt32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**InsecureTpc** | Pointer to **bool** |  | [optional] 
**Xen** | Pointer to **bool** |  | [optional] 
**Rpm** | Pointer to **string** |  | [optional] 
**Ro** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewISCSIExtent

`func NewISCSIExtent() *ISCSIExtent`

NewISCSIExtent instantiates a new ISCSIExtent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewISCSIExtentWithDefaults

`func NewISCSIExtentWithDefaults() *ISCSIExtent`

NewISCSIExtentWithDefaults instantiates a new ISCSIExtent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ISCSIExtent) GetName() int32`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ISCSIExtent) GetNameOk() (*int32, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ISCSIExtent) SetName(v int32)`

SetName sets Name field to given value.

### HasName

`func (o *ISCSIExtent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ISCSIExtent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ISCSIExtent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ISCSIExtent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ISCSIExtent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisk

`func (o *ISCSIExtent) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ISCSIExtent) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ISCSIExtent) SetDisk(v string)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ISCSIExtent) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *ISCSIExtent) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *ISCSIExtent) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetSerial

`func (o *ISCSIExtent) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ISCSIExtent) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ISCSIExtent) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ISCSIExtent) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *ISCSIExtent) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *ISCSIExtent) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetPath

`func (o *ISCSIExtent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ISCSIExtent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ISCSIExtent) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ISCSIExtent) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ISCSIExtent) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ISCSIExtent) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetFilesize

`func (o *ISCSIExtent) GetFilesize() int32`

GetFilesize returns the Filesize field if non-nil, zero value otherwise.

### GetFilesizeOk

`func (o *ISCSIExtent) GetFilesizeOk() (*int32, bool)`

GetFilesizeOk returns a tuple with the Filesize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesize

`func (o *ISCSIExtent) SetFilesize(v int32)`

SetFilesize sets Filesize field to given value.

### HasFilesize

`func (o *ISCSIExtent) HasFilesize() bool`

HasFilesize returns a boolean if a field has been set.

### GetBlocksize

`func (o *ISCSIExtent) GetBlocksize() int32`

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

`func (o *ISCSIExtent) GetBlocksizeOk() (*int32, bool)`

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

`func (o *ISCSIExtent) SetBlocksize(v int32)`

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *ISCSIExtent) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.

### GetPblocksize

`func (o *ISCSIExtent) GetPblocksize() bool`

GetPblocksize returns the Pblocksize field if non-nil, zero value otherwise.

### GetPblocksizeOk

`func (o *ISCSIExtent) GetPblocksizeOk() (*bool, bool)`

GetPblocksizeOk returns a tuple with the Pblocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPblocksize

`func (o *ISCSIExtent) SetPblocksize(v bool)`

SetPblocksize sets Pblocksize field to given value.

### HasPblocksize

`func (o *ISCSIExtent) HasPblocksize() bool`

HasPblocksize returns a boolean if a field has been set.

### GetAvailThreshold

`func (o *ISCSIExtent) GetAvailThreshold() int32`

GetAvailThreshold returns the AvailThreshold field if non-nil, zero value otherwise.

### GetAvailThresholdOk

`func (o *ISCSIExtent) GetAvailThresholdOk() (*int32, bool)`

GetAvailThresholdOk returns a tuple with the AvailThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailThreshold

`func (o *ISCSIExtent) SetAvailThreshold(v int32)`

SetAvailThreshold sets AvailThreshold field to given value.

### HasAvailThreshold

`func (o *ISCSIExtent) HasAvailThreshold() bool`

HasAvailThreshold returns a boolean if a field has been set.

### SetAvailThresholdNil

`func (o *ISCSIExtent) SetAvailThresholdNil(b bool)`

 SetAvailThresholdNil sets the value for AvailThreshold to be an explicit nil

### UnsetAvailThreshold
`func (o *ISCSIExtent) UnsetAvailThreshold()`

UnsetAvailThreshold ensures that no value is present for AvailThreshold, not even an explicit nil
### GetComment

`func (o *ISCSIExtent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ISCSIExtent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ISCSIExtent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ISCSIExtent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetInsecureTpc

`func (o *ISCSIExtent) GetInsecureTpc() bool`

GetInsecureTpc returns the InsecureTpc field if non-nil, zero value otherwise.

### GetInsecureTpcOk

`func (o *ISCSIExtent) GetInsecureTpcOk() (*bool, bool)`

GetInsecureTpcOk returns a tuple with the InsecureTpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureTpc

`func (o *ISCSIExtent) SetInsecureTpc(v bool)`

SetInsecureTpc sets InsecureTpc field to given value.

### HasInsecureTpc

`func (o *ISCSIExtent) HasInsecureTpc() bool`

HasInsecureTpc returns a boolean if a field has been set.

### GetXen

`func (o *ISCSIExtent) GetXen() bool`

GetXen returns the Xen field if non-nil, zero value otherwise.

### GetXenOk

`func (o *ISCSIExtent) GetXenOk() (*bool, bool)`

GetXenOk returns a tuple with the Xen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXen

`func (o *ISCSIExtent) SetXen(v bool)`

SetXen sets Xen field to given value.

### HasXen

`func (o *ISCSIExtent) HasXen() bool`

HasXen returns a boolean if a field has been set.

### GetRpm

`func (o *ISCSIExtent) GetRpm() string`

GetRpm returns the Rpm field if non-nil, zero value otherwise.

### GetRpmOk

`func (o *ISCSIExtent) GetRpmOk() (*string, bool)`

GetRpmOk returns a tuple with the Rpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpm

`func (o *ISCSIExtent) SetRpm(v string)`

SetRpm sets Rpm field to given value.

### HasRpm

`func (o *ISCSIExtent) HasRpm() bool`

HasRpm returns a boolean if a field has been set.

### GetRo

`func (o *ISCSIExtent) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *ISCSIExtent) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *ISCSIExtent) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *ISCSIExtent) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetEnabled

`func (o *ISCSIExtent) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ISCSIExtent) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ISCSIExtent) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ISCSIExtent) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


