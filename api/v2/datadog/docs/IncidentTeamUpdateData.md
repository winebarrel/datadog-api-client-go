# IncidentTeamUpdateData

## Properties

| Name              | Type                                                                           | Description                 | Notes                               |
| ----------------- | ------------------------------------------------------------------------------ | --------------------------- | ----------------------------------- |
| **Attributes**    | Pointer to [**IncidentTeamUpdateAttributes**](IncidentTeamUpdateAttributes.md) |                             | [optional]                          |
| **Id**            | Pointer to **string**                                                          | The incident team&#39;s ID. | [optional]                          |
| **Relationships** | Pointer to [**IncidentTeamRelationships**](IncidentTeamRelationships.md)       |                             | [optional]                          |
| **Type**          | [**IncidentTeamType**](IncidentTeamType.md)                                    |                             | [default to INCIDENTTEAMTYPE_TEAMS] |

## Methods

### NewIncidentTeamUpdateData

`func NewIncidentTeamUpdateData(type_ IncidentTeamType) *IncidentTeamUpdateData`

NewIncidentTeamUpdateData instantiates a new IncidentTeamUpdateData object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIncidentTeamUpdateDataWithDefaults

`func NewIncidentTeamUpdateDataWithDefaults() *IncidentTeamUpdateData`

NewIncidentTeamUpdateDataWithDefaults instantiates a new IncidentTeamUpdateData object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAttributes

`func (o *IncidentTeamUpdateData) GetAttributes() IncidentTeamUpdateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IncidentTeamUpdateData) GetAttributesOk() (*IncidentTeamUpdateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IncidentTeamUpdateData) SetAttributes(v IncidentTeamUpdateAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IncidentTeamUpdateData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetId

`func (o *IncidentTeamUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentTeamUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentTeamUpdateData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentTeamUpdateData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelationships

`func (o *IncidentTeamUpdateData) GetRelationships() IncidentTeamRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IncidentTeamUpdateData) GetRelationshipsOk() (*IncidentTeamRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IncidentTeamUpdateData) SetRelationships(v IncidentTeamRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IncidentTeamUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetType

`func (o *IncidentTeamUpdateData) GetType() IncidentTeamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentTeamUpdateData) GetTypeOk() (*IncidentTeamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentTeamUpdateData) SetType(v IncidentTeamType)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
