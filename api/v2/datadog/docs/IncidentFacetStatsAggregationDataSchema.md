# IncidentFacetStatsAggregationDataSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avg** | Pointer to **NullableFloat64** | The average value. | [optional] [readonly] 
**Count** | Pointer to **NullableFloat64** | The count of matches. | [optional] [readonly] 
**Max** | Pointer to **NullableFloat64** | The max value. | [optional] [readonly] 
**Min** | Pointer to **NullableFloat64** | The min value. | [optional] [readonly] 
**Sum** | Pointer to **NullableFloat64** | The sum of values. | [optional] [readonly] 

## Methods

### NewIncidentFacetStatsAggregationDataSchema

`func NewIncidentFacetStatsAggregationDataSchema() *IncidentFacetStatsAggregationDataSchema`

NewIncidentFacetStatsAggregationDataSchema instantiates a new IncidentFacetStatsAggregationDataSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentFacetStatsAggregationDataSchemaWithDefaults

`func NewIncidentFacetStatsAggregationDataSchemaWithDefaults() *IncidentFacetStatsAggregationDataSchema`

NewIncidentFacetStatsAggregationDataSchemaWithDefaults instantiates a new IncidentFacetStatsAggregationDataSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvg

`func (o *IncidentFacetStatsAggregationDataSchema) GetAvg() float64`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *IncidentFacetStatsAggregationDataSchema) GetAvgOk() (*float64, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *IncidentFacetStatsAggregationDataSchema) SetAvg(v float64)`

SetAvg sets Avg field to given value.

### HasAvg

`func (o *IncidentFacetStatsAggregationDataSchema) HasAvg() bool`

HasAvg returns a boolean if a field has been set.

### SetAvgNil

`func (o *IncidentFacetStatsAggregationDataSchema) SetAvgNil(b bool)`

 SetAvgNil sets the value for Avg to be an explicit nil

### UnsetAvg
`func (o *IncidentFacetStatsAggregationDataSchema) UnsetAvg()`

UnsetAvg ensures that no value is present for Avg, not even an explicit nil
### GetCount

`func (o *IncidentFacetStatsAggregationDataSchema) GetCount() float64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IncidentFacetStatsAggregationDataSchema) GetCountOk() (*float64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IncidentFacetStatsAggregationDataSchema) SetCount(v float64)`

SetCount sets Count field to given value.

### HasCount

`func (o *IncidentFacetStatsAggregationDataSchema) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *IncidentFacetStatsAggregationDataSchema) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *IncidentFacetStatsAggregationDataSchema) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetMax

`func (o *IncidentFacetStatsAggregationDataSchema) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *IncidentFacetStatsAggregationDataSchema) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *IncidentFacetStatsAggregationDataSchema) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *IncidentFacetStatsAggregationDataSchema) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *IncidentFacetStatsAggregationDataSchema) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *IncidentFacetStatsAggregationDataSchema) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetMin

`func (o *IncidentFacetStatsAggregationDataSchema) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *IncidentFacetStatsAggregationDataSchema) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *IncidentFacetStatsAggregationDataSchema) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *IncidentFacetStatsAggregationDataSchema) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *IncidentFacetStatsAggregationDataSchema) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *IncidentFacetStatsAggregationDataSchema) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetSum

`func (o *IncidentFacetStatsAggregationDataSchema) GetSum() float64`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *IncidentFacetStatsAggregationDataSchema) GetSumOk() (*float64, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *IncidentFacetStatsAggregationDataSchema) SetSum(v float64)`

SetSum sets Sum field to given value.

### HasSum

`func (o *IncidentFacetStatsAggregationDataSchema) HasSum() bool`

HasSum returns a boolean if a field has been set.

### SetSumNil

`func (o *IncidentFacetStatsAggregationDataSchema) SetSumNil(b bool)`

 SetSumNil sets the value for Sum to be an explicit nil

### UnsetSum
`func (o *IncidentFacetStatsAggregationDataSchema) UnsetSum()`

UnsetSum ensures that no value is present for Sum, not even an explicit nil

### AsIncidentFacetSchemaDataOneOf

`func (s *IncidentFacetStatsAggregationDataSchema) AsIncidentFacetSchemaDataOneOf() IncidentFacetSchemaDataOneOf`

Convenience method to wrap this instance of IncidentFacetStatsAggregationDataSchema in IncidentFacetSchemaDataOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


