# IncidentSelectionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IncidentSelection**](IncidentSelection.md) |  | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The User relationships. | [optional] [readonly] 

## Methods

### NewIncidentSelectionPayload

`func NewIncidentSelectionPayload(data IncidentSelection, ) *IncidentSelectionPayload`

NewIncidentSelectionPayload instantiates a new IncidentSelectionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSelectionPayloadWithDefaults

`func NewIncidentSelectionPayloadWithDefaults() *IncidentSelectionPayload`

NewIncidentSelectionPayloadWithDefaults instantiates a new IncidentSelectionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentSelectionPayload) GetData() IncidentSelection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentSelectionPayload) GetDataOk() (*IncidentSelection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentSelectionPayload) SetData(v IncidentSelection)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentSelectionPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentSelectionPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentSelectionPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentSelectionPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


