# ServiceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when the service was created. | [optional] [readonly] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when the service was modified. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the service. | [optional] 

## Methods

### NewServiceAttributes

`func NewServiceAttributes() *ServiceAttributes`

NewServiceAttributes instantiates a new ServiceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAttributesWithDefaults

`func NewServiceAttributesWithDefaults() *ServiceAttributes`

NewServiceAttributesWithDefaults instantiates a new ServiceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ServiceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServiceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServiceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ServiceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ServiceAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ServiceAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ServiceAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ServiceAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *ServiceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceAttributes) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


