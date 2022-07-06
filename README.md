# Go API client for truenas

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v2.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import truenas "github.com/dariusbakunas/truenas-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), truenas.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), truenas.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), truenas.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), truenas.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CronjobApi* | [**CreateCronJob**](docs/CronjobApi.md#createcronjob) | **Post** /cronjob | 
*CronjobApi* | [**DeleteCronJob**](docs/CronjobApi.md#deletecronjob) | **Delete** /cronjob/id/{id} | 
*CronjobApi* | [**GetCronJob**](docs/CronjobApi.md#getcronjob) | **Get** /cronjob/id/{id} | 
*CronjobApi* | [**UpdateCronJob**](docs/CronjobApi.md#updatecronjob) | **Put** /cronjob/id/{id} | 
*DatasetApi* | [**CreateDataset**](docs/DatasetApi.md#createdataset) | **Post** /pool/dataset | 
*DatasetApi* | [**DeleteDataset**](docs/DatasetApi.md#deletedataset) | **Delete** /pool/dataset/id/{id} | 
*DatasetApi* | [**GetDataset**](docs/DatasetApi.md#getdataset) | **Get** /pool/dataset/id/{id} | 
*DatasetApi* | [**ListDatasets**](docs/DatasetApi.md#listdatasets) | **Get** /pool/dataset | 
*DatasetApi* | [**UpdateDataset**](docs/DatasetApi.md#updatedataset) | **Put** /pool/dataset/id/{id} | 
*GroupApi* | [**CreateGroup**](docs/GroupApi.md#creategroup) | **Post** /group | 
*GroupApi* | [**DeleteGroup**](docs/GroupApi.md#deletegroup) | **Delete** /group/id/{id} | 
*GroupApi* | [**GetGroup**](docs/GroupApi.md#getgroup) | **Get** /group/id/{id} | 
*GroupApi* | [**ListGroups**](docs/GroupApi.md#listgroups) | **Get** /group | 
*GroupApi* | [**UpdateGroup**](docs/GroupApi.md#updategroup) | **Put** /group/id/{id} | 
*IscsiExtentApi* | [**CreateISCSIExtent**](docs/IscsiExtentApi.md#createiscsiextent) | **Post** /iscsi/extent | 
*IscsiExtentApi* | [**DeleteISCSIExtent**](docs/IscsiExtentApi.md#deleteiscsiextent) | **Delete** /iscsi/extent/id/{id} | 
*IscsiExtentApi* | [**GetISCSIExtent**](docs/IscsiExtentApi.md#getiscsiextent) | **Get** /iscsi/extent/id/{id} | 
*IscsiExtentApi* | [**GetISCSIExtents**](docs/IscsiExtentApi.md#getiscsiextents) | **Get** /iscsi/extent | 
*IscsiInitiatorApi* | [**CreateISCSIInitiator**](docs/IscsiInitiatorApi.md#createiscsiinitiator) | **Post** /iscsi/initiator | 
*IscsiInitiatorApi* | [**DeleteISCSIInitiator**](docs/IscsiInitiatorApi.md#deleteiscsiinitiator) | **Delete** /iscsi/initiator/id/{id} | 
*IscsiInitiatorApi* | [**GetISCSIInitiator**](docs/IscsiInitiatorApi.md#getiscsiinitiator) | **Get** /iscsi/initiator/id/{id} | 
*IscsiInitiatorApi* | [**GetISCSIInitiators**](docs/IscsiInitiatorApi.md#getiscsiinitiators) | **Get** /iscsi/initiator | 
*IscsiTargetApi* | [**CreateISCSITarget**](docs/IscsiTargetApi.md#createiscsitarget) | **Post** /iscsi/target | 
*IscsiTargetApi* | [**DeleteISCSITarget**](docs/IscsiTargetApi.md#deleteiscsitarget) | **Delete** /iscsi/target/id/{id} | 
*IscsiTargetApi* | [**GetISCSITarget**](docs/IscsiTargetApi.md#getiscsitarget) | **Get** /iscsi/target/id/{id} | 
*IscsiTargetApi* | [**GetISCSITargets**](docs/IscsiTargetApi.md#getiscsitargets) | **Get** /iscsi/target | 
*IscsiTargetextentApi* | [**CreateISCSITargetExtent**](docs/IscsiTargetextentApi.md#createiscsitargetextent) | **Post** /iscsi/targetextent | 
*IscsiTargetextentApi* | [**DeleteISCSITargetExtent**](docs/IscsiTargetextentApi.md#deleteiscsitargetextent) | **Delete** /iscsi/targetextent/id/{id} | 
*IscsiTargetextentApi* | [**GetISCSITargetExtent**](docs/IscsiTargetextentApi.md#getiscsitargetextent) | **Get** /iscsi/targetextent/id/{id} | 
*IscsiTargetextentApi* | [**GetISCSITargetExtents**](docs/IscsiTargetextentApi.md#getiscsitargetextents) | **Get** /iscsi/targetextent | 
*NetworkApi* | [**GetNetworkConfiguration**](docs/NetworkApi.md#getnetworkconfiguration) | **Get** /network/configuration | 
*NetworkApi* | [**GetNetworkSummary**](docs/NetworkApi.md#getnetworksummary) | **Get** /network/general/summary | 
*PoolApi* | [**ListPools**](docs/PoolApi.md#listpools) | **Get** /pool | 
*ServiceApi* | [**GetService**](docs/ServiceApi.md#getservice) | **Get** /service/id/{id} | 
*SharingApi* | [**CreateShareNFS**](docs/SharingApi.md#createsharenfs) | **Post** /sharing/nfs | 
*SharingApi* | [**CreateShareSMB**](docs/SharingApi.md#createsharesmb) | **Post** /sharing/smb | 
*SharingApi* | [**GetShareNFS**](docs/SharingApi.md#getsharenfs) | **Get** /sharing/nfs/id/{id} | 
*SharingApi* | [**GetShareSMB**](docs/SharingApi.md#getsharesmb) | **Get** /sharing/smb/id/{id} | 
*SharingApi* | [**ListSharesNFS**](docs/SharingApi.md#listsharesnfs) | **Get** /sharing/nfs | 
*SharingApi* | [**ListSharesSMB**](docs/SharingApi.md#listsharessmb) | **Get** /sharing/smb | 
*SharingApi* | [**RemoveShareNFS**](docs/SharingApi.md#removesharenfs) | **Delete** /sharing/nfs/id/{id} | 
*SharingApi* | [**RemoveShareSMB**](docs/SharingApi.md#removesharesmb) | **Delete** /sharing/smb/id/{id} | 
*SharingApi* | [**UpdateShareNFS**](docs/SharingApi.md#updatesharenfs) | **Put** /sharing/nfs/id/{id} | 
*SharingApi* | [**UpdateShareSMB**](docs/SharingApi.md#updatesharesmb) | **Put** /sharing/smb/id/{id} | 
*UserApi* | [**CreateUser**](docs/UserApi.md#createuser) | **Post** /user | 
*UserApi* | [**DeleteUser**](docs/UserApi.md#deleteuser) | **Delete** /user/id/{id} | 
*UserApi* | [**GetUser**](docs/UserApi.md#getuser) | **Get** /user/id/{id} | 
*UserApi* | [**GetUserShellChoices**](docs/UserApi.md#getusershellchoices) | **Post** /user/shell_choices | 
*UserApi* | [**ListUsers**](docs/UserApi.md#listusers) | **Get** /user | 
*UserApi* | [**UpdateUser**](docs/UserApi.md#updateuser) | **Put** /user/id/{id} | 
*VmApi* | [**GetVM**](docs/VmApi.md#getvm) | **Get** /vm/id/{id} | 
*VmApi* | [**ListVMS**](docs/VmApi.md#listvms) | **Get** /vm | 


## Documentation For Models

 - [CompositeValue](docs/CompositeValue.md)
 - [CreateCronjobParams](docs/CreateCronjobParams.md)
 - [CreateDatasetParams](docs/CreateDatasetParams.md)
 - [CreateDatasetParamsEncryptionOptions](docs/CreateDatasetParamsEncryptionOptions.md)
 - [CreateGroupParams](docs/CreateGroupParams.md)
 - [CreateISCSIExtentParams](docs/CreateISCSIExtentParams.md)
 - [CreateISCSITargetExtentParams](docs/CreateISCSITargetExtentParams.md)
 - [CreateISCSITargetParams](docs/CreateISCSITargetParams.md)
 - [CreateISCSITargetParamsGroupsInner](docs/CreateISCSITargetParamsGroupsInner.md)
 - [CreateShareNFSParams](docs/CreateShareNFSParams.md)
 - [CreateShareSMBParams](docs/CreateShareSMBParams.md)
 - [CreateUserParams](docs/CreateUserParams.md)
 - [CronJob](docs/CronJob.md)
 - [CronJobSchedule](docs/CronJobSchedule.md)
 - [Dataset](docs/Dataset.md)
 - [DeleteGroupParams](docs/DeleteGroupParams.md)
 - [DeleteISCSIExtentParams](docs/DeleteISCSIExtentParams.md)
 - [DeleteUserParams](docs/DeleteUserParams.md)
 - [Group](docs/Group.md)
 - [ISCSIExtent](docs/ISCSIExtent.md)
 - [ISCSIInitiator](docs/ISCSIInitiator.md)
 - [ISCSITarget](docs/ISCSITarget.md)
 - [ISCSITargetGroupsInner](docs/ISCSITargetGroupsInner.md)
 - [NetworkConfig](docs/NetworkConfig.md)
 - [NetworkConfigServiceAnnouncement](docs/NetworkConfigServiceAnnouncement.md)
 - [NetworkSummary](docs/NetworkSummary.md)
 - [NetworkSummaryIpsValue](docs/NetworkSummaryIpsValue.md)
 - [Pool](docs/Pool.md)
 - [Service](docs/Service.md)
 - [ShareNFS](docs/ShareNFS.md)
 - [ShareSMB](docs/ShareSMB.md)
 - [UpdateDatasetParams](docs/UpdateDatasetParams.md)
 - [UpdateUserParams](docs/UpdateUserParams.md)
 - [User](docs/User.md)
 - [UserGroup](docs/UserGroup.md)
 - [VM](docs/VM.md)
 - [VMDevicesInner](docs/VMDevicesInner.md)
 - [VMStatus](docs/VMStatus.md)


## Documentation For Authorization



### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



