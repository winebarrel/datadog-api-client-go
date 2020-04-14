# IncidentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**IncidentSummaryAttributes**](IncidentSummary_attributes.md) |  | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type | [optional] [default to "incidents_search_results"]

## Methods

### NewIncidentSummary

`func NewIncidentSummary() *IncidentSummary`

NewIncidentSummary instantiates a new IncidentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSummaryWithDefaults

`func NewIncidentSummaryWithDefaults() *IncidentSummary`

NewIncidentSummaryWithDefaults instantiates a new IncidentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *IncidentSummary) GetAttributes() IncidentSummaryAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentSummary) GetAttributesOk() (*IncidentSummaryAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentSummary) SetAttributes(v IncidentSummaryAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentSummary) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *IncidentSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentSummary) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


