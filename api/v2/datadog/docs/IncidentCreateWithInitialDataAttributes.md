# IncidentCreateWithInitialDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when an incident was created. | [optional] [readonly] 
**CustomerImpactEnd** | Pointer to [**NullableTime**](time.Time.md) | Timestamp of when cusomters/users were no longer impacted by the incident. | [optional] 
**CustomerImpactScope** | Pointer to **string** | The scope at which customers were impacted. | [optional] 
**CustomerImpactStart** | Pointer to [**NullableTime**](time.Time.md) | Timestamp of when cusomters/users were impacted by the incident. | [optional] 
**CustomerImpacted** | Pointer to **bool** | True if customers/users were impacted by the incident. | [optional] 
**Detected** | Pointer to [**NullableTime**](time.Time.md) | Timestamp of when an incident was detected. | [optional] 
**Duration** | Pointer to **int64** | Length of the incident representing the delta of customer_impact_end-customer_impact_start. | [optional] [readonly] 
**InitialCells** | Pointer to [**[]IncidentTimelineCellCommon**](IncidentTimelineCellCommon.md) | An array of initial timeline cell definitions to include. | [optional] 
**InitialUserDefinedFieldChoices** | Pointer to [**[]IncidentInitialUserDefinedFieldChoices**](IncidentInitialUserDefinedFieldChoices.md) | An array of initial choices when creating an incident. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when an incident was last modified. | [optional] [readonly] 
**PublicId** | Pointer to **int64** | Unique auto incrementing id for each created incident specific to the organization. | [optional] [readonly] 
**Resolved** | Pointer to [**NullableTime**](time.Time.md) | Timestamp of when an incident was resolved. | [optional] 
**State** | Pointer to **string** | The current state if an incident. | [optional] 
**TimeToDetect** | Pointer to **int64** | The amount of time in seconds to detect the incident. | [optional] [readonly] 
**TimeToInternalResponse** | Pointer to **int64** | The amount of time in seconds to detect the incidents occurence. Represents the delta of detected - created. | [optional] [readonly] 
**TimeToRepair** | Pointer to **int64** | The amount of time in seconds to resolve customer impact after detecting the issue. Represents the delta of customer_impact_end - detected. | [optional] [readonly] 
**Title** | Pointer to **string** | The title of the incident. | [optional] 

## Methods

### NewIncidentCreateWithInitialDataAttributes

`func NewIncidentCreateWithInitialDataAttributes() *IncidentCreateWithInitialDataAttributes`

NewIncidentCreateWithInitialDataAttributes instantiates a new IncidentCreateWithInitialDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentCreateWithInitialDataAttributesWithDefaults

`func NewIncidentCreateWithInitialDataAttributesWithDefaults() *IncidentCreateWithInitialDataAttributes`

NewIncidentCreateWithInitialDataAttributesWithDefaults instantiates a new IncidentCreateWithInitialDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IncidentCreateWithInitialDataAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentCreateWithInitialDataAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentCreateWithInitialDataAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentCreateWithInitialDataAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCustomerImpactEnd

`func (o *IncidentCreateWithInitialDataAttributes) GetCustomerImpactEnd() time.Time`

GetCustomerImpactEnd returns the CustomerImpactEnd field if non-nil, zero value otherwise.

### GetCustomerImpactEndOk

`func (o *IncidentCreateWithInitialDataAttributes) GetCustomerImpactEndOk() (*time.Time, bool)`

GetCustomerImpactEndOk returns a tuple with the CustomerImpactEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactEnd

`func (o *IncidentCreateWithInitialDataAttributes) SetCustomerImpactEnd(v time.Time)`

SetCustomerImpactEnd sets CustomerImpactEnd field to given value.

### HasCustomerImpactEnd

`func (o *IncidentCreateWithInitialDataAttributes) HasCustomerImpactEnd() bool`

HasCustomerImpactEnd returns a boolean if a field has been set.

### SetCustomerImpactEndNil

`func (o *IncidentCreateWithInitialDataAttributes) SetCustomerImpactEndNil(b bool)`

 SetCustomerImpactEndNil sets the value for CustomerImpactEnd to be an explicit nil

### UnsetCustomerImpactEnd
`func (o *IncidentCreateWithInitialDataAttributes) UnsetCustomerImpactEnd()`

UnsetCustomerImpactEnd ensures that no value is present for CustomerImpactEnd, not even an explicit nil
### GetCustomerImpactScope

`func (o *IncidentCreateWithInitialDataAttributes) GetCustomerImpactScope() string`

GetCustomerImpactScope returns the CustomerImpactScope field if non-nil, zero value otherwise.

### GetCustomerImpactScopeOk

`func (o *IncidentCreateWithInitialDataAttributes) GetCustomerImpactScopeOk() (*string, bool)`

