# IncidentConfigFieldChoicePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IncidentConfigFieldChoice**](IncidentConfigFieldChoice.md) |  | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The User relationships. | [optional] [readonly] 

## Methods

### NewIncidentConfigFieldChoicePayload

`func NewIncidentConfigFieldChoicePayload(data IncidentConfigFieldChoice, ) *IncidentConfigFieldChoicePayload`

NewIncidentConfigFieldChoicePayload instantiates a new IncidentConfigFieldChoicePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentConfigFieldChoicePayloadWithDefaults

`func NewIncidentConfigFieldChoicePayloadWithDefaults() *IncidentConfigFieldChoicePayload`

NewIncidentConfigFieldChoicePayloadWithDefaults instantiates a new IncidentConfigFieldChoicePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentConfigFieldChoicePayload) GetData() IncidentConfigFieldChoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentConfigFieldChoicePayload) GetDataOk() (*IncidentConfigFieldChoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentConfigFieldChoicePayload) SetData(v IncidentConfigFieldChoice)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentConfigFieldChoicePayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentConfigFieldChoicePayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentConfigFieldChoicePayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentConfigFieldChoicePayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


