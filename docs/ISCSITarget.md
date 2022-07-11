# ISCSITarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Alias** | **NullableString** |  | 
**Mode** | **string** |  | 
**Groups** | [**[]ISCSITargetGroupsInner**](ISCSITargetGroupsInner.md) |  | 

## Methods

### NewISCSITarget

`func NewISCSITarget(id int32, name string, alias NullableString, mode string, groups []ISCSITargetGroupsInner, ) *ISCSITarget`

NewISCSITarget instantiates a new ISCSITarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewISCSITargetWithDefaults

`func NewISCSITargetWithDefaults() *ISCSITarget`

NewISCSITargetWithDefaults instantiates a new ISCSITarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ISCSITarget) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ISCSITarget) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ISCSITarget) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ISCSITarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ISCSITarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ISCSITarget) SetName(v string)`

SetName sets Name field to given value.


### GetAlias

`func (o *ISCSITarget) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ISCSITarget) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ISCSITarget) SetAlias(v string)`

SetAlias sets Alias field to given value.


### SetAliasNil

`func (o *ISCSITarget) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *ISCSITarget) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetMode

`func (o *ISCSITarget) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ISCSITarget) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ISCSITarget) SetMode(v string)`

SetMode sets Mode field to given value.


### GetGroups

`func (o *ISCSITarget) GetGroups() []ISCSITargetGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ISCSITarget) GetGroupsOk() (*[]ISCSITargetGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ISCSITarget) SetGroups(v []ISCSITargetGroupsInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


