# IncidentPostmortemAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when an postmortem was created. | [optional] [readonly] 
**IncidentId** | Pointer to **string** | ID of the incident for this postmortem. | [optional] [readonly] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when an postmortem was modified. | [optional] [readonly] 
**State** | Pointer to **string** | Current state of the postmortem (i.e. not_started. draft). | [optional] [default to "not_started"]

## Methods

### NewIncidentPostmortemAttributes

`func NewIncidentPostmortemAttributes() *IncidentPostmortemAttributes`

NewIncidentPostmortemAttributes instantiates a new IncidentPostmortemAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentPostmortemAttributesWithDefaults

`func NewIncidentPostmortemAttributesWithDefaults() *IncidentPostmortemAttributes`

NewIncidentPostmortemAttributesWithDefaults instantiates a new IncidentPostmortemAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IncidentPostmortemAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentPostmortemAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentPostmortemAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentPostmortemAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetIncidentId

`func (o *IncidentPostmortemAttributes) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *IncidentPostmortemAttributes) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *IncidentPostmortemAttributes) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *IncidentPostmortemAttributes) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### GetModified

`func (o *IncidentPostmortemAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentPostmortemAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentPostmortemAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentPostmortemAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetState

`func (o *IncidentPostmortemAttributes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IncidentPostmortemAttributes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IncidentPostmortemAttributes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IncidentPostmortemAttributes) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


