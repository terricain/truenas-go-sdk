# ISCSITargetGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Portal** | **int32** |  | 
**Initiator** | Pointer to **int32** |  | [optional] 
**Auth** | Pointer to **map[string]interface{}** |  | [optional] 
**Authmethod** | **string** |  | 

## Methods

### NewISCSITargetGroupsInner

`func NewISCSITargetGroupsInner(portal int32, authmethod string, ) *ISCSITargetGroupsInner`

NewISCSITargetGroupsInner instantiates a new ISCSITargetGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewISCSITargetGroupsInnerWithDefaults

`func NewISCSITargetGroupsInnerWithDefaults() *ISCSITargetGroupsInner`

NewISCSITargetGroupsInnerWithDefaults instantiates a new ISCSITargetGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortal

`func (o *ISCSITargetGroupsInner) GetPortal() int32`

GetPortal returns the Portal field if non-nil, zero value otherwise.

### GetPortalOk

`func (o *ISCSITargetGroupsInner) GetPortalOk() (*int32, bool)`

GetPortalOk returns a tuple with the Portal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortal

`func (o *ISCSITargetGroupsInner) SetPortal(v int32)`

SetPortal sets Portal field to given value.


### GetInitiator

`func (o *ISCSITargetGroupsInner) GetInitiator() int32`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *ISCSITargetGroupsInner) GetInitiatorOk() (*int32, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *ISCSITargetGroupsInner) SetInitiator(v int32)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *ISCSITargetGroupsInner) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetAuth

`func (o *ISCSITargetGroupsInner) GetAuth() map[string]interface{}`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ISCSITargetGroupsInner) GetAuthOk() (*map[string]interface{}, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ISCSITargetGroupsInner) SetAuth(v map[string]interface{})`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ISCSITargetGroupsInner) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAuthmethod

`func (o *ISCSITargetGroupsInner) GetAuthmethod() string`

GetAuthmethod returns the Authmethod field if non-nil, zero value otherwise.

### GetAuthmethodOk

`func (o *ISCSITargetGroupsInner) GetAuthmethodOk() (*string, bool)`

GetAuthmethodOk returns a tuple with the Authmethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthmethod

`func (o *ISCSITargetGroupsInner) SetAuthmethod(v string)`

SetAuthmethod sets Authmethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


