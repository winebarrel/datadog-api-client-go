# Incident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**IncidentAttributes**](Incident_attributes.md) |  | [optional] 
**Id** | Pointer to **string** | The Incident ID. | [optional] [readonly] 
**Relationships** | Pointer to [**IncidentRelationships**](Incident_relationships.md) |  | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type | [optional] [default to "incidents"]

## Methods

### NewIncident

`func NewIncident() *Incident`

NewIncident instantiates a new Incident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentWithDefaults

`func NewIncidentWithDefaults() *Incident`

NewIncidentWithDefaults instantiates a new Incident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Incident) GetAttributes() IncidentAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Incident) GetAttributesOk() (*IncidentAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Incident) SetAttributes(v IncidentAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Incident) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *Incident) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Incident) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Incident) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Incident) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *Incident) GetRelationships() IncidentRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Incident) GetRelationshipsOk() (*IncidentRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Incident) SetRelationships(v IncidentRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Incident) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *Incident) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Incident) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Incident) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Incident) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


