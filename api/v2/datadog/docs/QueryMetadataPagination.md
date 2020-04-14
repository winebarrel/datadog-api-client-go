# QueryMetadataPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** | maximum number of pages to return | [optional] 
**NextOffset** | Pointer to **int64** | the next offset to retreive the next set of data | [optional] 
**Offset** | Pointer to **int64** | the index of the first element in the results | [optional] 

## Methods

### NewQueryMetadataPagination

`func NewQueryMetadataPagination() *QueryMetadataPagination`

NewQueryMetadataPagination instantiates a new QueryMetadataPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMetadataPaginationWithDefaults

`func NewQueryMetadataPaginationWithDefaults() *QueryMetadataPagination`

NewQueryMetadataPaginationWithDefaults instantiates a new QueryMetadataPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *QueryMetadataPagination) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QueryMetadataPagination) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QueryMetadataPagination) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QueryMetadataPagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextOffset

`func (o *QueryMetadataPagination) GetNextOffset() int64`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *QueryMetadataPagination) GetNextOffsetOk() (*int64, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *QueryMetadataPagination) SetNextOffset(v int64)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *QueryMetadataPagination) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetOffset

`func (o *QueryMetadataPagination) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QueryMetadataPagination) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QueryMetadataPagination) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QueryMetadataPagination) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


