# IncidentCreateWithInitialDataPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IncidentCreateWithInitialData**](IncidentCreateWithInitialData.md) |  | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The User relationships. | [optional] [readonly] 

## Methods

### NewIncidentCreateWithInitialDataPayload

`func NewIncidentCreateWithInitialDataPayload(data IncidentCreateWithInitialData, ) *IncidentCreateWithInitialDataPayload`

NewIncidentCreateWithInitialDataPayload instantiates a new IncidentCreateWithInitialDataPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentCreateWithInitialDataPayloadWithDefaults

`func NewIncidentCreateWithInitialDataPayloadWithDefaults() *IncidentCreateWithInitialDataPayload`

NewIncidentCreateWithInitialDataPayloadWithDefaults instantiates a new IncidentCreateWithInitialDataPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentCreateWithInitialDataPayload) GetData() IncidentCreateWithInitialData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentCreateWithInitialDataPayload) GetDataOk() (*IncidentCreateWithInitialData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentCreateWithInitialDataPayload) SetData(v IncidentCreateWithInitialData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentCreateWithInitialDataPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentCreateWithInitialDataPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentCreateWithInitialDataPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentCreateWithInitialDataPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


