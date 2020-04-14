# IncidentConfigFieldAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when the field was created. | [optional] [readonly] 
**DefaultChoiceId** | Pointer to **string** | The default Incident Config Choice/value for the provided field. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when the field was modified. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the field. | [optional] 
**TableId** | Pointer to [**IncidentConfigFieldTable**](IncidentConfigFieldTable.md) |  | [optional] 
**Type** | Pointer to [**IncidentConfigFieldType**](IncidentConfigFieldType.md) |  | [optional] 

## Methods

### NewIncidentConfigFieldAttributes

`func NewIncidentConfigFieldAttributes() *IncidentConfigFieldAttributes`

NewIncidentConfigFieldAttributes instantiates a new IncidentConfigFieldAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentConfigFieldAttributesWithDefaults

`func NewIncidentConfigFieldAttributesWithDefaults() *IncidentConfigFieldAttributes`

NewIncidentConfigFieldAttributesWithDefaults instantiates a new IncidentConfigFieldAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IncidentConfigFieldAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentConfigFieldAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentConfigFieldAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentConfigFieldAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDefaultChoiceId

`func (o *IncidentConfigFieldAttributes) GetDefaultChoiceId() string`

GetDefaultChoiceId returns the DefaultChoiceId field if non-nil, zero value otherwise.

### GetDefaultChoiceIdOk

`func (o *IncidentConfigFieldAttributes) GetDefaultChoiceIdOk() (*string, bool)`

GetDefaultChoiceIdOk returns a tuple with the DefaultChoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChoiceId

`func (o *IncidentConfigFieldAttributes) SetDefaultChoiceId(v string)`

SetDefaultChoiceId sets DefaultChoiceId field to given value.

### HasDefaultChoiceId

`func (o *IncidentConfigFieldAttributes) HasDefaultChoiceId() bool`

HasDefaultChoiceId returns a boolean if a field has been set.

### GetModified

`func (o *IncidentConfigFieldAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentConfigFieldAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentConfigFieldAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentConfigFieldAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetName

`func (o *IncidentConfigFieldAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentConfigFieldAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentConfigFieldAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentConfigFieldAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTableId

`func (o *IncidentConfigFieldAttributes) GetTableId() IncidentConfigFieldTable`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *IncidentConfigFieldAttributes) GetTableIdOk() (*IncidentConfigFieldTable, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *IncidentConfigFieldAttributes) SetTableId(v IncidentConfigFieldTable)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *IncidentConfigFieldAttributes) HasTableId() bool`

HasTableId returns a boolean if a field has been set.

### GetType

`func (o *IncidentConfigFieldAttributes) GetType() IncidentConfigFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentConfigFieldAttributes) GetTypeOk() (*IncidentConfigFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentConfigFieldAttributes) SetType(v IncidentConfigFieldType)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentConfigFieldAttributes) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


