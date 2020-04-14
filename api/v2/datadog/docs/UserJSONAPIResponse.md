# UserJSONAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**User**](User.md) |  | [optional] 
**Id** | Pointer to **string** | A unique identifier that represents the user. | [optional] 
**Relationships** | Pointer to [**UserJSONAPIResponseRelationships**](UserJSONAPIResponse_relationships.md) |  | [optional] 
**Type** | Pointer to **string** | JSONAPI Model Type | [optional] [default to "users"]

## Methods

### NewUserJSONAPIResponse

`func NewUserJSONAPIResponse() *UserJSONAPIResponse`

NewUserJSONAPIResponse instantiates a new UserJSONAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserJSONAPIResponseWithDefaults

`func NewUserJSONAPIResponseWithDefaults() *UserJSONAPIResponse`

NewUserJSONAPIResponseWithDefaults instantiates a new UserJSONAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UserJSONAPIResponse) GetAttributes() User`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserJSONAPIResponse) GetAttributesOk() (*User, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserJSONAPIResponse) SetAttributes(v User)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserJSONAPIResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *UserJSONAPIResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserJSONAPIResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserJSONAPIResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserJSONAPIResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *UserJSONAPIResponse) GetRelationships() UserJSONAPIResponseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *UserJSONAPIResponse) GetRelationshipsOk() (*UserJSONAPIResponseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *UserJSONAPIResponse) SetRelationships(v UserJSONAPIResponseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *UserJSONAPIResponse) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *UserJSONAPIResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserJSONAPIResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserJSONAPIResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserJSONAPIResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


