# IncidentSearchPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IncidentSearch**](IncidentSearch.md) |  | 
**Meta** | Pointer to [**QueryMetadata**](QueryMetadata.md) |  | [optional] 

## Methods

### NewIncidentSearchPayload

`func NewIncidentSearchPayload(data IncidentSearch, ) *IncidentSearchPayload`

NewIncidentSearchPayload instantiates a new IncidentSearchPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSearchPayloadWithDefaults

`func NewIncidentSearchPayloadWithDefaults() *IncidentSearchPayload`

NewIncidentSearchPayloadWithDefaults instantiates a new IncidentSearchPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentSearchPayload) GetData() IncidentSearch`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentSearchPayload) GetDataOk() (*IncidentSearch, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentSearchPayload) SetData(v IncidentSearch)`

SetData sets Data field to given value.


### GetMeta

`func (o *IncidentSearchPayload) GetMeta() QueryMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IncidentSearchPayload) GetMetaOk() (*QueryMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IncidentSearchPayload) SetMeta(v QueryMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IncidentSearchPayload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


