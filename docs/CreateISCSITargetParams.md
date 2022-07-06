# CreateISCSITargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Alias** | Pointer to **NullableString** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]CreateISCSITargetParamsGroupsInner**](CreateISCSITargetParamsGroupsInner.md) |  | [optional] 

## Methods

### NewCreateISCSITargetParams

`func NewCreateISCSITargetParams(name string, ) *CreateISCSITargetParams`

NewCreateISCSITargetParams instantiates a new CreateISCSITargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateISCSITargetParamsWithDefaults

`func NewCreateISCSITargetParamsWithDefaults() *CreateISCSITargetParams`

NewCreateISCSITargetParamsWithDefaults instantiates a new CreateISCSITargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateISCSITargetParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateISCSITargetParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateISCSITargetParams) SetName(v string)`

SetName sets Name field to given value.


### GetAlias

`func (o *CreateISCSITargetParams) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CreateISCSITargetParams) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CreateISCSITargetParams) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CreateISCSITargetParams) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *CreateISCSITargetParams) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *CreateISCSITargetParams) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetMode

`func (o *CreateISCSITargetParams) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateISCSITargetParams) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateISCSITargetParams) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateISCSITargetParams) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetGroups

`func (o *CreateISCSITargetParams) GetGroups() []CreateISCSITargetParamsGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateISCSITargetParams) GetGroupsOk() (*[]CreateISCSITargetParamsGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateISCSITargetParams) SetGroups(v []CreateISCSITargetParamsGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *CreateISCSITargetParams) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


