# IncidentPostmortemListPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]IncidentPostmortem**](IncidentPostmortem.md) | The Postmortem payloads. | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The User relationships. | [optional] [readonly] 
**Meta** | Pointer to [**QueryMetadata**](QueryMetadata.md) |  | [optional] 

## Methods

### NewIncidentPostmortemListPayload

`func NewIncidentPostmortemListPayload(data []IncidentPostmortem, ) *IncidentPostmortemListPayload`

NewIncidentPostmortemListPayload instantiates a new IncidentPostmortemListPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentPostmortemListPayloadWithDefaults

`func NewIncidentPostmortemListPayloadWithDefaults() *IncidentPostmortemListPayload`

NewIncidentPostmortemListPayloadWithDefaults instantiates a new IncidentPostmortemListPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentPostmortemListPayload) GetData() []IncidentPostmortem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentPostmortemListPayload) GetDataOk() (*[]IncidentPostmortem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentPostmortemListPayload) SetData(v []IncidentPostmortem)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentPostmortemListPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentPostmortemListPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentPostmortemListPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentPostmortemListPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *IncidentPostmortemListPayload) GetMeta() QueryMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IncidentPostmortemListPayload) GetMetaOk() (*QueryMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IncidentPostmortemListPayload) SetMeta(v QueryMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IncidentPostmortemListPayload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


