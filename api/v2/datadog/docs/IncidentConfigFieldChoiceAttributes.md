# IncidentConfigFieldChoiceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when the field was created | [optional] [readonly] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when the field was modified | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the field | [optional] 
**UserDefinedFieldId** | Pointer to **string** | The ID of the associated field | [optional] 
**Value** | Pointer to **string** | value of the field | [optional] 

## Methods

### NewIncidentConfigFieldChoiceAttributes

`func NewIncidentConfigFieldChoiceAttributes() *IncidentConfigFieldChoiceAttributes`

NewIncidentConfigFieldChoiceAttributes instantiates a new IncidentConfigFieldChoiceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentConfigFieldChoiceAttributesWithDefaults

`func NewIncidentConfigFieldChoiceAttributesWithDefaults() *IncidentConfigFieldChoiceAttributes`

NewIncidentConfigFieldChoiceAttributesWithDefaults instantiates a new IncidentConfigFieldChoiceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IncidentConfigFieldChoiceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentConfigFieldChoiceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentConfigFieldChoiceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentConfigFieldChoiceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *IncidentConfigFieldChoiceAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentConfigFieldChoiceAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentConfigFieldChoiceAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentConfigFieldChoiceAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *IncidentConfigFieldChoiceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentConfigFieldChoiceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentConfigFieldChoiceAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentConfigFieldChoiceAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFieldId

`func (o *IncidentConfigFieldChoiceAttributes) GetUserDefinedFieldId() string`

GetUserDefinedFieldId returns the UserDefinedFieldId field if non-nil, zero value otherwise.

### GetUserDefinedFieldIdOk

`func (o *IncidentConfigFieldChoiceAttributes) GetUserDefinedFieldIdOk() (*string, bool)`

GetUserDefinedFieldIdOk returns a tuple with the UserDefinedFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFieldId

`func (o *IncidentConfigFieldChoiceAttributes) SetUserDefinedFieldId(v string)`

SetUserDefinedFieldId sets UserDefinedFieldId field to given value.

### HasUserDefinedFieldId

`func (o *IncidentConfigFieldChoiceAttributes) HasUserDefinedFieldId() bool`

HasUserDefinedFieldId returns a boolean if a field has been set.

### GetValue

`func (o *IncidentConfigFieldChoiceAttributes) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IncidentConfigFieldChoiceAttributes) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IncidentConfigFieldChoiceAttributes) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IncidentConfigFieldChoiceAttributes) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


