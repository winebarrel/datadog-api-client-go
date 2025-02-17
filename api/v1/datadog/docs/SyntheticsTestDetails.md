# SyntheticsTestDetails

## Properties

| Name          | Type                                                                           | Description                                    | Notes                 |
| ------------- | ------------------------------------------------------------------------------ | ---------------------------------------------- | --------------------- |
| **Config**    | Pointer to [**SyntheticsTestConfig**](SyntheticsTestConfig.md)                 |                                                | [optional]            |
| **Creator**   | Pointer to [**Creator**](Creator.md)                                           |                                                | [optional]            |
| **Locations** | Pointer to **[]string**                                                        | Array of locations used to run the test.       | [optional]            |
| **Message**   | Pointer to **string**                                                          | Notification message associated with the test. | [optional]            |
| **MonitorId** | Pointer to **int64**                                                           | The associated monitor ID.                     | [optional] [readonly] |
| **Name**      | Pointer to **string**                                                          | Name of the test.                              | [optional]            |
| **Options**   | Pointer to [**SyntheticsTestOptions**](SyntheticsTestOptions.md)               |                                                | [optional]            |
| **PublicId**  | Pointer to **string**                                                          | The test public ID.                            | [optional] [readonly] |
| **Status**    | Pointer to [**SyntheticsTestPauseStatus**](SyntheticsTestPauseStatus.md)       |                                                | [optional]            |
| **Steps**     | Pointer to [**[]SyntheticsStep**](SyntheticsStep.md)                           | For browser test, the steps of the test.       | [optional]            |
| **Subtype**   | Pointer to [**SyntheticsTestDetailsSubType**](SyntheticsTestDetailsSubType.md) |                                                | [optional]            |
| **Tags**      | Pointer to **[]string**                                                        | Array of tags attached to the test.            | [optional]            |
| **Type**      | Pointer to [**SyntheticsTestDetailsType**](SyntheticsTestDetailsType.md)       |                                                | [optional]            |

## Methods

### NewSyntheticsTestDetails

`func NewSyntheticsTestDetails() *SyntheticsTestDetails`

NewSyntheticsTestDetails instantiates a new SyntheticsTestDetails object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsTestDetailsWithDefaults

`func NewSyntheticsTestDetailsWithDefaults() *SyntheticsTestDetails`

NewSyntheticsTestDetailsWithDefaults instantiates a new SyntheticsTestDetails object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetConfig

`func (o *SyntheticsTestDetails) GetConfig() SyntheticsTestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SyntheticsTestDetails) GetConfigOk() (*SyntheticsTestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SyntheticsTestDetails) SetConfig(v SyntheticsTestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SyntheticsTestDetails) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreator

`func (o *SyntheticsTestDetails) GetCreator() Creator`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *SyntheticsTestDetails) GetCreatorOk() (*Creator, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *SyntheticsTestDetails) SetCreator(v Creator)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *SyntheticsTestDetails) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetLocations

`func (o *SyntheticsTestDetails) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticsTestDetails) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticsTestDetails) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SyntheticsTestDetails) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetMessage

`func (o *SyntheticsTestDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyntheticsTestDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyntheticsTestDetails) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SyntheticsTestDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMonitorId

`func (o *SyntheticsTestDetails) GetMonitorId() int64`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *SyntheticsTestDetails) GetMonitorIdOk() (*int64, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *SyntheticsTestDetails) SetMonitorId(v int64)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *SyntheticsTestDetails) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsTestDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsTestDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsTestDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyntheticsTestDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *SyntheticsTestDetails) GetOptions() SyntheticsTestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SyntheticsTestDetails) GetOptionsOk() (*SyntheticsTestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SyntheticsTestDetails) SetOptions(v SyntheticsTestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SyntheticsTestDetails) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPublicId

`func (o *SyntheticsTestDetails) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SyntheticsTestDetails) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *SyntheticsTestDetails) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *SyntheticsTestDetails) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticsTestDetails) GetStatus() SyntheticsTestPauseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticsTestDetails) GetStatusOk() (*SyntheticsTestPauseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticsTestDetails) SetStatus(v SyntheticsTestPauseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticsTestDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSteps

`func (o *SyntheticsTestDetails) GetSteps() []SyntheticsStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SyntheticsTestDetails) GetStepsOk() (*[]SyntheticsStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SyntheticsTestDetails) SetSteps(v []SyntheticsStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *SyntheticsTestDetails) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetSubtype

`func (o *SyntheticsTestDetails) GetSubtype() SyntheticsTestDetailsSubType`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *SyntheticsTestDetails) GetSubtypeOk() (*SyntheticsTestDetailsSubType, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *SyntheticsTestDetails) SetSubtype(v SyntheticsTestDetailsSubType)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *SyntheticsTestDetails) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTags

`func (o *SyntheticsTestDetails) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyntheticsTestDetails) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SyntheticsTestDetails) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SyntheticsTestDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SyntheticsTestDetails) GetType() SyntheticsTestDetailsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticsTestDetails) GetTypeOk() (*SyntheticsTestDetailsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticsTestDetails) SetType(v SyntheticsTestDetailsType)`

SetType sets Type field to given value.

### HasType

`func (o *SyntheticsTestDetails) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
