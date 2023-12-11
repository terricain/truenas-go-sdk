# ShareNFS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Comment** | Pointer to **string** |  | [optional] 
**Hosts** | Pointer to **[]string** |  | [optional] 
**Alldirs** | Pointer to **bool** |  | [optional] 
**Ro** | Pointer to **bool** |  | [optional] 
**Quiet** | Pointer to **bool** |  | [optional] 
**MaprootUser** | Pointer to **string** |  | [optional] 
**MaprootGroup** | Pointer to **string** |  | [optional] 
**MapallUser** | Pointer to **string** |  | [optional] 
**MapallGroup** | Pointer to **string** |  | [optional] 
**Security** | Pointer to **[]string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Paths** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Networks** | Pointer to **[]string** |  | [optional] 

## Methods

### NewShareNFS

`func NewShareNFS(id int32, ) *ShareNFS`

NewShareNFS instantiates a new ShareNFS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareNFSWithDefaults

`func NewShareNFSWithDefaults() *ShareNFS`

NewShareNFSWithDefaults instantiates a new ShareNFS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShareNFS) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareNFS) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareNFS) SetId(v int32)`

SetId sets Id field to given value.


### GetComment

`func (o *ShareNFS) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ShareNFS) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ShareNFS) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ShareNFS) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetHosts

`func (o *ShareNFS) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ShareNFS) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ShareNFS) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ShareNFS) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetAlldirs

`func (o *ShareNFS) GetAlldirs() bool`

GetAlldirs returns the Alldirs field if non-nil, zero value otherwise.

### GetAlldirsOk

`func (o *ShareNFS) GetAlldirsOk() (*bool, bool)`

GetAlldirsOk returns a tuple with the Alldirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlldirs

`func (o *ShareNFS) SetAlldirs(v bool)`

SetAlldirs sets Alldirs field to given value.

### HasAlldirs

`func (o *ShareNFS) HasAlldirs() bool`

HasAlldirs returns a boolean if a field has been set.

### GetRo

`func (o *ShareNFS) GetRo() bool`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *ShareNFS) GetRoOk() (*bool, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *ShareNFS) SetRo(v bool)`

SetRo sets Ro field to given value.

### HasRo

`func (o *ShareNFS) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetQuiet

`func (o *ShareNFS) GetQuiet() bool`

GetQuiet returns the Quiet field if non-nil, zero value otherwise.

### GetQuietOk

`func (o *ShareNFS) GetQuietOk() (*bool, bool)`

GetQuietOk returns a tuple with the Quiet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiet

`func (o *ShareNFS) SetQuiet(v bool)`

SetQuiet sets Quiet field to given value.

### HasQuiet

`func (o *ShareNFS) HasQuiet() bool`

HasQuiet returns a boolean if a field has been set.

### GetMaprootUser

`func (o *ShareNFS) GetMaprootUser() string`

GetMaprootUser returns the MaprootUser field if non-nil, zero value otherwise.

### GetMaprootUserOk

`func (o *ShareNFS) GetMaprootUserOk() (*string, bool)`

GetMaprootUserOk returns a tuple with the MaprootUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaprootUser

`func (o *ShareNFS) SetMaprootUser(v string)`

SetMaprootUser sets MaprootUser field to given value.

### HasMaprootUser

`func (o *ShareNFS) HasMaprootUser() bool`

HasMaprootUser returns a boolean if a field has been set.

### GetMaprootGroup

`func (o *ShareNFS) GetMaprootGroup() string`

GetMaprootGroup returns the MaprootGroup field if non-nil, zero value otherwise.

### GetMaprootGroupOk

`func (o *ShareNFS) GetMaprootGroupOk() (*string, bool)`

GetMaprootGroupOk returns a tuple with the MaprootGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaprootGroup

`func (o *ShareNFS) SetMaprootGroup(v string)`

SetMaprootGroup sets MaprootGroup field to given value.

### HasMaprootGroup

`func (o *ShareNFS) HasMaprootGroup() bool`

HasMaprootGroup returns a boolean if a field has been set.

### GetMapallUser

`func (o *ShareNFS) GetMapallUser() string`

GetMapallUser returns the MapallUser field if non-nil, zero value otherwise.

### GetMapallUserOk

`func (o *ShareNFS) GetMapallUserOk() (*string, bool)`

GetMapallUserOk returns a tuple with the MapallUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapallUser

`func (o *ShareNFS) SetMapallUser(v string)`

SetMapallUser sets MapallUser field to given value.

### HasMapallUser

`func (o *ShareNFS) HasMapallUser() bool`

HasMapallUser returns a boolean if a field has been set.

### GetMapallGroup

`func (o *ShareNFS) GetMapallGroup() string`

GetMapallGroup returns the MapallGroup field if non-nil, zero value otherwise.

### GetMapallGroupOk

`func (o *ShareNFS) GetMapallGroupOk() (*string, bool)`

GetMapallGroupOk returns a tuple with the MapallGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapallGroup

`func (o *ShareNFS) SetMapallGroup(v string)`

SetMapallGroup sets MapallGroup field to given value.

### HasMapallGroup

`func (o *ShareNFS) HasMapallGroup() bool`

HasMapallGroup returns a boolean if a field has been set.

### GetSecurity

`func (o *ShareNFS) GetSecurity() []string`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ShareNFS) GetSecurityOk() (*[]string, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ShareNFS) SetSecurity(v []string)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *ShareNFS) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetEnabled

`func (o *ShareNFS) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ShareNFS) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ShareNFS) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ShareNFS) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLocked

`func (o *ShareNFS) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ShareNFS) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ShareNFS) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ShareNFS) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetPaths

`func (o *ShareNFS) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *ShareNFS) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *ShareNFS) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *ShareNFS) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetPath

`func (o *ShareNFS) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ShareNFS) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ShareNFS) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ShareNFS) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetNetworks

`func (o *ShareNFS) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ShareNFS) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ShareNFS) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ShareNFS) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


