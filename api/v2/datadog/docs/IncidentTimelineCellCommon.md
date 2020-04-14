# IncidentTimelineCellCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CellType** | Pointer to **string** | A string representing the cell content type. | 
**Content** | Pointer to [**IncidentTimelineCellCommonContentOneOf**](IncidentTimelineCellCommonContentOneOf.md) | The cell content. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when a cell was created. | [optional] [readonly] 
**Deleted** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when a cell was deleted. | [optional] [readonly] 
**DisplayTime** | Pointer to [**time.Time**](time.Time.md) | Timestamp of display time for a given cell. This is used to sort the timeline view. | [optional] [readonly] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when a cell was modified. | [optional] [readonly] 
**Source** | Pointer to **NullableString** | A string representing the source of the cell. | [optional] 

## Methods

### NewIncidentTimelineCellCommon

`func NewIncidentTimelineCellCommon(cellType string, content IncidentTimelineCellCommonContentOneOf, ) *IncidentTimelineCellCommon`

NewIncidentTimelineCellCommon instantiates a new IncidentTimelineCellCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTimelineCellCommonWithDefaults

`func NewIncidentTimelineCellCommonWithDefaults() *IncidentTimelineCellCommon`

NewIncidentTimelineCellCommonWithDefaults instantiates a new IncidentTimelineCellCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCellType

`func (o *IncidentTimelineCellCommon) GetCellType() string`

GetCellType returns the CellType field if non-nil, zero value otherwise.

### GetCellTypeOk

`func (o *IncidentTimelineCellCommon) GetCellTypeOk() (*string, bool)`

GetCellTypeOk returns a tuple with the CellType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellType

`func (o *IncidentTimelineCellCommon) SetCellType(v string)`

SetCellType sets CellType field to given value.


### GetContent

`func (o *IncidentTimelineCellCommon) GetContent() IncidentTimelineCellCommonContentOneOf`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IncidentTimelineCellCommon) GetContentOk() (*IncidentTimelineCellCommonContentOneOf, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IncidentTimelineCellCommon) SetContent(v IncidentTimelineCellCommonContentOneOf)`

SetContent sets Content field to given value.


### GetCreated

`func (o *IncidentTimelineCellCommon) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentTimelineCellCommon) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentTimelineCellCommon) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentTimelineCellCommon) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeleted

`func (o *IncidentTimelineCellCommon) GetDeleted() time.Time`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IncidentTimelineCellCommon) GetDeletedOk() (*time.Time, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IncidentTimelineCellCommon) SetDeleted(v time.Time)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IncidentTimelineCellCommon) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDisplayTime

`func (o *IncidentTimelineCellCommon) GetDisplayTime() time.Time`

GetDisplayTime returns the DisplayTime field if non-nil, zero value otherwise.

### GetDisplayTimeOk

`func (o *IncidentTimelineCellCommon) GetDisplayTimeOk() (*time.Time, bool)`

GetDisplayTimeOk returns a tuple with the DisplayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTime

`func (o *IncidentTimelineCellCommon) SetDisplayTime(v time.Time)`

SetDisplayTime sets DisplayTime field to given value.

### HasDisplayTime

`func (o *IncidentTimelineCellCommon) HasDisplayTime() bool`

HasDisplayTime returns a boolean if a field has been set.

### GetModified

`func (o *IncidentTimelineCellCommon) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentTimelineCellCommon) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentTimelineCellCommon) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentTimelineCellCommon) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetSource

`func (o *IncidentTimelineCellCommon) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IncidentTimelineCellCommon) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IncidentTimelineCellCommon) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IncidentTimelineCellCommon) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *IncidentTimelineCellCommon) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *IncidentTimelineCellCommon) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


