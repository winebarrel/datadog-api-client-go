# IncidentTodoAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to [**NullableTime**](time.Time.md) | Timestamp of when the todo was marked completed. | [optional] 
**Content** | Pointer to **string** | A string representing the content of the todo. | [optional] 
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when an todo was created. | [optional] [readonly] 
**IncidentId** | Pointer to **string** | ID of the incident for this todo. | [optional] [readonly] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when an todo was modified. | [optional] [readonly] 

## Methods

### NewIncidentTodoAttributes

`func NewIncidentTodoAttributes() *IncidentTodoAttributes`

NewIncidentTodoAttributes instantiates a new IncidentTodoAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTodoAttributesWithDefaults

`func NewIncidentTodoAttributesWithDefaults() *IncidentTodoAttributes`

NewIncidentTodoAttributesWithDefaults instantiates a new IncidentTodoAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *IncidentTodoAttributes) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *IncidentTodoAttributes) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *IncidentTodoAttributes) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *IncidentTodoAttributes) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### SetCompletedNil

`func (o *IncidentTodoAttributes) SetCompletedNil(b bool)`

 SetCompletedNil sets the value for Completed to be an explicit nil

### UnsetCompleted
`func (o *IncidentTodoAttributes) UnsetCompleted()`

UnsetCompleted ensures that no value is present for Completed, not even an explicit nil
### GetContent

`func (o *IncidentTodoAttributes) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IncidentTodoAttributes) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IncidentTodoAttributes) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *IncidentTodoAttributes) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreated

`func (o *IncidentTodoAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentTodoAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentTodoAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentTodoAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetIncidentId

`func (o *IncidentTodoAttributes) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *IncidentTodoAttributes) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *IncidentTodoAttributes) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *IncidentTodoAttributes) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### GetModified

`func (o *IncidentTodoAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentTodoAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentTodoAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentTodoAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


