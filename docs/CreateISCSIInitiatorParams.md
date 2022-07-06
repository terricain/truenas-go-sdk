# CreateISCSIInitiatorParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initiators** | Pointer to **[]string** |  | [optional] 
**AuthNetwork** | Pointer to **[]string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateISCSIInitiatorParams

`func NewCreateISCSIInitiatorParams() *CreateISCSIInitiatorParams`

NewCreateISCSIInitiatorParams instantiates a new CreateISCSIInitiatorParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateISCSIInitiatorParamsWithDefaults

`func NewCreateISCSIInitiatorParamsWithDefaults() *CreateISCSIInitiatorParams`

NewCreateISCSIInitiatorParamsWithDefaults instantiates a new CreateISCSIInitiatorParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiators

`func (o *CreateISCSIInitiatorParams) GetInitiators() []string`

GetInitiators returns the Initiators field if non-nil, zero value otherwise.

### GetInitiatorsOk

`func (o *CreateISCSIInitiatorParams) GetInitiatorsOk() (*[]string, bool)`

GetInitiatorsOk returns a tuple with the Initiators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiators

`func (o *CreateISCSIInitiatorParams) SetInitiators(v []string)`

SetInitiators sets Initiators field to given value.

### HasInitiators

`func (o *CreateISCSIInitiatorParams) HasInitiators() bool`

HasInitiators returns a boolean if a field has been set.

### GetAuthNetwork

`func (o *CreateISCSIInitiatorParams) GetAuthNetwork() []string`

GetAuthNetwork returns the AuthNetwork field if non-nil, zero value otherwise.

### GetAuthNetworkOk

`func (o *CreateISCSIInitiatorParams) GetAuthNetworkOk() (*[]string, bool)`

GetAuthNetworkOk returns a tuple with the AuthNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNetwork

`func (o *CreateISCSIInitiatorParams) SetAuthNetwork(v []string)`

SetAuthNetwork sets AuthNetwork field to given value.

### HasAuthNetwork

`func (o *CreateISCSIInitiatorParams) HasAuthNetwork() bool`

HasAuthNetwork returns a boolean if a field has been set.

### GetComment

`func (o *CreateISCSIInitiatorParams) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateISCSIInitiatorParams) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateISCSIInitiatorParams) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateISCSIInitiatorParams) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


