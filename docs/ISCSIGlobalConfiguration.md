# ISCSIGlobalConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Basename** | **string** |  | 
**IsnsServers** | **[]string** |  | 
**PoolAvailThreshold** | **NullableInt32** |  | 
**Alua** | **bool** |  | 

## Methods

### NewISCSIGlobalConfiguration

`func NewISCSIGlobalConfiguration(id int32, basename string, isnsServers []string, poolAvailThreshold NullableInt32, alua bool, ) *ISCSIGlobalConfiguration`

NewISCSIGlobalConfiguration instantiates a new ISCSIGlobalConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewISCSIGlobalConfigurationWithDefaults

`func NewISCSIGlobalConfigurationWithDefaults() *ISCSIGlobalConfiguration`

NewISCSIGlobalConfigurationWithDefaults instantiates a new ISCSIGlobalConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ISCSIGlobalConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ISCSIGlobalConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ISCSIGlobalConfiguration) SetId(v int32)`

SetId sets Id field to given value.


### GetBasename

`func (o *ISCSIGlobalConfiguration) GetBasename() string`

GetBasename returns the Basename field if non-nil, zero value otherwise.

### GetBasenameOk

`func (o *ISCSIGlobalConfiguration) GetBasenameOk() (*string, bool)`

GetBasenameOk returns a tuple with the Basename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasename

`func (o *ISCSIGlobalConfiguration) SetBasename(v string)`

SetBasename sets Basename field to given value.


### GetIsnsServers

`func (o *ISCSIGlobalConfiguration) GetIsnsServers() []string`

GetIsnsServers returns the IsnsServers field if non-nil, zero value otherwise.

### GetIsnsServersOk

`func (o *ISCSIGlobalConfiguration) GetIsnsServersOk() (*[]string, bool)`

GetIsnsServersOk returns a tuple with the IsnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsnsServers

`func (o *ISCSIGlobalConfiguration) SetIsnsServers(v []string)`

SetIsnsServers sets IsnsServers field to given value.


### GetPoolAvailThreshold

`func (o *ISCSIGlobalConfiguration) GetPoolAvailThreshold() int32`

GetPoolAvailThreshold returns the PoolAvailThreshold field if non-nil, zero value otherwise.

### GetPoolAvailThresholdOk

`func (o *ISCSIGlobalConfiguration) GetPoolAvailThresholdOk() (*int32, bool)`

GetPoolAvailThresholdOk returns a tuple with the PoolAvailThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAvailThreshold

`func (o *ISCSIGlobalConfiguration) SetPoolAvailThreshold(v int32)`

SetPoolAvailThreshold sets PoolAvailThreshold field to given value.


### SetPoolAvailThresholdNil

`func (o *ISCSIGlobalConfiguration) SetPoolAvailThresholdNil(b bool)`

 SetPoolAvailThresholdNil sets the value for PoolAvailThreshold to be an explicit nil

### UnsetPoolAvailThreshold
`func (o *ISCSIGlobalConfiguration) UnsetPoolAvailThreshold()`

UnsetPoolAvailThreshold ensures that no value is present for PoolAvailThreshold, not even an explicit nil
### GetAlua

`func (o *ISCSIGlobalConfiguration) GetAlua() bool`

GetAlua returns the Alua field if non-nil, zero value otherwise.

### GetAluaOk

`func (o *ISCSIGlobalConfiguration) GetAluaOk() (*bool, bool)`

GetAluaOk returns a tuple with the Alua field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlua

`func (o *ISCSIGlobalConfiguration) SetAlua(v bool)`

SetAlua sets Alua field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


