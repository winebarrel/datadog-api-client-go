# IncidentAttributes

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
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of when an incident was last modified. | [optional] [readonly] 
**PublicId** | Pointer to **int64** | Unique auto incrementing id for each created incident specific to the organization. | [optional] [readonly] 
**Resolved** | Pointer to [**NullableTime**](time.Time.md) | Timestamp of when an incident was resolved. | [optional] 
**State** | Pointer to **string** | The current state if an incident. | [optional] 
**TimeToDetect** | Pointer to **int64** | The amount of time in seconds to detect the incident. | [optional] [readonly] 
**TimeToInternalResponse** | Pointer to **int64** | The amount of time in seconds to detect the incidents occurence. Represents the delta of detected - created. | [optional] [readonly] 
**TimeToRepair** | Pointer to **int64** | The amount of time in seconds to resolve customer impact after detecting the issue. Represents the delta of customer_impact_end - detected. | [optional] [readonly] 
**Title** | Pointer to **string** | The title of the incident. | [optional] 

## Methods

### NewIncidentAttributes

`func NewIncidentAttributes() *IncidentAttributes`

NewIncidentAttributes instantiates a new IncidentAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentAttributesWithDefaults

`func NewIncidentAttributesWithDefaults() *IncidentAttributes`

NewIncidentAttributesWithDefaults instantiates a new IncidentAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IncidentAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IncidentAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IncidentAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IncidentAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCustomerImpactEnd

`func (o *IncidentAttributes) GetCustomerImpactEnd() time.Time`

GetCustomerImpactEnd returns the CustomerImpactEnd field if non-nil, zero value otherwise.

### GetCustomerImpactEndOk

`func (o *IncidentAttributes) GetCustomerImpactEndOk() (*time.Time, bool)`

GetCustomerImpactEndOk returns a tuple with the CustomerImpactEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactEnd

`func (o *IncidentAttributes) SetCustomerImpactEnd(v time.Time)`

SetCustomerImpactEnd sets CustomerImpactEnd field to given value.

### HasCustomerImpactEnd

`func (o *IncidentAttributes) HasCustomerImpactEnd() bool`

HasCustomerImpactEnd returns a boolean if a field has been set.

### SetCustomerImpactEndNil

`func (o *IncidentAttributes) SetCustomerImpactEndNil(b bool)`

 SetCustomerImpactEndNil sets the value for CustomerImpactEnd to be an explicit nil

### UnsetCustomerImpactEnd
`func (o *IncidentAttributes) UnsetCustomerImpactEnd()`

UnsetCustomerImpactEnd ensures that no value is present for CustomerImpactEnd, not even an explicit nil
### GetCustomerImpactScope

`func (o *IncidentAttributes) GetCustomerImpactScope() string`

GetCustomerImpactScope returns the CustomerImpactScope field if non-nil, zero value otherwise.

### GetCustomerImpactScopeOk

`func (o *IncidentAttributes) GetCustomerImpactScopeOk() (*string, bool)`

GetCustomerImpactScopeOk returns a tuple with the CustomerImpactScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactScope

`func (o *IncidentAttributes) SetCustomerImpactScope(v string)`

SetCustomerImpactScope sets CustomerImpactScope field to given value.

### HasCustomerImpactScope

`func (o *IncidentAttributes) HasCustomerImpactScope() bool`

HasCustomerImpactScope returns a boolean if a field has been set.

### GetCustomerImpactStart

`func (o *IncidentAttributes) GetCustomerImpactStart() time.Time`

GetCustomerImpactStart returns the CustomerImpactStart field if non-nil, zero value otherwise.

### GetCustomerImpactStartOk

`func (o *IncidentAttributes) GetCustomerImpactStartOk() (*time.Time, bool)`

GetCustomerImpactStartOk returns a tuple with the CustomerImpactStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpactStart

`func (o *IncidentAttributes) SetCustomerImpactStart(v time.Time)`

SetCustomerImpactStart sets CustomerImpactStart field to given value.

### HasCustomerImpactStart

`func (o *IncidentAttributes) HasCustomerImpactStart() bool`

HasCustomerImpactStart returns a boolean if a field has been set.

### SetCustomerImpactStartNil

`func (o *IncidentAttributes) SetCustomerImpactStartNil(b bool)`

 SetCustomerImpactStartNil sets the value for CustomerImpactStart to be an explicit nil

### UnsetCustomerImpactStart
`func (o *IncidentAttributes) UnsetCustomerImpactStart()`

UnsetCustomerImpactStart ensures that no value is present for CustomerImpactStart, not even an explicit nil
### GetCustomerImpacted

