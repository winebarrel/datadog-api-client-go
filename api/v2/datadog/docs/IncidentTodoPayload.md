# IncidentTodoPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IncidentTodo**](IncidentTodo.md) |  | 
**Included** | Pointer to [**[]UserJSONAPIResponse**](UserJSONAPIResponse.md) | JSON API Schema User relationships included. | [optional] [readonly] 

## Methods

### NewIncidentTodoPayload

`func NewIncidentTodoPayload(data IncidentTodo, ) *IncidentTodoPayload`

NewIncidentTodoPayload instantiates a new IncidentTodoPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTodoPayloadWithDefaults

`func NewIncidentTodoPayloadWithDefaults() *IncidentTodoPayload`

NewIncidentTodoPayloadWithDefaults instantiates a new IncidentTodoPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentTodoPayload) GetData() IncidentTodo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentTodoPayload) GetDataOk() (*IncidentTodo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentTodoPayload) SetData(v IncidentTodo)`

SetData sets Data field to given value.


### GetIncluded

`func (o *IncidentTodoPayload) GetIncluded() []UserJSONAPIResponse`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *IncidentTodoPayload) GetIncludedOk() (*[]UserJSONAPIResponse, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *IncidentTodoPayload) SetIncluded(v []UserJSONAPIResponse)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *IncidentTodoPayload) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


