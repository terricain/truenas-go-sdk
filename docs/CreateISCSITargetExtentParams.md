# CreateISCSITargetExtentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **int32** |  | 
**Lunid** | Pointer to **NullableInt32** |  | [optional] 
**Extent** | **int32** |  | 

## Methods

### NewCreateISCSITargetExtentParams

`func NewCreateISCSITargetExtentParams(target int32, extent int32, ) *CreateISCSITargetExtentParams`

NewCreateISCSITargetExtentParams instantiates a new CreateISCSITargetExtentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateISCSITargetExtentParamsWithDefaults

`func NewCreateISCSITargetExtentParamsWithDefaults() *CreateISCSITargetExtentParams`

NewCreateISCSITargetExtentParamsWithDefaults instantiates a new CreateISCSITargetExtentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *CreateISCSITargetExtentParams) GetTarget() int32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CreateISCSITargetExtentParams) GetTargetOk() (*int32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CreateISCSITargetExtentParams) SetTarget(v int32)`

SetTarget sets Target field to given value.


### GetLunid

`func (o *CreateISCSITargetExtentParams) GetLunid() int32`

GetLunid returns the Lunid field if non-nil, zero value otherwise.

### GetLunidOk

`func (o *CreateISCSITargetExtentParams) GetLunidOk() (*int32, bool)`

GetLunidOk returns a tuple with the Lunid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunid

`func (o *CreateISCSITargetExtentParams) SetLunid(v int32)`

SetLunid sets Lunid field to given value.

### HasLunid

`func (o *CreateISCSITargetExtentParams) HasLunid() bool`

HasLunid returns a boolean if a field has been set.

### SetLunidNil

`func (o *CreateISCSITargetExtentParams) SetLunidNil(b bool)`

 SetLunidNil sets the value for Lunid to be an explicit nil

### UnsetLunid
`func (o *CreateISCSITargetExtentParams) UnsetLunid()`

UnsetLunid ensures that no value is present for Lunid, not even an explicit nil
### GetExtent

`func (o *CreateISCSITargetExtentParams) GetExtent() int32`

GetExtent returns the Extent field if non-nil, zero value otherwise.

### GetExtentOk

`func (o *CreateISCSITargetExtentParams) GetExtentOk() (*int32, bool)`

GetExtentOk returns a tuple with the Extent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtent

`func (o *CreateISCSITargetExtentParams) SetExtent(v int32)`

SetExtent sets Extent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


