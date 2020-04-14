# IncidentSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**IncidentSearchAttributes**](IncidentSearch_attributes.md) |  | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type | [optional] [default to "incidents_search_results"]

## Methods

### NewIncidentSearch

`func NewIncidentSearch() *IncidentSearch`

NewIncidentSearch instantiates a new IncidentSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSearchWithDefaults

`func NewIncidentSearchWithDefaults() *IncidentSearch`

NewIncidentSearchWithDefaults instantiates a new IncidentSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentSearch) GetAttributes() IncidentSearchAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentSearch) GetAttributesOk() (*IncidentSearchAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentSearch) SetAttributes(v IncidentSearchAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentSearch) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *IncidentSearch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentSearch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentSearch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentSearch) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


