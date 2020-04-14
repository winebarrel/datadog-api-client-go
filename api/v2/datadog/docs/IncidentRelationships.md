# IncidentRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commander** | Pointer to [**UserJSONAPIID**](UserJSONAPIID.md) |  | [optional] 
**CreatedBy** | Pointer to [**UserJSONAPIID**](UserJSONAPIID.md) |  | [optional] 
**LastModifiedBy** | Pointer to [**UserJSONAPIID**](UserJSONAPIID.md) |  | [optional] 

## Methods

### NewIncidentRelationships

`func NewIncidentRelationships() *IncidentRelationships`

NewIncidentRelationships instantiates a new IncidentRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentRelationshipsWithDefaults

`func NewIncidentRelationshipsWithDefaults() *IncidentRelationships`

NewIncidentRelationshipsWithDefaults instantiates a new IncidentRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommander

`func (o *IncidentRelationships) GetCommander() UserJSONAPIID`

GetCommander returns the Commander field if non-nil, zero value otherwise.

### GetCommanderOk

`func (o *IncidentRelationships) GetCommanderOk() (*UserJSONAPIID, bool)`

GetCommanderOk returns a tuple with the Commander field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommander

`func (o *IncidentRelationships) SetCommander(v UserJSONAPIID)`

SetCommander sets Commander field to given value.

### HasCommander

`func (o *IncidentRelationships) HasCommander() bool`

HasCommander returns a boolean if a field has been set.

### GetCreatedBy

`func (o *IncidentRelationships) GetCreatedBy() UserJSONAPIID`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IncidentRelationships) GetCreatedByOk() (*UserJSONAPIID, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IncidentRelationships) SetCreatedBy(v UserJSONAPIID)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *IncidentRelationships) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *IncidentRelationships) GetLastModifiedBy() UserJSONAPIID`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *IncidentRelationships) GetLastModifiedByOk() (*UserJSONAPIID, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *IncidentRelationships) SetLastModifiedBy(v UserJSONAPIID)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *IncidentRelationships) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


