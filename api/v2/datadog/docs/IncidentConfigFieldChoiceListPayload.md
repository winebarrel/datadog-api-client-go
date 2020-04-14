# IncidentConfigFieldChoiceListPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]IncidentConfigFieldChoice**](IncidentConfigFieldChoice.md) | The Incident config field choices. | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The User relationships. | [optional] [readonly] 
**Meta** | Pointer to [**QueryMetadata**](QueryMetadata.md) |  | [optional] 

## Methods

### NewIncidentConfigFieldChoiceListPayload

`func NewIncidentConfigFieldChoiceListPayload(data []IncidentConfigFieldChoice, ) *IncidentConfigFieldChoiceListPayload`

NewIncidentConfigFieldChoiceListPayload instantiates a new IncidentConfigFieldChoiceListPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentConfigFieldChoiceListPayloadWithDefaults

`func NewIncidentConfigFieldChoiceListPayloadWithDefaults() *IncidentConfigFieldChoiceListPayload`

NewIncidentConfigFieldChoiceListPayloadWithDefaults instantiates a new IncidentConfigFieldChoiceListPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentConfigFieldChoiceListPayload) GetData() []IncidentConfigFieldChoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentConfigFieldChoiceListPayload) GetDataOk() (*[]IncidentConfigFieldChoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentConfigFieldChoiceListPayload) SetData(v []IncidentConfigFieldChoice)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentConfigFieldChoiceListPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentConfigFieldChoiceListPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentConfigFieldChoiceListPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentConfigFieldChoiceListPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *IncidentConfigFieldChoiceListPayload) GetMeta() QueryMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IncidentConfigFieldChoiceListPayload) GetMetaOk() (*QueryMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IncidentConfigFieldChoiceListPayload) SetMeta(v QueryMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IncidentConfigFieldChoiceListPayload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


