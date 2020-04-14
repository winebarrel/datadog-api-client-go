# IncidentPostmortemPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IncidentPostmortem**](IncidentPostmortem.md) |  | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The User relationships. | [optional] [readonly] 

## Methods

### NewIncidentPostmortemPayload

`func NewIncidentPostmortemPayload(data IncidentPostmortem, ) *IncidentPostmortemPayload`

NewIncidentPostmortemPayload instantiates a new IncidentPostmortemPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentPostmortemPayloadWithDefaults

`func NewIncidentPostmortemPayloadWithDefaults() *IncidentPostmortemPayload`

NewIncidentPostmortemPayloadWithDefaults instantiates a new IncidentPostmortemPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentPostmortemPayload) GetData() IncidentPostmortem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentPostmortemPayload) GetDataOk() (*IncidentPostmortem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentPostmortemPayload) SetData(v IncidentPostmortem)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentPostmortemPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentPostmortemPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentPostmortemPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentPostmortemPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


