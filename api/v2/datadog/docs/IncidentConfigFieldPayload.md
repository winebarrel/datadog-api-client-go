# IncidentConfigFieldPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IncidentConfigField**](IncidentConfigField.md) |  | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The Incident config fields. | [optional] [readonly] 

## Methods

### NewIncidentConfigFieldPayload

`func NewIncidentConfigFieldPayload(data IncidentConfigField, ) *IncidentConfigFieldPayload`

NewIncidentConfigFieldPayload instantiates a new IncidentConfigFieldPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentConfigFieldPayloadWithDefaults

`func NewIncidentConfigFieldPayloadWithDefaults() *IncidentConfigFieldPayload`

NewIncidentConfigFieldPayloadWithDefaults instantiates a new IncidentConfigFieldPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentConfigFieldPayload) GetData() IncidentConfigField`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentConfigFieldPayload) GetDataOk() (*IncidentConfigField, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentConfigFieldPayload) SetData(v IncidentConfigField)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentConfigFieldPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentConfigFieldPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentConfigFieldPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentConfigFieldPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


