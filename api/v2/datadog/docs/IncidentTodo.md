# IncidentTodo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**IncidentTodoAttributes**](IncidentTodo_attributes.md) |  | [optional] 
**Id** | Pointer to **string** | Incident Todo ID. | [optional] [readonly] 
**Relationships** | Pointer to [**IncidentTodoRelationships**](IncidentTodo_relationships.md) |  | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type. | [optional] 

## Methods

### NewIncidentTodo

`func NewIncidentTodo() *IncidentTodo`

NewIncidentTodo instantiates a new IncidentTodo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTodoWithDefaults

`func NewIncidentTodoWithDefaults() *IncidentTodo`

NewIncidentTodoWithDefaults instantiates a new IncidentTodo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentTodo) GetAttributes() IncidentTodoAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentTodo) GetAttributesOk() (*IncidentTodoAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentTodo) SetAttributes(v IncidentTodoAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentTodo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *IncidentTodo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentTodo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentTodo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentTodo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *IncidentTodo) GetRelationships() IncidentTodoRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentTodo) GetRelationshipsOk() (*IncidentTodoRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentTodo) SetRelationships(v IncidentTodoRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentTodo) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentTodo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentTodo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentTodo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentTodo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


