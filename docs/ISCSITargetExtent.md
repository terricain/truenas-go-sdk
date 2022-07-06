# ISCSITargetExtent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Target** | **int32** |  | 
**Lunid** | Pointer to **int32** |  | [optional] 
**Extent** | **int32** |  | 

## Methods

### NewISCSITargetExtent

`func NewISCSITargetExtent(id int32, target int32, extent int32, ) *ISCSITargetExtent`

NewISCSITargetExtent instantiates a new ISCSITargetExtent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewISCSITargetExtentWithDefaults

`func NewISCSITargetExtentWithDefaults() *ISCSITargetExtent`

NewISCSITargetExtentWithDefaults instantiates a new ISCSITargetExtent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ISCSITargetExtent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ISCSITargetExtent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ISCSITargetExtent) SetId(v int32)`

SetId sets Id field to given value.


### GetTarget

`func (o *ISCSITargetExtent) GetTarget() int32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ISCSITargetExtent) GetTargetOk() (*int32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ISCSITargetExtent) SetTarget(v int32)`

SetTarget sets Target field to given value.


### GetLunid

`func (o *ISCSITargetExtent) GetLunid() int32`

GetLunid returns the Lunid field if non-nil, zero value otherwise.

### GetLunidOk

`func (o *ISCSITargetExtent) GetLunidOk() (*int32, bool)`

GetLunidOk returns a tuple with the Lunid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunid

`func (o *ISCSITargetExtent) SetLunid(v int32)`

SetLunid sets Lunid field to given value.

### HasLunid

`func (o *ISCSITargetExtent) HasLunid() bool`

HasLunid returns a boolean if a field has been set.

### GetExtent

`func (o *ISCSITargetExtent) GetExtent() int32`

GetExtent returns the Extent field if non-nil, zero value otherwise.

### GetExtentOk

`func (o *ISCSITargetExtent) GetExtentOk() (*int32, bool)`

GetExtentOk returns a tuple with the Extent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtent

`func (o *ISCSITargetExtent) SetExtent(v int32)`

SetExtent sets Extent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


