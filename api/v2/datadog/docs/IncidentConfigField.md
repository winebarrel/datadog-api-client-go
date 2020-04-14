# IncidentConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**IncidentConfigFieldAttributes**](IncidentConfigField_attributes.md) |  | [optional] 
**Id** | Pointer to **string** | Unique identifier that represents the Incident Config Field. | [optional] 
**Relationships** | Pointer to [**IncidentConfigFieldRelationships**](IncidentConfigField_relationships.md) |  | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type | [optional] 

## Methods

### NewIncidentConfigField

`func NewIncidentConfigField() *IncidentConfigField`

NewIncidentConfigField instantiates a new IncidentConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentConfigFieldWithDefaults

`func NewIncidentConfigFieldWithDefaults() *IncidentConfigField`

NewIncidentConfigFieldWithDefaults instantiates a new IncidentConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentConfigField) GetAttributes() IncidentConfigFieldAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentConfigField) GetAttributesOk() (*IncidentConfigFieldAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentConfigField) SetAttributes(v IncidentConfigFieldAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentConfigField) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *IncidentConfigField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentConfigField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentConfigField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentConfigField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *IncidentConfigField) GetRelationships() IncidentConfigFieldRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentConfigField) GetRelationshipsOk() (*IncidentConfigFieldRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentConfigField) SetRelationships(v IncidentConfigFieldRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentConfigField) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentConfigField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentConfigField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentConfigField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentConfigField) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


