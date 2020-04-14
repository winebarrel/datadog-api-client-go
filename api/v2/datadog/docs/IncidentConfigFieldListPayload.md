# IncidentConfigFieldListPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]IncidentConfigField**](IncidentConfigField.md) | The Incident config fields. | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The User relationships. | [optional] [readonly] 
**Meta** | Pointer to [**QueryMetadata**](QueryMetadata.md) |  | [optional] 

## Methods

### NewIncidentConfigFieldListPayload

`func NewIncidentConfigFieldListPayload(data []IncidentConfigField, ) *IncidentConfigFieldListPayload`

NewIncidentConfigFieldListPayload instantiates a new IncidentConfigFieldListPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentConfigFieldListPayloadWithDefaults

`func NewIncidentConfigFieldListPayloadWithDefaults() *IncidentConfigFieldListPayload`

NewIncidentConfigFieldListPayloadWithDefaults instantiates a new IncidentConfigFieldListPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentConfigFieldListPayload) GetData() []IncidentConfigField`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentConfigFieldListPayload) GetDataOk() (*[]IncidentConfigField, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentConfigFieldListPayload) SetData(v []IncidentConfigField)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentConfigFieldListPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentConfigFieldListPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentConfigFieldListPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentConfigFieldListPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *IncidentConfigFieldListPayload) GetMeta() QueryMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IncidentConfigFieldListPayload) GetMetaOk() (*QueryMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IncidentConfigFieldListPayload) SetMeta(v QueryMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IncidentConfigFieldListPayload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


