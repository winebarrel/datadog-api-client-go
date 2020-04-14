# IncidentTimelineCellStatusChangeContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | The markdown content of the status change for the cell. | 
**StatusType** | Pointer to **string** | A string representing the status change type. | 

## Methods

### NewIncidentTimelineCellStatusChangeContent

`func NewIncidentTimelineCellStatusChangeContent(content string, statusType string, ) *IncidentTimelineCellStatusChangeContent`

NewIncidentTimelineCellStatusChangeContent instantiates a new IncidentTimelineCellStatusChangeContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTimelineCellStatusChangeContentWithDefaults

`func NewIncidentTimelineCellStatusChangeContentWithDefaults() *IncidentTimelineCellStatusChangeContent`

NewIncidentTimelineCellStatusChangeContentWithDefaults instantiates a new IncidentTimelineCellStatusChangeContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *IncidentTimelineCellStatusChangeContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IncidentTimelineCellStatusChangeContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IncidentTimelineCellStatusChangeContent) SetContent(v string)`

SetContent sets Content field to given value.


### GetStatusType

`func (o *IncidentTimelineCellStatusChangeContent) GetStatusType() string`

GetStatusType returns the StatusType field if non-nil, zero value otherwise.

### GetStatusTypeOk

`func (o *IncidentTimelineCellStatusChangeContent) GetStatusTypeOk() (*string, bool)`

GetStatusTypeOk returns a tuple with the StatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusType

`func (o *IncidentTimelineCellStatusChangeContent) SetStatusType(v string)`

SetStatusType sets StatusType field to given value.



### AsIncidentTimelineCellCommonContentOneOf

`func (s *IncidentTimelineCellStatusChangeContent) AsIncidentTimelineCellCommonContentOneOf() IncidentTimelineCellCommonContentOneOf`

Convenience method to wrap this instance of IncidentTimelineCellStatusChangeContent in IncidentTimelineCellCommonContentOneOf

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


