# IncidentSummaryAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facets** | Pointer to [**map[string]IncidentFacetSchema**](IncidentFacetSchema.md) | A mapping of search facets by facet name. | [optional] [readonly] 
**Total** | Pointer to **int64** | The total count of matched incidents. | [optional] [readonly] 

## Methods

### NewIncidentSummaryAttributes

`func NewIncidentSummaryAttributes() *IncidentSummaryAttributes`

NewIncidentSummaryAttributes instantiates a new IncidentSummaryAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSummaryAttributesWithDefaults

`func NewIncidentSummaryAttributesWithDefaults() *IncidentSummaryAttributes`

NewIncidentSummaryAttributesWithDefaults instantiates a new IncidentSummaryAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacets

`func (o *IncidentSummaryAttributes) GetFacets() map[string]IncidentFacetSchema`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *IncidentSummaryAttributes) GetFacetsOk() (*map[string]IncidentFacetSchema, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *IncidentSummaryAttributes) SetFacets(v map[string]IncidentFacetSchema)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *IncidentSummaryAttributes) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### SetFacetsNil

`func (o *IncidentSummaryAttributes) SetFacetsNil(b bool)`

 SetFacetsNil sets the value for Facets to be an explicit nil

### UnsetFacets
`func (o *IncidentSummaryAttributes) UnsetFacets()`

UnsetFacets ensures that no value is present for Facets, not even an explicit nil
### GetTotal

`func (o *IncidentSummaryAttributes) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IncidentSummaryAttributes) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IncidentSummaryAttributes) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IncidentSummaryAttributes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


