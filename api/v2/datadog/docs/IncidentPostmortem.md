# IncidentPostmortem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**IncidentPostmortemAttributes**](IncidentPostmortem_attributes.md) |  | [optional] 
**Id** | Pointer to **string** | The Postmortem ID. | [optional] [readonly] 
**Relationships** | Pointer to [**IncidentConfigFieldRelationships**](IncidentConfigField_relationships.md) |  | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type | [optional] 

## Methods

### NewIncidentPostmortem

`func NewIncidentPostmortem() *IncidentPostmortem`

NewIncidentPostmortem instantiates a new IncidentPostmortem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentPostmortemWithDefaults

`func NewIncidentPostmortemWithDefaults() *IncidentPostmortem`

NewIncidentPostmortemWithDefaults instantiates a new IncidentPostmortem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentPostmortem) GetAttributes() IncidentPostmortemAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentPostmortem) GetAttributesOk() (*IncidentPostmortemAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentPostmortem) SetAttributes(v IncidentPostmortemAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentPostmortem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *IncidentPostmortem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentPostmortem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentPostmortem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentPostmortem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *IncidentPostmortem) GetRelationships() IncidentConfigFieldRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentPostmortem) GetRelationshipsOk() (*IncidentConfigFieldRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentPostmortem) SetRelationships(v IncidentConfigFieldRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentPostmortem) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentPostmortem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentPostmortem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentPostmortem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentPostmortem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


