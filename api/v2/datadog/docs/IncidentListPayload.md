# IncidentListPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Incident**](Incident.md) | The Incident payloads. | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The User relationships. | [optional] [readonly] 
**Meta** | Pointer to [**QueryMetadata**](QueryMetadata.md) |  | [optional] 

## Methods

### NewIncidentListPayload

`func NewIncidentListPayload(data []Incident, ) *IncidentListPayload`

NewIncidentListPayload instantiates a new IncidentListPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentListPayloadWithDefaults

`func NewIncidentListPayloadWithDefaults() *IncidentListPayload`

NewIncidentListPayloadWithDefaults instantiates a new IncidentListPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentListPayload) GetData() []Incident`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentListPayload) GetDataOk() (*[]Incident, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentListPayload) SetData(v []Incident)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentListPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentListPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentListPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentListPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *IncidentListPayload) GetMeta() QueryMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IncidentListPayload) GetMetaOk() (*QueryMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IncidentListPayload) SetMeta(v QueryMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IncidentListPayload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


