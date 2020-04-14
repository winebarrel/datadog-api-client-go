# IncidentSearchAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facets** | Pointer to [**map[string]IncidentFacetSchema**](IncidentFacetSchema.md) | A mapping of search facets by facet name. | [optional] [readonly] 
**Incidents** | Pointer to [**[]IncidentPayload**](IncidentPayload.md) | The matched incidents if any. | [optional] [readonly] 
**Total** | Pointer to **int64** | The total count of matched incidents. | [optional] [readonly] 

## Methods

### NewIncidentSearchAttributes

`func NewIncidentSearchAttributes() *IncidentSearchAttributes`

NewIncidentSearchAttributes instantiates a new IncidentSearchAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSearchAttributesWithDefaults

`func NewIncidentSearchAttributesWithDefaults() *IncidentSearchAttributes`

NewIncidentSearchAttributesWithDefaults instantiates a new IncidentSearchAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacets

`func (o *IncidentSearchAttributes) GetFacets() map[string]IncidentFacetSchema`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *IncidentSearchAttributes) GetFacetsOk() (*map[string]IncidentFacetSchema, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *IncidentSearchAttributes) SetFacets(v map[string]IncidentFacetSchema)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *IncidentSearchAttributes) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### SetFacetsNil

`func (o *IncidentSearchAttributes) SetFacetsNil(b bool)`

 SetFacetsNil sets the value for Facets to be an explicit nil

### UnsetFacets
`func (o *IncidentSearchAttributes) UnsetFacets()`

UnsetFacets ensures that no value is present for Facets, not even an explicit nil
### GetIncidents

`func (o *IncidentSearchAttributes) GetIncidents() []IncidentPayload`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *IncidentSearchAttributes) GetIncidentsOk() (*[]IncidentPayload, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *IncidentSearchAttributes) SetIncidents(v []IncidentPayload)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *IncidentSearchAttributes) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetTotal

`func (o *IncidentSearchAttributes) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IncidentSearchAttributes) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IncidentSearchAttributes) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IncidentSearchAttributes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


