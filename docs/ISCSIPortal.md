# ISCSIPortal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Tag** | **int32** |  | 
**Comment** | **string** |  | 
**DiscoveryAuthmethod** | **string** |  | 
**DiscoveryAuthgroup** | **NullableInt32** |  | 
**Listen** | [**[]ISCSIPortalListenInner**](ISCSIPortalListenInner.md) |  | 

## Methods

### NewISCSIPortal

`func NewISCSIPortal(id int32, tag int32, comment string, discoveryAuthmethod string, discoveryAuthgroup NullableInt32, listen []ISCSIPortalListenInner, ) *ISCSIPortal`

NewISCSIPortal instantiates a new ISCSIPortal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewISCSIPortalWithDefaults

`func NewISCSIPortalWithDefaults() *ISCSIPortal`

NewISCSIPortalWithDefaults instantiates a new ISCSIPortal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ISCSIPortal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ISCSIPortal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ISCSIPortal) SetId(v int32)`

SetId sets Id field to given value.


### GetTag

`func (o *ISCSIPortal) GetTag() int32`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ISCSIPortal) GetTagOk() (*int32, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ISCSIPortal) SetTag(v int32)`

SetTag sets Tag field to given value.


### GetComment

`func (o *ISCSIPortal) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ISCSIPortal) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ISCSIPortal) SetComment(v string)`

SetComment sets Comment field to given value.


### GetDiscoveryAuthmethod

`func (o *ISCSIPortal) GetDiscoveryAuthmethod() string`

GetDiscoveryAuthmethod returns the DiscoveryAuthmethod field if non-nil, zero value otherwise.

### GetDiscoveryAuthmethodOk

`func (o *ISCSIPortal) GetDiscoveryAuthmethodOk() (*string, bool)`

GetDiscoveryAuthmethodOk returns a tuple with the DiscoveryAuthmethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryAuthmethod

`func (o *ISCSIPortal) SetDiscoveryAuthmethod(v string)`

SetDiscoveryAuthmethod sets DiscoveryAuthmethod field to given value.


### GetDiscoveryAuthgroup

`func (o *ISCSIPortal) GetDiscoveryAuthgroup() int32`

GetDiscoveryAuthgroup returns the DiscoveryAuthgroup field if non-nil, zero value otherwise.

### GetDiscoveryAuthgroupOk

`func (o *ISCSIPortal) GetDiscoveryAuthgroupOk() (*int32, bool)`

GetDiscoveryAuthgroupOk returns a tuple with the DiscoveryAuthgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryAuthgroup

`func (o *ISCSIPortal) SetDiscoveryAuthgroup(v int32)`

SetDiscoveryAuthgroup sets DiscoveryAuthgroup field to given value.


### SetDiscoveryAuthgroupNil

`func (o *ISCSIPortal) SetDiscoveryAuthgroupNil(b bool)`

 SetDiscoveryAuthgroupNil sets the value for DiscoveryAuthgroup to be an explicit nil

### UnsetDiscoveryAuthgroup
`func (o *ISCSIPortal) UnsetDiscoveryAuthgroup()`

UnsetDiscoveryAuthgroup ensures that no value is present for DiscoveryAuthgroup, not even an explicit nil
### GetListen

`func (o *ISCSIPortal) GetListen() []ISCSIPortalListenInner`

GetListen returns the Listen field if non-nil, zero value otherwise.

### GetListenOk

`func (o *ISCSIPortal) GetListenOk() (*[]ISCSIPortalListenInner, bool)`

GetListenOk returns a tuple with the Listen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListen

`func (o *ISCSIPortal) SetListen(v []ISCSIPortalListenInner)`

SetListen sets Listen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


