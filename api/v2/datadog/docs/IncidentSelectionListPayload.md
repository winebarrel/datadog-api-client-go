# IncidentSelectionListPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]IncidentSelection**](IncidentSelection.md) | The Incident Selection payloads. | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | The User relationships. | [optional] [readonly] 
**Meta** | Pointer to [**QueryMetadata**](QueryMetadata.md) |  | [optional] 

## Methods

### NewIncidentSelectionListPayload

`func NewIncidentSelectionListPayload(data []IncidentSelection, ) *IncidentSelectionListPayload`

NewIncidentSelectionListPayload instantiates a new IncidentSelectionListPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSelectionListPayloadWithDefaults

`func NewIncidentSelectionListPayloadWithDefaults() *IncidentSelectionListPayload`

NewIncidentSelectionListPayloadWithDefaults instantiates a new IncidentSelectionListPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentSelectionListPayload) GetData() []IncidentSelection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentSelectionListPayload) GetDataOk() (*[]IncidentSelection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentSelectionListPayload) SetData(v []IncidentSelection)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentSelectionListPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentSelectionListPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentSelectionListPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentSelectionListPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetMeta

`func (o *IncidentSelectionListPayload) GetMeta() QueryMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IncidentSelectionListPayload) GetMetaOk() (*QueryMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IncidentSelectionListPayload) SetMeta(v QueryMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IncidentSelectionListPayload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


