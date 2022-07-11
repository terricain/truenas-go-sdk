# ISCSITarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **NullableString** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]ISCSITargetGroupsInner**](ISCSITargetGroupsInner.md) |  | [optional] 

## Methods

### NewISCSITarget

`func NewISCSITarget() *ISCSITarget`

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

### HasId

`func (o *ISCSITarget) HasId() bool`

HasId returns a boolean if a field has been set.

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

### HasName

`func (o *ISCSITarget) HasName() bool`

HasName returns a boolean if a field has been set.

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

### HasAlias

`func (o *ISCSITarget) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

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

### HasMode

`func (o *ISCSITarget) HasMode() bool`

HasMode returns a boolean if a field has been set.

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

### HasGroups

`func (o *ISCSITarget) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


