# IncidentFacetTermsAggregationDataSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to [**NullableIncidentFacetTermsAggregationDataSchemaCountOneOf**](IncidentFacetTermsAggregationDataSchemaCountOneOf.md) | A count of the terms. | [optional] [readonly] 
**Name** | Pointer to [**IncidentFacetTermsAggregationDataSchemaNameOneOf**](IncidentFacetTermsAggregationDataSchemaNameOneOf.md) | The key value for the term. | [optional] [readonly] 

## Methods

### NewIncidentFacetTermsAggregationDataSchema

`func NewIncidentFacetTermsAggregationDataSchema() *IncidentFacetTermsAggregationDataSchema`

NewIncidentFacetTermsAggregationDataSchema instantiates a new IncidentFacetTermsAggregationDataSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentFacetTermsAggregationDataSchemaWithDefaults

`func NewIncidentFacetTermsAggregationDataSchemaWithDefaults() *IncidentFacetTermsAggregationDataSchema`

NewIncidentFacetTermsAggregationDataSchemaWithDefaults instantiates a new IncidentFacetTermsAggregationDataSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IncidentFacetTermsAggregationDataSchema) GetCount() IncidentFacetTermsAggregationDataSchemaCountOneOf`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IncidentFacetTermsAggregationDataSchema) GetCountOk() (*IncidentFacetTermsAggregationDataSchemaCountOneOf, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IncidentFacetTermsAggregationDataSchema) SetCount(v IncidentFacetTermsAggregationDataSchemaCountOneOf)`

SetCount sets Count field to given value.

### HasCount

`func (o *IncidentFacetTermsAggregationDataSchema) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *IncidentFacetTermsAggregationDataSchema) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *IncidentFacetTermsAggregationDataSchema) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetName

`func (o *IncidentFacetTermsAggregationDataSchema) GetName() IncidentFacetTermsAggregationDataSchemaNameOneOf`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentFacetTermsAggregationDataSchema) GetNameOk() (*IncidentFacetTermsAggregationDataSchemaNameOneOf, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentFacetTermsAggregationDataSchema) SetName(v IncidentFacetTermsAggregationDataSchemaNameOneOf)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentFacetTermsAggregationDataSchema) HasName() bool`

HasName returns a boolean if a field has been set.


### AsIncidentFacetSchemaDataOneOf

`func (s *IncidentFacetTermsAggregationDataSchema) AsIncidentFacetSchemaDataOneOf() IncidentFacetSchemaDataOneOf`

Convenience method to wrap this instance of IncidentFacetTermsAggregationDataSchema in IncidentFacetSchemaDataOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


