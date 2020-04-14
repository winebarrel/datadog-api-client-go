# IncidentSelectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when an selection was created. | [optional] [readonly] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when an selection was modified. | [optional] [readonly] 
**ObjectId** | Pointer to **string** | Represents the incident ID. | [optional] [readonly] 

## Methods

### NewIncidentSelectionAttributes

`func NewIncidentSelectionAttributes() *IncidentSelectionAttributes`

NewIncidentSelectionAttributes instantiates a new IncidentSelectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSelectionAttributesWithDefaults

`func NewIncidentSelectionAttributesWithDefaults() *IncidentSelectionAttributes`

NewIncidentSelectionAttributesWithDefaults instantiates a new IncidentSelectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IncidentSelectionAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentSelectionAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentSelectionAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentSelectionAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *IncidentSelectionAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentSelectionAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentSelectionAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentSelectionAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetObjectId

`func (o *IncidentSelectionAttributes) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *IncidentSelectionAttributes) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *IncidentSelectionAttributes) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *IncidentSelectionAttributes) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