GetCustomerImpactScopeOk returns a tuple with the CustomerImpactScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactScope

`func (o *IncidentCreateWithInitialDataAttributes) SetCustomerImpactScope(v string)`

SetCustomerImpactScope sets CustomerImpactScope field to given value.

### HasCustomerImpactScope

`func (o *IncidentCreateWithInitialDataAttributes) HasCustomerImpactScope() bool`

HasCustomerImpactScope returns a boolean if a field has been set.

### GetCustomerImpactStart

`func (o *IncidentCreateWithInitialDataAttributes) GetCustomerImpactStart() time.Time`

GetCustomerImpactStart returns the CustomerImpactStart field if non-nil, zero value otherwise.

### GetCustomerImpactStartOk

`func (o *IncidentCreateWithInitialDataAttributes) GetCustomerImpactStartOk() (*time.Time, bool)`

GetCustomerImpactStartOk returns a tuple with the CustomerImpactStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactStart

`func (o *IncidentCreateWithInitialDataAttributes) SetCustomerImpactStart(v time.Time)`

SetCustomerImpactStart sets CustomerImpactStart field to given value.

### HasCustomerImpactStart

`func (o *IncidentCreateWithInitialDataAttributes) HasCustomerImpactStart() bool`

HasCustomerImpactStart returns a boolean if a field has been set.

### SetCustomerImpactStartNil

`func (o *IncidentCreateWithInitialDataAttributes) SetCustomerImpactStartNil(b bool)`

 SetCustomerImpactStartNil sets the value for CustomerImpactStart to be an explicit nil

### UnsetCustomerImpactStart
`func (o *IncidentCreateWithInitialDataAttributes) UnsetCustomerImpactStart()`

UnsetCustomerImpactStart ensures that no value is present for CustomerImpactStart, not even an explicit nil
### GetCustomerImpacted

`func (o *IncidentCreateWithInitialDataAttributes) GetCustomerImpacted() bool`

GetCustomerImpacted returns the CustomerImpacted field if non-nil, zero value otherwise.

### GetCustomerImpactedOk

`func (o *IncidentCreateWithInitialDataAttributes) GetCustomerImpactedOk() (*bool, bool)`

GetCustomerImpactedOk returns a tuple with the CustomerImpacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpacted

`func (o *IncidentCreateWithInitialDataAttributes) SetCustomerImpacted(v bool)`

SetCustomerImpacted sets CustomerImpacted field to given value.

### HasCustomerImpacted

`func (o *IncidentCreateWithInitialDataAttributes) HasCustomerImpacted() bool`

HasCustomerImpacted returns a boolean if a field has been set.

### GetDetected

`func (o *IncidentCreateWithInitialDataAttributes) GetDetected() time.Time`

GetDetected returns the Detected field if non-nil, zero value otherwise.

### GetDetectedOk

`func (o *IncidentCreateWithInitialDataAttributes) GetDetectedOk() (*time.Time, bool)`

GetDetectedOk returns a tuple with the Detected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetected

`func (o *IncidentCreateWithInitialDataAttributes) SetDetected(v time.Time)`

SetDetected sets Detected field to given value.

### HasDetected

`func (o *IncidentCreateWithInitialDataAttributes) HasDetected() bool`

HasDetected returns a boolean if a field has been set.

### SetDetectedNil

`func (o *IncidentCreateWithInitialDataAttributes) SetDetectedNil(b bool)`

 SetDetectedNil sets the value for Detected to be an explicit nil

### UnsetDetected
`func (o *IncidentCreateWithInitialDataAttributes) UnsetDetected()`

UnsetDetected ensures that no value is present for Detected, not even an explicit nil
### GetDuration

