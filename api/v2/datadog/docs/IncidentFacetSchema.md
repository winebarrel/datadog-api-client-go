# IncidentFacetSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IncidentFacetSchemaDataOneOf**](IncidentFacetSchemaDataOneOf.md) | A facet definition. | [optional] 
**Type** | Pointer to [**IncidentFacetType**](IncidentFacetType.md) |  | [optional] 

## Methods

### NewIncidentFacetSchema

`func NewIncidentFacetSchema() *IncidentFacetSchema`

NewIncidentFacetSchema instantiates a new IncidentFacetSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentFacetSchemaWithDefaults

`func NewIncidentFacetSchemaWithDefaults() *IncidentFacetSchema`

NewIncidentFacetSchemaWithDefaults instantiates a new IncidentFacetSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IncidentFacetSchema) GetData() IncidentFacetSchemaDataOneOf`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IncidentFacetSchema) GetDataOk() (*IncidentFacetSchemaDataOneOf, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IncidentFacetSchema) SetData(v IncidentFacetSchemaDataOneOf)`

SetData sets Data field to given value.

### HasData

`func (o *IncidentFacetSchema) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *IncidentFacetSchema) GetType() IncidentFacetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentFacetSchema) GetTypeOk() (*IncidentFacetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentFacetSchema) SetType(v IncidentFacetType)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentFacetSchema) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


