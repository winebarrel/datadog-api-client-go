# IncidentSelectionRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choice** | Pointer to [**IncidentSelectionRelationshipsChoice**](IncidentSelection_relationships_choice.md) |  | [optional] 
**Field** | Pointer to [**IncidentSelectionRelationshipsField**](IncidentSelection_relationships_field.md) |  | [optional] 
**LastModifiedBy** | Pointer to [**UserJSONAPIID**](UserJSONAPIID.md) |  | [optional] 

## Methods

### NewIncidentSelectionRelationships

`func NewIncidentSelectionRelationships() *IncidentSelectionRelationships`

NewIncidentSelectionRelationships instantiates a new IncidentSelectionRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentSelectionRelationshipsWithDefaults

`func NewIncidentSelectionRelationshipsWithDefaults() *IncidentSelectionRelationships`

NewIncidentSelectionRelationshipsWithDefaults instantiates a new IncidentSelectionRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoice

`func (o *IncidentSelectionRelationships) GetChoice() IncidentSelectionRelationshipsChoice`

GetChoice returns the Choice field if non-nil, zero value otherwise.

### GetChoiceOk

`func (o *IncidentSelectionRelationships) GetChoiceOk() (*IncidentSelectionRelationshipsChoice, bool)`

GetChoiceOk returns a tuple with the Choice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoice

`func (o *IncidentSelectionRelationships) SetChoice(v IncidentSelectionRelationshipsChoice)`

SetChoice sets Choice field to given value.

### HasChoice

`func (o *IncidentSelectionRelationships) HasChoice() bool`

HasChoice returns a boolean if a field has been set.

### GetField

`func (o *IncidentSelectionRelationships) GetField() IncidentSelectionRelationshipsField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *IncidentSelectionRelationships) GetFieldOk() (*IncidentSelectionRelationshipsField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *IncidentSelectionRelationships) SetField(v IncidentSelectionRelationshipsField)`

SetField sets Field field to given value.

### HasField

`func (o *IncidentSelectionRelationships) HasField() bool`

HasField returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *IncidentSelectionRelationships) GetLastModifiedBy() UserJSONAPIID`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *IncidentSelectionRelationships) GetLastModifiedByOk() (*UserJSONAPIID, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *IncidentSelectionRelationships) SetLastModifiedBy(v UserJSONAPIID)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *IncidentSelectionRelationships) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


