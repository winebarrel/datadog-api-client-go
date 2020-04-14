# UserJSONAPIIDData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier that represents the user. | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type | [optional] [default to "users"]

## Methods

### NewUserJSONAPIIDData

`func NewUserJSONAPIIDData() *UserJSONAPIIDData`

NewUserJSONAPIIDData instantiates a new UserJSONAPIIDData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserJSONAPIIDDataWithDefaults

`func NewUserJSONAPIIDDataWithDefaults() *UserJSONAPIIDData`

NewUserJSONAPIIDDataWithDefaults instantiates a new UserJSONAPIIDData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserJSONAPIIDData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserJSONAPIIDData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserJSONAPIIDData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserJSONAPIIDData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserJSONAPIIDData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserJSONAPIIDData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserJSONAPIIDData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserJSONAPIIDData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


