# IncidentPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Incident**](Incident.md) |  | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The User relationships. | [optional] [readonly] 

## Methods

### NewIncidentPayload

`func NewIncidentPayload(data Incident, ) *IncidentPayload`

NewIncidentPayload instantiates a new IncidentPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentPayloadWithDefaults

`func NewIncidentPayloadWithDefaults() *IncidentPayload`

NewIncidentPayloadWithDefaults instantiates a new IncidentPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentPayload) GetData() Incident`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentPayload) GetDataOk() (*Incident, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentPayload) SetData(v Incident)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