`func (o *IncidentAttributes) GetCustomerImpacted() bool`

GetCustomerImpacted returns the CustomerImpacted field if non-nil, zero value otherwise.

### GetCustomerImpactedOk

`func (o *IncidentAttributes) GetCustomerImpactedOk() (*bool, bool)`

GetCustomerImpactedOk returns a tuple with the CustomerImpacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerImpacted

`func (o *IncidentAttributes) SetCustomerImpacted(v bool)`

SetCustomerImpacted sets CustomerImpacted field to given value.

### HasCustomerImpacted

`func (o *IncidentAttributes) HasCustomerImpacted() bool`

HasCustomerImpacted returns a boolean if a field has been set.

### GetDetected

`func (o *IncidentAttributes) GetDetected() time.Time`

GetDetected returns the Detected field if non-nil, zero value otherwise.

### GetDetectedOk

`func (o *IncidentAttributes) GetDetectedOk() (*time.Time, bool)`

GetDetectedOk returns a tuple with the Detected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetected

`func (o *IncidentAttributes) SetDetected(v time.Time)`

SetDetected sets Detected field to given value.

### HasDetected

`func (o *IncidentAttributes) HasDetected() bool`

HasDetected returns a boolean if a field has been set.

### SetDetectedNil

`func (o *IncidentAttributes) SetDetectedNil(b bool)`

 SetDetectedNil sets the value for Detected to be an explicit nil

### UnsetDetected
`func (o *IncidentAttributes) UnsetDetected()`

UnsetDetected ensures that no value is present for Detected, not even an explicit nil
### GetDuration

`func (o *IncidentAttributes) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *IncidentAttributes) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *IncidentAttributes) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *IncidentAttributes) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetModified

`func (o *IncidentAttributes) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IncidentAttributes) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IncidentAttributes) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IncidentAttributes) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetPublicId

`func (o *IncidentAttributes) GetPublicId() int64`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *IncidentAttributes) GetPublicIdOk() (*int64, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *IncidentAttributes) SetPublicId(v int64)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *IncidentAttributes) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetResolved

`func (o *IncidentAttributes) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *IncidentAttributes) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *IncidentAttributes) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *IncidentAttributes) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *IncidentAttributes) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *IncidentAttributes) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetState

`func (o *IncidentAttributes) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IncidentAttributes) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IncidentAttributes) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IncidentAttributes) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeToDetect

`func (o *IncidentAttributes) GetTimeToDetect() int64`

GetTimeToDetect returns the TimeToDetect field if non-nil, zero value otherwise.

### GetTimeToDetectOk

`func (o *IncidentAttributes) GetTimeToDetectOk() (*int64, bool)`

GetTimeToDetectOk returns a tuple with the TimeToDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToDetect

`func (o *IncidentAttributes) SetTimeToDetect(v int64)`

SetTimeToDetect sets TimeToDetect field to given value.

### HasTimeToDetect

`func (o *IncidentAttributes) HasTimeToDetect() bool`

HasTimeToDetect returns a boolean if a field has been set.

### GetTimeToInternalResponse

`func (o *IncidentAttributes) GetTimeToInternalResponse() int64`

GetTimeToInternalResponse returns the TimeToInternalResponse field if non-nil, zero value otherwise.

### GetTimeToInternalResponseOk

`func (o *IncidentAttributes) GetTimeToInternalResponseOk() (*int64, bool)`

GetTimeToInternalResponseOk returns a tuple with the TimeToInternalResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToInternalResponse

`func (o *IncidentAttributes) SetTimeToInternalResponse(v int64)`

SetTimeToInternalResponse sets TimeToInternalResponse field to given value.

### HasTimeToInternalResponse

`func (o *IncidentAttributes) HasTimeToInternalResponse() bool`

HasTimeToInternalResponse returns a boolean if a field has been set.

### GetTimeToRepair

`func (o *IncidentAttributes) GetTimeToRepair() int64`

GetTimeToRepair returns the TimeToRepair field if non-nil, zero value otherwise.

### GetTimeToRepairOk

`func (o *IncidentAttributes) GetTimeToRepairOk() (*int64, bool)`

GetTimeToRepairOk returns a tuple with the TimeToRepair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToRepair

`func (o *IncidentAttributes) SetTimeToRepair(v int64)`

SetTimeToRepair sets TimeToRepair field to given value.

### HasTimeToRepair

`func (o *IncidentAttributes) HasTimeToRepair() bool`

HasTimeToRepair returns a boolean if a field has been set.

### GetTitle

`func (o *IncidentAttributes) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IncidentAttributes) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IncidentAttributes) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IncidentAttributes) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


