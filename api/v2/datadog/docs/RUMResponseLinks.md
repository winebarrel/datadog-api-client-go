# RUMResponseLinks

## Properties

| Name     | Type                  | Description                                                                                       | Notes      |
| -------- | --------------------- | ------------------------------------------------------------------------------------------------- | ---------- |
| **Next** | Pointer to **string** | Link for the next set of results. Note that the request can also be made using the POST endpoint. | [optional] |

## Methods

### NewRUMResponseLinks

`func NewRUMResponseLinks() *RUMResponseLinks`

NewRUMResponseLinks instantiates a new RUMResponseLinks object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRUMResponseLinksWithDefaults

`func NewRUMResponseLinksWithDefaults() *RUMResponseLinks`

NewRUMResponseLinksWithDefaults instantiates a new RUMResponseLinks object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetNext

`func (o *RUMResponseLinks) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *RUMResponseLinks) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *RUMResponseLinks) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *RUMResponseLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