`func (o *IncidentCreateWithInitialDataAttributes) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *IncidentCreateWithInitialDataAttributes) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *IncidentCreateWithInitialDataAttributes) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *IncidentCreateWithInitialDataAttributes) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetInitialCells

`func (o *IncidentCreateWithInitialDataAttributes) GetInitialCells() []IncidentTimelineCellCommon`

GetInitialCells returns the InitialCells field if non-nil, zero value otherwise.

### GetInitialCellsOk

`func (o *IncidentCreateWithInitialDataAttributes) GetInitialCellsOk() (*[]IncidentTimelineCellCommon, bool)`

GetInitialCellsOk returns a tuple with the InitialCells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCells

`func (o *IncidentCreateWithInitialDataAttributes) SetInitialCells(v []IncidentTimelineCellCommon)`

SetInitialCells sets InitialCells field to given value.

### HasInitialCells

`func (o *IncidentCreateWithInitialDataAttributes) HasInitialCells() bool`

HasInitialCells returns a boolean if a field has been set.

### GetInitialUserDefinedFieldChoices

`func (o *IncidentCreateWithInitialDataAttributes) GetInitialUserDefinedFieldChoices() []IncidentInitialUserDefinedFieldChoices`

GetInitialUserDefinedFieldChoices returns the InitialUserDefinedFieldChoices field if non-nil, zero value otherwise.

### GetInitialUserDefinedFieldChoicesOk

`func (o *IncidentCreateWithInitialDataAttributes) GetInitialUserDefinedFieldChoicesOk() (*[]IncidentInitialUserDefinedFieldChoices, bool)`

GetInitialUserDefinedFieldChoicesOk returns a tuple with the InitialUserDefinedFieldChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialUserDefinedFieldChoices

`func (o *IncidentCreateWithInitialDataAttributes) SetInitialUserDefinedFieldChoices(v []IncidentInitialUserDefinedFieldChoices)`

SetInitialUserDefinedFieldChoices sets InitialUserDefinedFieldChoices field to given value.

### HasInitialUserDefinedFieldChoices

`func (o *IncidentCreateWithInitialDataAttributes) HasInitialUserDefinedFieldChoices() bool`

HasInitialUserDefinedFieldChoices returns a boolean if a field has been set.

### GetModified

`func (o *IncidentCreateWithInitialDataAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentCreateWithInitialDataAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentCreateWithInitialDataAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentCreateWithInitialDataAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetPublicId

`func (o *IncidentCreateWithInitialDataAttributes) GetPublicId() int64`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *IncidentCreateWithInitialDataAttributes) GetPublicIdOk() (*int64, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *IncidentCreateWithInitialDataAttributes) SetPublicId(v int64)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *IncidentCreateWithInitialDataAttributes) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetResolved

`func (o *IncidentCreateWithInitialDataAttributes) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *IncidentCreateWithInitialDataAttributes) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *IncidentCreateWithInitialDataAttributes) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *IncidentCreateWithInitialDataAttributes) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *IncidentCreateWithInitialDataAttributes) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *IncidentCreateWithInitialDataAttributes) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetState

`func (o *IncidentCreateWithInitialDataAttributes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IncidentCreateWithInitialDataAttributes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IncidentCreateWithInitialDataAttributes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IncidentCreateWithInitialDataAttributes) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeToDetect

`func (o *IncidentCreateWithInitialDataAttributes) GetTimeToDetect() int64`

GetTimeToDetect returns the TimeToDetect field if non-nil, zero value otherwise.

### GetTimeToDetectOk

`func (o *IncidentCreateWithInitialDataAttributes) GetTimeToDetectOk() (*int64, bool)`

GetTimeToDetectOk returns a tuple with the TimeToDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToDetect

`func (o *IncidentCreateWithInitialDataAttributes) SetTimeToDetect(v int64)`

SetTimeToDetect sets TimeToDetect field to given value.

### HasTimeToDetect

`func (o *IncidentCreateWithInitialDataAttributes) HasTimeToDetect() bool`

HasTimeToDetect returns a boolean if a field has been set.

### GetTimeToInternalResponse

`func (o *IncidentCreateWithInitialDataAttributes) GetTimeToInternalResponse() int64`

GetTimeToInternalResponse returns the TimeToInternalResponse field if non-nil, zero value otherwise.

### GetTimeToInternalResponseOk

`func (o *IncidentCreateWithInitialDataAttributes) GetTimeToInternalResponseOk() (*int64, bool)`

GetTimeToInternalResponseOk returns a tuple with the TimeToInternalResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToInternalResponse

`func (o *IncidentCreateWithInitialDataAttributes) SetTimeToInternalResponse(v int64)`

SetTimeToInternalResponse sets TimeToInternalResponse field to given value.

### HasTimeToInternalResponse

`func (o *IncidentCreateWithInitialDataAttributes) HasTimeToInternalResponse() bool`

HasTimeToInternalResponse returns a boolean if a field has been set.

### GetTimeToRepair

`func (o *IncidentCreateWithInitialDataAttributes) GetTimeToRepair() int64`

GetTimeToRepair returns the TimeToRepair field if non-nil, zero value otherwise.

### GetTimeToRepairOk

`func (o *IncidentCreateWithInitialDataAttributes) GetTimeToRepairOk() (*int64, bool)`

GetTimeToRepairOk returns a tuple with the TimeToRepair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToRepair

`func (o *IncidentCreateWithInitialDataAttributes) SetTimeToRepair(v int64)`

SetTimeToRepair sets TimeToRepair field to given value.

### HasTimeToRepair

`func (o *IncidentCreateWithInitialDataAttributes) HasTimeToRepair() bool`

HasTimeToRepair returns a boolean if a field has been set.

### GetTitle

`func (o *IncidentCreateWithInitialDataAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IncidentCreateWithInitialDataAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IncidentCreateWithInitialDataAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IncidentCreateWithInitialDataAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


