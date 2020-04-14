# IncidentSummaryPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IncidentSummary**](IncidentSummary.md) |  | 
**Meta** | Pointer to [**QueryMetadata**](QueryMetadata.md) |  | [optional] 

## Methods

### NewIncidentSummaryPayload

`func NewIncidentSummaryPayload(data IncidentSummary, ) *IncidentSummaryPayload`

NewIncidentSummaryPayload instantiates a new IncidentSummaryPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSummaryPayloadWithDefaults

`func NewIncidentSummaryPayloadWithDefaults() *IncidentSummaryPayload`

NewIncidentSummaryPayloadWithDefaults instantiates a new IncidentSummaryPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentSummaryPayload) GetData() IncidentSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentSummaryPayload) GetDataOk() (*IncidentSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentSummaryPayload) SetData(v IncidentSummary)`

SetData sets Data field to given value.


### GetMeta

`func (o *IncidentSummaryPayload) GetMeta() QueryMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IncidentSummaryPayload) GetMetaOk() (*QueryMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IncidentSummaryPayload) SetMeta(v QueryMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IncidentSummaryPayload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


