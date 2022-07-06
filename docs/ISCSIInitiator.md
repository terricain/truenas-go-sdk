# ISCSIInitiator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Initiators** | **[]map[string]interface{}** |  | 
**AuthNetwork** | **[]string** |  | 
**Comment** | **string** |  | 

## Methods

### NewISCSIInitiator

`func NewISCSIInitiator(id int32, initiators []map[string]interface{}, authNetwork []string, comment string, ) *ISCSIInitiator`

NewISCSIInitiator instantiates a new ISCSIInitiator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewISCSIInitiatorWithDefaults

`func NewISCSIInitiatorWithDefaults() *ISCSIInitiator`

NewISCSIInitiatorWithDefaults instantiates a new ISCSIInitiator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ISCSIInitiator) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ISCSIInitiator) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ISCSIInitiator) SetId(v int32)`

SetId sets Id field to given value.


### GetInitiators

`func (o *ISCSIInitiator) GetInitiators() []map[string]interface{}`

GetInitiators returns the Initiators field if non-nil, zero value otherwise.

### GetInitiatorsOk

`func (o *ISCSIInitiator) GetInitiatorsOk() (*[]map[string]interface{}, bool)`

GetInitiatorsOk returns a tuple with the Initiators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiators

`func (o *ISCSIInitiator) SetInitiators(v []map[string]interface{})`

SetInitiators sets Initiators field to given value.


### GetAuthNetwork

`func (o *ISCSIInitiator) GetAuthNetwork() []string`

GetAuthNetwork returns the AuthNetwork field if non-nil, zero value otherwise.

### GetAuthNetworkOk

`func (o *ISCSIInitiator) GetAuthNetworkOk() (*[]string, bool)`

GetAuthNetworkOk returns a tuple with the AuthNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNetwork

`func (o *ISCSIInitiator) SetAuthNetwork(v []string)`

SetAuthNetwork sets AuthNetwork field to given value.


### GetComment

`func (o *ISCSIInitiator) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ISCSIInitiator) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ISCSIInitiator) SetComment(v string)`

SetComment sets Comment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


