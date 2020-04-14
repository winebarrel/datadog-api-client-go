# IncidentConfigFieldChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**IncidentConfigFieldChoiceAttributes**](IncidentConfigFieldChoice_attributes.md) |  | [optional] 
**Id** | Pointer to **string** | Unique identifier that represents the Incident Config Field Choice | [optional] 
**Relationships** | Pointer to [**IncidentConfigFieldRelationships**](IncidentConfigField_relationships.md) |  | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type | [optional] 

## Methods

### NewIncidentConfigFieldChoice

`func NewIncidentConfigFieldChoice() *IncidentConfigFieldChoice`

NewIncidentConfigFieldChoice instantiates a new IncidentConfigFieldChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentConfigFieldChoiceWithDefaults

`func NewIncidentConfigFieldChoiceWithDefaults() *IncidentConfigFieldChoice`

NewIncidentConfigFieldChoiceWithDefaults instantiates a new IncidentConfigFieldChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentConfigFieldChoice) GetAttributes() IncidentConfigFieldChoiceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentConfigFieldChoice) GetAttributesOk() (*IncidentConfigFieldChoiceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentConfigFieldChoice) SetAttributes(v IncidentConfigFieldChoiceAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentConfigFieldChoice) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *IncidentConfigFieldChoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentConfigFieldChoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentConfigFieldChoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentConfigFieldChoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *IncidentConfigFieldChoice) GetRelationships() IncidentConfigFieldRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentConfigFieldChoice) GetRelationshipsOk() (*IncidentConfigFieldRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentConfigFieldChoice) SetRelationships(v IncidentConfigFieldRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentConfigFieldChoice) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentConfigFieldChoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentConfigFieldChoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentConfigFieldChoice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentConfigFieldChoice) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


