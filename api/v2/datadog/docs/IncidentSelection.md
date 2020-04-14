# IncidentSelection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**IncidentSelectionAttributes**](IncidentSelection_attributes.md) |  | [optional] 
**Id** | Pointer to **string** | The Incident Selection ID. | [optional] [readonly] 
**Relationships** | Pointer to [**IncidentSelectionRelationships**](IncidentSelection_relationships.md) |  | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type | [optional] 

## Methods

### NewIncidentSelection

`func NewIncidentSelection() *IncidentSelection`

NewIncidentSelection instantiates a new IncidentSelection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSelectionWithDefaults

`func NewIncidentSelectionWithDefaults() *IncidentSelection`

NewIncidentSelectionWithDefaults instantiates a new IncidentSelection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentSelection) GetAttributes() IncidentSelectionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentSelection) GetAttributesOk() (*IncidentSelectionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentSelection) SetAttributes(v IncidentSelectionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentSelection) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *IncidentSelection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentSelection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentSelection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentSelection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *IncidentSelection) GetRelationships() IncidentSelectionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentSelection) GetRelationshipsOk() (*IncidentSelectionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentSelection) SetRelationships(v IncidentSelectionRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentSelection) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentSelection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentSelection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentSelection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentSelection) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


