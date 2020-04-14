# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**ServiceAttributes**](Service_attributes.md) |  | [optional] 
**Id** | Pointer to **string** | Unique identifier that represents the Incident Config Field Choice. | [optional] 
**Relationships** | Pointer to [**IncidentConfigFieldRelationships**](IncidentConfigField_relationships.md) |  | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type | [optional] [default to "service"]

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Service) GetAttributes() ServiceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Service) GetAttributesOk() (*ServiceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Service) SetAttributes(v ServiceAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Service) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Service) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *Service) GetRelationships() IncidentConfigFieldRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Service) GetRelationshipsOk() (*IncidentConfigFieldRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Service) SetRelationships(v IncidentConfigFieldRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Service) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *Service) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Service) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Service) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Service) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


